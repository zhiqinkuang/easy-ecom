package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/controller"
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

	order := r.Group("/order")
	{
		// 获取指定用户和状态的订单
		order.GET("", controller.NewOrderController().GetOrdersHandler)
		// 根据订单ID获取订单
		order.GET("/:orderId", controller.NewOrderController().GetOrderByIdHandler)
		// 更新订单状态
		order.PUT("/:orderId", controller.NewOrderController().UpdateOrderStatusHandler)
		// 创建新订单
		order.POST("", controller.NewOrderController().CreateNewOrderHandler)
	}

	goodAtr := r.Group("/goodsAtr")
	{
		goodAtr.GET("/:goodsId", controller.GoodAtrController{}.GetGoodsAtr)
	}

	goodType := r.Group("/goodsType")
	{
		goodType.GET("/parentId/:parent", controller.GoodsCategoryController{}.GetGoodsTypeByParent)
		goodType.GET("/levelId/:level", controller.GoodsCategoryController{}.GetGoodsTypeByLevel)
	}

	good := r.Group("/goods")
	{
		good.GET("/:categoryId", controller.GoodController{}.GetGoods)
	}
	return r
}
