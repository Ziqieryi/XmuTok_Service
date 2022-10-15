package controller

import (
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/service/userService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")
	var response response.PublishAction
	userService.PublishAction(data, token, title, &response)
	c.JSON(http.StatusOK, response)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	var response response.PublishList
	userService.PublishList(token, userId, &response)
	c.JSON(http.StatusOK, response)
}
