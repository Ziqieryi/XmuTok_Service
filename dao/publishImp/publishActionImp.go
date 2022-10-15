package publishImp

import (
	"douyin/dao"
	"douyin/module"
	"time"
)

func QueryUserId(userId int64, usertable *module.UserTable) (err error) {
	err = dao.Db.Where("user_id = ?", userId).Find(&usertable).Error
	return err
}

// insert data to video
func InsertData(userId int64, play_url string, cover_url string, title string) error {
	temp := module.VideoTable{
		VideoId:    0,
		AuthorId:   userId,
		PlayUrl:    play_url,
		CoverUrl:   cover_url,
		VideoTitle: title,
		FavCount:   0,
		ComCount:   0,
		UploadDate: time.Now().Unix(),
	}
	result := dao.Db.Create(&temp)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}
