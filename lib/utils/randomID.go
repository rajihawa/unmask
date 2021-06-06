package utils

import (
	"encoding/hex"
	"math/rand"
)

// GenerateRandomID - generate safe unique ID
func GenerateRandomID() string {
	size := 15
	bs := make([]byte, size)
	rand.Read(bs)
	return hex.EncodeToString(bs)
}
