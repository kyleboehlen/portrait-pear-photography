package util

import (
	"encoding/json"
	"friday/api/response"
	"io"
	"net/http"
)

func ReadRequestBody[R any](w http.ResponseWriter, r *http.Request, req *R) (bool, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WriteJSONErrorResponse(w, "Failed to read request body", response.ErrorCodeRequestBodyUnreadable)
		return false, err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		response.WriteJSONErrorResponse(w, "Failed to unmarshal request body", response.ErrorCodeRequestJSONUnmarshalFailed)
		return false, err
	}

	return true, nil
}
