package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieMgr, err := utils.GetCookie(utils.AccessTokenCookieName, r)
		if err != nil {
			utils.HttpError(w, err, http.StatusUnauthorized, "Unauthorized.")
			return
		}

		claims, errMsg, errCode, err := utils.ValidateToken(cookieMgr.Cookie.Value)
		if err != nil {
			utils.HttpError(w, err, errCode, errMsg)
			return
		}
		fmt.Println(claims)

		client := r.Context().Value(utils.ContextClientKey).(domain.Client)
		userID := claims["user_id"].(string)

		user, err := usecases.NewUsersUsecase(repository.NewUsersRepository()).GetUser(userID, client, domain.GetUsersOpts{})
		if err != nil {
			utils.HttpError(w, err, http.StatusBadRequest, "Can't get user.")
			return
		}

		ctx := context.WithValue(r.Context(), utils.ContextUserKey, *user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
