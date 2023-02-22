package mydb

import (
	"context"
	"testing"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := mydb.CreateUserParams{
		Name:     "testuser",
		Password: "123456",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

}

func TestGetUser(t *testing.T) {
	testname := "baobao"
	testpassword := "123456"
	user, err := testQueries.GetUser(context.Background(), testname)

	require.NoError(t, err)
	require.Equal(t, user.Name, testname)
	require.Equal(t, user.Password, testpassword)
}
