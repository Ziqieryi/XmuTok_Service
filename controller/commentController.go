package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"douyin/service/visitorService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	module.Response
	CommentList []module.Comment `json:"comment_list,omitempty"`
}

func CommentAction(c *gin.Context) {
	//uid := c.Query("user_id")
	token := c.Query("token")
	vid := c.Query("video_id")
	actiontype := c.Query("action_type")

	var response response.CommentActionResponse
	//userId, err := strconv.ParseInt(uid, 10, 64)
	//if err != nil {
	//	response.StatusCode = -1
	//	response.StatusMsg = "userId Decoding failure"
	//	c.JSON(http.StatusOK, response)
	//	return
	//}
	videoId, _ := strconv.ParseInt(vid, 10, 64)

	if actiontype == "1" {
		comment_text := c.Query("comment_text")
		//commentService.AddComment(userId, token, videoId, comment_text, &response)
		userService.AddComment(token, videoId, comment_text, &response)
		c.JSON(http.StatusOK, response)
	} else if actiontype == "2" {
		cid := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(cid, 10, 64)
		userService.DeleteComment(token, commentId, &response)
		c.JSON(http.StatusOK, response)
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("video_id")
	videoId, err := strconv.ParseInt(id, 10, 64)
	var response response.CommentList
	if err != nil && id != "" {
		response.StatusCode = -1
		response.StatusMsg = "userId Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	if token == "" {
		//游客身份
		//调用游客Feed流服务装填response
		visitorService.ComList(videoId, &response)
		c.JSON(http.StatusOK, response)
	} else {
		//用户身份
		//调用用户Feed流服务装填response
		userService.ComList(videoId, &response)
		c.JSON(http.StatusOK, response)
	}
}
