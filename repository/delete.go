package repository

func (*DbCommentDao) DeleteCommentByCommentId(videoId int64, commentId string) {
	db.Where("video_id = ? AND comment_id = ?", videoId, commentId).Delete(&DbComment{})
}
