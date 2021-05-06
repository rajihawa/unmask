package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rajihawa/unmask/app/domain"
)

var (
	MySQL *sql.DB
)

type MySqlDB struct {
	Config    domain.DatabaseConfig
	Migration *migrate.Migrate
}

func NewMySqlDB(conf domain.DatabaseConfig) *MySqlDB {
	return &MySqlDB{
		Config: conf,
	}
}

func (mdb *MySqlDB) Init() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mdb.Config.Username, mdb.Config.Password, mdb.Config.Host, mdb.Config.Port, mdb.Config.Database)
	db, err := sql.Open(mdb.Config.Driver, url)
	if err != nil {
		log.Println("Error while opening connection with mysql database.")
		panic(err)
	}
	MySQL = db
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
	migrationsPath := pwd
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s/migrations", migrationsPath), mdb.Config.Driver, driver)
	if err != nil {
		log.Println("Error while getting migration instance.")
		panic(err)
	}
	mdb.Migration = m
	err = m.Up()
	if err != nil && !strings.Contains(err.Error(), "no change") {
		log.Println("Error while migrating up.")
		panic(err)
	}
}

func (mdb *MySqlDB) Clear() {
	err := mdb.Migration.Down()
	if err != nil {
		log.Println("Error while migrating up.")
		panic(err)
	}
}
func (mdb *MySqlDB) Close() {
	// mdb.DB.Close()
}
