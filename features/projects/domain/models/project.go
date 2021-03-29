package models

import "time"

// Project ...
type Project struct {
	ID        string    `json:"id,omitempty" rethinkdb:"id,omitempty"`
	Name      string    `json:"name" rethinkdb:"name"`
	UpdatedAt time.Time `json:"-" rethinkdb:"updated_at"`
	CreatedAt time.Time `json:"-" rethinkdb:"created_at"`
}

// ProjectRepository - the project's repository
type ProjectRepository interface {
	GetAll() ([]Project, error)
}

// ProjectUsercase - the project's usecases
type ProjectUsercase interface {
	Insert(newProject Project) error
}
