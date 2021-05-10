package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/rajihawa/unmask/app/data"
	"github.com/rajihawa/unmask/app/domain"
)

type UserMySqlRepo struct {
	db *sql.DB
}

func NewUserMySqlRepo() domain.UserRepo {
	return &UserMySqlRepo{
		db: data.MySQL,
	}
}

func (c UserMySqlRepo) GetOne(id string) (*domain.User, error) {
	clientQuery := squirrel.Select("*").From("users").Where(squirrel.Eq{"id": id})
	rows, err := clientQuery.RunWith(c.db).Query()
	if err != nil {
		return nil, err
	}
	emptyUser := domain.User{}
	for rows.Next() {
		var attrs string
		// err := rows.Scan(&emptyUser.ID, &emptyUser.Secret, &emptyUser.Name, &emptyUser.Description, &emptyUser.HomeURL, &emptyUser.CallbackURL, &emptyUser.Privileges, &emptyUser.AutoVerify, &emptyUser.CreatedAt, &emptyUser.UpdatedAt, &emptyUser.ProjectID)
		err := rows.Scan(&emptyUser.ID, &emptyUser.Username, &emptyUser.Email, &emptyUser.PasswordHash, &attrs, &emptyUser.Verified, &emptyUser.CreatedAt, &emptyUser.UpdatedAt, &emptyUser.ProjectID, &emptyUser.ClientID)
		if err != nil {
			return nil, err
		}
		var attrsMap map[string]interface{}
		if err := json.Unmarshal([]byte(attrs), &attrsMap); err != nil {
			return nil, err
		}
		emptyUser.Attributes = attrsMap
	}
	return &emptyUser, nil
}

func (c UserMySqlRepo) GetAll(limit int, offset int) ([]domain.User, error) {
	usersQuery := squirrel.Select("*").From("users").Limit(uint64(limit)).Offset(uint64(offset))
	rows, err := usersQuery.RunWith(c.db).Query()
	if err != nil {
		return nil, err
	}
	emptyUsers := []domain.User{}
	for rows.Next() {
		emptyUser := domain.User{}
		var attrs string
		err := rows.Scan(&emptyUser.ID, &emptyUser.Username, &emptyUser.Email, &emptyUser.PasswordHash, &attrs, &emptyUser.Verified, &emptyUser.CreatedAt, &emptyUser.UpdatedAt, &emptyUser.ProjectID, &emptyUser.ClientID)
		if err != nil {
			return nil, err
		}
		var attrsMap map[string]interface{}
		if err := json.Unmarshal([]byte(attrs), &attrsMap); err != nil {
			return nil, err
		}
		emptyUser.Attributes = attrsMap

		emptyUsers = append(emptyUsers, emptyUser)
	}
	return emptyUsers, nil
}

func (c UserMySqlRepo) UpdateOne(id string, newUser domain.User) error {
	attrs, err := json.Marshal(newUser.Attributes)
	if err != nil {
		return err
	}
	updateQuery := squirrel.Update("users").Set("username", newUser.Username).Set("email", newUser.Email).Set("password_hash", newUser.PasswordHash).Set("attributes", string(attrs)).Set("verified", newUser.Verified).Where(squirrel.Eq{"id": id})
	_, err = updateQuery.RunWith(c.db).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (c UserMySqlRepo) CreateOne(newUser domain.User) error {
	attrs, err := json.Marshal(newUser.Attributes)
	if err != nil {
		return err
	}
	createQuery := squirrel.Insert("users").Columns("id", "username", "email", "password_hash", "attributes", "verified", "project_id", "client_id").Values(newUser.ID, newUser.Username, newUser.Email, newUser.PasswordHash, string(attrs), newUser.Verified, newUser.ProjectID, newUser.ClientID)
	_, err = createQuery.RunWith(c.db).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (c UserMySqlRepo) DeleteOne(id string) error {
	deleteQuery := squirrel.Delete("users").Where(squirrel.Eq{"id": id})
	_, err := deleteQuery.RunWith(c.db).Exec()
	if err != nil {
		return err
	}
	return nil
}
