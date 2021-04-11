package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Declare a new project struct
	project := domain.Project{}

	// TODO: make validator layer
	// Try to decode the request body into struct
	// If there is an error, respond to client with 400 status code
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't create project.")
		return
	}

	err = usecases.NewProjectUsecase(repository.NewProjectRepository()).CreateProject(project)
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't create project.")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {

	projects, err := usecases.NewProjectUsecase(repository.NewProjectRepository()).GetAll(domain.GetProjectOpts{})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't get projects.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
func GetProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["project"]

	projects, err := usecases.NewProjectUsecase(repository.NewProjectRepository()).GetProject(projectID, domain.GetProjectOpts{})
	if err != nil {
		utils.HttpError(w, err, http.StatusBadRequest, "Couldn't get projects.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
