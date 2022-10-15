package response

import "douyin/module"

type FollowerList struct {
	module.Response
	List []module.User `json:"user_list"`
}
