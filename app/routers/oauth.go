package routers

import (
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/handlers"
)

func OauthRouter(router *mux.Router) {
	subRouter := router.PathPrefix("/oauth").Subrouter()

	subRouter.HandleFunc("/connect", handlers.Oauthorize)
	subRouter.HandleFunc("/verify", handlers.Overify)
	subRouter.HandleFunc("/current", handlers.CurrentSession)
	subRouter.HandleFunc("/{client}/client", handlers.Oclient)

}
