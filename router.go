package main

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/routers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	routers.AdminRouter(router)
	routers.ProjectRouter(router)
	routers.ClientRouter(router)
	routers.HealthRouter(router)
	return router
}
