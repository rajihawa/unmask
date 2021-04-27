package repository

import (
	"github.com/rajihawa/unmask/app/database"
	"github.com/rajihawa/unmask/domain"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type ProjectRepository struct {
	db   r.Term
	sess *r.Session
}

// NewProjectRepository will create an object that represent the project.Repository interface
func NewProjectRepository() domain.ProjectRepository {
	return &ProjectRepository{
		db:   r.Table(database.ProjectsTableName),
		sess: database.Session,
	}
}

func filterProjectOpts(baseTerm r.Term, opts domain.GetProjectOpts) (term r.Term) {
	term = baseTerm
	// if opts.GetClients {
	// 	term = term.Merge(func(p r.Term) interface{} {
	// 		return map[string]interface{}{
	// 			"client_ids": r.Table(utils.ClientsTableName).GetAll(r.Args(p.Field("client_ids"))).CoerceTo("array"),
	// 		}
	// 	})
	// } else {
	// 	term = term.Without("client_id")
	// }
	// if opts.GetUsers {
	// 	term = term.Merge(func(p r.Term) interface{} {
	// 		return map[string]interface{}{
	// 			"user_ids": r.Table(utils.UsersTableName).GetAll(r.Args(p.Field("user_ids"))).CoerceTo("array"),
	// 		}
	// 	})
	// } else {
	// 	term = term.Without("project_id")
	// }
	return
}

func (p *ProjectRepository) GetAll(opts domain.GetProjectOpts) ([]domain.Project, error) {
	term := p.db
	term = filterProjectOpts(term, opts)
	cur, err := term.Run(p.sess)
	if err != nil {
		return nil, err
	}

	var projects []domain.Project

	err = cur.All(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
func (p *ProjectRepository) Get(id string, opts domain.GetProjectOpts) (*domain.Project, error) {
	term := p.db.Get(id)
	term = filterProjectOpts(term, opts)
	cur, err := term.Run(p.sess)
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

func (p *ProjectRepository) SetUserCount(projectID string, newCount int) error {
	return p.db.Get(projectID).Update(map[string]interface{}{
		"users_count": newCount,
	}).Exec(p.sess)
}
