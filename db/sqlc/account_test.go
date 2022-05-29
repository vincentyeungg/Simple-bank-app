package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vincentyeungg/Simple-bank-app/util"
)

func createRandomAccount(t *testing.T) Account {
	// test arguments
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	// using package testify to handle error nil checking, and stops tests when error detected
	require.NoError(t, err)

	// verify the returned account is not an empty object
	require.NotEmpty(t, account)

	// verify the newly created rcord has fields matching the argument
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	// ensure the account has an id field generated from postgres
	require.NotZero(t, account.ID)
	// ensure the account record also has a generated time stamp
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// to get an account, create an account first
	account1 := createRandomAccount(t)

	// attempt to retrieve the created account
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	// ensure an account is retrieved
	require.NotEmpty(t, account2)

	// assert the fields are equal
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// assert the time of creation are the same (within 1 second for testing purposes)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}

	// attempt to update new account
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	// try querying for an account with the original account1.Id, should not be there
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// should return an error if the row doesn't exist
	require.Error(t, err)
	// ensure the error is no rows selected
	require.EqualError(t, err, sql.ErrNoRows.Error())
	// ensure account2 obj is empty
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	// generate a list of random accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	// since offset and limit 5, there should still be 5 accounts
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	// ensure 5 accounts selected
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}