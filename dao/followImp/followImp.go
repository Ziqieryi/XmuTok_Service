package followImp

import (
	"douyin/dao"
	"douyin/module"
)

func Insert(followId int64, followerId int64) (message string) {
	temp := module.FollowTable{
		FollowId:   followId,
		FollowerId: followerId,
	}
	result := dao.Db.Create(&temp)
	if result.RowsAffected > 0 {
		return ""
	}
	return result.Error.Error()
}
func Delete(followId int64, followerId int64) (message string) {
	temp := module.FollowTable{}
	result := dao.Db.Where("follow_id =? and follower_id =?", followId, followerId).Delete(&temp)
	if result.RowsAffected > 0 {
		return ""
	}
	return result.Error.Error()
}
