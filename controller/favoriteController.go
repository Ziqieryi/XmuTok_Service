package controller

import (
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"douyin/service/visitorService"
	"douyin/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavouriteAction(c *gin.Context) {
	//首先拿到必要四个参数,并且声明response
	userIdString := c.Query("user_id")
	token := c.Query("token")
	videoIdString := c.Query("video_id")
	actionTypeString := c.Query("action_type")
	var response response.Favourite
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil && userIdString != "" {
		response.StatusCode = -1
		response.StatusMsg = "userId Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	videoId, err := strconv.ParseInt(videoIdString, 10, 64)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "videoId Decoding failure"
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
	//解码token验证和userId是否吻合
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
	//调用点赞服务
	userService.UserFav(user.UserId, videoId, actionType, &response)
	c.JSON(http.StatusOK, response)
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	var response response.FavouriteList
	if err != nil && id != "" {
		response.StatusCode = -1
		response.StatusMsg = "userId Decoding failure"
		c.JSON(http.StatusOK, response)
		return
	}
	if token == "" {
		//游客身份
		//调用游客Feed流服务装填response
		visitorService.FavList(userId, &response)
		c.JSON(http.StatusOK, response)
	} else {
		//用户身份
		//调用用户Feed流服务装填response
		userService.FavList(userId, token, &response)
		c.JSON(http.StatusOK, response)
	}
}
