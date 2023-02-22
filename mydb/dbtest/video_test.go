package mydb

import (
	"context"
	"fmt"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetVideo(t *testing.T) {
	var test_id int64
	test_id = 1
	video, err := testQueries.GetVideoById(context.Background(), test_id)

	require.NoError(t, err)
	require.Equal(t, "http://81.68.118.43:8088/video/1_VIDEO_20230213_183226847.mp4", video[0].PlayUrl)
	// require.Equal(t, user.Password, testpassword)
}

func TestCreateVideo(t *testing.T) {
	arg := mydb.CreateVideoParams{
		Author:   1,
		PlayUrl:  "localhost:8888/play",
		CoverUrl: "localhost:8888/cover",
		Title:    "test",
	}

	video, err := testQueries.CreateVideo(context.Background(), arg)

	require.NoError(t, err)

	r, err := video.LastInsertId()

	require.NoError(t, err)

	fmt.Printf("id: %d\n", r)
}
