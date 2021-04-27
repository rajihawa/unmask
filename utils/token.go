package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rajihawa/unmask/config"
	"github.com/rajihawa/unmask/domain"
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

func GenerateToken(claims jwt.MapClaims) string {
	// Create a token with the claims
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	// Sign the token send the tokenString
	tokenString, err := token.SignedString(JwtSigningKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func GenerateAuthToken(userID string) string {
	exp := config.AccessTokenExpire()
	claims := jwt.MapClaims{
		"user": userID,
		"exp":  fmt.Sprintf("%d", exp.Unix()),
	}

	return GenerateToken(claims)
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func GenerateAuthResponse(user domain.User) (authResponse, error) {
	accessExp := config.AccessTokenExpire()
	refreshExp := config.RefreshTokenExpire()

	accessClaims := jwt.MapClaims{
		"user": user,
		"exp":  fmt.Sprintf("%d", accessExp.Unix()),
	}

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     fmt.Sprintf("%d", refreshExp.Unix()),
	}

	return authResponse{
		AccessToken:  GenerateToken(accessClaims),
		RefreshToken: GenerateToken(refreshClaims),
		ExpiresIn:    accessExp.Unix(),
		Type:         "Bearer",
	}, nil
}

func VerifyToken(authBearer string) (jwt.MapClaims, error) {
	token := strings.Split(authBearer, "Bearer ")[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	exp, err := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
	if err != nil {
		return nil, err
	}
	currTime := time.Now().Unix()

	if currTime > exp {
		return nil, errors.New("expired token")
	}
	return claims, nil
}
