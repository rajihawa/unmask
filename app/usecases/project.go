package usecases

import (
	"github.com/rajihawa/unmask/domain"
)

type ProjectUsecases struct {
	repo domain.ProjectRepository
}

func NewRethinkProjectUsecase(repo domain.ProjectRepository) domain.ProjectUsecases {
	return &ProjectUsecases{
		repo: repo,
	}
}

func (pu *ProjectUsecases) CreateProject(newProject domain.Project) error {
	return pu.repo.Insert(newProject)
}

func (pu *ProjectUsecases) GetAll() ([]domain.Project, error) {
	return pu.repo.GetAll()
}
