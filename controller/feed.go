package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/repository"
	"time"
)

type FeedResponse struct {
	repository.Response
	VideoList []repository.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	repository.NewUserDaoInstance().QueryUsersByTokenFromMemory(token)
	repository.NewVideoDaoInstance().QueryVideoByCreateTimeFromDb(token)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  repository.Response{StatusCode: 0},
		VideoList: repository.PreVideos,
		NextTime:  time.Now().Unix(),
	})
}
