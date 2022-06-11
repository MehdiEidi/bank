package password

import (
	"testing"

	"github.com/mehdieidi/bank/pkg/rand"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := rand.RandomString(6)

	hashedPassword1, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = Check(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := rand.RandomString(6)
	err = Check(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
