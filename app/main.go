package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/rajihawa/unmask/app/domain"
	"github.com/rajihawa/unmask/app/repository"
	usecases "github.com/rajihawa/unmask/app/usecases"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Database string
	Port     string
	Username string
	Password string
}

type AppConfig struct {
	DB DatabaseConfig
}

type App struct {
	Project domain.ProjectUsecases
}

func InitApp(conf AppConfig) App {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.DB.Username, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Database)
	db, err := sql.Open(conf.DB.Driver, url)
	if err != nil {
		log.Println("Error while opening connection with mysql database.")
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error while getting driver instance.")
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file///migrations", conf.DB.Driver, driver)
	if err != nil {
		log.Println("Error while getting migration instance.")
		panic(err)
	}
	err = m.Up()
	if err != nil {
		log.Println("Error while migrating up.")
		panic(err)
	}
	return App{
		Project: usecases.NewProjectUsecases(repository.NewProjectRepo()),
	}
}
