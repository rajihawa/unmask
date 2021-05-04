package project_test

import (
	"log"
	"testing"

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
	Env: "testing"}

func TestClientUsecases(t *testing.T) {
	app := app.InitApp(appConfig)
	defer app.Close()

	project := &domain.Project{
		Name:        "test project",
		Description: "test description of project",
	}
	log.Println("Testing CreateProject")
	projectID, err := app.Project.CreateProject(*project)
	if err != nil {
		t.Errorf("Failed to create project %+v", err)
	}

	client := &domain.Client{
		Name:        "test client",
		Description: "test description of client",
		HomeURL:     "http://test.test",
		CallbackURL: "http://test.test/auth",
		Privileges:  "all",
		AutoVerify:  true,
	}

	log.Println("Testing CreateClient")
	newId, err := app.Client.CreateClient(projectID, *client)
	if err != nil {
		t.Errorf("Failed to create client %+v", err)
	}

	log.Println("Testing GetClients")
	clients, err := app.Client.GetClients(0, 0)
	if err != nil {
		t.Errorf("Failed to get clients %+v", err)
	}
	clientsCount := len(clients)
	if clientsCount != 1 {
		t.Errorf("Clients length expected %d but got %d", 1, clientsCount)
	}

	createdClient := clients[0]

	if createdClient.ID != newId {
		t.Errorf("Client ID expected %s but got %s", newId, createdClient.ID)
	}

	newClient := createdClient
	newClient.Name = "updated test project"

	log.Println("Testing UpdateProject")
	err = app.Client.UpdateClient(createdClient.ID, newClient)
	if err != nil {
		t.Errorf("Failed to update client %+v", err)
	}

	log.Println("Testing GetClient")
	fetchedClient, err := app.Client.GetClient(createdClient.ID)
	if err != nil {
		t.Errorf("Failed to get client %s, %+v", createdClient.ID, err)
	}
	if fetchedClient.Name != newClient.Name {
		t.Errorf("Client name expected %s but got %s", newClient.Name, fetchedClient.Name)
	}

	log.Println("Testing DeleteClient")
	err = app.Client.DeleteClient(fetchedClient.ID)
	if err != nil {
		t.Errorf("Failed to delete client %s, %+v", fetchedClient.ID, err)
	}
}
