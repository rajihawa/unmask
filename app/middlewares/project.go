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

func ProjectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projectID := mux.Vars(r)["project"]
		project, err := usecases.NewProjectUsecase(repository.NewProjectRepository()).GetProject(projectID, domain.GetProjectOpts{})
		if err != nil {
			utils.HttpError(w, err, http.StatusUnauthorized, "Can't get project.")
			return
		}
		ctx := context.WithValue(r.Context(), utils.ContextProjectKey, *project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
