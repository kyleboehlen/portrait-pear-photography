package routing_test

import (
	"friday/api/routing"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	// Check current route count
	ogRouteCount := routing.GetRegisteredRoutesCount()

	// Define two test Route
	testRoutes := []routing.Route{
		{
			Method: "GET",
			Path:   "/test1",
			Handle: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			Method: "POST",
			Path:   "/test2",
			Handle: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
			},
		},
	}

	// Register test routes
	routing.RegisterRoutes(testRoutes)

	// Uses the helper function to check the route count (registeredRoutes is private)
	newCount := routing.GetRegisteredRoutesCount()
	expectedCount := ogRouteCount + len(testRoutes)

	if newCount != expectedCount {
		t.Errorf("Expected %d routes, got %d", expectedCount, newCount)
	}
}

func TestGetNewRouterIsNotNil(t *testing.T) {
	mux := routing.GetNewRouter()

	if mux == nil {
		t.Error("Expected non-nil ServeMux")
	}
}

func TestRouterMethodNotAllowed(t *testing.T) {
	// Define a test route that returns 200 OK for GET requests only
	testRoutes := []routing.Route{
		{
			Method: "GET",
			Path:   "/test-method",
			Handle: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
	}

	// Register route and get the multiplexer
	routing.RegisterRoutes(testRoutes)
	mux := routing.GetNewRouter()

	// Test correct method (GET)
	req := httptest.NewRequest("GET", "/test-method", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d for GET, got %d", http.StatusOK, w.Code)
	}

	// Test incorrect method (POST)
	req = httptest.NewRequest("POST", "/test-method", nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for POST, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
