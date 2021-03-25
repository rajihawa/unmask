package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/rajihawa/unmask/config"
	"github.com/rajihawa/unmask/lib"
)

type adminLoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminLoginHandler - admin login handler
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Declare a new admin login struct
	var loginData adminLoginStruct

	// Try to decode the request body into struct
	// If there is an error, respond to client with 400 status code
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate username and password
	if loginData.Username != config.AdminUsername || loginData.Password != config.AdminPassword {
		lib.HttpError(w, err, http.StatusUnauthorized, "Username or Password doesn't match.")
		return
	}

	// Create expire date
	exp := config.JwtAdminExpire()

	// Create claims for the jwt
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = fmt.Sprintf("%d", exp.Unix())

	// Create a token with the claims
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	// Sign the token send the tokenString
	tokenString, err := token.SignedString(config.JwtSigningKey)
	if err != nil {
		lib.HttpError(w, err, http.StatusInternalServerError, "Couldn't sign jwt token.")
		return
	}

	// Create a secure cookie
	cookie := &http.Cookie{
		Name:     config.JwtCookieName,
		Expires:  exp,
		Value:    tokenString,
		HttpOnly: true,
		Secure:   config.IsProd,
	}

	// Set the cookie
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Person: %+v,\nJWT: %s\n", loginData, tokenString)
}
