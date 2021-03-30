package domain

import "time"

// Project ...
type Project struct {
	ID        string    `json:"id,omitempty" rethinkdb:"id,omitempty"`
	Name      string    `json:"name" rethinkdb:"name"`
	UpdatedAt time.Time `json:"-" rethinkdb:"updated_at"`
	CreatedAt time.Time `json:"-" rethinkdb:"created_at"`
}

// ProjectEntity - the project's repository
type ProjectRepository interface {
	GetAll() ([]Project, error)
	Insert(project Project) error
}

// ProjectUsecases - the project's usecases
type ProjectUsecases interface {
	GetAll() ([]Project, error)
	CreateProject(newProject Project) error
}
