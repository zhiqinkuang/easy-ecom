package repository

import (
	"errors"
	"github.com/zhiqinkuang/easy-ecom/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

// 查看用户的order信息
type Order struct {
	ID          int       `gorm:"column:id"`          // 主键 ID，自增
	UserID      string    `gorm:"column:user_id"`     // 用户ID
	OrderID     string    `gorm:"column:order_id"`    // 订单ID，设置了唯一索引
	OrderStatus int       `gorm:"column:orderstatus"` // 订单状态
	TotalAmount float64   `gorm:"column:totalamount"` // 订单总金额
	UpdateTime  time.Time `gorm:"column:updatetime"`  // 更新时间，默认当前时间戳
	List        []byte    `gorm:"column:list"`        // JSON类型字段，存储列表数据，这里用[]byte接收
	UserMsg     []byte    `gorm:"column:user_msg"`    // JSON类型字段，存储用户相关消息，用[]byte接收
	Note        string    `gorm:"column:note"`        // 备注信息
}

func (Order) TableName() string {
	return "order"
}

type OrderDao struct {
}

var orderDao *OrderDao
var orderOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewOrderDaoInstance() *OrderDao {
	orderOnce.Do(
		func() {
			orderDao = &OrderDao{}
		})
	return orderDao
}

// 创建order
func (o *OrderDao) CreateOrder(order *Order) error {
	// 如果没有重复，继续创建
	err := db.Create(&order).Error
	if err != nil {
		util.Logger.Error("create collect err: " + err.Error())
		return err
	}
	return nil
}

// 根据 userId,status 查询符合条件的order
func (o *OrderDao) GetOrders(userId string, status int) ([]Order, error) {
	var orders []Order
	err := db.Where("user_id = ? and orderstatus = ?", userId, status).Find(&orders).Error
	if err != nil {
		util.Logger.Error("get orders err: " + err.Error())
		return nil, err
	}
	return orders, nil
}

// 根据 orderId 查询符合条件的order
func (o *OrderDao) GetOrderById(orderId string) (*Order, error) {
	var order Order
	err := db.Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		// 如果是not found错误，则返回nil和nil
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Error("get order err: " + err.Error())
			return nil, err
		}
		util.Logger.Error("get order err: " + err.Error())
		return nil, err
	}
	return &order, nil
}

// 通过修改orderstatus来修改订单状态
func (o *OrderDao) UpdateOrderStatus(orderId string, status int) error {
	err := db.Model(&Order{}).Where("order_id = ?", orderId).Update("orderstatus", status).Error
	if err != nil {
		util.Logger.Error("update order status err: " + err.Error())
		return err
	}
	return nil
}
