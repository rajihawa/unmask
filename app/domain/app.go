package domain

import "log"

type AppConfig struct {
	DB  DatabaseConfig
	Env Env
}

type Env struct {
	Stage  string
	Domain string
}

type App struct {
	Project ProjectUsecases
	Client  ClientUsecases
	User    UserUsecases
	Env     Env
	DB      Database
}

func (a *App) Close() {
	if a.Env.Stage == "testing" {
		log.Println("Clearing db...")
		a.DB.Clear()
	}
	a.DB.Close()
}
