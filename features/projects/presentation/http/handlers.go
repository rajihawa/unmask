package http

import (
	"encoding/json"
	"net/http"

	"github.com/rajihawa/unmask/core/lib"
	"github.com/rajihawa/unmask/features/projects/domain"
	"github.com/rajihawa/unmask/features/projects/repository"
	"github.com/rajihawa/unmask/features/projects/usecases"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Declare a new project struct
	project := domain.Project{}

	// Try to decode the request body into struct
	// If there is an error, respond to client with 400 status code
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		lib.HttpError(w, err, http.StatusBadRequest, "Couldn't create project.")
		return
	}

	err = usecases.NewRethinkProjectUsecase().Insert(project)
	if err != nil {
		lib.HttpError(w, err, http.StatusBadRequest, "Couldn't create project.")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {

	projects, err := repository.NewRethinkProjectRepository().GetAll()
	if err != nil {
		lib.HttpError(w, err, http.StatusBadRequest, "Couldn't get projects.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
