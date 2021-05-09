package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateClientID(domain string) string {
	c := 15
	b := make([]byte, c)
	rand.Read(b)
	return fmt.Sprintf("%s.%s", hex.EncodeToString(b), domain)
}

func GenerateClientSecret() string {
	c := 64
	b := make([]byte, c)
	rand.Read(b)
	return hex.EncodeToString(b)
}
