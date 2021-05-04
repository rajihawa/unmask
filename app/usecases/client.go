package usecases

import (
	"github.com/rajihawa/unmask/app/domain"
)

type ClientUsecases struct {
	repo domain.ClientRepo
}

func NewClientUsecases(repo domain.ClientRepo) domain.ClientUsecases {
	return &ClientUsecases{
		repo: repo,
	}
}

func (cu ClientUsecases) GetClient(id string) (*domain.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (cu ClientUsecases) GetClients(limit int, offset int) ([]domain.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (cu ClientUsecases) UpdateClient(id string, newClient domain.Client) error {
	panic("not implemented") // TODO: Implement
}

func (cu ClientUsecases) CreateClient(projectID string, newClient domain.Client) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (cu ClientUsecases) DeleteClient(id string) error {
	panic("not implemented") // TODO: Implement
}
