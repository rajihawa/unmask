package domain

import "time"

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	Attributes   map[string]interface{}
	Verified     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProjectID    string
	ClientID     string
}

type UserSignup struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Attributes      map[string]interface{}
	Scope           string
}

type UserRepo interface {
	GetOne(id string) (*User, error)
	GetAll(limit int, offset int) ([]User, error)
	UpdateOne(id string, newUser User) error
	CreateOne(newUser User) error
	DeleteOne(id string) error
}

type UserUsecases interface {
	GetUser(id string) (*User, error)
	GetUsers(limit int, offset int) ([]User, error)
	UpdateUser(id string, newUser User) error
	CreateUser(projectID string, client Client, newUser UserSignup) (string, error)
	DeleteUser(id string) error
}
