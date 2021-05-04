package repository

import (
	"database/sql"

	"github.com/rajihawa/unmask/app/data"
	"github.com/rajihawa/unmask/app/domain"
)

type ClientMySqlRepo struct {
	db *sql.DB
}

func NewClientMySqlRepo() domain.ClientRepo {
	return &ClientMySqlRepo{
		db: data.MySQL,
	}
}

func (cu ClientMySqlRepo) GetOne(id string) (*domain.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (cu ClientMySqlRepo) GetAll(limit int, offset int) ([]domain.Client, error) {
	panic("not implemented") // TODO: Implement
}

func (cu ClientMySqlRepo) UpdateOne(id string, newClient domain.Client) error {
	panic("not implemented") // TODO: Implement
}

func (cu ClientMySqlRepo) CreateOne(newClient domain.Client) error {
	panic("not implemented") // TODO: Implement
}

func (cu ClientMySqlRepo) DeleteOne(id string) error {
	panic("not implemented") // TODO: Implement
}
