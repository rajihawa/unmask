package utils

import (
	"log"
	"os"
	"strings"
	"time"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	databaseUrl       = os.Getenv("DATABASE_URL")
	databaseName      = "unmask"
	databaseUsername  = os.Getenv("DATABASE_USERNAME")
	databasePassword  = os.Getenv("DATABASE_PASSWORD")
	Session           *r.Session
	ProjectsTableName = "projects"
	tables            = map[string](r.TableCreateOpts){
		ProjectsTableName: r.TableCreateOpts{
			PrimaryKey: "id",
		},
	}
)

func InitDatabase() {
	session, err := r.Connect(r.ConnectOpts{
		Address:      databaseUrl,
		Database:     databaseName,
		Username:     databaseUsername,
		Password:     databasePassword,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
		Timeout:      3 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for tableName, tableOpts := range tables {
		err = r.TableCreate(tableName, tableOpts).Exec(session)
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			log.Fatalln(err)
		}
	}

	Session = session

}
