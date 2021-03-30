package utils

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func HttpError(w http.ResponseWriter, err error, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	httpErr := httpError{
		Message: message,
		Error:   err.Error(),
		Code:    code,
	}
	w.WriteHeader(httpErr.Code)
	json.NewEncoder(w).Encode(httpErr)
}
