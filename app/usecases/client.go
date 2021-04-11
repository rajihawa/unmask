package usecases

import (
	"crypto/sha256"
	"fmt"

	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
	"github.com/rs/xid"
)

type ClientUsecases struct {
	repo domain.ClientRepository
}

func NewClientUsecase(repo domain.ClientRepository) domain.ClientUsecases {
	return &ClientUsecases{
		repo: repo,
	}
}

func (cu *ClientUsecases) GetAll(project domain.Project, opts domain.GetClientOpts) ([]domain.Client, error) {
	return cu.repo.GetAll(project.ID, opts)
}

func (cu *ClientUsecases) GetClient(clientID string, opts domain.GetClientOpts) (*domain.Client, error) {
	return cu.repo.Get(clientID, opts)
}

func (cu *ClientUsecases) CreateClient(newClient *domain.Client) error {
	newClient.ID = xid.New().String()
	if !newClient.Frontend {
		secret := fmt.Sprintf("%x", sha256.Sum256(xid.New().Bytes()))
		newClient.ClientSecret = utils.EncryptGCM(secret)
	}
	return cu.repo.Insert(*newClient)
}

func (cu *ClientUsecases) GetClientSecret(clientID string) (string, error) {
	return cu.repo.GetClientSecret(clientID)
}
