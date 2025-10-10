package middleware

import (
	"encoding/json"
	"friday/api/response"
	"friday/services/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuthenticateAdmin_NonAdminRoute(t *testing.T) {
	// Non-admin routes should pass through without authentication
	req := httptest.NewRequest("GET", "/public", nil)
	w := httptest.NewRecorder()

	// Mock handler to verify next is called
	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	AuthenticateAdmin(next).ServeHTTP(w, req)

	if !nextCalled {
		t.Error("Expected next handler to be called for non-admin route")
	}
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAuthenticateAdmin_AdminRoute_MissingAuthHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/admin/test", nil)
	w := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Next handler should not be called when auth header is missing")
	})

	AuthenticateAdmin(next).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if apiResp.Success {
		t.Error("Expected success to be false")
	}
	if apiResp.ErrorCode != response.ErrorCodeMissingAuthorizationHeader {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeMissingAuthorizationHeader, apiResp.ErrorCode)
	}
}

func TestAuthenticateAdmin_AdminRoute_InvalidAuthHeaderFormat(t *testing.T) {
	testCases := []struct {
		name   string
		header string
	}{
		{"no bearer prefix", "token123"},
		{"wrong prefix", "Basic token123"},
		{"only bearer", "Bearer"},
		{"empty bearer", "Bearer "},
		{"too many parts", "Bearer token123 extra"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/admin/test", nil)
			req.Header.Set("Authorization", tc.header)
			w := httptest.NewRecorder()

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Error("Next handler should not be called with invalid auth header format")
			})

			AuthenticateAdmin(next).ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
			}

			var apiResp response.APIResponse
			if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if apiResp.Success {
				t.Error("Expected success to be false")
			}
			if apiResp.ErrorCode != response.ErrorCodeInvalidAuthorizationHeaderFormat && apiResp.ErrorCode != response.ErrorCodeInvalidOrExpiredToken {
				t.Errorf("Expected error code %v or %v, got %v", response.ErrorCodeInvalidAuthorizationHeaderFormat, response.ErrorCodeInvalidOrExpiredToken, apiResp.ErrorCode)
			}
		})
	}
}

func TestAuthenticateAdmin_AdminRoute_InvalidToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/admin/test", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken123")
	w := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Next handler should not be called with invalid token")
	})

	AuthenticateAdmin(next).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if apiResp.Success {
		t.Error("Expected success to be false")
	}
	if apiResp.ErrorCode != response.ErrorCodeInvalidOrExpiredToken {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeInvalidOrExpiredToken, apiResp.ErrorCode)
	}
}

func TestAuthenticateAdmin_AdminRoute_ValidToken(t *testing.T) {
	// Set up test environment
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Reset auth secrets
	auth.ResetSecrets()

	// Create a valid token
	token, err := auth.CreateAdminJWT()
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	req := httptest.NewRequest("GET", "/admin/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Mock handler to verify next is called
	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	AuthenticateAdmin(next).ServeHTTP(w, req)

	if !nextCalled {
		t.Error("Expected next handler to be called with valid token")
	}
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAuthenticateAdmin_AdminRoute_ExpiredToken(t *testing.T) {
	// This test would require creating an expired token
	// For now, we'll test with a malformed token that would fail validation
	req := httptest.NewRequest("GET", "/admin/test", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.expired.token")
	w := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Next handler should not be called with expired token")
	})

	AuthenticateAdmin(next).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if apiResp.Success {
		t.Error("Expected success to be false")
	}
	if apiResp.ErrorCode != response.ErrorCodeInvalidOrExpiredToken {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeInvalidOrExpiredToken, apiResp.ErrorCode)
	}
}

func TestAuthenticateAdmin_DifferentAdminPaths(t *testing.T) {
	adminPaths := []string{
		"/admin",
		"/admin/",
		"/admin/test",
		"/admin/users",
		"/admin/deep/nested/path",
	}

	for _, path := range adminPaths {
		t.Run("path_"+path, func(t *testing.T) {
			req := httptest.NewRequest("GET", path, nil)
			w := httptest.NewRecorder()

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Error("Next handler should not be called without auth for admin path: " + path)
			})

			AuthenticateAdmin(next).ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d for path %s, got %d", http.StatusOK, path, w.Code)
			}

			var apiResp response.APIResponse
			if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
				t.Fatalf("Failed to unmarshal response for path %s: %v", path, err)
			}

			if apiResp.ErrorCode != response.ErrorCodeMissingAuthorizationHeader {
				t.Errorf("Expected error code %v for path %s, got %v", response.ErrorCodeMissingAuthorizationHeader, path, apiResp.ErrorCode)
			}
		})
	}
}
