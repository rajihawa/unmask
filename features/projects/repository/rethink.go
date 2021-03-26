package repository

import (
	"github.com/rajihawa/unmask/core/database/rethink"
	"github.com/rajihawa/unmask/features/projects/domain"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type rethinkProjectRepo struct {
	db r.Term
}

func NewRethinkProjectRepository() domain.ProjectRepository {
	return &rethinkProjectRepo{
		db: r.Table(rethink.ProjectsTableName),
	}
}

func (p *rethinkProjectRepo) GetAll() ([]domain.Project, error) {
	cur, err := p.db.Run(rethink.Session)
	if err != nil {
		return nil, err
	}

	var projects []domain.Project

	err = cur.All(&projects)
	defer cur.Close()
	if err != nil {
		return nil, err
	}

	return projects, nil
}
