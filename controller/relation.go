package controller

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//校验token令牌
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//获取关注者follower_id，被关注者followed_id
	followerId := claim.UserId
	followedId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "strconv to_user_id error. " + err.Error(),
		})
		return
	}
	// 禁止自己关注自己
	if followedId == followerId {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "cannot follow yourself.",
		})
		return
	}
	//获取action动作 1-关注，2-取消关注
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "strconv action_type error. " + err.Error(),
		})
		return
	}

	query := GetConn()
	arg := mydb.GetRelationByIDParams{
		FollowerID: followerId,
		FollowedID: followedId,
	}
	//检查曾经是否已经关注过
	isDeleted, err := query.GetRelationByID(context.Background(), arg)
	if err == nil { //曾经关注过
		var relationInfo mydb.UpdateRelationParams
		switch {
		// 关注，曾经取关
		case actionType == 1 && isDeleted == 1:
			relationInfo = mydb.UpdateRelationParams{
				Deleted:    0,
				FollowerID: followerId,
				FollowedID: followedId,
			}
		// 取消关注，曾经关注
		case actionType == 2 && isDeleted == 0:
			relationInfo = mydb.UpdateRelationParams{
				Deleted:    1,
				FollowerID: followerId,
				FollowedID: followedId,
			}
		default:
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  "invalid operation",
			})
			return
		}
		if err = query.UpdateRelation(context.Background(), relationInfo); err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else { //没关注过
		if actionType != 1 {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  "invalid operation",
			})
			return
		}
		relationInfo := mydb.CreateRelationParams{
			FollowerID: followerId,
			FollowedID: followedId,
		}
		if err = query.CreateRelation(context.Background(), relationInfo); err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "OK",
	})
	go UpdateRelationCount(followerId, followedId, query)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	//验证token
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	// 获取查询者id
	userID := claim.UserId
	query := GetConn()
	_, FollowedList, err := GetIdAndList(query.GetFollowedIdByFollower, c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	// if len(FollowedList) == 0 {
	// 	c.JSON(http.StatusOK, UserListResponse{
	// 		Response: Response{
	// 			StatusCode: 0,
	// 		},
	// 		UserList: []User{},
	// 	})
	// 	return
	// }

	user_list := make([]User, len(FollowedList))
	for i, followedId := range FollowedList {
		user, err := query.GetUserById(context.Background(), followedId)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		user_list[i] = User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount.Int64,
			FollowerCount: user.FollowerCount.Int64,
			IsFollow:      IsFollowUser(userID, followedId),
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user_list,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	//验证token
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	// 获取查询者id
	userID := claim.UserId
	query := GetConn()
	_, FollowerList, err := GetIdAndList(query.GetFollowerIdByFollowed, c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	user_list := make([]User, len(FollowerList))
	for i, followerId := range FollowerList {
		user, err := query.GetUserById(context.Background(), followerId)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		user_list[i] = User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount.Int64,
			FollowerCount: user.FollowerCount.Int64,
			IsFollow:      IsFollowUser(userID, followerId),
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user_list,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	//验证token
	_, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	query := GetConn()
	followerId, FollowedList, err := GetIdAndList(query.GetFollowedIdByFollower, c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	user_list := []User{}
	for _, followedId := range FollowedList {
		if is_follow := IsFollowUser(followedId, followerId); !is_follow {
			// followedId未关注followerId， 双方不是好友
			continue
		}
		// 是好友，加入列表
		user, err := query.GetUserById(context.Background(), followedId)
		if err != nil {
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		user_list = append(user_list, User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount.Int64,
			FollowerCount: user.FollowerCount.Int64,
			IsFollow:      true,
		})
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user_list,
	})
}

// 判断followerId 是否关注followedId
func IsFollowUser(followerID int64, followedID int64) bool {
	query := GetConn()
	arg := mydb.GetRelationByIDParams{
		FollowerID: followerID,
		FollowedID: followedID,
	}
	//检查曾经是否已经关注过
	isDeleted, err := query.GetRelationByID(context.Background(), arg)
	if err == nil {
		//曾经关注过
		return isDeleted == 0
	}
	//没关注过
	return false
}

// 给定字符串形式id和方法，获取id和列表
func GetIdAndList(queryFunc func(ctx context.Context, ID int64) ([]int64, error), value string) (int64, []int64, error) {
	Id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1, nil, err
	}

	List, err := queryFunc(context.Background(), Id)
	if err != nil {
		return Id, nil, err
	}
	return Id, List, nil
}

// 更新followerID的关注数，更新followedID的粉丝数
func UpdateRelationCount(followerID int64, followedID int64, query *mydb.Queries) {
	follow_count, _ := query.GetFollowedCount(context.Background(), followerID)
	argFollowCount := mydb.UpdateFollowCountParams{
		FollowCount: sql.NullInt64{
			Int64: follow_count,
			Valid: true,
		},
		UserID: followerID,
	}
	query.UpdateFollowCount(context.Background(), argFollowCount)

	follower_count, _ := query.GetFollowerCount(context.Background(), followedID)
	argFollowerCount := mydb.UpdateFollowerCountParams{
		FollowerCount: sql.NullInt64{
			Int64: follower_count,
			Valid: true,
		},
		UserID: followedID,
	}
	query.UpdateFollowerCount(context.Background(), argFollowerCount)
}
