package response

import "douyin/module"

type Register struct {
	module.Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
