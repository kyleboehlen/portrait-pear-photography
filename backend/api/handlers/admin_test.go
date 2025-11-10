package handlers

import (
	"bytes"
	"embed"
	"encoding/json"
	"friday/api/response"
	"friday/database/models"
	"friday/database/repository"
	"friday/services/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")

	// Initialize the repository before running tests
	if err := repository.Setup(embed.FS{}); err != nil {
		panic("Failed to setup test database: " + err.Error())
	}

	// Run tests
	code := m.Run()

	// Cleanup
	os.Exit(code)
}

func TestAuthenticateRoute_Success(t *testing.T) {
	// Set TEST environment variable for test mode
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Reset auth secrets and once
	auth.ResetSecrets()
	auth.SetSecretKey()

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
func TestUpsertShootRoute_CreateSuccess(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request body with new shoot
	reqBody := map[string]interface{}{
		"name": "Test Shoot",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/admin/shoot", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpsertShootRoute.Handle(w, req)

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
}

func TestUpsertShootRoute_MissingRequiredFields(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request without name or ID
	reqBody := map[string]interface{}{
		"slug": "asdfasdfasdf",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/admin/shoot", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpsertShootRoute.Handle(w, req)

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
	if apiResp.ErrorCode != response.ErrorCodeShootsMissingRequiredFields {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeShootsMissingRequiredFields, apiResp.ErrorCode)
	}
}

func TestGetShootsRoute_Success(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	req := httptest.NewRequest("GET", "/admin/shoots", nil)
	w := httptest.NewRecorder()

	ListShootsRoute.Handle(w, req)

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
}

func TestDeleteShootRoute_Success(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request with shoot ID to delete
	reqBody := DeleteShootRequest{ID: 1}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("DELETE", "/admin/unshoot", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	DeleteShootRoute.Handle(w, req)

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
}

func TestDeleteShootRoute_MissingID(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request without ID
	reqBody := DeleteShootRequest{ID: 0}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("DELETE", "/admin/unshoot", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	DeleteShootRoute.Handle(w, req)

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
	if apiResp.ErrorCode != response.ErrorCodeShootsMissingRequiredFields {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodeShootsMissingRequiredFields, apiResp.ErrorCode)
	}
}

func TestDeletePhotoRoute_Success(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	db := repository.Get()

	_ = db.CreatePhoto(&models.Photo{
		ID:         1,
		Guid:       "test-guid",
		Favorite:   false,
		ShootID:    1,
		Categories: []int{},
	})

	// Create request with shoot ID to delete
	reqBody := DeletePhotoRequest{ID: 1}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("DELETE", "/admin/delete-photo", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	DeletePhotoRoute.Handle(w, req)

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
}

func TestDeletePhotoRoute_MissingID(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	// Create request without ID
	reqBody := DeletePhotoRequest{ID: 0}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("DELETE", "/admin/delete-photo", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	DeletePhotoRoute.Handle(w, req)

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
	if apiResp.ErrorCode != response.ErrorCodePhotosMissingRequiredFields {
		t.Errorf("Expected error code %v, got %v", response.ErrorCodePhotosMissingRequiredFields, apiResp.ErrorCode)
	}
}

func TestUpsertShootRoute_UpdateSuccess(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	db := repository.Get()

	// Create a shoot first
	shoot := &models.Shoot{
		Name: "Original Shoot",
		Slug: "original-slug",
	}
	_ = db.CreateShoot(shoot)

	// Update the shoot
	reqBody := map[string]interface{}{
		"id":   shoot.ID,
		"name": "Updated Shoot",
		"slug": "updated-slug",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/admin/shoot", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpsertShootRoute.Handle(w, req)

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

	// Verify the shoot was updated
	var updatedShoot models.Shoot
	dataBytes, _ := json.Marshal(apiResp.Content)
	if err := json.Unmarshal(dataBytes, &updatedShoot); err != nil {
		t.Fatalf("Failed to unmarshal shoot response: %v", err)
	}

	if updatedShoot.Name != "Updated Shoot" {
		t.Errorf("Expected name to be 'Updated Shoot', got '%s'", updatedShoot.Name)
	}
	if updatedShoot.Slug != "updated-slug" {
		t.Errorf("Expected slug to be 'updated-slug', got '%s'", updatedShoot.Slug)
	}
}

func TestUpdatePhotoRoute_Success(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	db := repository.Get()

	// Create a photo first
	photo := &models.Photo{
		Guid:       "test-guid-update",
		Favorite:   false,
		ShootID:    1,
		Categories: []int{1, 2},
	}
	_ = db.CreatePhoto(photo)

	// Update the photo
	reqBody := UpdatePhotoRequest{
		Photo: models.Photo{
			ID:         photo.ID,
			Guid:       photo.Guid,
			Favorite:   true,
			ShootID:    2,
			Categories: []int{3, 4, 5},
		},
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/admin/update-photo", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdatePhotoRoute.Handle(w, req)

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

	// Verify the photo was updated
	var updatedPhoto models.Photo
	dataBytes, _ := json.Marshal(apiResp.Content)
	if err := json.Unmarshal(dataBytes, &updatedPhoto); err != nil {
		t.Fatalf("Failed to unmarshal photo response: %v", err)
	}

	if !updatedPhoto.Favorite {
		t.Error("Expected favorite to be true")
	}
	if updatedPhoto.ShootID != 2 {
		t.Errorf("Expected shoot_id to be 2, got %d", updatedPhoto.ShootID)
	}
	if len(updatedPhoto.Categories) != 3 {
		t.Errorf("Expected 3 categories, got %d", len(updatedPhoto.Categories))
	}
}

func TestListShootsRoute_Success(t *testing.T) {
	// Set TEST environment variable
	_ = os.Setenv("TEST", "true")
	defer func() {
		_ = os.Unsetenv("TEST")
	}()

	db := repository.Get()

	// Create some test shoots
	_ = db.CreateShoot(&models.Shoot{Name: "Shoot 1", Slug: "shoot-1"})
	_ = db.CreateShoot(&models.Shoot{Name: "Shoot 2", Slug: "shoot-2"})

	req := httptest.NewRequest("GET", "/admin/shoots", nil)
	w := httptest.NewRecorder()

	ListShootsRoute.Handle(w, req)

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

	// Verify shoots are returned
	var shoots []models.Shoot
	dataBytes, _ := json.Marshal(apiResp.Content)
	if err := json.Unmarshal(dataBytes, &shoots); err != nil {
		t.Fatalf("Failed to unmarshal shoots response: %v", err)
	}

	if len(shoots) < 2 {
		t.Errorf("Expected at least 2 shoots, got %d", len(shoots))
	}
}
