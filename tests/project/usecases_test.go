package project_test

import (
	"log"
	"testing"

	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/domain"
)

var appConfig = app.AppConfig{DB: domain.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost",
	Database: "db",
	Port:     "3306",
	Username: "user",
	Password: "password",
},
	Env: "testing"}

func TestProjectUsecases(t *testing.T) {
	app := app.InitApp(appConfig)
	defer app.Close()

	project := &domain.Project{
		Name:        "test project",
		Description: "test description of project",
	}
	log.Println("Testing CreateProject")
	err := app.Project.CreateProject(*project)
	if err != nil {
		t.Errorf("Failed to create project %+v", err)
	}

	log.Println("Testing GetProjects")
	projects, err := app.Project.GetProjects(10, 0)
	if err != nil {
		t.Errorf("Failed to get projects %+v", err)
	}
	projectsCount := len(projects)
	if projectsCount != 1 {
		t.Errorf("Projects length expected %d but got %d", 1, projectsCount)
	}

	createdProject := projects[0]
	newProject := createdProject
	newProject.Name = "updated test project"

	log.Println("Testing UpdateProject")
	err = app.Project.UpdateProject(createdProject.ID, newProject)
	if err != nil {
		t.Errorf("Failed to update project %+v", err)
	}

	log.Println("Testing GetProject")
	fetchedProject, err := app.Project.GetProject(createdProject.ID)
	if err != nil {
		t.Errorf("Failed to get project %s, %+v", createdProject.ID, err)
	}
	if fetchedProject.Name != newProject.Name {
		t.Errorf("Project name expected %s but got %s", newProject.Name, fetchedProject.Name)
	}

	log.Println("Testing DeleteProject")
	err = app.Project.DeleteProject(fetchedProject.ID)
	if err != nil {
		t.Errorf("Failed to delete project %s, %+v", fetchedProject.ID, err)
	}
}
