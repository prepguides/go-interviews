package controllers

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webserverv1alpha1 "github.com/webserver/webserver-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// WebserverReconciler reconciles a Webserver object
type WebserverReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webserver.io,resources=webservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webserver.io,resources=webservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webserver.io,resources=webservers/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *WebserverReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the Webserver instance
	webserver := &webserverv1alpha1.Webserver{}
	err := r.Get(ctx, req.NamespacedName, webserver)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Return and don't requeue
			log.Info("Webserver resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get Webserver")
		return ctrl.Result{}, err
	}

	// Set default values
	if webserver.Spec.Replicas == 0 {
		webserver.Spec.Replicas = 1
	}
	if webserver.Spec.Image == "" {
		webserver.Spec.Image = "nginx:1.25"
	}
	if webserver.Spec.Port == 0 {
		webserver.Spec.Port = 80
	}
	if webserver.Spec.ServiceType == "" {
		webserver.Spec.ServiceType = "ClusterIP"
	}
	if webserver.Spec.Config.Title == "" {
		webserver.Spec.Config.Title = "Webserver Operator Demo"
	}
	if webserver.Spec.Config.Message == "" {
		webserver.Spec.Config.Message = "Welcome to the Webserver Operator Demo!"
	}
	if webserver.Spec.Config.Color == "" {
		webserver.Spec.Config.Color = "#f0f0f0"
	}

	// Update the status
	webserver.Status.ObservedGeneration = webserver.Generation
	webserver.Status.Phase = "Reconciling"

	// Create or update the deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webserver.Name + "-deployment",
			Namespace: webserver.Namespace,
		},
	}

	op, err := ctrl.CreateOrUpdate(ctx, r.Client, deployment, func() error {
		return r.mutateDeployment(deployment, webserver)
	})
	if err != nil {
		log.Error(err, "Failed to create or update deployment")
		return ctrl.Result{}, err
	}

	if op != controllerutil.OperationResultNone {
		log.Info("Deployment operation", "operation", op)
	}

	// Create or update the configmap
	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webserver.Name + "-config",
			Namespace: webserver.Namespace,
		},
	}

	op, err = ctrl.CreateOrUpdate(ctx, r.Client, configmap, func() error {
		return r.mutateConfigMap(configmap, webserver)
	})
	if err != nil {
		log.Error(err, "Failed to create or update configmap")
		return ctrl.Result{}, err
	}

	if op != controllerutil.OperationResultNone {
		log.Info("ConfigMap operation", "operation", op)
	}

	// Create or update the service
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webserver.Name + "-service",
			Namespace: webserver.Namespace,
		},
	}

	op, err = ctrl.CreateOrUpdate(ctx, r.Client, service, func() error {
		return r.mutateService(service, webserver)
	})
	if err != nil {
		log.Error(err, "Failed to create or update service")
		return ctrl.Result{}, err
	}

	if op != controllerutil.OperationResultNone {
		log.Info("Service operation", "operation", op)
	}

	// Update status with deployment information
	if err := r.updateStatus(ctx, webserver); err != nil {
		log.Error(err, "Failed to update status")
		return ctrl.Result{}, err
	}

	// Update the final status
	webserver.Status.Phase = "Ready"
	if err := r.Status().Update(ctx, webserver); err != nil {
		log.Error(err, "Failed to update Webserver status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: time.Minute * 5}, nil
}

// mutateDeployment creates or updates the deployment
func (r *WebserverReconciler) mutateDeployment(deployment *appsv1.Deployment, webserver *webserverv1alpha1.Webserver) error {
	// Set the owner reference
	if err := ctrl.SetControllerReference(webserver, deployment, r.Scheme); err != nil {
		return err
	}

	// Set labels
	deployment.Labels = map[string]string{
		"app":        "webserver",
		"instance":   webserver.Name,
		"managed-by": "webserver-operator",
	}

	// Set spec
	deployment.Spec = appsv1.DeploymentSpec{
		Replicas: &webserver.Spec.Replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app":      "webserver",
				"instance": webserver.Name,
			},
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app":      "webserver",
					"instance": webserver.Name,
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "webserver",
						Image: webserver.Spec.Image,
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: webserver.Spec.Port,
								Name:          "http",
							},
						},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      "html-content",
								MountPath: "/usr/share/nginx/html",
								ReadOnly:  true,
							},
						},
					},
				},
				Volumes: []corev1.Volume{
					{
						Name: "html-content",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: webserver.Name + "-config",
								},
							},
						},
					},
				},
			},
		},
	}

	return nil
}

// mutateService creates or updates the service
func (r *WebserverReconciler) mutateService(service *corev1.Service, webserver *webserverv1alpha1.Webserver) error {
	// Set the owner reference
	if err := ctrl.SetControllerReference(webserver, service, r.Scheme); err != nil {
		return err
	}

	// Set labels
	service.Labels = map[string]string{
		"app":        "webserver",
		"instance":   webserver.Name,
		"managed-by": "webserver-operator",
	}

	// Set spec
	service.Spec = corev1.ServiceSpec{
		Selector: map[string]string{
			"app":      "webserver",
			"instance": webserver.Name,
		},
		Ports: []corev1.ServicePort{
			{
				Port:       80,
				TargetPort: intstr.FromInt(int(webserver.Spec.Port)),
				Name:       "http",
			},
		},
		Type: corev1.ServiceType(webserver.Spec.ServiceType),
	}

	return nil
}

// mutateConfigMap creates or updates the configmap with HTML content
func (r *WebserverReconciler) mutateConfigMap(configmap *corev1.ConfigMap, webserver *webserverv1alpha1.Webserver) error {
	// Set the owner reference
	if err := ctrl.SetControllerReference(webserver, configmap, r.Scheme); err != nil {
		return err
	}

	// Set labels
	configmap.Labels = map[string]string{
		"app":        "webserver",
		"instance":   webserver.Name,
		"managed-by": "webserver-operator",
	}

	// Generate HTML content
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: %s;
            margin: 0;
            padding: 20px;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }
        .container {
            text-align: center;
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            max-width: 600px;
        }
        h1 {
            color: #333;
            margin-bottom: 20px;
        }
        p {
            color: #666;
            font-size: 18px;
            line-height: 1.6;
        }
        .info {
            margin-top: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 5px;
            text-align: left;
        }
        .info h3 {
            margin-top: 0;
            color: #495057;
        }
        .info ul {
            color: #6c757d;
        }
        .status {
            margin-top: 20px;
            padding: 10px;
            background: #d4edda;
            border: 1px solid #c3e6cb;
            border-radius: 5px;
            color: #155724;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        <p>%s</p>
        
        <div class="info">
            <h3>Webserver Operator Demo</h3>
            <ul>
                <li><strong>Instance:</strong> %s</li>
                <li><strong>Namespace:</strong> %s</li>
                <li><strong>Replicas:</strong> %d</li>
                <li><strong>Image:</strong> %s</li>
                <li><strong>Port:</strong> %d</li>
                <li><strong>Service Type:</strong> %s</li>
                <li><strong>Generated:</strong> %s</li>
            </ul>
        </div>
        
        <div class="status">
            ✅ Web server is running successfully!
        </div>
    </div>
</body>
</html>`,
		webserver.Spec.Config.Title,
		webserver.Spec.Config.Color,
		webserver.Spec.Config.Title,
		webserver.Spec.Config.Message,
		webserver.Name,
		webserver.Namespace,
		webserver.Spec.Replicas,
		webserver.Spec.Image,
		webserver.Spec.Port,
		webserver.Spec.ServiceType,
		time.Now().Format("2006-01-02 15:04:05 MST"))

	// Set the HTML content
	configmap.Data = map[string]string{
		"index.html": htmlContent,
	}

	return nil
}

// updateStatus updates the status of the Webserver resource
func (r *WebserverReconciler) updateStatus(ctx context.Context, webserver *webserverv1alpha1.Webserver) error {
	// Get the deployment
	deployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{
		Name:      webserver.Name + "-deployment",
		Namespace: webserver.Namespace,
	}, deployment)
	if err != nil {
		return err
	}

	// Update ready replicas
	webserver.Status.ReadyReplicas = deployment.Status.ReadyReplicas

	// Add conditions
	condition := metav1.Condition{
		Type:               "Ready",
		Status:             metav1.ConditionTrue,
		LastTransitionTime: metav1.Now(),
		Reason:             "ReconciliationSucceeded",
		Message:            "Webserver is ready",
	}

	if deployment.Status.ReadyReplicas != webserver.Spec.Replicas {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "ReplicasNotReady"
		condition.Message = fmt.Sprintf("Expected %d replicas, got %d", webserver.Spec.Replicas, deployment.Status.ReadyReplicas)
	}

	// Update conditions
	webserver.Status.Conditions = []metav1.Condition{condition}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WebserverReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webserverv1alpha1.Webserver{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
