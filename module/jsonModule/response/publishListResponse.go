package response

import "douyin/module"

type PublishList struct {
	module.Response
	VideoList []module.Video `json:"video_list"`
}
