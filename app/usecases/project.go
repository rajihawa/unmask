package usecase

import "github.com/rajihawa/unmask/domain"

type ProjectUsecases struct {
	repo domain.ProjectRepo
}

func NewProjectUsecases(repo domain.ProjectRepo) domain.ProjectUsecases {
	return &ProjectUsecases{
		repo: repo,
	}
}
func (u *ProjectUsecases) GetProject(id string) (*domain.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (u *ProjectUsecases) GetProjects(limit int, offset int) ([]domain.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (u *ProjectUsecases) UpdateProject(id string, newProject domain.Project) error {
	panic("not implemented") // TODO: Implement
}

func (u *ProjectUsecases) CreateProject(newProject domain.Project) error {
	panic("not implemented") // TODO: Implement
}

func (u *ProjectUsecases) DeleteProject(id string) error {
	panic("not implemented") // TODO: Implement
}
