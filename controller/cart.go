package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type UpdateCartQuery struct {
	UserId    string `json:"uid" binding:"required"`
	CartNum   string `json:"cart_num" binding:"required"`
	SelectNum int    `json:"select_num" binding:"required"`
}
type DeleteCartQuery struct {
	UserId  string `json:"uid" binding:"required"`
	CartNum string `json:"cart_num" binding:"required"`
}

// 建立一个cartcontroller
type CartController struct{}

// 用uid获得购物车列表
func (cart CartController) GetCartList(c *gin.Context) {
	uid := c.Param("uid")
	carts, err := service.GetCart(uid)
	if err != nil {
		ErrorJson(c, "no found uid")
		return
	}
	SuccessJson(c, "成功获得用户信息", carts, int64(len(carts)))
}

func (cart CartController) AddCart(c *gin.Context) {
	var cartQuery service.CartService

	// 绑定请求数据到 CollectQuery
	if err := c.ShouldBindJSON(&cartQuery); err != nil {
		RequestErrorJson(c, "请求参数错误："+err.Error())
		return
	}

	// 调用 GetCartItem 检查购物车中是否已存在该商品
	_, err := service.GetCartItem(cartQuery.UserID, cartQuery.CartID)
	if err == nil {
		ErrorJson(c, "购物车中已存在该商品")
		return
	}

	// 调用服务层逻辑创建 Cart
	err2 := service.CreateCartItem(cartQuery)
	if err2 != nil {
		ErrorJson(c, "创建失败："+err2.Error())
		return
	}

	// 返回成功响应
	ChangeJson(c, "成功创建")
}

// 根据输入更新cart 信息
func (cart CartController) UpdateCart(c *gin.Context) {
	var cartQuery UpdateCartQuery
	if err := c.ShouldBindJSON(&cartQuery); err != nil {
		RequestErrorJson(c, "请求参数错误："+err.Error())
		return
	}

	// 调用getcartItem 实现检查是否存在,不存在则报错
	_, err := service.GetCartItem(cartQuery.UserId, cartQuery.CartNum)
	if err != nil {
		ErrorJson(c, "购物车中不存在该商品")
		return
	}
	err = service.UpdateCartItem(cartQuery.UserId, cartQuery.CartNum, cartQuery.SelectNum)
	if err != nil {
		ErrorJson(c, "更新失败："+err.Error())
		return
	}
	ChangeJson(c, "购物车成功更新")
}

// 根据输入删除 cart 信息
func (cart CartController) DeleteCart(c *gin.Context) {
	var deleteCartQuery DeleteCartQuery
	if err := c.ShouldBindJSON(&deleteCartQuery); err != nil {
		RequestErrorJson(c, "请求参数错误："+err.Error())
		return
	}
	// 调用getcartItem 实现检查是否存在,不存在则报错
	_, err := service.GetCartItem(deleteCartQuery.UserId, deleteCartQuery.CartNum)
	if err != nil {
		ErrorJson(c, "购物车中不存在该商品")
		return
	}

	err = service.DeleteCart(deleteCartQuery.UserId, deleteCartQuery.CartNum)
	if err != nil {
		ErrorJson(c, "删除失败："+err.Error())
		return
	}
	ChangeJson(c, "购物车成功删除")
}
