package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/handlers"
)

func mainRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/").Subrouter()
	subRouter.HandleFunc("/health", handlers.HealthHandler)
}
