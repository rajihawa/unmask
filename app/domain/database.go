package domain

import "database/sql"

type DatabaseConfig struct {
	Driver   string
	Host     string
	Database string
	Port     string
	Username string
	Password string
}

type DatabaseSessions struct {
	MySQL *sql.DB
}

type Database interface {
	Init(sess *DatabaseSessions)
	Clear()
	Close()
}
