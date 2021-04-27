package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	OauthRouter(apiRouter)
	AdminRouter(apiRouter)
	ProjectRouter(apiRouter)
	ClientRouter(apiRouter)
	HealthRouter(apiRouter)
	UserRouter(apiRouter)
	router.PathPrefix("/").Handler(handlers.Spa)

	return router
}
