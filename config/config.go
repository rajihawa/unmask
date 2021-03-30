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
