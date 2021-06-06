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
func NewProject(name, description string) *Project {
	p := &Project{
		ID:          GenerateRandomID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return p
}

// AddUser - adds a user to the project
func (p *Project) AddUser(u User) {
	p.Users = append(p.Users, u)
}

// GetUser - find a user in the project
func (p *Project) GetUser(id string) (*User, error) {
	for i, u := range p.Users {
		if u.ID == id {
			return &p.Users[i], nil
		}
	}
	return nil, ErrNotFound
}

// RemoveUser - removes a user from the project
func (p *Project) RemoveUser(id string) error {
	for i, u := range p.Users {
		if u.ID == id {
			p.Users = append(p.Users[:i], p.Users[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
