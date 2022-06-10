package db

import (
	"context"
	"testing"
	"time"

	"github.com/mehdieidi/bank/pkg/password"
	"github.com/mehdieidi/bank/pkg/rand"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, arg, err := createRandomUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
}

func TestGetUser(t *testing.T) {
	user1, _, _ := createRandomUser()
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func createRandomUser() (User, CreateUserParams, error) {
	hashedPassword, _ := password.Hash(rand.RandomString(6))

	arg := CreateUserParams{
		Username:       rand.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       rand.RandomOwner(),
		Email:          rand.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	return user, arg, err
}
