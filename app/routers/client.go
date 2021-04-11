package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
	"github.com/rajihawa/unmask/app/middlewares"
)

func ClientRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/{project}/clients").Subrouter()
	subRouter.Use(middlewares.ProjectMiddleware)
	authSubRouter := subRouter.NewRoute().Subrouter()
	authSubRouter.Use(middlewares.AdminMiddleware)

	authSubRouter.HandleFunc("", handlers.GetAllClients).Methods("GET").Queries("show_projects", "{show_projects}")

	authSubRouter.HandleFunc("", handlers.CreateNewClients).Methods("POST")

	authSubRouter.HandleFunc("/{client}/secret", handlers.GetClientSecret).Methods("GET")
}
