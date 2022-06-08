package repository

import "strconv"

func (*DbUserDao) CreateNewUserDb(dbuser DbUser) error {
	result := db.Create(&dbuser)
	if result.Error == nil {
		UserIndexMap[dbuser.Token] = dbuser
	}
	return result.Error
}

func (*DbVideoDao) CreateNewVideoDb(dbvideo *DbVideo, user *DbUser) error {
	result := db.Create(dbvideo)
	Id := strconv.FormatInt(dbvideo.Id, 10)
	NewUserDaoInstance().UpdateUserPublishList(user, Id)
	if result.Error == nil {
		NewVideoDaoInstance().QueryVideoByCreateTimeFromDb(dbvideo.Token)
	}
	return result.Error
}

func (*DbCommentDao) CreateNewCommentDb(videoId int64, content string, createDate string, token string) error {
	var count int64
	count = 0
	db.Model(&DbComment{}).Where("video_id = ?", videoId).Count(&count)
	comment := &DbComment{
		VideoId:    videoId,
		CommentId:  count + 1,
		Content:    content,
		Token:      token,
		CreateDate: createDate,
	}
	result := db.Create(&comment)
	return result.Error
}
