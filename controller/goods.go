package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type GoodController struct{}

// GetGoodsAtr 获取商品属性
func (g GoodController) GetGoods(c *gin.Context) {
	// 从请求中获取 goodsId 参数
	categoryId := c.Param("categoryId")
	// 调用 service 层获取商品属性
	good, err := service.GetGoodByCategory(categoryId)
	if err != nil {
		// 返回错误响应
		ErrorJson(c, err.Error()) // 调用 ErrorJson 返回错误格式
		return
	}
	// 返回成功响应
	SuccessJson(c, "获取商品属性成功", good, 1) // count 设置为 1
}
