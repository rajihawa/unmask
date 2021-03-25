package handlers

import (
	"io"
	"net/http"
)

func AdminMeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"authorized": true}`)
}
