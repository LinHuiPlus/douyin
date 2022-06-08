package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/repository"
	"strconv"
	"time"
)

type CommentListResponse struct {
	repository.Response
	CommentList []repository.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	repository.Response
	Comment repository.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := repository.UserIndexMap[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			//时间戳
			chuo := time.Now().Unix()
			//时间日期格式模板
			timeTemplate := "2006-01-02 15:04:05"
			tm := time.Unix(int64(chuo), 0)
			timeStr := tm.Format(timeTemplate)
			videoId := c.Query("video_id")
			id, _ := strconv.ParseInt(videoId, 10, 64)
			repository.NewCommentDaoInstance().CreateNewCommentDb(id, text, timeStr, token)
			repository.NewVideoDaoInstance().UpdateVideoCommentCount(videoId, actionType)
			c.JSON(http.StatusOK, CommentActionResponse{Response: repository.Response{StatusCode: 0},
				Comment: repository.Comment{
					Id:         id,
					User:       user.User,
					Content:    text,
					CreateDate: timeStr, //"05-01",
				}})
			//return
		} else {
			videoId := c.Query("video_id")
			id, _ := strconv.ParseInt(videoId, 10, 64)
			commentId := c.Query("comment_id")
			repository.NewCommentDaoInstance().DeleteCommentByCommentId(id, commentId)
			repository.NewVideoDaoInstance().UpdateVideoCommentCount(videoId, actionType)
		}
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	repository.NewCommentDaoInstance().QueryCommentsByVideoId(c.Query("video_id"))
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    repository.Response{StatusCode: 0},
		CommentList: repository.Comments,
	})
}
