package mydb

import (
	"context"
	"fmt"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComment(t *testing.T) {
	arg := mydb.CreateCommentParams{
		UserID:  1,
		VideoID: 1,
		Content: "test comment",
	}

	comment, err := testQueries.CreateComment(context.Background(), arg)

	require.NoError(t, err)

	r, err := comment.RowsAffected()

	require.NoError(t, err)

	fmt.Printf("rows: %d\n", r)

}
