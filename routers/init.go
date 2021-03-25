package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	mainRouter(router)
	adminRouter(router)
	return router
}
