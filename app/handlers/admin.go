package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/rajihawa/unmask/config"
	"github.com/rajihawa/unmask/utils"
)

var (
	adminUsername = os.Getenv("ADMIN_USERNAME")
	adminPassword = os.Getenv("ADMIN_PASSWORD")
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
	if loginData.Username != adminUsername || loginData.Password != adminPassword {
		utils.HttpError(w, err, http.StatusUnauthorized, "Username or Password doesn't match.")
		return
	}

	// Create expire date
	exp := config.AdminTokenExpire()

	// Create claims for the jwt
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = fmt.Sprintf("%d", exp.Unix())

	tokenString, errMsg, errCode, err := utils.CreateToken(claims)

	if err != nil {
		utils.HttpError(w, err, errCode, errMsg)
		return
	}

	cookieMgr := utils.CreateJwtCookie(tokenString, exp)
	cookieMgr.SetCookie(w)

	fmt.Fprintf(w, "Person: %+v,\nJWT: %s\n", loginData, tokenString)
}

func AdminMeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"authorized": true}`)
}
