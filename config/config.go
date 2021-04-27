package config

import (
	"os"
	"time"
)

// global env vars
var (
	IsProd = getEnvBool("PRODUCTION")
)

func getEnvBool(env string) bool {
	return os.Getenv(env) == "true"
}

func AdminTokenExpire() time.Time {
	return time.Now().Add(time.Minute * 60)
}

func AccessTokenExpire() time.Time {
	return time.Now().Add(time.Hour * 24 * 60)
}

func RefreshTokenExpire() time.Time {
	return time.Now().Add(time.Hour * 24 * 365)
}

func AuthTokenExpire() time.Time {
	return time.Now().Add(time.Minute * 5)
}
