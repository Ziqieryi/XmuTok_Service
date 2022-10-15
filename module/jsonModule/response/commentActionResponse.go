package response

import "douyin/module"

type CommentActionResponse struct {
	module.Response
	module.User
	CommentId  int64  `gorm:"column:comment_id"`
	Content    string `gorm:"column:content"`
	CreateDate string `gorm:"column:create_date"`
}
