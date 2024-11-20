package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/controller"
	"net/http"
)

// 路由初始化函数
func Router() *gin.Engine {
	r := gin.Default()

	// 用户相关路由
	user := r.Group("/user")
	{
		user.GET("/info", controller.UserController{}.GetUserInfo)
		user.GET("/list", controller.UserController{}.GetList)    // 获取用户列表
		user.GET("/:id", controller.UserController{}.GetUserByID) // 根据 ID 获取用户
		user.POST("", createUser)                                 // 创建新用户
		user.PUT("/:id", updateUserByID)                          // 更新用户信息
		user.DELETE("/:id", deleteUserByID)                       // 删除用户
	}

	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}
	return r
}

// 获取用户列表
func getUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "用户列表"})
}

// 根据 ID 获取用户
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "获取用户", "id": id})
}

// 创建用户
func createUser(c *gin.Context) {
	// 解析请求体（假设请求体包含 JSON 格式的用户数据）
	var user map[string]interface{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "用户已创建", "user": user})
}

// 更新用户信息
func updateUserByID(c *gin.Context) {
	id := c.Param("id")
	var user map[string]interface{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户信息已更新", "id": id, "user": user})
}

// 删除用户
func deleteUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "用户已删除", "id": id})
}
