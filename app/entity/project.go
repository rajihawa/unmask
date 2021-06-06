package entity

import (
	"time"

	"github.com/rajihawa/mask-off/lib/utils"
)

// Project - projects are a collection of users and clients that share certain properties.
type Project struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewProject returns new project
func NewProject(name, description string) *Project {
	p := &Project{
		ID:          utils.GenerateRandomID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return p
}
