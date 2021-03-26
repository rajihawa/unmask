package config

import (
	"os"
	"time"
)

var (
	AdminUsername    = os.Getenv("ADMIN_USERNAME")
	AdminPassword    = os.Getenv("ADMIN_PASSWORD")
	JwtCookieName    = "qid"
	JwtSigningKey    = []byte(os.Getenv("JWT_SIGN_KEY"))
	IsProd           = getEnvBool("PRODUCTION")
	DatabaseUrl      = os.Getenv("DATABASE_URL")
	DatabaseName     = "unmask"
	DatabaseUsername = os.Getenv("DATABASE_USERNAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
)

func getEnvBool(env string) bool {
	return os.Getenv(env) == "true"
}

func JwtAdminExpire() time.Time {
	return time.Now().Add(time.Minute * 60)
}
