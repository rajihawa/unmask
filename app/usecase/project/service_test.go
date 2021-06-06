package project_test

import (
	"testing"
	"time"

	"github.com/rajihawa/mask-off/app/entity"
	"github.com/rajihawa/mask-off/app/usecase/project"
	"github.com/stretchr/testify/assert"
)

func newTestProject() *entity.Project {
	return &entity.Project{
		ID:          entity.GenerateRandomID(),
		Name:        "test project",
		Description: "test project description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestCrud(t *testing.T) {
	repo := project.NewInmem()
	s := project.NewService(repo)
	p := newTestProject()
	id, err := s.CreateProject(p.Name, p.Description)
	assert.Nil(t, err)
	saved, err := s.GetProject(id)
	savedUpdated := saved.UpdatedAt
	assert.Nil(t, err)
	saved.Name = "test2"
	assert.Nil(t, s.UpdateProject(saved))
	updated, err := s.GetProject(id)
	assert.Nil(t, err)
	assert.Equal(t, updated.Name, "test2")
	assert.NotEqual(t, updated.UpdatedAt, savedUpdated)
	pl, err := s.ListProjects()
	assert.Nil(t, err)
	assert.Equal(t, len(pl), 1)
	assert.Nil(t, s.DeleteProject(id))
	npl, err := s.ListProjects()
	assert.Nil(t, err)
	assert.Equal(t, len(npl), 0)
}
