package response

import "douyin/module"

type FollowList struct {
	module.Response
	UserList []module.User `json:"user_list"`
}
