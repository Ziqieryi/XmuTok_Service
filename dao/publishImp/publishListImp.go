package publishImp

import (
	"douyin/dao"
	"douyin/module"
	"gorm.io/gorm"
)

func QueryVideoByUserId(userId string, videoList *[]*module.VideoTable) (err error) {
	err = dao.Db.Where("author_id = ?", userId).Order("upload_date desc").Find(&videoList).Error
	return
}

func QueryAuthorByUserId(userId string, author *module.UserTable) (err error) {
	err = dao.Db.Where("user_id = ?", userId).First(&author).Error
	return
}

func IsFollow(followerId int64, followId string, follow *module.FollowTable) (isFollow bool, err error) {
	result := dao.Db.Where("follower_id = ? AND follow_id = ?", followerId, followId).First(&follow)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		isFollow = true
	} else {
		isFollow = false
	}
	return
}

func IsFavorite(followerId int64, videoList []*module.VideoTable, fav []*module.FavTable) (isFavList []bool, err error) {
	for i := 0; i < len(videoList); i++ {
		err := dao.Db.Where("fav_user_id = ? AND fav_video_id = ?", followerId, videoList[i].VideoId).First(&fav).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				isFavList = append(isFavList, false)
				err = nil
			} else {
				return nil, err
			}
		} else {
			isFavList = append(isFavList, true)
		}
	}
	return
}
