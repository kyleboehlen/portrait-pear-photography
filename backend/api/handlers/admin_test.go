package handlers

import (
	"bytes"
	"encoding/json"
	"friday/api/response"
	"friday/services/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAuthenticateRoute_Success(t *testing.T) {
	// Set TEST environment variable for test mode
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Reset auth secrets and once
	auth.ResetSecrets()

	// Create request body with test password
	reqBody := AdminAuthRequest{Password: "password"}
	jsonBody, _ := json.Marshal(reqBody)

	// Create request
	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Execute handler
	AuthenticateRoute.Handle(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse response
	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !apiResp.Success {
		t.Errorf("Expected success to be true, got false")
	}

	// Check that token is returned
	var authResp AdminAuthResponse
	dataBytes, _ := json.Marshal(apiResp.Content)
	if err := json.Unmarshal(dataBytes, &authResp); err != nil {
		t.Fatalf("Failed to unmarshal auth response: %v", err)
	}

	if authResp.Token == "" {
		t.Error("Expected non-empty token")
	}

	// Verify token is valid
	valid, err := auth.IsJWTValid(authResp.Token)
	if err != nil {
		t.Errorf("Token validation failed: %v", err)
	}
	if !valid {
		t.Error("Expected token to be valid")
	}
}

func TestAuthenticateRoute_MissingPassword(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request with empty password
	reqBody := AdminAuthRequest{Password: ""}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	AuthenticateRoute.Handle(w, req)

	// Check error response
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
	if apiResp.ErrorCode != response.ErrorCodeMissingAdminPassword {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeMissingAdminPassword, apiResp.ErrorCode)
	}
}

func TestAuthenticateRoute_InvalidPassword(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	auth.ResetSecrets()

	// Create request with wrong password
	reqBody := AdminAuthRequest{Password: "wrongpassword"}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/authenticate", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	AuthenticateRoute.Handle(w, req)

	// Check error response
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
	if apiResp.ErrorCode != response.ErrorCodeInvalidPassword {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeInvalidPassword, apiResp.ErrorCode)
	}
}

func TestAuthenticateRoute_InvalidJSON(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create malformed JSON
	invalidJSON := `{"password":"test"`

	req := httptest.NewRequest("POST", "/authenticate", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	AuthenticateRoute.Handle(w, req)

	// Should return JSON unmarshal error
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
	if apiResp.ErrorCode != response.ErrorCodeRequestJSONUnmarshalFailed {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeRequestJSONUnmarshalFailed, apiResp.ErrorCode)
	}
}

func TestTestRoute_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/admin/test", nil)
	w := httptest.NewRecorder()

	TestRoute.Handle(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !apiResp.Success {
		t.Error("Expected success to be true")
	}

	// Check response data
	var testResp AdminTestResponse
	dataBytes, _ := json.Marshal(apiResp.Content)
	if err := json.Unmarshal(dataBytes, &testResp); err != nil {
		t.Fatalf("Failed to unmarshal test response: %v", err)
	}

	if testResp.Test != "asdf" {
		t.Errorf("Expected test field to be 'asdf', got '%s'", testResp.Test)
	}
}
