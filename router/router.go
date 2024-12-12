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
	userMsg := r.Group("/usermsg")
	{
		userMsg.GET("/:uid", controller.UserMsgController{}.GetUserInfo)
		// 建立信息
		userMsg.POST("", controller.UserMsgController{}.CreateUserInfo)
		// 更新信息
		userMsg.PUT("/:uid", controller.UserMsgController{}.UpdateUserInfo)
	}

	collect := r.Group("/collect")
	{
		collect.GET("/:uid", controller.CollectController{}.GetCollectInfo)
		// 添加一条收藏
		collect.POST("", controller.CollectController{}.CreateCollectInfo)
		// 软删除
		collect.PUT("", controller.CollectController{}.UpdateCollectInfo)
	}

	cart := r.Group("/cart")
	{ // 获取购物车列表
		cart.GET("/:uid", controller.CartController{}.GetCartList)
		// 更新购物车
		cart.PUT("", controller.CartController{}.UpdateCart)
		// 删除购物车
		cart.DELETE("", controller.CartController{}.DeleteCart)
		// 添加购物车
		cart.POST("", controller.CartController{}.AddCart)
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
