package entity

import "time"

// Project - projects are a collection of users and clients that share certain properties.
type Project struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
