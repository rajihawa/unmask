package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtSigningKey = []byte(os.Getenv("JWT_SIGN_KEY"))
)

type Token struct {
}

func CreateToken(claims jwt.MapClaims) (string, string, int, error) {
	// Create a token with the claims
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	// Sign the token send the tokenString
	tokenString, err := token.SignedString(JwtSigningKey)
	if err != nil {
		return "", "Can't sign token", http.StatusInternalServerError, err
	}
	return tokenString, "", http.StatusOK, nil
}

func ValidateToken(token string) (jwt.MapClaims, string, int, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSigningKey, nil
	})
	if err != nil {
		return nil, "Cant validate token.", http.StatusInternalServerError, err
	}
	exp, err := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
	if err != nil {
		return nil, "Cant parse expiration date.", http.StatusInternalServerError, err
	}
	currTime := time.Now().Unix()

	if currTime > exp {
		var emptyErr error
		return nil, "Token expired.", http.StatusUnauthorized, emptyErr
	}
	return claims, "", http.StatusOK, nil
}
