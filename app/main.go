package app

import (
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/repository"
	usecases "github.com/rajihawa/unmask/app/usecases"
)

type AppConfig struct {
}

type App struct {
	Project domain.ProjectUsecases
}

func InitApp(conf AppConfig) App {
	return App{
		Project: usecases.NewProjectUsecases(repository.NewProjectRepo()),
	}
}
