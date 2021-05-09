package usecases

import (
	"github.com/google/uuid"
	"github.com/rajihawa/unmask/app/domain"
)

type ProjectUsecases struct {
	repo domain.ProjectRepo
	env  domain.Env
}

func NewProjectUsecases(repo domain.ProjectRepo, env domain.Env) domain.ProjectUsecases {
	return &ProjectUsecases{
		repo: repo,
		env:  env,
	}
}
func (u *ProjectUsecases) GetProject(id string) (*domain.Project, error) {
	return u.repo.GetOne(id)
}

func (u *ProjectUsecases) GetProjects(limit int, offset int) ([]domain.Project, error) {
	return u.repo.GetAll(limit, offset)
}

func (u *ProjectUsecases) UpdateProject(id string, newProject domain.Project) error {
	return u.repo.UpdateOne(id, newProject)
}

func (u *ProjectUsecases) CreateProject(newProject domain.Project) (string, error) {
	id := uuid.New().String()
	project := &newProject
	project.ID = id
	if err := u.repo.CreateOne(*project); err != nil {
		return "", err
	}
	return id, nil
}

func (u *ProjectUsecases) DeleteProject(id string) error {
	return u.repo.DeleteOne(id)
}
