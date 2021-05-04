package domain

import "time"

// Project - the project type for the app
type Project struct {
	ID          string
	Name        string
	Description string
	UserCount   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// The ProjectRepo will have all the database functions for the project type
type ProjectRepo interface {
	GetOne(id string) (*Project, error)
	GetAll(limit int, offset int) ([]Project, error)
	UpdateOne(id string, newProject Project) error
	CreateOne(newProject Project) error
	DeleteOne(id string) error
}

// The ProjectUsecases will contain all the business logic of the projects
type ProjectUsecases interface {
	GetProject(id string) (*Project, error)
	GetProjects(limit int, offset int) ([]Project, error)
	UpdateProject(id string, newProject Project) error
	CreateProject(newProject Project) (string, error)
	DeleteProject(id string) error
}
