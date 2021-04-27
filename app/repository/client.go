package repository

import (
	"errors"

	"github.com/rajihawa/unmask/app/database"
	"github.com/rajihawa/unmask/domain"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type ClientRepository struct {
	db   r.Term
	sess *r.Session
}

// NewRethinkProjectRepository will create an object that represent the project.Repository interface
func NewClientRepository() domain.ClientRepository {
	return &ClientRepository{
		db:   r.Table(database.ClientsTableName),
		sess: database.Session,
	}
}

func (c *ClientRepository) GetAll(projectID string, opts domain.GetClientOpts) ([]domain.Client, error) {
	term := c.db.GetAllByIndex("project_id", projectID).Without("client_secret")
	if opts.GetProjects {
		term = term.Merge(func(p r.Term) interface{} {
			return map[string]interface{}{
				"project_id": r.Table(database.ProjectsTableName).Get(p.Field("project_id")),
			}
		})
	} else {
		term = term.Without("project_id")
	}

	cur, err := term.Run(c.sess)
	if err != nil {
		return nil, err
	}

	var clients []domain.Client

	err = cur.All(&clients)
	if err != nil {
		return nil, err
	}

	return clients, nil
}
func (c *ClientRepository) Get(clientID string, opts domain.GetClientOpts) (*domain.Client, error) {

	term := c.db.Get(clientID).Without("client_secret")
	if opts.GetProjects {
		term = term.Merge(func(p r.Term) interface{} {
			return map[string]interface{}{
				"project_id": r.Table(database.ProjectsTableName).Get(p.Field("project_id")),
			}
		})
	} else {
		term = term.Without("project_id")
	}

	cur, err := term.Run(c.sess)

	if err == r.ErrEmptyResult {
		return nil, errors.New("client does not exist")
	}
	if err != nil {

		return nil, err
	}

	var client domain.Client

	err = cur.One(&client)
	if err != nil {

		return nil, err
	}

	return &client, nil
}

func (c *ClientRepository) Insert(client domain.Client) error {
	err := c.db.Insert(client).Exec(c.sess)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientRepository) GetClientSecret(clientID string) (string, error) {
	cur, err := c.db.Get(clientID).Pluck("client_secret").Run(c.sess)
	if err != nil {
		return "", err
	}
	var data map[string]string
	err = cur.One(&data)
	if err != nil {
		return "", err
	}
	return data["client_secret"], nil
}
