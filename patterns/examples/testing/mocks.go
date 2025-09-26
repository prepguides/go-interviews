package testing

import (
	"context"
	"fmt"
	"time"
)

// MockLogger implements the Logger interface for testing
type MockLogger struct {
	DebugCalls []LogCall
	InfoCalls  []LogCall
	WarnCalls  []LogCall
	ErrorCalls []LogCall
}

type LogCall struct {
	Msg  string
	Args []interface{}
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		DebugCalls: make([]LogCall, 0),
		InfoCalls:  make([]LogCall, 0),
		WarnCalls:  make([]LogCall, 0),
		ErrorCalls: make([]LogCall, 0),
	}
}

func (m *MockLogger) Debug(msg string, keysAndValues ...interface{}) {
	m.DebugCalls = append(m.DebugCalls, LogCall{Msg: msg, Args: keysAndValues})
}

func (m *MockLogger) Info(msg string, keysAndValues ...interface{}) {
	m.InfoCalls = append(m.InfoCalls, LogCall{Msg: msg, Args: keysAndValues})
}

func (m *MockLogger) Warn(msg string, keysAndValues ...interface{}) {
	m.WarnCalls = append(m.WarnCalls, LogCall{Msg: msg, Args: keysAndValues})
}

func (m *MockLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ErrorCalls = append(m.ErrorCalls, LogCall{Msg: msg, Args: keysAndValues})
}

func (m *MockLogger) WithValues(keysAndValues ...interface{}) Logger {
	return m // Simplified for testing
}

func (m *MockLogger) WithName(name string) Logger {
	return m // Simplified for testing
}

// MockResourceManager implements the ResourceManager interface for testing
type MockResourceManager struct {
	Resources map[string]interface{}
	GetError  error
	CreateError error
	UpdateError error
	DeleteError error
	ListError  error
}

func NewMockResourceManager() *MockResourceManager {
	return &MockResourceManager{
		Resources: make(map[string]interface{}),
	}
}

func (m *MockResourceManager) Get(ctx context.Context, key string) (interface{}, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	return m.Resources[key], nil
}

func (m *MockResourceManager) Create(ctx context.Context, obj interface{}) error {
	if m.CreateError != nil {
		return m.CreateError
	}
	// Simulate creating a resource with a key
	key := fmt.Sprintf("resource-%d", time.Now().UnixNano())
	m.Resources[key] = obj
	return nil
}

func (m *MockResourceManager) Update(ctx context.Context, obj interface{}) error {
	if m.UpdateError != nil {
		return m.UpdateError
	}
	// Simplified update - just store the object
	key := fmt.Sprintf("resource-%d", time.Now().UnixNano())
	m.Resources[key] = obj
	return nil
}

func (m *MockResourceManager) Delete(ctx context.Context, key string) error {
	if m.DeleteError != nil {
		return m.DeleteError
	}
	delete(m.Resources, key)
	return nil
}

func (m *MockResourceManager) List(ctx context.Context, selector string) ([]interface{}, error) {
	if m.ListError != nil {
		return nil, m.ListError
	}
	results := make([]interface{}, 0, len(m.Resources))
	for _, resource := range m.Resources {
		results = append(results, resource)
	}
	return results, nil
}

// MockHealthChecker implements the HealthChecker interface for testing
type MockHealthChecker struct {
	Healthy bool
	Message string
	Details map[string]interface{}
	CheckError error
}

func NewMockHealthChecker() *MockHealthChecker {
	return &MockHealthChecker{
		Healthy: true,
		Message: "healthy",
		Details: make(map[string]interface{}),
	}
}

func (m *MockHealthChecker) CheckHealth(ctx context.Context) error {
	return m.CheckError
}

func (m *MockHealthChecker) IsHealthy() bool {
	return m.Healthy
}

func (m *MockHealthChecker) GetHealthStatus() HealthStatus {
	return HealthStatus{
		Healthy: m.Healthy,
		Message: m.Message,
		Details: m.Details,
	}
}

// TestHelper provides common testing utilities
type TestHelper struct {
	MockLogger         *MockLogger
	MockResourceManager *MockResourceManager
	MockHealthChecker  *MockHealthChecker
}

func NewTestHelper() *TestHelper {
	return &TestHelper{
		MockLogger:         NewMockLogger(),
		MockResourceManager: NewMockResourceManager(),
		MockHealthChecker:  NewMockHealthChecker(),
	}
}

// AssertLogContains checks if a log call contains expected content
func (th *TestHelper) AssertLogContains(level string, expectedMsg string) bool {
	var calls []LogCall
	
	switch level {
	case "debug":
		calls = th.MockLogger.DebugCalls
	case "info":
		calls = th.MockLogger.InfoCalls
	case "warn":
		calls = th.MockLogger.WarnCalls
	case "error":
		calls = th.MockLogger.ErrorCalls
	default:
		return false
	}
	
	for _, call := range calls {
		if call.Msg == expectedMsg {
			return true
		}
	}
	return false
}

// AssertResourceExists checks if a resource exists in the mock manager
func (th *TestHelper) AssertResourceExists(key string) bool {
	_, exists := th.MockResourceManager.Resources[key]
	return exists
}

// AssertResourceCount checks if the resource count matches expected
func (th *TestHelper) AssertResourceCount(expected int) bool {
	return len(th.MockResourceManager.Resources) == expected
}
