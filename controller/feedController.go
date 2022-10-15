package controller

import (
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"douyin/service/visitorService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Feed(c *gin.Context) {
	token := c.Query("token")
	Time := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(Time, 10, 64)
	if err != nil {
		latestTime = 0
	}
	var response response.Feed
	if token == "" {
		//游客身份
		//调用游客Feed流服务装填response
		visitorService.Feed(latestTime, &response)
		c.JSON(http.StatusOK, response)
	} else {
		//用户身份
		//调用用户Feed流服务装填response
		userService.Feed(latestTime, token, &response)
		c.JSON(http.StatusOK, response)
	}
}
