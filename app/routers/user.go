package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
	"github.com/rajihawa/unmask/app/middlewares"
)

func UserRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/{client}/users").Subrouter()
	subRouter.Use(middlewares.ClientMiddleware)

	subRouter.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	subRouter.HandleFunc("/signup", handlers.SignupUser).Methods("POST")

	userAuthSubRouter := subRouter.NewRoute().Subrouter()
	userAuthSubRouter.Use(middlewares.UserMiddleware)

	userAuthSubRouter.HandleFunc("/me", handlers.CurrentUser).Methods("GET")

	// adminAuthSubRouter := subRouter.NewRoute().Subrouter()
	// adminAuthSubRouter.Use(middlewares.AdminMiddleware)

	// authSubRouter.HandleFunc("", handlers.GetAllClients).Methods("GET").Queries("show_projects", "{show_projects}")

	// adminAuthSubRouter.HandleFunc("/create", handlers.AdminCreateUser).Methods("POST")

	// authSubRouter.HandleFunc("/{client}/secret", handlers.GetClientSecret).Methods("GET")
}
