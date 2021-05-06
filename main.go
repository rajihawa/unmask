package main

import (
	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/domain"
)

var appConfig = app.AppConfig{DB: domain.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost",
	Database: "db",
	Port:     "3306",
	Username: "user",
	Password: "password",
},
	Env: "development"}

func main() {
	app := app.InitApp(appConfig)
	defer app.Close()
}
