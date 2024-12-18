package service

import (
	"errors"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
	"time"
)

// OrderService 订单服务
type OrderService struct {
	orderDao *repository.OrderDao
}

// NewOrderService 创建订单服务对象
func NewOrderService() *OrderService {
	return &OrderService{
		orderDao: repository.NewOrderDaoInstance(),
	}
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(order *repository.Order) error {
	// 在创建订单之前，可以进行一些数据校验等操作
	if order.UserID == "" || order.OrderID == "" {
		return errors.New("userID and orderID cannot be empty")
	}

	// 调用 DAO 层方法保存订单
	err := s.orderDao.CreateOrder(order)
	if err != nil {
		util.Logger.Error("创建订单失败: " + err.Error())
		return err
	}
	return nil
}

// GetOrders 获取指定用户和状态的订单
func (s *OrderService) GetOrders(userId string, status int) ([]repository.Order, error) {
	if userId == "" {
		return nil, errors.New("userId cannot be empty")
	}
	// 调用 DAO 层获取订单数据
	orders, err := s.orderDao.GetOrders(userId, status)
	if err != nil {
		util.Logger.Error("获取订单失败: " + err.Error())
		return nil, err
	}
	return orders, nil
}

// GetOrderById 根据订单ID获取订单
func (s *OrderService) GetOrderById(orderId string) (*repository.Order, error) {
	if orderId == "" {
		return nil, errors.New("orderId cannot be empty")
	}

	// 调用 DAO 层根据订单ID获取订单
	order, err := s.orderDao.GetOrderById(orderId)
	if err != nil {
		util.Logger.Error("根据订单ID获取订单失败: " + err.Error())
		return nil, err
	}
	return order, nil
}

// UpdateOrderStatus 更新订单状态
func (s *OrderService) UpdateOrderStatus(orderId string, status int) error {
	if orderId == "" {
		return errors.New("orderId cannot be empty")
	}

	// 调用 DAO 层更新订单状态
	err := s.orderDao.UpdateOrderStatus(orderId, status)
	if err != nil {
		util.Logger.Error("更新订单状态失败: " + err.Error())
		return err
	}
	return nil
}

// CreateNewOrder 创建新的订单，简化参数传递
func (s *OrderService) CreateNewOrder(userID, orderID string, totalAmount float64, list []byte, userMsg []byte, note string) error {
	order := &repository.Order{
		UserID:      userID,
		OrderID:     orderID,
		OrderStatus: 1, // 假设新订单的状态是 1
		TotalAmount: totalAmount,
		UpdateTime:  time.Now(),
		List:        list,
		UserMsg:     userMsg,
		Note:        note,
	}

	// 调用 CreateOrder 方法保存订单
	return s.CreateOrder(order)
}
