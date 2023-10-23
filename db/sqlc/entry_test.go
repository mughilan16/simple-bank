package db

import (
	"context"
	"testing"
	"time"

	"github.com/mughilan16/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, acc Account) Entry {
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, acc.ID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.Equal(t, entry1.ID, entry1.ID)
	require.Equal(t, entry1.AccountID, entry1.AccountID)
	require.Equal(t, entry1.Amount, entry1.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntires(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntiresParams{
		AccountID: account.ID,
		Offset:    5,
		Limit:     5,
	}
	entires, err := testQueries.ListEntires(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entires, 5)

	for _, entry := range entires {
		require.NotEmpty(t, entry)
	}
}
