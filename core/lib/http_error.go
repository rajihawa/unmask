package lib

import (
	"encoding/json"
	"net/http"

	domain "github.com/rajihawa/unmask/core/domain"
)

func HttpError(w http.ResponseWriter, err error, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	httpErr := domain.HttpError{
		Message: message,
		Error:   err.Error(),
		Code:    code,
	}
	w.WriteHeader(httpErr.Code)
	json.NewEncoder(w).Encode(httpErr)
}
