package usecases

import (
	"errors"

	"github.com/rajihawa/unmask/domain"
	"golang.org/x/crypto/bcrypt"
)

type UsersUsecases struct {
	repo domain.UsersRepository
}

func NewUsersUsecase(repo domain.UsersRepository) domain.UsersUsecases {
	return &UsersUsecases{
		repo: repo,
	}
}

func (uu *UsersUsecases) GetAll(client domain.Client, opts domain.GetUsersOpts) ([]domain.User, error) {
	return uu.repo.GetAll(client.Project.ID, opts)
}

func (uu *UsersUsecases) GetUser(id string, opts domain.GetUsersOpts) (*domain.User, error) {
	return uu.repo.Get(id, opts)
}

func (uu *UsersUsecases) CheckUserLogin(userLogin domain.UserLogin, client domain.Client) (*domain.User, error) {

	users, err := uu.repo.GetByUsername(userLogin.Username, client.Project.ID, domain.GetUsersOpts{GetProject: false})
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, errors.New("user does not exist")
	}
	user := users[0]
	if user.Disabled {
		return nil, errors.New("user is disabled")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLogin.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return &user, nil
}

func (uu *UsersUsecases) SignupUser(client domain.Client, user *domain.User) error {
	if !client.Signup {
		return errors.New("client can't register users")
	}

	if user.Password != user.PasswordConfirm {
		return errors.New("passwords don't match")
	}

	users, err := uu.repo.GetByUsername(user.Username, client.Project.ID, domain.GetUsersOpts{GetProject: false})
	if err != nil {
		return err
	}
	if len(users) > 1 {
		return errors.New("user already exist")
	}

	user.Confirmed = !client.Project.VerifyEmail

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordConfirm), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)

	user.Project = client.Project
	user.Client = &client

	return uu.repo.Insert(*user)
}
