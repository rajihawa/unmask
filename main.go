package main

import (
	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/domain"
)

var appConfig = domain.AppConfig{DB: domain.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost",
	Database: "db",
	Port:     "3306",
	Username: "user",
	Password: "password",
},
	Env: domain.Env{
		Stage:  "development",
		Domain: "unmask.local.com",
	}}

func main() {
	app := app.InitApp(appConfig)
	defer app.Close()
	client := domain.Client{
		Name:        "test client",
		Description: "test description of client",
		HomeURL:     "http://test.test",
		CallbackURL: "http://test.test/auth",
		Privileges:  "all",
		AutoVerify:  true,
	}
	app.Client.CreateClient("fdgds", client)
}
