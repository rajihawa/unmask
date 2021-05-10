package usecases

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rajihawa/unmask/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecases struct {
	repo domain.UserRepo
	env  domain.Env
}

func NewUserUsecases(repo domain.UserRepo, env domain.Env) domain.UserUsecases {
	return &UserUsecases{
		repo: repo,
		env:  env,
	}
}

func (uu *UserUsecases) GetUser(id string) (*domain.User, error) {
	return uu.repo.GetOne(id)
}

func (uu *UserUsecases) GetUsers(limit int, offset int) ([]domain.User, error) {
	return uu.repo.GetAll(limit, offset)
}

func (uu *UserUsecases) UpdateUser(id string, newUser domain.User) error {
	return uu.repo.UpdateOne(id, newUser)
}

func (uu *UserUsecases) CreateUser(projectID string, client domain.Client, newUser domain.UserSignup) (string, error) {
	id := uuid.New().String()
	if newUser.Password != newUser.PasswordConfirm {
		return "", errors.New("passwords don't match")
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &domain.User{
		ID:           id,
		Username:     newUser.Email,
		Email:        newUser.Email,
		PasswordHash: fmt.Sprint(hashBytes),
		Attributes:   newUser.Attributes,
		Verified:     client.AutoVerify,
		ProjectID:    projectID,
		ClientID:     client.ID,
	}

	if err := uu.repo.CreateOne(*user); err != nil {
		return "", err
	}
	return id, nil
}

func (uu *UserUsecases) DeleteUser(id string) error {
	return uu.repo.DeleteOne(id)
}
