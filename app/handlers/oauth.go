package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

type OauthData struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

func Oauthorize(w http.ResponseWriter, r *http.Request) {
	var oauthData OauthData
	if r.Header["Content-Type"][0] == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		oauthData = OauthData{
			ClientID:     r.PostFormValue("client_id"),
			ClientSecret: r.PostFormValue("client_secret"),
			Username:     r.PostFormValue("username"),
			Password:     r.PostFormValue("password"),
			GrantType:    r.PostFormValue("grant_type"),
			Scope:        r.PostFormValue("scope"),
		}
		fmt.Println(oauthData)
	} else {
		err := json.NewDecoder(r.Body).Decode(&oauthData)
		if err != nil {
			log.Println("err")
			fmt.Println(err)
			utils.HttpError(w, err, http.StatusBadRequest, "Wrong data provided.")
			return
		}
	}

	client, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetClient(oauthData.ClientID, domain.GetClientOpts{GetProjects: true})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "invalid client.")
		return
	}

	user, err := usecases.NewUsersUsecase(repository.NewUsersRepository()).CheckUserLogin(domain.UserLogin{Username: oauthData.Username, Password: oauthData.Password}, *client)
	if err != nil {
		utils.HttpError(w, err, http.StatusUnauthorized, "Wrong username or password.")
		return
	}

	// Redirect to callback
	token := utils.GenerateAuthToken(user.ID)
	// http.Redirect(w, r, fmt.Sprintf("%s?token=%s", client.CallbackURL, token), http.StatusSeeOther)

	// log.Println(*user)

	// res, err := utils.GenerateAuthResponse(*user)
	// if err != nil {
	// 	utils.HttpError(w, err, http.StatusInternalServerError, "Can't generate response token.")
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func Overify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Println("here")
	token := r.Header["Authorization"][0]
	claims, err := utils.VerifyToken(token)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Invalid token.")
		return
	}
	user, err := usecases.NewUsersUsecase(repository.NewUsersRepository()).GetUser(claims["user"].(string), domain.GetUsersOpts{})
	if err != nil {
		utils.HttpError(w, err, http.StatusUnauthorized, "Wrong username or password.")
		return
	}
	res, err := utils.GenerateAuthResponse(*user)
	if err != nil {
		utils.HttpError(w, err, http.StatusInternalServerError, "Can't generate response token.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

func CurrentSession(w http.ResponseWriter, r *http.Request) {
	token := r.Header["Authorization"][0]
	claims, err := utils.VerifyToken(token)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Invalid token.")
		return
	}
	fmt.Println(claims)
	user, err := usecases.NewUsersUsecase(repository.NewUsersRepository()).GetUser(claims["user"].(map[string]interface{})["id"].(string), domain.GetUsersOpts{})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Can't get user.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*user)
}

type clientRes struct {
	Name        string `json:"name"`
	CallbackURL string `json:"callback_url"`
	Signup      bool   `json:"signup"`
}

func Oclient(w http.ResponseWriter, r *http.Request) {
	clientID := mux.Vars(r)["client"]
	client, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetClient(clientID, domain.GetClientOpts{GetProjects: true})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Can't get client.")
		return
	}

	res := &clientRes{
		Name:        client.Project.Name,
		CallbackURL: client.CallbackURL,
		Signup:      client.Signup,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*res)
}
