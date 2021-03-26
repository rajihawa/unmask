package rethink

import (
	"log"
	"strings"
	"time"

	"github.com/rajihawa/unmask/core/config"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
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
		Address:      config.DatabaseUrl,
		Database:     config.DatabaseName,
		Username:     config.DatabaseUsername,
		Password:     config.DatabasePassword,
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
