package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/repository"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	if user, exist := repository.UserIndexMap[token]; exist {
		action_type := c.Query("action_type")
		repository.NewUserDaoInstance().UpdateUserFavoriteList(&user, videoId, action_type)
		//repository.NewVideoDaoInstance().UpdateVideoInfo(videoId, action_type)
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := repository.NewUserDaoInstance().QueryUsersByTokenFromMemory(token); exist {
		if user, exist := repository.UserIndexMap[token]; exist {
			repository.NewVideoDaoInstance().QueryFavoriteVideoByTokenFromDb(&user)
		}
		//这里还需要考虑用户不存在的处理，待续
		c.JSON(http.StatusOK, VideoListResponse{
			Response: repository.Response{
				StatusCode: 0,
			},
			VideoList: repository.FavoriteVideos,
		})
	}
}
