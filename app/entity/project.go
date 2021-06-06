package entity

import (
	"time"
)

// Project - projects are a collection of users and clients that share certain properties.
type Project struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Users       []User
}

// NewProject returns new project
func NewProject(name, description string) (*Project, error) {
	p := &Project{
		ID:          GenerateRandomID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return p, nil
}

// AddUser - adds a user to the project
func (p *Project) AddUser(u User) error {
	_, err := p.findUserIndex(u.ID)
	if err == ErrNotFound {
		p.Users = append(p.Users, u)
		return nil
	}
	return ErrUserAlreadyExists
}

// GetUser - find a user in the project
func (p *Project) GetUser(id string) (*User, error) {
	uIndex, err := p.findUserIndex(id)
	if err != nil {
		return nil, err
	}
	return &(p.Users[uIndex]), nil
}

// RemoveUser - removes a user from the project
func (p *Project) RemoveUser(id string) error {
	uIndex, err := p.findUserIndex(id)
	if err != nil {
		return err
	}
	p.Users = append(p.Users[:uIndex], p.Users[uIndex+1:]...)
	return nil
}

// private functions
func (p *Project) findUserIndex(id string) (int, error) {
	for i, u := range p.Users {
		if u.ID == id {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
