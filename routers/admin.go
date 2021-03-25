package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/handlers"
	"github.com/rajihawa/unmask/middlewares"
)

func adminRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/admin").Subrouter()
	subRouter.HandleFunc("/login", handlers.AdminLoginHandler).Methods("POST")

	authSubRouter := subRouter.NewRoute().Subrouter()
	authSubRouter.Use(middlewares.AdminMiddleware)
	authSubRouter.HandleFunc("/me", handlers.AdminMeHandler).Methods("GET", "POST")
}
