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

func IdTokenExpire() time.Time {
	return time.Now().Add(time.Hour * 24 * 365)
}
