package user_test

import (
	"log"
	"testing"

	"github.com/rajihawa/unmask/app"
	"github.com/rajihawa/unmask/tests"
)

func TestUserUsecases(t *testing.T) {
	app := app.InitApp(tests.AppConfig)
	defer app.Close()

	log.Println("Testing CreateProject")
	projectID, err := app.Project.CreateProject(*tests.NewProject)
	if err != nil {
		t.Errorf("Failed to create project %+v", err)
	}

	log.Println("Testing CreateClient")
	newId, err := app.Client.CreateClient(projectID, *tests.NewClient)
	if err != nil {
		t.Errorf("Failed to create client %+v", err)
	}

	log.Println("Testing GetClient")
	fetchedClient, err := app.Client.GetClient(newId)
	if err != nil {
		t.Errorf("Failed to get client %s, %+v", newId, err)
	}

	log.Println("Testing CreateUser")
	newId, err = app.User.CreateUser(projectID, *fetchedClient, *tests.NewUser)
	if err != nil {
		t.Errorf("Failed to create user %+v", err)
	}

	log.Println("Testing GetUsers")
	users, err := app.User.GetUsers(10, 0)
	if err != nil {
		t.Errorf("Failed to get users %+v", err)
	}

	usersCount := len(users)
	if usersCount != 1 {
		t.Errorf("Users length expected %d but got %d", 1, usersCount)
	}

	createdUser := users[0]

	if createdUser.ID != newId {
		t.Errorf("User ID expected %s but got %s", newId, createdUser.ID)
	}

	newUser := createdUser
	newUser.Username = "test2"

	log.Println("Testing UpdateProject")
	err = app.User.UpdateUser(createdUser.ID, newUser)
	if err != nil {
		t.Errorf("Failed to update user %+v", err)
	}

	log.Println("Testing GetUser")
	fetchedUser, err := app.User.GetUser(createdUser.ID)
	if err != nil {
		t.Errorf("Failed to get user %s, %+v", createdUser.ID, err)
	}
	if fetchedUser.Username != newUser.Username {
		t.Errorf("User name expected %s but got %s", newUser.Username, fetchedUser.Username)
	}

	log.Println("Testing DeleteUser")
	err = app.User.DeleteUser(fetchedUser.ID)
	if err != nil {
		t.Errorf("Failed to delete user %s, %+v", fetchedUser.ID, err)
	}
}
