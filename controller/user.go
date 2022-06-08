package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/repository"
	"simple-demo/service"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
/*var usersLoginInfo = map[string]repository.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
*/

type UserLoginResponse struct {
	repository.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	repository.Response
	User repository.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := service.QueryUserInfo(token); exist == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&repository.Count, 1)
		newUser := repository.DbUser{
			//Id:   userIdSequence,
			User:  repository.User{Name: username},
			Token: token,
		}
		if err := repository.NewUserDaoInstance().CreateNewUserDb(newUser); err == nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: repository.Response{StatusCode: 0},
				UserId:   newUser.Id,
				Token:    username + password,
			})
		} else {
			c.JSON(http.StatusOK, UserResponse{
				Response: repository.Response{StatusCode: 1, StatusMsg: "register failed"},
			})
		}

	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, err := service.QueryUserInfo(token); err == nil {
		repository.NewVideoDaoInstance().QueryVideoByTokenFromDb(user.User)
		repository.NewVideoDaoInstance().QueryFavoriteVideoByTokenFromDb(user.User)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 0},
			UserId:   user.User.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := service.QueryUserInfo(token); exist == nil {
		//repository.NewVideoDaoInstance().QueryVideoByTokenFromDb(token)
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 0},
			User:     user.User.User,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
