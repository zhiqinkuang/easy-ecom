package controller

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	SuccessJson(c, 200, "success", "你好", 1)
}

func GetEmptyList(c *gin.Context) {
	ErrorJson(c, 400, "null")
}
