package repository

import "github.com/rajihawa/unmask/domain"

type ProjectRepo struct {
}

func (p *ProjectRepo) GetOne(id string) (*domain.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (p *ProjectRepo) GetAll(limit int, offset int) ([]domain.Project, error) {
	panic("not implemented") // TODO: Implement
}

func (p *ProjectRepo) UpdateOne(id string, newProject domain.Project) error {
	panic("not implemented") // TODO: Implement
}

func (p *ProjectRepo) CreateOne(newProject domain.Project) error {
	panic("not implemented") // TODO: Implement
}

func (p *ProjectRepo) DeleteOne(id string) error {
	panic("not implemented") // TODO: Implement
}

func NewProjectRepo() domain.ProjectRepo {
	return &ProjectRepo{}
}
