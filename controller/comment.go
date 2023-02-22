package controller

import (
	"context"
	"net/http"
	"strconv"
	"sync/atomic"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction handle comment request: create and delete comment
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	claim, err := util.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Token parse error",
			},
		})
	}

	userID := (*claim).UserId

	if actionType == "1" {
		CreateComment(userID, c)
	} else if actionType == "2" {
		DeleteComment(userID, c)
	}
}

// CommentList return comment list desc
func CommentList(c *gin.Context) {

	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	var queries = GetConn()

	allComment, err := queries.GetCommentsById(context.Background(), videoID)

	if err != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get comment list error",
			},
		})
		return
	}

	var comments []Comment
	for _, dbComment := range allComment {
		user, _ := queries.GetUserById(context.Background(), dbComment.UserID)

		comment := Comment{
			Id: dbComment.CommentID,
			User: User{
				Id:   user.UserID,
				Name: user.Name,
			},
			VideoID:   dbComment.VideoID,
			Content:   dbComment.Content,
			CreatedAt: dbComment.CreatedAt.Time.Format("01-02"),
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

func CreateComment(userID int64, c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	content := c.Query("comment_text")

	var queries = GetConn()
	commentID, err := queries.MaxCommentID(context.Background())

	if err != nil {
		// c.JSON(http.StatusOK, CommentActionResponse{
		// 	Response: Response{
		// 		StatusCode: 1,
		// 		StatusMsg:  "Get comment id error",
		// 	},
		// })
		// return
		commentID = 0
	}

	atomic.AddInt64(&commentID, 1)

	arg := mydb.CreateCommentParams{
		CommentID: commentID,
		UserID:    userID,
		VideoID:   videoID,
		Content:   content,
	}

	_, err = queries.CreateComment(context.Background(), arg)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Create comment error",
			},
		})
		return
	}

	comment, _ := queries.GetComment(context.Background(), commentID)
	user, _ := queries.GetUserById(context.Background(), userID)

	queries.AddCommentCount(context.Background(), videoID)

	c.JSON(http.StatusOK, CommentActionResponse{
		Response: Response{
			StatusCode: 0,
		},
		Comment: Comment{
			Id: comment.CommentID,
			User: User{
				Id:   user.UserID,
				Name: user.Name,
			},
			VideoID:   comment.VideoID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Time.Format("01-02"),
		},
	})

}

func DeleteComment(userID int64, c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	var queries = GetConn()

	comment, err := queries.GetComment(context.Background(), commentID)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Delete comment id error",
			},
		})
	}

	if userID != comment.UserID {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Bad user id",
			},
		})
		return
	}

	err = queries.DeleteComment(context.Background(), commentID)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Delete comment error",
			},
		})
		return
	}
	queries.DelCommentCount(context.Background(), videoID)

	c.JSON(http.StatusOK, CommentActionResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "Delete success",
		},
	})
}
