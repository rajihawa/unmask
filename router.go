package main

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/routers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	routers.AdminRouter(router)
	routers.ProjectRouter(router)
	routers.ClientRouter(router)
	routers.HealthRouter(router)
	routers.UserRouter(router)
	return router
}
