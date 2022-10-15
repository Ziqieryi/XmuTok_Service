package feedImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func Feed3(userId int64, videoId int64) (exist bool, message string) {
	//根据用户id和视频id查询用户是否点赞过该视频
	temp := module.FavTable{}
	temp.FavVideoId = 10
	temp.FavUserId = 10
	err := dao.Db.Where("fav_user_id=? and fav_video_id=?", userId, videoId).Take(&temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, ""
		}
		return false, "select exception"
	}
	if temp.FavUserId == userId && temp.FavVideoId == videoId {
		return true, ""
	}
	return false, ""
}
func Feed4(authorId int64, userId int64) (exist bool, message string) {
	//根据用户id和作者id查询用户是否关注了该作者
	temp := module.FollowTable{}
	err := dao.Db.Where("follow_id=? and follower_id=?", authorId, userId).Take(&temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, ""
		}
		return false, "select exception"
	}
	if temp.FollowId == authorId && temp.FollowerId == userId {
		return true, ""
	}
	return false, ""
}
func Feed2(latestTime int64, data *[]module.VideoWithAuthor) (message string) {
	//有限制时间戳,返回视频表里比给定时间戳更小即更早的五个视频
	//装填data,如果有错误信息,通过写入message直接返回
	err := dao.Db.Where("upload_date<?", latestTime).Order("upload_date desc").Limit(2).Offset(0).Find(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Cant find any result"
		}
		return err.Error()
	}
	return ""
}
func Feed1(data *[]module.VideoWithAuthor) (message string) {
	//没有限制时间戳,返回视频表里时间最早的五个视频
	//装填data,如果有错误信息,通过写入message直接返回
	err := dao.Db.Order("upload_date desc").Limit(2).Offset(0).Find(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Cant find any result"
		}
		return err.Error()
	}
	return ""
}
