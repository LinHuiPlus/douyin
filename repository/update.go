package repository

import (
	"gorm.io/gorm"
	"strings"
)

func (*DbUserDao) UpdateUserFavoriteList(user *DbUser, videoId string, actionType string) {
	if actionType == "1" {
		if pos := strings.LastIndex(user.FavoriteList, videoId); pos != -1 {
			return
		}
		user.FavoriteList += videoId
		user.FavoriteList += ","
	} else {
		if pos := strings.LastIndex(user.FavoriteList, videoId); pos == -1 {
			return
		}
		pos := strings.LastIndex(user.FavoriteList, videoId)
		user.FavoriteList = user.FavoriteList[:pos] + user.FavoriteList[pos+len(videoId)+1:]
	}
	db.Model(&DbUser{}).Where("token = ?", user.Token).Update("favorite_list", user.FavoriteList)
	NewVideoDaoInstance().UpdateVideoInfo(videoId, actionType)
}

func (*DbUserDao) UpdateUserPublishList(user *DbUser, video_id string) {
	if video_id == "" {
		return
	}
	user.PublishList += video_id
	user.PublishList += ","
	db.Model(&DbUser{}).Where("token = ?", user.Token).Update("publish_list", user.PublishList)
	NewVideoDaoInstance().QueryVideoByTokenFromDb(user)
}

func (*DbVideoDao) UpdateVideoInfo(videoId string, actionType string) {
	if actionType == "1" {
		db.Model(&DbVideo{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	} else {
		db.Model(&DbVideo{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	}
}

func (*DbVideoDao) UpdateVideoCommentCount(videoId string, actionType string) {
	if actionType == "1" {
		db.Model(&DbVideo{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	} else {
		db.Model(&DbVideo{}).Where("id = ?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	}
}

func (*DbUserDao) UpdateUserFollowList(user DbUser, actionType string, userId string, toUserId string) {
	if actionType == "1" {
		if pos := strings.LastIndex(user.FollowList, toUserId); pos != -1 {
			return
		}
		user.FollowList += toUserId
		user.FollowList += ","
	} else {
		if pos := strings.LastIndex(user.FollowList, toUserId); pos == -1 {
			return
		}
		pos := strings.LastIndex(user.FollowList, toUserId)
		user.FollowList = user.FollowList[:pos] + user.FollowList[pos+len(toUserId)+1:]
	}
	db.Model(&DbUser{}).Where("id = ?", userId).Update("follow_list", user.FollowList)
	NewUserDaoInstance().UpdateUserFollowCount(userId, actionType)
}

func (*DbUserDao) UpdateUserFollowerList(actionType string, userId string, toUserId string) {
	user := DbUser{}
	db.Model(&DbUser{}).Where("id = ?", toUserId).Find(&user)
	if actionType == "1" {
		if pos := strings.LastIndex(user.FollowerList, userId); pos != -1 {
			return
		}
		user.FollowerList += userId
		user.FollowerList += ","
	} else {
		if pos := strings.LastIndex(user.FollowerList, userId); pos == -1 {
			return
		}
		pos := strings.LastIndex(user.FollowerList, userId)
		user.FollowerList = user.FollowerList[:pos] + user.FollowerList[pos+len(userId)+1:]
	}
	db.Model(&DbUser{}).Where("id = ?", toUserId).Update("follower_list", user.FollowerList)
	NewUserDaoInstance().UpdateUserFollowerCount(toUserId, actionType)
}

func (*DbUserDao) UpdateUserFollowCount(userId string, actionType string) {
	if actionType == "1" {
		db.Model(&DbUser{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
	} else {
		db.Model(&DbUser{}).Where("id = ?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
	}
}

func (*DbUserDao) UpdateUserFollowerCount(userId string, actionType string) {
	if actionType == "1" {
		db.Model(&DbUser{}).Where("id = ?", userId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
	} else {
		db.Model(&DbUser{}).Where("id = ?", userId).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
	}
}
