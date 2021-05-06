package repository

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
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

func (c ClientMySqlRepo) GetOne(id string) (*domain.Client, error) {
	clientQuery := squirrel.Select("*").From("clients").Where(squirrel.Eq{"id": id})
	rows, err := clientQuery.RunWith(c.db).Query()
	if err != nil {
		return nil, err
	}
	emptyClient := domain.Client{}
	for rows.Next() {
		err := rows.Scan(&emptyClient.ID, &emptyClient.Secret, &emptyClient.Name, &emptyClient.Description, &emptyClient.HomeURL, &emptyClient.CallbackURL, &emptyClient.Privileges, &emptyClient.AutoVerify, &emptyClient.CreatedAt, &emptyClient.UpdatedAt, &emptyClient.ProjectID)
		if err != nil {
			return nil, err
		}
	}
	return &emptyClient, nil
}

func (c ClientMySqlRepo) GetAll(limit int, offset int) ([]domain.Client, error) {
	clientsQuery := squirrel.Select("*").From("clients").Limit(uint64(limit)).Offset(uint64(offset))
	rows, err := clientsQuery.RunWith(c.db).Query()
	if err != nil {
		return nil, err
	}
	emptyClients := []domain.Client{}
	for rows.Next() {
		emptyClient := domain.Client{}
		err := rows.Scan(&emptyClient.ID, &emptyClient.Secret, &emptyClient.Name, &emptyClient.Description, &emptyClient.HomeURL, &emptyClient.CallbackURL, &emptyClient.Privileges, &emptyClient.AutoVerify, &emptyClient.CreatedAt, &emptyClient.UpdatedAt, &emptyClient.ProjectID)
		if err != nil {
			return nil, err
		}
		emptyClients = append(emptyClients, emptyClient)
	}
	return emptyClients, nil
}

func (c ClientMySqlRepo) UpdateOne(id string, newClient domain.Client) error {
	updateQuery := squirrel.Update("clients").Set("name", newClient.Name).Set("description", newClient.Description).Set("home_url", newClient.HomeURL).Set("callback_url", newClient.CallbackURL).Set("privileges", newClient.Privileges).Set("auto_verify", newClient.AutoVerify).Where(squirrel.Eq{"id": id})
	_, err := updateQuery.RunWith(c.db).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (c ClientMySqlRepo) CreateOne(newClient domain.Client) error {
	panic("not implemented") // TODO: Implement
}

func (c ClientMySqlRepo) DeleteOne(id string) error {
	panic("not implemented") // TODO: Implement
}
