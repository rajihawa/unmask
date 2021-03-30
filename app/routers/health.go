package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
)

func HealthRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/").Subrouter()
	subRouter.HandleFunc("/health", handlers.HealthHandler)
}
