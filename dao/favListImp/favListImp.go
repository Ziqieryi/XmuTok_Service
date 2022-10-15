package favListImp

import (
	"douyin/dao"
	"douyin/module"
	"fmt"
	"gorm.io/gorm"
)

func GetVideoList(userId int64, data *[]module.UserLikeVideoList) (message string) {
	//根据userId,装填用户id为userId喜欢的所有视频进data
	err := dao.Db.Where("fav_user_id=?", userId).Find(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Cant find any result"
		}
		return err.Error()
	}
	fmt.Println("走了这里", userId)
	return ""
}
func IsFav(userId int64, videoId int64) (exist bool, message string) {
	//根据用户id和视频id查询用户是否点赞过该视频
	temp := module.FavTable{}
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
func IsFollow(authorId int64, userId int64) (exist bool, message string) {
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
