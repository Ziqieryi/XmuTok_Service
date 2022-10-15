package response

import "douyin/module"

type Feed struct {
	module.Response
	NextTime int64          `json:"next_time"`
	List     []module.Video `json:"video_list"`
}
