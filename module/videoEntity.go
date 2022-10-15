package module

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	VideoTitle    string `json:"title"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}
type VideoTable struct {
	VideoId    int64  `gorm:"column:video_id"`
	AuthorId   int64  `gorm:"column:author_id"`
	PlayUrl    string `gorm:"column:play_url"`
	CoverUrl   string `gorm:"column:cover_url"`
	VideoTitle string `gorm:"column:video_title"`
	FavCount   int64  `gorm:"column:favourite_count"`
	ComCount   int64  `gorm:"column:comment_count"`
	UploadDate int64  `gorm:"column:upload_date"`
}

func (u VideoTable) TableName() string {
	// 绑定 Mysql 表名
	return "video"
}
