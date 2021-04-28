package project_test

import (
	"testing"

	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/domain"
)

func TestProjectUsecases(t *testing.T) {
	app := app.InitApp(app.AppConfig{DB: app.DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Database: "db",
		Port:     "3306",
		Username: "user",
		Password: "password",
	}})

	project := &domain.Project{
		Name:        "test project",
		Description: "test description of project",
	}

	err := app.Project.CreateProject(*project)
	if err != nil {
		t.Errorf("Failed to create project %+v", err)
	}

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

	err = app.Project.UpdateProject(createdProject.ID, newProject)
	if err != nil {
		t.Errorf("Failed to update project %+v", err)
	}

	fetchedProject, err := app.Project.GetProject(createdProject.ID)
	if err != nil {
		t.Errorf("Failed to get project %s, %+v", createdProject.ID, err)
	}
	if fetchedProject.Name != newProject.Name {
		t.Errorf("Project name expected %s but got %s", newProject.Name, fetchedProject.Name)
	}

	err = app.Project.DeleteProject(fetchedProject.ID)
	if err != nil {
		t.Errorf("Failed to delete project %s, %+v", fetchedProject.ID, err)
	}
}
