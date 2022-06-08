package repository

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (*DbUserDao) QueryUserCountFromDb() {
	db.Model(&DbUser{}).Count(&Count)
}

func (*DbUserDao) QueryUserByTokenFromDb(token string) error {
	var Usertmp DbUser
	result := db.Model(&DbUser{}).Where("token = ?", token).Find(&Usertmp)
	if result.Error != nil {
		fmt.Errorf("result.Error")
	}
	if Usertmp.Token != "" {
		UserIndexMap[token] = Usertmp
	} else {
		return fmt.Errorf("no user")
	}
	return nil
}

func (*DbUserDao) QueryUsersByTokenFromMemory(token string) (DbUser, bool) {
	if UserIndexMap[token].Token == token {
		NewUserDaoInstance().QueryUserFollowListByToken(token)
		return UserIndexMap[token], true
	} else {
		NewUserDaoInstance().QueryUserByTokenFromDb(token)
		user, exist := UserIndexMap[token]
		return user, exist
	}
}

func (*DbVideoDao) QueryVideoByTokenFromDb(user *DbUser) bool {
	VideoIdList := strings.Split(user.PublishList, ",")
	if len(VideoIdList) == 0 {
		return true
	}
	videos := make([]Video, len(VideoIdList)-1)
	for i := 0; i < len(VideoIdList)-1; i++ {
		video := DbVideo{}
		//video_user := DbUser{}
		db.Model(&DbVideo{}).Where("id = ?", VideoIdList[i]).Find(&video)
		//db.Model(&DbUser{}).Where("token = ?", video.Token).Find(&video_user)
		videos[i] = Video{
			Id:            video.Id,
			Author:        user.User,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    video.IsFavorite,
		}
	}
	UserVideos = videos
	return true
}

func (*DbVideoDao) QueryFavoriteVideoByTokenFromDb(user *DbUser) bool {
	//db.Model(&DbVideo{}).Where("token = ?", user.Token).Find(&UserTmpVideos)
	VideoIdList := strings.Split(user.FavoriteList, ",")
	if len(VideoIdList) == 0 {
		return true
	}
	videos := make([]Video, len(VideoIdList)-1)

	for i := 0; i < len(VideoIdList)-1; i++ {
		video := DbVideo{}
		videoUser := DbUser{}
		db.Model(&DbVideo{}).Where("id = ?", VideoIdList[i]).Find(&video)
		db.Model(&DbUser{}).Where("token = ?", video.Token).Find(&videoUser)
		videos[i] = Video{
			Id:            video.Id,
			Author:        videoUser.User,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    video.IsFavorite,
		}
	}
	FavoriteVideos = videos
	return true
}

func (*DbVideoDao) QueryVideoByCreateTimeFromDb(token string) {
	db.Model(&DbVideo{}).Order("create_time desc").Limit(preVideoCount).Find(&PreTmpVideos)
	if len(PreTmpVideos) == 0 {
		return
	}
	videos := make([]Video, len(PreTmpVideos))
	pos := 0

	sort.Sort(VideoList(PreTmpVideos))
	user := UserIndexMap[token]
	for _, video := range PreTmpVideos {
		newUser := DbUser{}
		db.Model(&DbUser{}).Where("token = ?", video.Token).Find(&newUser)
		IsFavorite := false
		if token != "" {
			temp := strconv.FormatInt(video.Id, 10)
			if strings.LastIndex(user.FavoriteList, temp) != -1 {
				IsFavorite = true
			}
		}
		videos[pos] = Video{
			Id:            video.Id,
			Author:        newUser.User,
			PlayUrl:       Ngork + video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			CommentCount:  video.CommentCount,
			FavoriteCount: video.FavoriteCount,
			IsFavorite:    IsFavorite, //video.IsFavorite,
		}
		pos++
	}
	PreVideos = videos
	for i, _ := range Follows {
		if strings.LastIndex(user.FollowList, strconv.FormatInt(Follows[i].Id, 10)) != -1 {
			Follows[i].IsFollow = true
		} else {
			Follows[i].IsFollow = false
		}
	}

}

func (*DbVideoDao) QueryVideoByTokenFromMemory(token string) ([]*Video, bool) {
	user, exist := NewUserDaoInstance().QueryUsersByTokenFromMemory(token)
	if exist == false {
		return nil, false
	}

	length := 0
	for _, video := range PreTmpVideos {
		if token == video.Token {
			length++
		}
	}
	videos := make([]*Video, length)
	pos := 0

	for _, video := range PreTmpVideos {
		if token == video.Token {
			videos[pos] = &Video{
				Id:            video.Id,
				Author:        user.User,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				CommentCount:  video.CommentCount,
				FavoriteCount: video.FavoriteCount,
				IsFavorite:    video.IsFavorite,
			}
			pos++
		}
	}
	return videos, true
}

func (*DbVideoDao) QueryVideoIsAlreadyExist(name string) bool {
	var video DbVideo
	result := db.Model(&DbVideo{}).Where("play_url = ?", name).Find(&video)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (*DbCommentDao) QueryCommentsByVideoId(videoId string) {
	db.Model(&DbComment{}).Where("video_id = ?", videoId).Find(&DbComments)
	if len(DbComments) == 0 {
		Comments = []Comment{}
		return
	}
	CommentsTmp := make([]Comment, len(DbComments))
	for i, comment := range DbComments {
		userTmp := DbUser{}
		db.Model(&DbUser{}).Where("token = ?", comment.Token).Find(&userTmp)
		CommentsTmp[i] = Comment{
			Id:         comment.CommentId,
			User:       userTmp.User,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
	}
	Comments = CommentsTmp
}

func (*DbUserDao) QueryUserFollowListByToken(token string) {
	var user DbUser
	var exist bool
	if user, exist = UserIndexMap[token]; exist == false {
		NewUserDaoInstance().QueryUserByTokenFromDb(token)
		if user, exist = UserIndexMap[token]; exist == false {
			//处理待续
			return
		}
	}
	FollowList := strings.Split(user.FollowList, ",")
	if len(FollowList) == 0 {
		return
	}
	usersTmp := make([]User, len(FollowList)-1)
	for i := 0; i < len(FollowList)-1; i++ {
		dbUser := DbUser{}
		db.Model(&DbUser{}).Where("id = ?", FollowList[i]).Find(&dbUser)
		usersTmp[i] = User{
			Id:            dbUser.Id,
			Name:          dbUser.Name,
			FollowCount:   dbUser.FollowCount,
			FollowerCount: dbUser.FollowerCount,
			IsFollow:      true,
		}
	}
	Follows = usersTmp
}

func (*DbUserDao) QueryUserFollowerListByToken(token string) {
	var user DbUser
	var exist bool
	if user, exist = UserIndexMap[token]; exist {
		NewUserDaoInstance().QueryUserByTokenFromDb(token)
		if user, exist = UserIndexMap[token]; exist == false {
			//处理待续
			return
		}
	}
	FollowerList := strings.Split(user.FollowerList, ",")
	if len(FollowerList) == 0 {
		return
	}
	usersTmp := make([]User, len(FollowerList)-1)
	for i := 0; i < len(FollowerList)-1; i++ {
		dbUser := DbUser{}
		db.Model(&DbUser{}).Where("id = ?", FollowerList[i]).Find(&dbUser)
		isFollow := false
		if strings.LastIndex(user.FollowList, FollowerList[i]) != -1 {
			isFollow = true
		}
		usersTmp[i] = User{
			Id:            dbUser.Id,
			Name:          dbUser.Name,
			FollowCount:   dbUser.FollowCount,
			FollowerCount: dbUser.FollowerCount,
			IsFollow:      isFollow,
		}
	}
	Followers = usersTmp
}
