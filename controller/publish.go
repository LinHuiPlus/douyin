package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"simple-demo/repository"
	"simple-demo/service"
	"time"
)

type VideoListResponse struct {
	repository.Response
	VideoList []repository.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := repository.UserIndexMap[token]; !exist {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := repository.UserIndexMap[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	if repository.NewVideoDaoInstance().QueryVideoIsAlreadyExist(finalName) == true {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  "已经发布过了！",
		})
		return
	}
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	newVideo := repository.DbVideo{
		Token:      token,
		PlayUrl:    finalName,
		IsFavorite: false,
		Title:      c.PostForm("title"),
		CreateTime: time.Now().Unix(),
	}
	if err := repository.NewVideoDaoInstance().CreateNewVideoDb(&newVideo, &user); err == nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded successfully",
		})
	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := service.QueryUserInfo(token); exist == nil {
		repository.NewVideoDaoInstance().QueryVideoByTokenFromDb(user.User)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: repository.Response{
				StatusCode: 0,
			},
			VideoList: repository.UserVideos,
		})
	}
}
