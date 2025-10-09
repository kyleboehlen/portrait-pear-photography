package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCORS_ProductionMode(t *testing.T) {
	// Ensure DEBUG_CORS is not set
	_ = os.Unsetenv("DEBUG_CORS")

	// Create a test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	})

	// Wrap with CORS middleware
	handler := CORS(testHandler)

	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Execute request
	handler.ServeHTTP(w, req)

	// Check headers
	if origin := w.Header().Get("Access-Control-Allow-Origin"); origin != "https://portraitpear.photography" {
		t.Errorf("Expected origin to be 'https://portraitpear.photography', got '%s'", origin)
	}
	if vary := w.Header().Get("Vary"); vary != "Origin" {
		t.Errorf("Expected Vary header to be 'Origin', got '%s'", vary)
	}
	if creds := w.Header().Get("Access-Control-Allow-Credentials"); creds != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials to be 'true', got '%s'", creds)
	}
	if methods := w.Header().Get("Access-Control-Allow-Methods"); methods != "GET,POST,DELETE" {
		t.Errorf("Expected Access-Control-Allow-Methods to be 'GET,POST,DELETE', got '%s'", methods)
	}
	if headers := w.Header().Get("Access-Control-Allow-Headers"); headers != "Content-Type, Authorization" {
		t.Errorf("Expected Access-Control-Allow-Headers to be 'Content-Type, Authorization', got '%s'", headers)
	}

	// Check that the handler was called
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	if body := w.Body.String(); body != "test response" {
		t.Errorf("Expected body to be 'test response', got '%s'", body)
	}
}

func TestCORS_DebugMode(t *testing.T) {
	// Set DEBUG_CORS environment variable
	_ = os.Setenv("DEBUG_CORS", "true")
	defer func() {
		_ = os.Unsetenv("DEBUG_CORS")
	}()

	// Create a test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	})

	// Wrap with CORS middleware
	handler := CORS(testHandler)

	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Execute request
	handler.ServeHTTP(w, req)

	// Check that origin is set to localhost in debug mode
	if origin := w.Header().Get("Access-Control-Allow-Origin"); origin != "http://localhost:5173" {
		t.Errorf("Expected origin to be 'http://localhost:5173' in debug mode, got '%s'", origin)
	}

	// Check that the handler was called
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
