package main

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/routers"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	routers.AdminRouter(router)
	routers.AdminRouter(router)
	routers.HealthRouter(router)
	return router
}
