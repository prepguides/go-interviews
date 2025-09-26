package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WebserverSpec defines the desired state of Webserver
type WebserverSpec struct {
	// Replicas is the number of desired replicas for the web server deployment
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the container image to use for the web server
	Image string `json:"image,omitempty"`

	// Port is the port the web server listens on
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	// ServiceType is the type of Kubernetes service to create
	ServiceType string `json:"serviceType,omitempty"`

	// Config contains configuration options for the web server
	Config WebserverConfig `json:"config,omitempty"`
}

// WebserverConfig defines configuration options for the web server
type WebserverConfig struct {
	// Title is the title displayed on the web page
	Title string `json:"title,omitempty"`

	// Message is the message displayed on the web page
	Message string `json:"message,omitempty"`

	// Color is the background color of the web page
	Color string `json:"color,omitempty"`

	// Features enables/disables specific features
	Features map[string]bool `json:"features,omitempty"`
}

// WebserverStatus defines the observed state of Webserver
type WebserverStatus struct {
	// Conditions represent the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed Webserver resource
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// ReadyReplicas is the number of ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Phase represents the current phase of the Webserver deployment
	Phase string `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.readyReplicas
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Ready",type="integer",JSONPath=".status.readyReplicas"
//+kubebuilder:printcolumn:name="Desired",type="integer",JSONPath=".spec.replicas"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Webserver is the Schema for the webservers API
type Webserver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebserverSpec   `json:"spec,omitempty"`
	Status WebserverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WebserverList contains a list of Webserver
type WebserverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webserver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Webserver{}, &WebserverList{})
}
