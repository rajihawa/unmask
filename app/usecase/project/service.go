package project

import (
	"time"

	"github.com/rajihawa/mask-off/app/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetProject(id string) (*entity.Project, error) {
	return s.repo.Get(id)
}

func (s *Service) ListProjects() ([]*entity.Project, error) {
	return s.repo.List()
}

func (s *Service) CreateProject(name, description string) (string, error) {
	p, err := entity.NewProject(name, description)
	if err != nil {
		return "", err
	}
	return s.repo.Create(p)
}

func (s *Service) UpdateProject(p *entity.Project) error {
	_, err := s.GetProject(p.ID)
	if err != nil {
		return err
	}
	p.UpdatedAt = time.Now()
	return s.repo.Update(p)
}

func (s *Service) DeleteProject(id string) error {
	_, err := s.GetProject(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
