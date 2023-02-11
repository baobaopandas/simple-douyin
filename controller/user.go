package controller

import (
	"context"
	"net/http"
	"strconv"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var queries = GetConn()
	//验证用户名是否存在
	_, err := queries.GetUser(context.Background(), username)
	if err != nil {
		arg := mydb.CreateUserParams{
			Name:     username,
			Password: password,
		}
		_, err := queries.CreateUser(context.Background(), arg)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Create user error"},
			})
		} else {
			user, err := queries.GetUser(context.Background(), username)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "Get user_id error"},
				})
				return
			}
			//生成token
			token, err := util.GenerateToken(&user)
			if err != nil {
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 1, StatusMsg: "generate token error"},
				})
			} else {
				id := user.UserID
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{StatusCode: 0},
					UserId:   id,
					Token:    token,
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})

	}
}

func Login(c *gin.Context) {
	var queries = GetConn()
	username := c.Query("username")
	password := c.Query("password")
	user, err := queries.GetUser(context.Background(), username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else if user.Password != password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "password error"},
		})
	} else {
		//重新生成token
		token, err := util.GenerateToken(&user)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "generate token error"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserID,
			Token:    token,
		})
	}

}

// TODO：关注以及存缓存
func UserInfo(c *gin.Context) {
	var queries = GetConn()
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//先查询缓存中的user信息
	//失败回库

	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	// token := c.Query("token")

	user, err := queries.GetUserById(context.Background(), id)
	//ToDO: 是否关注后面再做
	var User_return = User{Id: user.UserID, Name: user.Name, FollowCount: user.FollowCount.Int64, FollowerCount: user.FollowerCount.Int64, IsFollow: true}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     User_return,
	})
	//再存入缓存

}
