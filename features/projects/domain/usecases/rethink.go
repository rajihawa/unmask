package usecases

import (
	"github.com/rajihawa/unmask/core/database/rethink"
	"github.com/rajihawa/unmask/features/projects/domain/models"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type rethinkProjectUsecase struct {
	db r.Term
}

func NewRethinkProjectUsecase() models.ProjectUsercase {
	return &rethinkProjectUsecase{
		db: r.Table(rethink.ProjectsTableName),
	}
}

func (p *rethinkProjectUsecase) Insert(newProject models.Project) error {
	err := p.db.Insert(newProject).Exec(rethink.Session)
	if err != nil {
		return err
	}
	return nil
}
