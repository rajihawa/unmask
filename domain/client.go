package domain

import "time"

type Client struct {
	ID           string    `json:"id,omitempty" rethinkdb:"id,omitempty"`
	Name         string    `json:"name" rethinkdb:"name"`
	HomeURL      string    `json:"home_url" rethinkdb:"home_url"`
	Description  string    `json:"description" rethinkdb:"description"`
	CallbackURL  string    `json:"callback_url" rethinkdb:"callback_url"`
	Frontend     bool      `json:"frontend" rethinkdb:"frontend"`
	Signup       bool      `json:"signup" rethinkdb:"signup"`
	Disabled     bool      `json:"disabled" rethinkdb:"disabled"`
	ClientSecret string    `json:"-" rethinkdb:"client_secret,omitempty"`
	UpdatedAt    time.Time `json:"-" rethinkdb:"updated_at"`
	CreatedAt    time.Time `json:"-" rethinkdb:"created_at"`
	Project      *Project  `json:"project,omitempty" rethinkdb:"project_id,reference" rethinkdb_ref:"id"`
}

type GetClientOpts struct {
	GetProjects bool
}

type ClientRepository interface {
	GetAll(projectID string, opts GetClientOpts) ([]Client, error)
	Insert(client Client) error
	GetClientSecret(clientID string) (string, error)
	Get(clientID string, opts GetClientOpts) (*Client, error)
}

type ClientUsecases interface {
	GetAll(project Project, opts GetClientOpts) ([]Client, error)
	CreateClient(newClient *Client) error
	GetClientSecret(clientID string) (string, error)
	GetClient(clientID string, opts GetClientOpts) (*Client, error)
}
