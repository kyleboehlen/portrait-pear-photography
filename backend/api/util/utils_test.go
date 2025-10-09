package util

import (
	"bytes"
	"encoding/json"
	"friday/api/response"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReadRequestBody_Success(t *testing.T) {
	// Create a valid JSON request body
	jsonBody := `{"name":"John","age":30}`
	req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()

	var testReq TestRequest
	success, err := ReadRequestBody(w, req, &testReq)

	if !success {
		t.Errorf("Expected success to be true, got false")
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if testReq.Name != "John" {
		t.Errorf("Expected name to be 'John', got '%s'", testReq.Name)
	}
	if testReq.Age != 30 {
		t.Errorf("Expected age to be 30, got %d", testReq.Age)
	}
}

func TestReadRequestBody_UnknownFields(t *testing.T) {
	// Create JSON with unknown fields (DisallowUnknownFields will catch this)
	jsonBody := `{"name":"John","age":30,"unknown":"field"}`
	req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()

	var testReq TestRequest
	success, err := ReadRequestBody(w, req, &testReq)

	if success {
		t.Errorf("Expected success to be false, got true")
	}
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Parse the response body to check the error code
	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.ErrorCode != response.ErrorCodeRequestJSONUnmarshalFailed {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeRequestJSONUnmarshalFailed, apiResp.ErrorCode)
	}
}

func TestReadRequestBody_MalformedJSON(t *testing.T) {
	// Create malformed JSON
	jsonBody := `{"name":"John","age":}`
	req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()

	var testReq TestRequest
	success, err := ReadRequestBody(w, req, &testReq)

	if success {
		t.Errorf("Expected success to be false, got true")
	}
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Parse the response body to check the error code
	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.ErrorCode != response.ErrorCodeRequestJSONUnmarshalFailed {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeRequestJSONUnmarshalFailed, apiResp.ErrorCode)
	}
}

func TestReadRequestBody_WrongType(t *testing.T) {
	// Create JSON with wrong type for age field
	jsonBody := `{"name":"John","age":"thirty"}`
	req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
	w := httptest.NewRecorder()

	var testReq TestRequest
	success, err := ReadRequestBody(w, req, &testReq)

	if success {
		t.Errorf("Expected success to be false, got true")
	}
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Parse the response body to check the error code
	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.ErrorCode != response.ErrorCodeRequestJSONUnmarshalFailed {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeRequestJSONUnmarshalFailed, apiResp.ErrorCode)
	}
}

func TestReadRequestBody_EmptyBody(t *testing.T) {
	// Create request with empty body
	req := httptest.NewRequest("POST", "/test", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	var testReq TestRequest
	success, err := ReadRequestBody(w, req, &testReq)

	if success {
		t.Errorf("Expected success to be false, got true")
	}
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Parse the response body to check the error code
	var apiResp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.ErrorCode != response.ErrorCodeRequestJSONUnmarshalFailed {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeRequestJSONUnmarshalFailed, apiResp.ErrorCode)
	}
}
