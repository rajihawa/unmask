package entity_test

import (
	"testing"

	"github.com/rajihawa/mask-off/app/entity"
	"github.com/stretchr/testify/assert"
)

var name = "New project"
var description = "a new project for testing."

func TestNewProject(t *testing.T) {
	p, err := entity.NewProject(name, description)
	assert.Nil(t, err)
	assert.Equal(t, p.Name, name)
	assert.Equal(t, p.Description, description)
	assert.NotNil(t, p.ID)
	assert.Equal(t, 0, len(p.Users))
}

func TestAddUser(t *testing.T) {
	p, _ := entity.NewProject(name, description)
	u, _ := entity.NewUser(username, password)
	p.AddUser(*u)
	assert.Equal(t, 1, len(p.Users))
}
