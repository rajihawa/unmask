package usecases

import (
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/utils"
)

type ClientUsecases struct {
	repo domain.ClientRepo
	env  domain.Env
}

func NewClientUsecases(repo domain.ClientRepo, env domain.Env) domain.ClientUsecases {
	return &ClientUsecases{
		repo: repo,
		env:  env,
	}
}

func (cu ClientUsecases) GetClient(id string) (*domain.Client, error) {
	return cu.repo.GetOne(id)
}

func (cu ClientUsecases) GetClients(limit int, offset int) ([]domain.Client, error) {
	return cu.repo.GetAll(limit, offset)
}

func (cu ClientUsecases) UpdateClient(id string, newClient domain.Client) error {
	return cu.repo.UpdateOne(id, newClient)
}

func (cu ClientUsecases) CreateClient(projectID string, newClient domain.Client) (string, error) {
	id := utils.GenerateClientID(cu.env.Domain)
	client := &newClient
	client.ID = id
	client.ProjectID = projectID
	if !client.Public {
		client.Secret = utils.GenerateClientSecret()
	}
	if err := cu.repo.CreateOne(*client); err != nil {
		return "", err
	}
	return id, nil
}

func (cu ClientUsecases) DeleteClient(id string) error {
	return cu.repo.DeleteOne(id)
}
