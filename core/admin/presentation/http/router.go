package http

import "github.com/gorilla/mux"

func Router(router *mux.Router) {
	subRouter := router.PathPrefix("/admin").Subrouter()
	subRouter.HandleFunc("/login", AdminLoginHandler).Methods("POST")

	authSubRouter := subRouter.NewRoute().Subrouter()
	authSubRouter.Use(AdminMiddleware)
	authSubRouter.HandleFunc("/me", AdminMeHandler).Methods("GET", "POST")
}
