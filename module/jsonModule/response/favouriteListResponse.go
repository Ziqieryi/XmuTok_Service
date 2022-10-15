package response

import "douyin/module"

type FavouriteList struct {
	module.Response
	List []module.Video `json:"video_list"`
}
