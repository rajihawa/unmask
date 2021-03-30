package middlewares

import (
	"net/http"

	"github.com/rajihawa/unmask/utils"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieMgr, err := utils.GetCookie(r)
		if err != nil {
			utils.HttpError(w, err, http.StatusUnauthorized, "Unauthorized.")
			return
		}

		errMsg, errCode, err := utils.ValidateToken(cookieMgr.Cookie.Value)
		if err != nil {
			utils.HttpError(w, err, errCode, errMsg)
			return
		}
		next.ServeHTTP(w, r)
	})
}
