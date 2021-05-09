package app

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rajihawa/unmask/app/data"
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/repository"
	"github.com/rajihawa/unmask/app/usecases"
)

func InitApp(conf domain.AppConfig) domain.App {
	var db domain.Database
	if conf.DB.Driver == "mysql" {
		db = data.NewMySqlDB(conf.DB)
		var dbSess domain.DatabaseSessions
		db.Init(&dbSess)

		return domain.App{
			Project: usecases.NewProjectUsecases(repository.NewProjectMySqlRepo(dbSess.MySQL), conf.Env),
			Client:  usecases.NewClientUsecases(repository.NewClientMySqlRepo(dbSess.MySQL), conf.Env),
			DB:      db,
			Env:     conf.Env,
		}
	}
	return domain.App{}
}
