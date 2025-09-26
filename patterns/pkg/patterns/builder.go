package patterns

import (
	"context"
	"fmt"
	"time"
)

// Builder pattern implementation - common in Go for complex object construction
// This demonstrates method chaining and fluent interfaces

// WebServerConfig represents the configuration for a web server
type WebServerConfig struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxConnections int
	EnableTLS      bool
	CertFile       string
	KeyFile        string
	Middlewares    []string
}

// WebServerConfigBuilder provides a fluent interface for building WebServerConfig
type WebServerConfigBuilder struct {
	config *WebServerConfig
}

// NewWebServerConfigBuilder creates a new builder
func NewWebServerConfigBuilder() *WebServerConfigBuilder {
	return &WebServerConfigBuilder{
		config: &WebServerConfig{
			Host:           "localhost",
			Port:           8080,
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxConnections: 1000,
			EnableTLS:      false,
		},
	}
}

// Host sets the host
func (b *WebServerConfigBuilder) Host(host string) *WebServerConfigBuilder {
	b.config.Host = host
	return b
}

// Port sets the port
func (b *WebServerConfigBuilder) Port(port int) *WebServerConfigBuilder {
	b.config.Port = port
	return b
}

// ReadTimeout sets the read timeout
func (b *WebServerConfigBuilder) ReadTimeout(timeout time.Duration) *WebServerConfigBuilder {
	b.config.ReadTimeout = timeout
	return b
}

// WriteTimeout sets the write timeout
func (b *WebServerConfigBuilder) WriteTimeout(timeout time.Duration) *WebServerConfigBuilder {
	b.config.WriteTimeout = timeout
	return b
}

// MaxConnections sets the maximum number of connections
func (b *WebServerConfigBuilder) MaxConnections(max int) *WebServerConfigBuilder {
	b.config.MaxConnections = max
	return b
}

// WithTLS enables TLS with the given certificate files
func (b *WebServerConfigBuilder) WithTLS(certFile, keyFile string) *WebServerConfigBuilder {
	b.config.EnableTLS = true
	b.config.CertFile = certFile
	b.config.KeyFile = keyFile
	return b
}

// AddMiddleware adds a middleware to the configuration
func (b *WebServerConfigBuilder) AddMiddleware(middleware string) *WebServerConfigBuilder {
	b.config.Middlewares = append(b.config.Middlewares, middleware)
	return b
}

// Build returns the built configuration
func (b *WebServerConfigBuilder) Build() *WebServerConfig {
	return b.config
}

// Validate validates the configuration
func (b *WebServerConfigBuilder) Validate() error {
	if b.config.Host == "" {
		return fmt.Errorf("host cannot be empty")
	}
	if b.config.Port <= 0 || b.config.Port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	if b.config.EnableTLS && (b.config.CertFile == "" || b.config.KeyFile == "") {
		return fmt.Errorf("certificate and key files are required when TLS is enabled")
	}
	return nil
}

// WebServer represents a web server that can be configured using the builder
type WebServer struct {
	config *WebServerConfig
}

// NewWebServer creates a new web server with the given configuration
func NewWebServer(config *WebServerConfig) *WebServer {
	return &WebServer{
		config: config,
	}
}

// Start starts the web server
func (ws *WebServer) Start(ctx context.Context) error {
	// Simulate starting the web server
	fmt.Printf("Starting web server on %s:%d\n", ws.config.Host, ws.config.Port)
	if ws.config.EnableTLS {
		fmt.Printf("TLS enabled with cert: %s, key: %s\n", ws.config.CertFile, ws.config.KeyFile)
	}
	fmt.Printf("Read timeout: %v, Write timeout: %v\n", ws.config.ReadTimeout, ws.config.WriteTimeout)
	fmt.Printf("Max connections: %d\n", ws.config.MaxConnections)
	if len(ws.config.Middlewares) > 0 {
		fmt.Printf("Middlewares: %v\n", ws.config.Middlewares)
	}
	return nil
}
