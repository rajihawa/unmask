package main

import (
	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/app/domain"
)

var appConfig = app.AppConfig{DB: domain.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost",
	Database: "db",
	Port:     "3306",
	Username: "user",
	Password: "password",
},
	Env: "development"}

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

func main() {
	// s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	// v := reflect.ValueOf(&s).Elem()
	// typeOfS := v.Type()

	// for i := 0; i < v.NumField(); i++ {

	// 	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	// }
	// return
	app := app.InitApp(appConfig)
	defer app.Close()
	client := domain.Client{
		Name:        "test client",
		Description: "test description of client",
		HomeURL:     "http://test.test",
		CallbackURL: "http://test.test/auth",
		Privileges:  "all",
		AutoVerify:  true,
	}
	app.Client.CreateClient("fdgds", client)
}
