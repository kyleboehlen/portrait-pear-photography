package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"friday/api/response"
	"io"
	"net/http"
	"strings"
)

func ReadRequestBody[R any](w http.ResponseWriter, r *http.Request, req *R) (bool, error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // json.Unmarshal is lenient by default, this makes it strict
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	if err := decoder.Decode(&req); err != nil {
		var unmarshalErr *json.UnmarshalTypeError
		var syntaxErr *json.SyntaxError

		switch {
		case errors.As(err, &unmarshalErr):
			response.WriteJSONErrorResponse(w,
				fmt.Sprintf("Invalid type for field '%s'", unmarshalErr.Field),
				response.ErrorCodeRequestJSONUnmarshalFailed)
		case errors.As(err, &syntaxErr):
			response.WriteJSONErrorResponse(w,
				"Malformed JSON",
				response.ErrorCodeRequestJSONUnmarshalFailed)
		case strings.Contains(err.Error(), "unknown field"):
			response.WriteJSONErrorResponse(w,
				"Request contains unknown fields",
				response.ErrorCodeRequestJSONUnmarshalFailed)
		default:
			response.WriteJSONErrorResponse(w,
				"Failed to unmarshal request body",
				response.ErrorCodeRequestJSONUnmarshalFailed)
		}
		return false, err
	}

	return true, nil
}
