package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

// type UserINFO struct {
// 	id int64
// 	name string
// 	follow_count int64
// 	follower_count int64
// 	is_follow bool
// }

// type VideoLike struct {
// 	VideoID int64
// 	Author UserINFO
// 	play_url string
// 	cover_url string
// 	favorite_count int64
// 	comment_count int64
// 	is_favorite bool
// 	title string
// }

type UserFavoriteResponse struct {
	Response
}

type UserFavoriteListResponse struct {
	Response
	Videolist []Video
}

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
		c.JSON(http.StatusOK, UserFavoriteResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get an illegal videoId"},
		})
	}

	user_id := claim.UserId
	var queries = GetConn()
	if action_type == "1" {
		arg := mydb.GetInfoParams{
			UserID:  user_id,
			VideoID: video_id,
		}
		Info, err := queries.GetInfo(context.Background(), arg)
		if err != nil {
			arg := mydb.AddFavoriteParams{
				UserID:  user_id,
				VideoID: video_id,
			}
			err := queries.AddFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserFavoriteResponse{
					Response: Response{StatusCode: 3, StatusMsg: "Add Favorite Failed"},
				})
			}
		} else {
			arg := mydb.UpdateFavoriteParams{
				UserID:  Info.UserID,
				VideoID: Info.VideoID,
			}
			err := queries.UpdateFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserFavoriteResponse{
					Response: Response{StatusCode: 2, StatusMsg: "Update Failed"},
				})
			}
		}
		err = queries.AddFavoriteCount(context.Background(), video_id)
		if err != nil {
			c.JSON(http.StatusOK, UserFavoriteResponse{
				Response: Response{StatusCode: 3, StatusMsg: "Update Favorite Count Failed"},
			})
		}
		c.JSON(http.StatusOK, UserFavoriteResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Update Successed"},
		})
	} else {
		arg := mydb.GetInfoParams{
			UserID:  user_id,
			VideoID: video_id,
		}
		Info, err := queries.GetInfo(context.Background(), arg)
		if err != nil {
			c.JSON(http.StatusOK, UserFavoriteResponse{
				Response: Response{StatusCode: 4, StatusMsg: "No relevant record"},
			})
		} else {
			arg := mydb.DeleteFavoriteParams{
				UserID:  Info.UserID,
				VideoID: Info.VideoID,
			}
			err := queries.DeleteFavorite(context.Background(), arg)
			if err != nil {
				c.JSON(http.StatusOK, UserFavoriteResponse{
					Response: Response{StatusCode: 2, StatusMsg: "Update Failed"},
				})
			} else {
				err = queries.DelFavoriteCount(context.Background(), video_id)
				if err != nil {
					c.JSON(http.StatusOK, UserFavoriteResponse{
						Response: Response{StatusCode: 3, StatusMsg: "Update Favorite Count Failed"},
					})
				}
				c.JSON(http.StatusOK, UserFavoriteResponse{
					Response: Response{StatusCode: 0, StatusMsg: "Update Successed"},
				})
			}
		}

	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	User_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserFavoriteResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get an illegal user_id"},
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
	if err != nil {
		c.JSON(http.StatusOK, UserFavoriteResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get an illegal user_id"},
		})
	}

	var favoriteVideo []Video
	for _, video := range VideoList {
		user_info, err := queries.GetUserById(context.Background(), User_id)
		if err != nil {
			c.JSON(http.StatusOK, UserFavoriteResponse{
				Response: Response{StatusCode: 10, StatusMsg: "Get user_info Failed"},
			})
		}
		video_info, err := queries.GetVideo(context.Background(), video.VideoID)
		if err != nil {
			c.JSON(http.StatusOK, UserFavoriteResponse{
				Response: Response{StatusCode: 10, StatusMsg: "Get video_info Failed"},
			})
		}
		INFO := Video{
			Id:            video.VideoID,
			PlayUrl:       video_info.PlayUrl,
			CoverUrl:      video_info.CoverUrl,
			FavoriteCount: video_info.FavoriteCount.Int64,
			CommentCount:  video_info.CommentCount.Int64,
			IsFavorite:    video.Statement,
			Title:         video_info.Title,
			Author: User{
				Id:            user_info.UserID,
				Name:          user_info.Name,
				FollowCount:   user_info.FollowCount.Int64,
				FollowerCount: user_info.FollowerCount.Int64,
				IsFollow:      IsFollowUser(user_info.UserID, User_id),
			},
		}
		favoriteVideo = append(favoriteVideo, INFO)
		fmt.Println(favoriteVideo)
	}
	c.JSON(http.StatusOK, UserFavoriteListResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Get user favorite list"},
		Videolist: favoriteVideo,
	})
}
