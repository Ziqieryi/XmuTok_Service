package module

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Signature       string `json:"signature"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
}
type UserTable struct {
	UserId          int64  `gorm:"column:user_id"`
	Username        string `gorm:"column:user_name"`
	Password        string `gorm:"column:account_password"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	Signature       string `gorm:"column:signature"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
}

func (u UserTable) TableName() string {
	// 绑定 Mysql 表名
	return "user_table"
}
