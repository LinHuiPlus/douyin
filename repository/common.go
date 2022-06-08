package repository

import "sync"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type DbVideo struct {
	Token         string
	Id            int64  //`json:"id,omitempty"`
	PlayUrl       string //`json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string //`json:"cover_url,omitempty"`
	FavoriteCount int64  //`json:"favorite_count,omitempty"`
	CommentCount  int64  //`json:"comment_count,omitempty"`
	IsFavorite    bool   //`json:"is_favorite,omitempty"`
	CreateTime    int64
	Title         string
}

type DbUser struct {
	User
	Token        string
	FavoriteList string
	PublishList  string
	FollowList   string
	FollowerList string
}

type DbComment struct {
	Token      string
	VideoId    int64
	CommentId  int64
	Content    string
	CreateDate string
}

type DbFavoriteList struct {
	token        string
	FavoriteList string
}

type DbVideoDao struct {
}

type DbUserDao struct {
}

type DbCommentDao struct {
}

var (
	userDao     *DbUserDao
	videoDao    *DbVideoDao
	commentDao  *DbCommentDao
	videoOnce   sync.Once
	userOnce    sync.Once
	commentOnce sync.Once
)

func NewVideoDaoInstance() *DbVideoDao {
	videoOnce.Do(
		func() {
			videoDao = &DbVideoDao{}
		})
	return videoDao
}

func NewUserDaoInstance() *DbUserDao {
	userOnce.Do(
		func() {
			userDao = &DbUserDao{}
		})
	return userDao
}

func NewCommentDaoInstance() *DbCommentDao {
	commentOnce.Do(
		func() {
			commentDao = &DbCommentDao{}
		})
	return commentDao
}
