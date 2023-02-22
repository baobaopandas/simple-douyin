package mydb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComment(t *testing.T) {
	//arg := mydb.CreateCommentParams{
	//	UserID:  1,
	//	VideoID: 1,
	//	Content: "test comment",
	//}

	count, err := testQueries.TotalComment(context.Background(), 1)

	//atomic.AddInt64(&id, 1)
	//arg := mydb.CreateCommentParams{
	//	CommentID: id,
	//	UserID:    1,
	//	VideoID:   1,
	//	Content:   "comment get",
	//}
	//
	//_, err := testQueries.CreateComment(context.Background(), arg)

	require.NoError(t, err)

	fmt.Printf("count: %d\n", count)

}
