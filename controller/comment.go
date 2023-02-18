package controller

import (
	"context"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
// new fix: add comment creation
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	claim, err := util.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			Comment: Comment{},
		})
	}

	userID := (*claim).UserId

	if actionType == "1" {
		CreateComment(userID, c)
	} else {
		DeleteComment(c)
	}
	GetCommentList(c)
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	GetCommentList(c)
}

// CreateComment create a comment
func CreateComment(userID int64, c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	content := c.Query("comment_text")

	var queries = GetConn()

	arg := mydb.CreateCommentParams{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}

	_, err := queries.CreateComment(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusOK, "create comment error")
	}

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get comment error",
			},
			Comment: Comment{},
		})
	}

	c.JSON(http.StatusOK, CommentActionResponse{
		Response: Response{
			StatusCode: 0,
		},
		Comment: Comment{},
	})

}

func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	var queries = GetConn()

	comment, err := queries.GetComment(context.Background(), commentID)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "delete comment id error",
			},
			Comment: Comment{},
		})
	}

	if userID != comment.CommentID {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "bad user_id",
			},
			Comment: Comment{},
		})
	}

	err = queries.DeleteComment(context.Background(), commentID)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "delete comment error",
			},
			Comment: Comment{},
		})
	}

	c.JSON(http.StatusOK, CommentActionResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "delete success",
		},
		Comment: Comment{},
	})

}

func GetCommentList(c *gin.Context) {
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{

			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get video_id error",
			},
			CommentList: nil,
		})
	}

	var queries = GetConn()

	allComment, err := queries.GetCommentsById(context.Background(), videoID)

	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get comment list error",
			},
			CommentList: nil,
		})
	}
	var comments []Comment

	for _, dbComment := range allComment {

		user, _ := queries.GetUserById(context.Background(), dbComment.UserID)

		comment := Comment{
			Id: dbComment.CommentID,
			User: User{
				Id:            user.UserID,
				Name:          user.Name,
				FollowCount:   user.FollowCount.Int64,
				FollowerCount: user.FollowerCount.Int64,
				IsFollow:      true,
			},
			VideoID: dbComment.VideoID,
			Content: dbComment.Content,
		}
		comments = append(comments, comment)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
		},
		CommentList: comments,
	})
}
