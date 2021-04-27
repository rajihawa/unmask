package main

import (
	"os"

	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/database"
)

var (
	// PORT - port of the server
	PORT             = os.Getenv("PORT")
	databaseUrl      = os.Getenv("DATABASE_URL")
	databaseName     = os.Getenv("DATABASE_NAME")
	databaseUsername = os.Getenv("DATABASE_USERNAME")
	databasePassword = os.Getenv("DATABASE_PASSWORD")
)

func main() {
	app.StartApp(app.AppConfig{
		Port: PORT,
		RethinkDB: database.RethinkConfig{
			DatabaseURL:      databaseUrl,
			DatabaseName:     databaseName,
			DatabaseUsername: databaseUsername,
			DatabasePassword: databasePassword,
		},
	})
}
