package utils

import (
	"net/http"
	"time"

	"github.com/rajihawa/unmask/config"
)

var (
	AdminTokenCookieName  = "qid"
	AccessTokenCookieName = "access_token"
	IdTokenCookieName     = "id_token"
)

type CookieManager struct {
	Cookie *http.Cookie
}

func (c *CookieManager) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, c.Cookie)
}

func GetCookie(name string, r *http.Request) (*CookieManager, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return &CookieManager{
		Cookie: cookie,
	}, nil
}

func CreateCookie(name string, value string, exp time.Time) *CookieManager {
	cookie := &http.Cookie{
		Name:     name,
		Expires:  exp,
		Value:    value,
		HttpOnly: true,
		Secure:   config.IsProd,
		Path:     "/",
	}

	return &CookieManager{
		Cookie: cookie,
	}
}

func CreateJwtCookie(name string, value string, exp time.Time) *CookieManager {
	return CreateCookie(name, value, exp)
}
