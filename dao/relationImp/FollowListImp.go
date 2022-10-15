package relationImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func QueryUserById(userId int64, userList *[]module.UserTable) (err error) {
	err = dao.Db.Debug().Where("user_id IN (?)", dao.Db.Where("follower_id = ?", userId).Table("user_follow").Select("follow_id")).Find(&userList).Error
	return
}

func IsFollow(followerId int64, userList []module.UserTable) (isFolList []bool, err error) {
	var fol module.FollowTable
	for i := 0; i < len(userList); i++ {
		err = dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, userList[i].UserId).First(&fol).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				isFolList = append(isFolList, false)
				err = nil
			} else {
				return nil, err
			}
		} else {
			isFolList = append(isFolList, true)
		}
	}
	return
}
