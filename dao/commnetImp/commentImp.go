package commentImp

import (
	"douyin/dao"
	"douyin/module"
)

// Judge user is it exist
func QueryUserId(userId int64, usertable *module.UserTable) (err error) {
	err = dao.Db.Where("user_id = ?", userId).Find(&usertable).Error
	return
}

// make commentMsg insert database
func InsertCommentMsg(vedioId int64, userId int64, content string, createDate string) (err error) {
	err = dao.Db.Select("com_video_id", "com_user_id", "content", "create_date").Create(module.CommentTable{ComVideoId: vedioId, ComUserId: userId, Content: content, CreateDate: createDate}).Error
	return
}

// Query commentMsg
func QueryCommentMsgRes(vedioId int64, userId int64, createDate string, commentTable *module.CommentTable) (err error) {
	err = dao.Db.Where(module.CommentTable{ComVideoId: vedioId, ComUserId: userId, CreateDate: createDate}).Find(&commentTable).Error
	return
}

// Delete comment
func DeleteCommentImp(commentId int64, CommentTable *module.CommentTable) (err error) {
	err = dao.Db.Where("comment_id", commentId).Delete(&CommentTable).Error
	return
}

// Query commentList
func GetCommentList(videoId int64, data *[]module.CommentTable) (err error) {
	err = dao.Db.Where("com_video_id = ?", videoId).Order("create_date desc").Find(&data).Error
	return
}
