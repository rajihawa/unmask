package app

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rajihawa/unmask/app/data"
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/repository"
	usecases "github.com/rajihawa/unmask/app/usecases"
)

type AppConfig struct {
	DB  domain.DatabaseConfig
	Env string
}

type App struct {
	Project domain.ProjectUsecases
	Env     string
	DB      domain.Database
}

func InitApp(conf AppConfig) App {
	var db domain.Database
	if conf.DB.Driver == "mysql" {
		db = data.NewMySqlDB(conf.DB)
		db.Init()

		return App{
			Project: usecases.NewProjectUsecases(repository.NewProjectMySqlRepo()),
			Env:     conf.Env,
			DB:      db,
		}
	}
	return App{}
}

func (a *App) Close() {
	if a.Env == "testing" {
		a.DB.Clear()
	}
	a.DB.Close()
}
