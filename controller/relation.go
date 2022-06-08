package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/repository"
	"strconv"
)

type UserListResponse struct {
	repository.Response
	UserList []repository.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	if user, exist := repository.UserIndexMap[token]; exist {
		actionType := c.Query("action_type")
		//userId := c.Query("user_id")
		userId := strconv.FormatInt(user.Id, 10)
		toUserId := c.Query("to_user_id")
		repository.NewUserDaoInstance().UpdateUserFollowList(user, actionType, userId, toUserId)
		repository.NewUserDaoInstance().UpdateUserFollowerList(actionType, userId, toUserId)
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	repository.NewUserDaoInstance().QueryUserFollowListByToken(token)
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: repository.Follows, //[]repository.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	repository.NewUserDaoInstance().QueryUserFollowerListByToken(token)
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: repository.Followers, //[]repository.User{DemoUser},
	})
}
