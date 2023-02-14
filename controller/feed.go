package controller

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// token := c.Query("token")
	latest_time := c.Query("latest_time")
	var latest_Time sql.NullTime
	latest_Time.Valid = true
	if latest_time == "" {
		latest_Time.Time = time.Now()
	} else {
		temp_time, _ := strconv.ParseInt(latest_time, 10, 64)
		latest_Time.Time = time.Unix(temp_time, 0)
	}
	//TODO鉴权外加判断关注和喜爱
	var timeLayoutStr = "2006-01-02 15:04:05"
	query := GetConn()
	videosall, _ := query.ListVideos(context.Background(), latest_Time.Time.Format(timeLayoutStr))
	video_list := []Video{}
	var nexttime time.Time
	nexttime = time.Now()
	for _, video := range videosall {
		user_id := video.Author
		user, _ := query.GetUserById(context.Background(), user_id)
		author := User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount.Int64,
			FollowerCount: user.FollowerCount.Int64,
			IsFollow:      true,
		}
		var controller_video Video
		controller_video.Id = video.VideoID
		controller_video.Author = author
		controller_video.PlayUrl = video.PlayUrl
		controller_video.CoverUrl = video.CoverUrl
		controller_video.FavoriteCount = video.FavoriteCount.Int64
		controller_video.CommentCount = video.CommentCount.Int64
		controller_video.IsFavorite = true
		controller_video.Title = video.Title
		video_list = append(video_list, controller_video)
		nexttime = video.CreatedAt.Time
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: video_list,
		NextTime:  nexttime.Unix(),
	})
}
