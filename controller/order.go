package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

// OrderController 订单控制器
type OrderController struct {
	orderService *service.OrderService
}

// NewOrderController 创建订单控制器对象
func NewOrderController() *OrderController {
	return &OrderController{
		orderService: service.NewOrderService(),
	}
}

// GetOrdersHandler 获取指定用户和状态的订单
func (c *OrderController) GetOrdersHandler(ctx *gin.Context) {
	userId := ctx.Query("userId")
	status := ctx.Query("status")
	state := parseStatus(status)
	if userId == "" || state == 0 {
		RequestErrorJson(ctx, "input error")
		return
	}
	orders, err := c.orderService.GetOrders(userId, state)
	if err != nil {
		ErrorJson(ctx, "Failed to fetch orders: "+err.Error())
		return
	}
	SuccessJson(ctx, "Orders fetched successfully", orders, int64(len(orders)))
}

// GetOrderByIdHandler 根据订单ID获取订单
func (c *OrderController) GetOrderByIdHandler(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	if orderId == "" {
		RequestErrorJson(ctx, "orderId is required")
		return
	}

	order, err := c.orderService.GetOrderById(orderId)
	if err != nil {
		ErrorJson(ctx, "Failed to fetch order: "+err.Error())
		return
	}

	SuccessJson(ctx, "Order fetched successfully", order, 1)
}

// UpdateOrderStatusHandler 更新订单状态
func (c *OrderController) UpdateOrderStatusHandler(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	var req struct {
		Status int `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		RequestErrorJson(ctx, "Invalid request data: "+err.Error())
		return
	}

	err := c.orderService.UpdateOrderStatus(orderId, req.Status)
	if err != nil {
		ErrorJson(ctx, "Failed to update order status: "+err.Error())
		return
	}

	ChangeJson(ctx, "Order status updated successfully")
}

// CreateNewOrderHandler 创建新的订单
func (c *OrderController) CreateNewOrderHandler(ctx *gin.Context) {
	var req struct {
		UserID      string  `json:"userId"`
		OrderID     string  `json:"orderId"`
		TotalAmount float64 `json:"totalAmount"`
		List        []byte  `json:"list"`
		UserMsg     []byte  `json:"userMsg"`
		Note        string  `json:"note"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		RequestErrorJson(ctx, "Invalid request data: "+err.Error())
		return
	}

	err := c.orderService.CreateNewOrder(req.UserID, req.OrderID, req.TotalAmount, req.List, req.UserMsg, req.Note)
	if err != nil {
		ErrorJson(ctx, "Failed to create new order: "+err.Error())
		return
	}

	ChangeJson(ctx, "New order created successfully")
}

// parseStatus 辅助函数，用于解析状态
func parseStatus(status string) int {
	if status == "" {
		return 0 // 默认值，例如 0 表示所有状态
	}

	// 转换字符串为整数状态值
	parsedStatus, err := strconv.Atoi(status)
	if err != nil {
		return 0
	}
	return parsedStatus
}
