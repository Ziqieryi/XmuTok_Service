package controller

import (
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"douyin/service/visitorService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var response response.Register
	//获取request中的信息
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		//request里的信息为空,无需调用登录服务,直接装填response
		response.StatusCode = -1
		response.StatusMsg = "Required information is NULL"
		c.JSON(http.StatusOK, response)
		return
	} else {
		userService.Register(username, password, &response)
		c.JSON(http.StatusOK, response)
	}
}

func Login(c *gin.Context) {
	//拿request里給的信息
	username := c.Query("username")
	password := c.Query("password")
	//声明此次请求需要返回的response
	var response response.Login
	// Judge whether the username and password are
	if username == "" || password == "" {
		//request里的信息为空,无需调用登录服务,直接装填response
		response.StatusCode = -1
		response.StatusMsg = "Required information is NULL"
		c.JSON(http.StatusOK, response)
		return
	} else {
		visitorService.Login(username, password, &response)
		c.JSON(http.StatusOK, response)
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var response response.UserInfo
	if token == "" || userId == "" {
		response.StatusCode = -1
		response.StatusMsg = "Required information is NULL"
		c.JSON(http.StatusOK, response)
		return
	} else {
		//进入逻辑层
		userService.UserInfo(token, userId, &response)
		c.JSON(http.StatusOK, response)
	}
}
