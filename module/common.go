package module

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type FollowTable struct {
	FollowId   int64 `gorm:"column:follow_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}
type FavTable struct {
	FavUserId  int64 `gorm:"column:fav_user_id"`
	FavVideoId int64 `gorm:"column:fav_video_id"`
}

type VideoWithAuthor struct {
	VideoTable
	UserTable
}
type UserLikeVideoList struct {
	FavTable
	VideoWithAuthor
}

func (u UserLikeVideoList) TableName() string {
	// 绑定 Mysql 表名
	return "user_like_videoList"
}
func (u VideoWithAuthor) TableName() string {
	return "video_with_author"
}
func (u FollowTable) TableName() string {
	// 绑定 Mysql 表名
	return "user_follow"
}
func (u FavTable) TableName() string {
	// 绑定 Mysql 表名
	return "user_favourite"
}
