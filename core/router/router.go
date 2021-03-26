package router

import (
	"github.com/gorilla/mux"
	admin "github.com/rajihawa/unmask/core/admin/presentation/http"
	health "github.com/rajihawa/unmask/core/health"
	projects "github.com/rajihawa/unmask/features/projects/presentation/http"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	health.Router(router)
	admin.Router(router)
	projects.Router(router)
	return router
}
