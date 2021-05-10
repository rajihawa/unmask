package tests

import "github.com/rajihawa/unmask/app/domain"

var AppConfig = domain.AppConfig{
	DB: domain.DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Database: "db",
		Port:     "3306",
		Username: "user",
		Password: "password",
	},
	Env: domain.Env{
		Stage:  "testing",
		Domain: "unmask.test.com",
	},
}

var NewProject = &domain.Project{
	Name:        "test project",
	Description: "test description of project",
}

var NewClient = &domain.Client{
	Name:        "test client",
	Description: "test description of client",
	HomeURL:     "http://test.test",
	CallbackURL: "http://test.test/auth",
	Privileges:  "all",
	AutoVerify:  true,
}
var NewUser = &domain.UserSignup{
	Username:        "test",
	Email:           "test@test.com",
	Password:        "test123",
	PasswordConfirm: "test123",
	Attributes: map[string]interface{}{
		"age": 13,
	},
}
