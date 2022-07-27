package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	var password string = RandomString(6)

	var hashedPassword1 string
	var err error

	hashedPassword1, err = HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	// var wrongPassword string
	// err = CheckPassword(password, wrongPassword)
	// require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	var hashedPassword2 string

	hashedPassword2, err = HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
