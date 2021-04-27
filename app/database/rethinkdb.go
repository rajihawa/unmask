package database

import (
	"log"
	"strings"
	"time"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	Session           *r.Session
	ProjectsTableName = "projects"
	ClientsTableName  = "clients"
	UsersTableName    = "users"
	tables            = map[string](r.TableCreateOpts){
		ProjectsTableName: r.TableCreateOpts{
			PrimaryKey: "id",
		},
		ClientsTableName: r.TableCreateOpts{
			PrimaryKey: "id",
		},
		UsersTableName: r.TableCreateOpts{
			PrimaryKey: "id",
		},
	}
	indexes = map[string]([]string){
		ClientsTableName: {"project_id"},
		UsersTableName:   {"project_id", "username"},
	}
)

func ResetDatabase() {
	for tableName, tableOpts := range tables {
		err := r.TableDrop(tableName, tableOpts).Exec(Session)
		if err != nil && !strings.Contains(err.Error(), "exists") {
			log.Fatalln(err)
		}
	}
}

type RethinkConfig struct {
	DatabaseURL      string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
}

func InitRethinkSession(conf RethinkConfig) {
	session, err := r.Connect(r.ConnectOpts{
		Address:      conf.DatabaseURL,
		Database:     conf.DatabaseName,
		Username:     conf.DatabaseUsername,
		Password:     conf.DatabasePassword,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
		Timeout:      3 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	Session = session
}

func InitDatabase() {

	for tableName, tableOpts := range tables {
		err := r.TableCreate(tableName, tableOpts).Exec(Session)
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			log.Fatalln(err)
		}
	}
	for tableName, indexNames := range indexes {
		for _, indexName := range indexNames {
			err := r.Table(tableName).IndexCreate(indexName).Exec(Session)
			if err != nil && !strings.Contains(err.Error(), "already exists") {
				log.Fatalln(err)
			}
		}
	}
}
