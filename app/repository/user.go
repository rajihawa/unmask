package repository

import (
	"github.com/rajihawa/unmask/app/database"
	"github.com/rajihawa/unmask/domain"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type UsersRepository struct {
	db   r.Term
	sess *r.Session
}

// NewRethinkProjectRepository will create an object that represent the project.Repository interface
func NewUsersRepository() domain.UsersRepository {
	return &UsersRepository{
		db:   r.Table(database.UsersTableName),
		sess: database.Session,
	}
}

func filterUserOpts(baseTerm r.Term, opts domain.GetUsersOpts) (term r.Term) {
	term = baseTerm
	if opts.GetClient {
		term = term.Merge(func(p r.Term) interface{} {
			return map[string]interface{}{
				"client_id": r.Table(database.ClientsTableName).Get(p.Field("client_id")),
			}
		})
	} else {
		term = term.Without("client_id")
	}
	if opts.GetProject {
		term = term.Merge(func(p r.Term) interface{} {
			return map[string]interface{}{
				"project_id": r.Table(database.ProjectsTableName).Get(p.Field("project_id")),
			}
		})
	} else {
		term = term.Without("project_id")
	}
	return
}

func (u *UsersRepository) GetAll(projectID string, opts domain.GetUsersOpts) ([]domain.User, error) {
	term := u.db.GetAllByIndex("project_id", projectID)
	term = filterUserOpts(term, opts)

	cur, err := term.Run(u.sess)
	if err != nil {
		return nil, err
	}

	var users []domain.User

	err = cur.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UsersRepository) Get(id string, opts domain.GetUsersOpts) (*domain.User, error) {
	term := u.db.Get(id)
	term = filterUserOpts(term, opts)
	cur, err := term.Run(u.sess)
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = cur.One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepository) GetByUsername(username string, projectID string, opts domain.GetUsersOpts) ([]domain.User, error) {

	term := u.db.GetAllByIndex("username", username).Filter(map[string]interface{}{
		"project_id": projectID,
	})
	term = filterUserOpts(term, opts)

	cur, err := term.Run(u.sess)

	if err != nil {
		return nil, err
	}
	var users []domain.User
	err = cur.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UsersRepository) Insert(user domain.User) error {
	err := u.db.Insert(user).Exec(u.sess)
	if err != nil {
		return err
	}
	return nil
}
