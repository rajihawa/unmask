package health

import "github.com/gorilla/mux"

func Router(router *mux.Router) {
	subRouter := router.PathPrefix("/").Subrouter()
	subRouter.HandleFunc("/health", HealthHandler)
}
