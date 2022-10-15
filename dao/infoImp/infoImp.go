package infoImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func SelectAuthorByUserId(userId string, author *module.UserTable) (err error) {
	err = dao.Db.Where("user_id = ?", userId).Find(&author).Error
	return
}

func IsFollow(followerId int64, followId string, fol *module.FollowTable) (isFolList bool, err error) {
	err = dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, followId).Find(&fol).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isFolList = false
			err = nil
		} else {
			return false, err
		}
	} else {
		isFolList = true
	}
	return
}
