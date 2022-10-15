package response

import "douyin/module"

type CommentList struct {
	module.Response
	List []module.Comment `json:"comment_list"`
}
