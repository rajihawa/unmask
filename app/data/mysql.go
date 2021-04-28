package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rajihawa/unmask/app/domain"
)

type MySqlDB struct {
	config    domain.DatabaseConfig
	db        *sql.DB
	migration *migrate.Migrate
}

func NewMySqlDB(conf domain.DatabaseConfig) domain.Database {
	return &MySqlDB{
		config: conf,
	}
}

func (mdb *MySqlDB) Init() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mdb.config.Username, mdb.config.Password, mdb.config.Host, mdb.config.Port, mdb.config.Database)
	db, err := sql.Open(mdb.config.Driver, url)
	if err != nil {
		log.Println("Error while opening connection with mysql database.")
		panic(err)
	}
	mdb.db = db
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error while getting driver instance.")
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Error while getting current path")
		panic(err)
	}
	migrationsPath := path.Join(path.Dir(pwd), "../")
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/migrations", migrationsPath), mdb.config.Driver, driver)
	if err != nil {
		log.Println("Error while getting migration instance.")
		panic(err)
	}
	mdb.migration = m
	err = m.Up()
	if err != nil {
		log.Println("Error while migrating up.")
		panic(err)
	}
}

func (mdb *MySqlDB) Clear() {
	err := mdb.migration.Down()
	if err != nil {
		log.Println("Error while migrating up.")
		panic(err)
	}
}
