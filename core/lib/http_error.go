package lib

import (
	"encoding/json"
	"net/http"

	"github.com/rajihawa/unmask/core/models"
)

func HttpError(w http.ResponseWriter, err error, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	httpErr := models.HttpError{
		Message: message,
		Error:   err.Error(),
		Code:    code,
	}
	w.WriteHeader(httpErr.Code)
	json.NewEncoder(w).Encode(httpErr)
}
