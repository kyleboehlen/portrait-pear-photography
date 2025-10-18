package handlers_test

import (
	"encoding/json"
	_ "friday/api/handlers"
	"friday/api/response"
	"friday/api/routing"
	"net/http/httptest"
	"testing"
)

// TestHeartbeat all we're testing is that we get the heartbeat back, just an acceptance test
func TestHealthEndpoint(t *testing.T) {
	mux := routing.GetNewRouter()
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	var resp response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Content != "Heartbeat" {
		t.Errorf("Expected content 'Heartbeat', got %v", resp.Content)
	}
}
