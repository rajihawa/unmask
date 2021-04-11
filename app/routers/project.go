package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
	"github.com/rajihawa/unmask/app/middlewares"
)

func ProjectRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/projects").Subrouter()
	authSubRouter := subRouter.NewRoute().Subrouter()
	authSubRouter.Use(middlewares.AdminMiddleware)

	authSubRouter.HandleFunc("", handlers.GetAllProjects).Methods("GET")
	authSubRouter.HandleFunc("/{project}", handlers.GetProject).Methods("GET")

	authSubRouter.HandleFunc("", handlers.CreateProject).Methods("POST")
}
