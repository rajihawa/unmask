package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User - the User struct will hold all the user info..
type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser - create a new user
func NewUser(username, password string) (*User, error) {
	u := &User{
		ID:        GenerateRandomID(),
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	pwdHash := hashPassword(password)
	u.Password = pwdHash
	if err := u.validatePassword(pwdHash); err != nil {
		panic("Password hashing not working correctly.")
	}
	return u, nil
}

// private functions
func hashPassword(raw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(raw), 10)
	return string(hash)
}

func (u *User) validatePassword(password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return
}
