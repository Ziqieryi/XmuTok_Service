package relationImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
	"sort"
)

func QueryFollwerListUserById(userId int64, userList *[]module.UserTable) (err error) {
	err = dao.Db.Debug().Where("user_id IN (?)", dao.Db.Where("follow_id = ?", userId).Table("user_follow").Select("follower_id")).Find(&userList).Error
	return
}

func IsFollwer(followerId int64, userList []module.UserTable) (isFolList []bool, err error) {
	var fol []module.FollowTable
	//for i := 0; i < len(userList); i++ {
	//err = dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, userList[i].UserId).First(&fol).Error
	//	if err != nil {
	//		if err == gorm.ErrRecordNotFound {
	//			isFolList = append(isFolList, false)
	//			err = nil
	//		} else {
	//			return nil, err
	//		}
	//	} else {
	//		isFolList = append(isFolList, true)
	//	}
	//}
	//return
	//查询followerId关注的用户
	err = dao.Db.Debug().Where("follower_id = ?", followerId).Find(&fol).Error
	followIds := make([]int, len(fol))
	for i := 0; i < len(fol); i++ {
		followIds[i] = int(fol[i].FollowId)
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isFolList = append(isFolList, false)
			err = nil
		} else {
			return nil, err
		}
	} else {
		//排序
		sort.Ints(followIds)
		for i := 0; i < len(userList); i++ {
			//二分查找followIds对应的userList[i].UserId的值是否存在
			index := twoPoints(followIds, int(userList[i].UserId), len(followIds)-1)
			if index != -1 { //存在
				isFolList = append(isFolList, true)
			} else { //不存在
				isFolList = append(isFolList, false)
			}
		}
	}

	return
}

//二分查找
func twoPoints(nums []int, key int, n int) int {
	var low, high, mid = 0, n, 0
	for {
		mid = (high + low) / 2
		if key < nums[mid] {
			high = mid - 1
		} else if key > nums[mid] {
			low = mid + 1
		} else {
			return nums[mid]
		}
		if high < low {
			return -1
		}
	}
	return -1
}
