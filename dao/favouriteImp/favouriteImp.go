package favouriteImp

import (
	"douyin/dao"
	"douyin/module"
)

func Insert(userId int64, videoId int64) (message string) {
	temp := module.FavTable{
		FavUserId:  userId,
		FavVideoId: videoId,
	}
	result := dao.Db.Create(&temp)
	if result.RowsAffected > 0 {
		return ""
	}
	return result.Error.Error()
}
func Delete(userId int64, videoId int64) (message string) {
	temp := module.FavTable{}
	result := dao.Db.Where("fav_user_id =? and fav_video_id =?", userId, videoId).Delete(&temp)
	if result.RowsAffected > 0 {
		return ""
	}
	return result.Error.Error()
}
