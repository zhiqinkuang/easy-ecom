package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type GoodAtrController struct{}

// GetGoodsAtr 获取商品属性
func (g GoodAtrController) GetGoodsAtr(c *gin.Context) {
	// 从请求中获取 goodsId 参数
	goodsId := c.Param("goodsId")
	// 调用 service 层获取商品属性
	goodAtr, err := service.GetGoodAtr(goodsId)
	if err != nil {
		// 返回错误响应
		ErrorJson(c, err.Error()) // 调用 ErrorJson 返回错误格式
		return
	}

	// 返回成功响应
	SuccessJson(c, "获取商品属性成功", goodAtr, 1) // count 设置为 1
}
