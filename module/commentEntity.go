package module

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type CommentTable struct {
	CommentId  int64  `gorm:"column:comment_id"`
	ComVideoId int64  `gorm:"column:com_video_id"`
	ComUserId  int64  `gorm:"column:com_user_id"`
	Content    string `gorm:"column:content"`
	CreateDate string `gorm:"column:create_date"`
}

func (u CommentTable) TableName() string {
	// 绑定 Mysql 表名
	return "comment"
}
