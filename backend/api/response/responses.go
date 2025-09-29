package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"error_message"`
	ErrorCode    ErrorCode   `json:"error_code"`
	Content      interface{} `json:"content"`
}

func WriteJSONSuccessResponse(w http.ResponseWriter, content interface{}) {
	writeJSONResponse(w, APIResponse{
		Success:      true,
		ErrorMessage: "",
		ErrorCode:    ErrorCodeNone,
		Content:      content,
	})
}

func WriteJSONErrorResponse(w http.ResponseWriter, errorMessage string, errorCode ErrorCode) {
	writeJSONResponse(w, APIResponse{
		Success:      false,
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
		Content:      nil,
	})
}

func writeJSONResponse(w http.ResponseWriter, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
