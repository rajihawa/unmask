package app

import (
	"github.com/rajihawa/unmask/app/repository"
	usecases "github.com/rajihawa/unmask/app/usecases"
	"github.com/rajihawa/unmask/domain"
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
