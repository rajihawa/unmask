package config

import (
	"os"
	"time"
)

var (
	AdminUsername = os.Getenv("ADMIN_USERNAME")
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	JwtCookieName = "qid"
	JwtSigningKey = []byte(os.Getenv("JWT_SIGN_KEY"))
	IsProd        = getEnvBool("PRODUCTION")
)

func getEnvBool(env string) bool {
	return os.Getenv(env) == "true"
}

func JwtAdminExpire() time.Time {
	return time.Now().Add(time.Minute * 60)
}
