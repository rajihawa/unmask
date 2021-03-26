package main

import (
	"github.com/gorilla/mux"
	admin "github.com/rajihawa/unmask/core/admin/presentation/http"
	health "github.com/rajihawa/unmask/core/health"
	projects "github.com/rajihawa/unmask/features/projects/presentation/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	admin.Router(router)
	projects.Router(router)
	health.Router(router)
	return router
}
