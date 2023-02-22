package controller

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/RaymondCode/simple-demo/config"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	//校验token令牌
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//获取上传者信息
	user_id := claim.UserId
	title := c.PostForm("title")

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user_id, filename)
	saveFile := filepath.Join("./public/video/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	covername := strings.TrimRight(finalName, ".mp4")

	covername = fmt.Sprintf("%s.jpg", covername)
	saveCover := filepath.Join("./public/cover/", covername)

	err = util.GetFrame(saveFile, saveCover)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	play_url := fmt.Sprintf("http://%s:%s/video/%s", config.CONFIG.VideoConfig.Host, config.CONFIG.VideoConfig.Port, finalName)
	cover_url := fmt.Sprintf("http://%s:%s/cover/%s", config.CONFIG.VideoConfig.Host, config.CONFIG.VideoConfig.Port, covername)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var queries = GetConn()
	args := mydb.CreateVideoParams{
		Author:   user_id,
		PlayUrl:  play_url,
		CoverUrl: cover_url,
		Title:    title,
	}
	_, err = queries.CreateVideo(context.Background(), args)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "create video error",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	//验证token
	_, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// self_id := claim.UserId
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	var video_list = []Video{}
	query := GetConn()
	user, _ := query.GetUserById(context.Background(), user_id)

	//判断是否自己是否已关注
	is_follow := IsFollowUser(user_id, user.UserID)

	author := User{
		Id:            user.UserID,
		Name:          user.Name,
		FollowCount:   user.FollowCount.Int64,
		FollowerCount: user.FollowerCount.Int64,
		IsFollow:      is_follow,
	}
	//TODO判断是否点赞该视频
	videos, _ := query.GetVideoById(context.Background(), user_id)
	for _, video := range videos {
		var controller_video Video
		controller_video.Id = video.VideoID
		controller_video.Author = author
		controller_video.PlayUrl = video.PlayUrl
		controller_video.CoverUrl = video.CoverUrl
		controller_video.FavoriteCount = video.FavoriteCount.Int64
		controller_video.CommentCount = video.CommentCount.Int64
		controller_video.IsFavorite = IsFavoriteVideo(user_id, video.VideoID)
		controller_video.Title = video.Title
		video_list = append(video_list, controller_video)

	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video_list,
	})
}
