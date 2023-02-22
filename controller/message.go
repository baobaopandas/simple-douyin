package controller

import (
	"context"
	"fmt"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []mydb.Message `json:"message_list"`
}

var ctx = context.Background()

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	//将toUserId转换为数字
	toUserIdNum, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	content := c.Query("content")

	// 校验token令牌
	claim, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		})
		return
	}
	//获得userId,prev
	formUserId := claim.UserId
	//根据userId和toUserID存入数据库
	queries := GetConn()
	createMessageParams := mydb.CreateMessageParams{formUserId, toUserIdNum, content, time.Now().UnixMilli()}
	//将消息放入数据库
	_, err = queries.CreateMessage(ctx, createMessageParams)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//插入成功
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
	})

	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int64(userIdB))
	//
	//	atomic.AddInt64(&messageIdSequence, 1)
	//	curMessage := Message{
	//		Id:         messageIdSequence,
	//		Content:    content,
	//		CreateTime: time.Now().Format(time.Kitchen),
	//	}
	//
	//	if messages, exist := tempChat[chatKey]; exist {
	//		tempChat[chatKey] = append(messages, curMessage)
	//	} else {
	//		tempChat[chatKey] = []Message{curMessage}
	//	}
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	// 校验token令牌
	claim, err := util.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	toUserId := c.Query("to_user_id")
	//将toUserId转换为数字
	toUserIdNum, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//得到fromUserId
	fromUserId := claim.UserId
	//获取上次聊天时间
	pre := c.Query("pre_msg_time")
	preMsgTime, _ := strconv.ParseInt(pre, 10, 64)
	//查询聊天记录
	listMessagesParams := mydb.ListMessagesParams{fromUserId, toUserIdNum, toUserIdNum, fromUserId, preMsgTime}
	queries := GetConn()
	listMessages, err := queries.ListMessages(ctx, listMessagesParams)
	fmt.Println("========5")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//返回聊天记录
	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: listMessages})

	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int64(userIdB))
	//
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

//func genChatKey(userIdA int64, userIdB int64) string {
//	if userIdA > userIdB {
//		return fmt.Sprintf("%d_%d", userIdB, userIdA)
//	}
//	return fmt.Sprintf("%d_%d", userIdA, userIdB)
//}
