package domain

import "time"

// Project ...
type Project struct {
	ID          string    `json:"id,omitempty" rethinkdb:"id,omitempty"`
	Name        string    `json:"name" rethinkdb:"name"`
	Description string    `json:"description" rethinkdb:"description"`
	Email       bool      `json:"email" rethinkdb:"email"`
	VerifyEmail bool      `json:"verify_email" rethinkdb:"verify_email"`
	UsersCount  int       `json:"users_count" rethinkdb:"users_count"`
	UpdatedAt   time.Time `json:"-" rethinkdb:"updated_at"`
	CreatedAt   time.Time `json:"-" rethinkdb:"created_at"`
}

type GetProjectOpts struct {
}

// ProjectEntity - the project's repository
type ProjectRepository interface {
	GetAll(opts GetProjectOpts) ([]Project, error)
	Get(id string, opts GetProjectOpts) (*Project, error)
	Insert(project Project) error
	SetUserCount(projectID string, newCount int) error
}

// ProjectUsecases - the project's usecases
type ProjectUsecases interface {
	GetAll(opts GetProjectOpts) ([]Project, error)
	GetProject(id string, opts GetProjectOpts) (*Project, error)
	CreateProject(newProject Project) error
	AddUserCount(projectID string) error
}
