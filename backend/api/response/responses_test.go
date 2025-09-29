package response_test

import (
	"encoding/json"
	"friday/api/response"
	"net/http/httptest"
	"testing"
)

type TestContent struct {
	DoesIt bool   `json:"does_it"`
	Word   string `json:"word"`
	Number int    `json:"number"`
}

func TestWriteJSONSuccessResponse(t *testing.T) {
	w := httptest.NewRecorder()
	content := &TestContent{
		DoesIt: true,
		Word:   "test",
		Number: 42,
	}

	response.WriteJSONSuccessResponse(w, content)

	// Assert status code is OK
	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Assert content type is application/json
	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected content type %s, got %s", expectedContentType, contentType)
	}

	// Attempt to parse response, if it is not valid json it will fail
	var aR response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &aR); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// The success field should always be true from the success response function
	if !aR.Success {
		t.Error("Expected success to be true")
	}

	// The error message should always be empty from the success response function
	if aR.ErrorMessage != "" {
		t.Errorf("Expected empty error message, got %s", aR.ErrorMessage)
	}

	// An error code of 0 should always be returned from the success response function
	if aR.ErrorCode != 0 {
		t.Errorf("Expected error code 0, got %d", aR.ErrorCode)
	}

	// The content of the response should deserialize correctly back to the original content

	// Marshal the content back to JSON
	contentJSON, err := json.Marshal(aR.Content)
	if err != nil {
		t.Fatalf("Failed to marshal content: %v", err)
	}

	// Unmarshal into TestContent struct
	var deserializedContent TestContent
	if err := json.Unmarshal(contentJSON, &deserializedContent); err != nil {
		t.Fatalf("Failed to unmarshal content to TestContent: %v", err)
	}

	// Now compare the fields - boolean
	if !deserializedContent.DoesIt {
		t.Errorf("Expected DoesIt %v, got %v", content.DoesIt, deserializedContent.DoesIt)
	} // string
	if deserializedContent.Word != content.Word {
		t.Errorf("Expected Word %s, got %s", content.Word, deserializedContent.Word)
	} // integer
	if deserializedContent.Number != content.Number {
		t.Errorf("Expected Number %d, got %d", content.Number, deserializedContent.Number)
	}
}

func TestWriteJSONErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	expectedErrorMessage := "Something went wrong"
	expectedErrorCode := response.ErrorCodeTest

	response.WriteJSONErrorResponse(w, expectedErrorMessage, expectedErrorCode)

	// Assert status code is OK
	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Assert content type is application/json
	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected content type %s, got %s", expectedContentType, contentType)
	}

	// Attempt to parse response, if it is not valid json it will fail
	var aR response.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &aR); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// The success field should always be false from the error response function
	if aR.Success {
		t.Error("Expected success to be false")
	}

	// The error message should match what was passed in
	if aR.ErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message %s, got %s", expectedErrorMessage, aR.ErrorMessage)
	}

	// The error code should match what was passed in
	if aR.ErrorCode != expectedErrorCode {
		t.Errorf("Expected error code %d, got %d", expectedErrorCode, aR.ErrorCode)
	}

	// The content should be nil from the error response function
	if aR.Content != nil {
		t.Errorf("Expected content to be nil, got %v", aR.Content)
	}
}
