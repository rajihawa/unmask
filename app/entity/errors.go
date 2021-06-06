package entity

import "errors"

// ErrNotFound - not found
var ErrNotFound = errors.New("not found")

// ErrCannotBeDeleted - cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")
