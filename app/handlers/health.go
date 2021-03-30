package handlers

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Healthy bool `json:"healthy"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health{Healthy: true})
}
