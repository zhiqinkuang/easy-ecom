package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
	"github.com/zhiqinkuang/easy-ecom/util"
	"strconv"
)

type GoodsCategoryController struct{}

func (g GoodsCategoryController) GetGoodsTypeByLevel(c *gin.Context) {
	// 从请求中获取 goodsId 参数
	level := c.Param("level")
	l, err := strconv.Atoi(level)
	if err != nil {
		util.Logger.Error("参数错误")
		ErrorJson(c, "参数错误")
	}
	// 调用 service 层获取商品属性
	goodType, err2 := service.GetCategoriesByActiveAndLevel(l)
	if err != nil {
		// 返回错误响应
		ErrorJson(c, err2.Error()) // 调用 ErrorJson 返回错误格式
		return
	}
	// 返回成功响应
	SuccessJson(c, "获取商品类型成功", goodType, 1) // count 设置为 1
}

func (g GoodsCategoryController) GetGoodsTypeByParent(c *gin.Context) {
	// 从请求中获取 goodsId 参数
	P := c.Param("parent")
	// 调用 service 层获取商品属性
	goodType, err := service.GetCategoriesByParent(P)
	if err != nil {
		// 返回错误响应
		ErrorJson(c, err.Error()) // 调用 ErrorJson 返回错误格式
		return
	}
	// 返回成功响应
	SuccessJson(c, "获取商品类型成功", goodType, 1) // count 设置为 1
}
