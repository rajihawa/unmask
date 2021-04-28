package domain

type DatabaseConfig struct {
	Driver   string
	Host     string
	Database string
	Port     string
	Username string
	Password string
}

type Database interface {
	Init()
	Clear()
	Close()
}
