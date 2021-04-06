package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := &domain.Project{
		ID: vars["project"],
	}

	opts := domain.GetClientOpts{
		GetProjects: vars["show_projects"] == "true",
	}

	clients, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetAll(*project, opts)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't get clients.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)

}

func CreateNewClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	client := &domain.Client{}
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client.Project = &domain.Project{ID: vars["project"]}
	err = usecases.NewClientUsecase(repository.NewClientRepository()).CreateClient(client)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't create client.")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetClientSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID := vars["client"]
	secret, err := usecases.NewClientUsecase(repository.NewClientRepository()).GetClientSecret(clientID)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't get client secret.")
		return
	}
	fmt.Fprint(w, secret)
}
