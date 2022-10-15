package response

import "douyin/module"

type UserInfo struct {
	module.Response
	module.User `json:"user"`
}
