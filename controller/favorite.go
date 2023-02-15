package controller

import (
	"context"
	"net/http"
	"strconv"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	action_type := c.Query("action_type")
	video := c.Query("video_id")
	// 校验token令牌
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	video_id, err := strconv.ParseInt(video, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get an illegal videoId"},
		})
	}

	user_id := claim.UserId
	var queries = GetConn()
	if action_type == "1" {
		arg := mydb.GetInfoParams{
			UserID: user_id,
			VideoID: video_id,
		}
		Info, err := queries.GetInfo(context.Background(), arg)
		if err != nil {
			arg := mydb.AddFavoriteParams {
				UserID: user_id,
				VideoID: video_id,
			}
			err := queries.AddFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 3, StatusMsg: "Add Favorite Failed"},
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Update Successed"},
				})
			}
		} else {
			arg := mydb.UpdateFavoriteParams {
				UserID: Info.UserID,
				VideoID: Info.VideoID,
			}
			err := queries.UpdateFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 2, StatusMsg: "Update Failed"},
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Update Successed"},
				})
			}
		}
	} else {
		arg := mydb.GetInfoParams {
			UserID: user_id,
			VideoID: video_id,
		}
		Info, err := queries.GetInfo(context.Background(), arg)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 4, StatusMsg: "No relevant record"},
			})
		} else {
			arg := mydb.DeleteFavoriteParams {
				UserID: Info.UserID,
				VideoID: Info.VideoID,
			}
			err := queries.DeleteFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 2, StatusMsg: "Update Failed"},
				})
			} else {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Update Successed"},
				})
			}
		}

	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	User_id, err := strconv.ParseInt(c.Query("action_type") , 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get an illegal videoId"},
		})
	}
	// 校验token令牌
	_, err = util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var queries = GetConn()
	VideoList, err := queries.GetUserLike(context.Background(), User_id)

	type UserINFO struct {
		id int64
		name string
		follow_count int64
		follower_count int64
		is_follow bool
	}

	type VideoLike struct {
		VideoID int64
		Author UserINFO
		play_url string
		cover_url string
		favorite_count int64
		comment_count int64
		is_favorite bool
		title string
	}
	// TODO is follow
	for _, video := range VideoList {
		user_info, err = queries.GetUserById(context.Background(), User_id)
		video_info, err = queries.GetVideo(context.Background(), video.VideoID)
	}
}
