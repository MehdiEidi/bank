package db

import (
	"context"
	"testing"
	"time"

	"github.com/mehdieidi/bank/pkg/rand"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(account1, account2 Account) (Transfer, CreateTransferParams, error) {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        rand.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	return transfer, arg, err
}

func TestCreateTransfer(t *testing.T) {
	account1, _, _ := createRandomAccount()
	account2, _, _ := createRandomAccount()

	transfer, arg, err := createRandomTransfer(account1, account2)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	account1, _, _ := createRandomAccount()
	account2, _, _ := createRandomAccount()

	transfer1, _, _ := createRandomTransfer(account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1, _, _ := createRandomAccount()
	account2, _, _ := createRandomAccount()

	for i := 0; i < 5; i++ {
		createRandomTransfer(account1, account2)
		createRandomTransfer(account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
