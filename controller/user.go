package controller

import "github.com/gin-gonic/gin"

// 为防止重名函数的干扰建立一个struct
type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	SuccessJson(c, 200, "success", "你好", 1)
}

func (u UserController) GetList(c *gin.Context) {
	ErrorJson(c, 400, "null")
}

func (u UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	SuccessJson(c, 200, "success", id, 1)
}
