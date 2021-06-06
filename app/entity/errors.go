package entity

import "errors"

// ErrNotFound - not found
var ErrNotFound = errors.New("not found")

// ErrCannotBeDeleted - cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

// ErrUserAlreadyExists = user already exists in the project
var ErrUserAlreadyExists = errors.New("user already exists")
