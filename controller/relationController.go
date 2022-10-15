package controller

import (
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"douyin/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	//首先拿到必要四个参数,并且声明response
	userIdString := c.Query("user_id")
	token := c.Query("token")
	followIdString := c.Query("to_user_id")
	actionTypeString := c.Query("action_type")
	var response response.Follow
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil && userIdString != "" {
		response.StatusCode = -1
		response.StatusMsg = "userId Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	followId, err := strconv.ParseInt(followIdString, 10, 64)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "followId Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeString, 10, 64)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "actionType Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	//无误拿到userId,videoId,actionType
	//检查token
	if token == "" {
		//如果没有token
		response.StatusCode = -1
		response.StatusMsg = "missing token required"
		c.JSON(http.StatusOK, response)
		return
	}
	//拿到token
	//解码token
	user, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Encryption failed"
		c.JSON(http.StatusOK, response)
		return
	}
	////验证是否匹配
	//if user.UserId != userId {
	//	response.StatusCode = -1
	//	response.StatusMsg = "The token and user ID do not match"
	//	return
	//}
	fmt.Println(userId)
	//调用关注服务
	userService.UserFol(followId, user.UserId, actionType, &response)
	c.JSON(http.StatusOK, response)
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	var response response.FollowList
	userService.FollowList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	var response response.FollowList
	userService.FollowerList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}
