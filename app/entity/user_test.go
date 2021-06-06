package entity_test

import (
	"testing"

	"github.com/rajihawa/mask-off/app/entity"
	"github.com/stretchr/testify/assert"
)

var username = "username1"
var password = "password1"

func TestNewUser(t *testing.T) {
	u := entity.NewUser(username, password)
	assert.Equal(t, u.Username, username)
	assert.NotNil(t, u.ID)
	assert.NotEqual(t, u.Password, password)
}
