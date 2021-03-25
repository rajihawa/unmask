package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rajihawa/unmask/config"
	"github.com/rajihawa/unmask/lib"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(config.JwtCookieName)
		if err != nil {
			lib.HttpError(w, err, http.StatusUnauthorized, "Unauthorized.")
			return
		}
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSigningKey, nil
		})
		if err != nil {
			lib.HttpError(w, err, http.StatusInternalServerError, "Cant validate token.")
			return
		}

		exp, err := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
		if err != nil {
			lib.HttpError(w, err, http.StatusInternalServerError, "Cant parse expiration date.")
			return
		}

		currTime := time.Now().Unix()

		if currTime > exp {
			lib.HttpError(w, err, http.StatusUnauthorized, "Unauthorized.")
			return
		}
		next.ServeHTTP(w, r)
	})
}
