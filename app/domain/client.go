package domain

import "time"

// Client - the client type for the app
type Client struct {
	ID          string
	Secret      string
	Name        string
	Description string
	HomeURL     string
	CallbackURL string
	Privileges  string
	AutoVerify  bool
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProjectID   string
}

// The ClientRepo will have all the database functions for the client type
type ClientRepo interface {
	GetOne(id string) (*Client, error)
	GetAll(limit int, offset int) ([]Client, error)
	UpdateOne(id string, newClient Client) error
	CreateOne(newClient Client) error
	DeleteOne(id string) error
}

// The ClientUsecases will contain all the business logic of the projects
type ClientUsecases interface {
	GetClient(id string) (*Client, error)
	GetClients(limit int, offset int) ([]Client, error)
	UpdateClient(id string, newClient Client) error
	CreateClient(projectID string, newClient Client) (string, error)
	DeleteClient(id string) error
}
