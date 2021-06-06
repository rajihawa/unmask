package project

import "github.com/rajihawa/mask-off/app/entity"

type Reader interface {
	Get(id string) (*entity.Project, error)
	List() ([]*entity.Project, error)
}

type Writer interface {
	Create(p *entity.Project) (string, error)
	Update(p *entity.Project) error
	Delete(id string) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetProject(id string) (*entity.Project, error)
	ListProjects() ([]*entity.Project, error)
	CreateProject(name, description string) (string, error)
	UpdateProject(p *entity.Project) error
	DeleteProject(id string) error
}
