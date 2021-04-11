package middlewares

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

func ClientMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := mux.Vars(r)["client"]
		client, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetClient(clientID, domain.GetClientOpts{GetProjects: true})
		if err != nil {
			utils.HttpError(w, err, http.StatusBadRequest, "Can't get client.")
			return
		}
		if client.Disabled {
			utils.HttpError(w, err, http.StatusBadRequest, "Client is disabled.")
			return
		}
		ctx := context.WithValue(r.Context(), utils.ContextClientKey, *client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
