package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/config"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID := vars["client"]

	client, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetClient(clientID, domain.GetClientOpts{GetProjects: true})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't find client.")
		return
	}

	var user domain.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Wrong user data.")
		return
	}

	err = usecases.NewUsersUsecase(repository.NewUsersRepository()).SignupUser(*client, &user)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't signup user.")
		return
	}

	err = usecases.NewProjectUsecase(repository.NewProjectRepository()).AddUserCount(client.Project.ID)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't increament project users.")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// func AdminCreateUser(w http.ResponseWriter, r *http.Request) {

// 	var user *domain.User
// 	err := json.NewDecoder(r.Body).Decode(user)
// 	if err != nil {
// 		utils.HttpError(w, err, http.StatusBadRequest, "Wrong user data.")
// 		return
// 	}

// 	err = usecases.NewUsersUsecase(repository.NewUsersRepository()).AdminCreateUser(*user)
// 	if err != nil {
// 		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't signup user.")
// 		return
// 	}

// 	err = usecases.NewProjectUsecase(repository.NewProjectRepository()).AddUserCount(user.Project.ID)
// 	if err != nil {
// 		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't increament project users.")
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin domain.UserLogin
	client := r.Context().Value(utils.ContextClientKey).(domain.Client)

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Wrong user data.")
		return
	}

	user, err := usecases.NewUsersUsecase(repository.NewUsersRepository()).CheckUserLogin(userLogin, client)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Wrong username or password.")
		return
	}

	exp := config.AccessTokenExpire()

	// userJson, _ := json.Marshal(&user)
	claims := jwt.MapClaims{
		"user_id": &user.ID,
		"exp":     fmt.Sprintf("%d", exp.Unix()),
	}

	tokenString, errMsg, errCode, err := utils.CreateToken(claims)
	if err != nil {
		utils.HttpError(w, err, errCode, errMsg)
		return
	}

	cookieMgr := utils.CreateJwtCookie(utils.AccessTokenCookieName, tokenString, exp)
	cookieMgr.SetCookie(w)

	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(utils.ContextUserKey).(domain.User)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
