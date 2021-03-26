package http

import (
	"github.com/gorilla/mux"
	admin "github.com/rajihawa/unmask/core/admin/presentation/http"
)

func Router(router *mux.Router) {
	subRouter := router.PathPrefix("/projects").Subrouter()
	authSubRouter := subRouter.NewRoute().Subrouter()
	authSubRouter.Use(admin.AdminMiddleware)

	authSubRouter.HandleFunc("", GetAllProjects).Methods("GET")

	authSubRouter.HandleFunc("", CreateProject).Methods("POST")
}
