package app

import (
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/repository"
	usecases "github.com/rajihawa/unmask/app/usecases"
)

type AppConfig struct {
}

type App struct {
	project domain.ProjectUsecases
}

func StartApp(conf AppConfig) App {
	return App{
		project: usecases.NewProjectUsecases(repository.NewProjectRepo()),
	}
}
