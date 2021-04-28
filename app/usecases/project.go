package usecases

import (
	"github.com/google/uuid"
	"github.com/rajihawa/unmask/app/domain"
)

type ProjectUsecases struct {
	repo domain.ProjectRepo
}

func NewProjectUsecases(repo domain.ProjectRepo) domain.ProjectUsecases {
	return &ProjectUsecases{
		repo: repo,
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

func (u *ProjectUsecases) CreateProject(newProject domain.Project) error {
	id := uuid.New().String()
	project := domain.Project{
		ID:          id,
		Name:        newProject.Name,
		Description: newProject.Description,
	}
	return u.repo.CreateOne(project)
}

func (u *ProjectUsecases) DeleteProject(id string) error {
	return u.repo.DeleteOne(id)
}
