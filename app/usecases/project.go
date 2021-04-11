package usecases

import (
	"github.com/rajihawa/unmask/domain"
)

type ProjectUsecases struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(repo domain.ProjectRepository) domain.ProjectUsecases {
	return &ProjectUsecases{
		repo: repo,
	}
}

func (pu *ProjectUsecases) CreateProject(newProject domain.Project) error {
	return pu.repo.Insert(newProject)
}

func (pu *ProjectUsecases) GetAll(opts domain.GetProjectOpts) ([]domain.Project, error) {
	return pu.repo.GetAll(opts)
}

func (pu *ProjectUsecases) GetProject(id string, opts domain.GetProjectOpts) (*domain.Project, error) {
	return pu.repo.Get(id, opts)
}

func (pu *ProjectUsecases) AddUserCount(projectID string) error {
	project, err := pu.repo.Get(projectID, domain.GetProjectOpts{})
	if err != nil {
		return err
	}
	return pu.repo.SetUserCount(project.ID, project.UsersCount+1)
}
