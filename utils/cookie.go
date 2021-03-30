package utils

import (
	"net/http"
	"time"

	"github.com/rajihawa/unmask/config"
)

var (
	jwtCookieName = "qid"
)

type CookieManager struct {
	Cookie *http.Cookie
}

func (c *CookieManager) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, c.Cookie)
}

func GetCookie(r *http.Request) (*CookieManager, error) {
	cookie, err := r.Cookie(jwtCookieName)
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

func CreateJwtCookie(value string, exp time.Time) *CookieManager {
	return CreateCookie(jwtCookieName, value, exp)
}
