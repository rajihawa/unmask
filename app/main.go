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
		db.Init()

		return domain.App{
			Project: usecases.NewProjectUsecases(repository.NewProjectMySqlRepo(), conf.Env),
			Client:  usecases.NewClientUsecases(repository.NewClientMySqlRepo(), conf.Env),
			User:    usecases.NewUserUsecases(repository.NewUserMySqlRepo(), conf.Env),
			DB:      db,
			Env:     conf.Env,
		}
	}
	return domain.App{}
}
