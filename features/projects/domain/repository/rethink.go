package repository

import (
	"github.com/rajihawa/unmask/core/database/rethink"
	"github.com/rajihawa/unmask/features/projects/domain/models"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type rethinkProjectRepo struct {
	db r.Term
}

func NewRethinkProjectRepository() models.ProjectRepository {
	return &rethinkProjectRepo{
		db: r.Table(rethink.ProjectsTableName),
	}
}

func (p *rethinkProjectRepo) GetAll() ([]models.Project, error) {
	cur, err := p.db.Run(rethink.Session)
	if err != nil {
		return nil, err
	}

	var projects []models.Project

	err = cur.All(&projects)
	defer cur.Close()
	if err != nil {
		return nil, err
	}

	return projects, nil
}
