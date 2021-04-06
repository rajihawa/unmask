package repository

import (
	"github.com/rajihawa/unmask/domain"
	"github.com/rajihawa/unmask/utils"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type ProjectRepository struct {
	db   r.Term
	sess *r.Session
}

// NewProjectRepository will create an object that represent the project.Repository interface
func NewProjectRepository() domain.ProjectRepository {
	return &ProjectRepository{
		db:   r.Table(utils.ProjectsTableName),
		sess: utils.Session,
	}
}

func (p *ProjectRepository) GetAll() ([]domain.Project, error) {
	cur, err := p.db.Run(p.sess)
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
func (p *ProjectRepository) Get(id string) (*domain.Project, error) {
	cur, err := p.db.Get(id).Run(p.sess)
	if err != nil {
		return nil, err
	}

	var project domain.Project

	err = cur.One(&project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *ProjectRepository) Insert(project domain.Project) error {
	err := p.db.Insert(project).Exec(p.sess)
	if err != nil {
		return err
	}
	return nil
}
