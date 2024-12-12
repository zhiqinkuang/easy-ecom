package repository

import (
	"errors"
	"fmt"
	"github.com/zhiqinkuang/easy-ecom/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Cart struct {
	ID         int64     `gorm:"column:id;primaryKey"`         // 主键 ID
	CartID     string    `gorm:"column:cart_id"`               // 购物车 ID
	UserID     string    `gorm:"column:user_id"`               // 用户 ID
	ProductID  string    `gorm:"column:product_id"`            // 商品 ID
	SelectNum  int       `gorm:"column:select_num;default:1"`  // 选择的商品数量，默认 1
	Price      float64   `gorm:"column:price"`                 // 商品价格
	ImageURL   string    `gorm:"column:image_url"`             // 商品图片 URL
	IsActive   bool      `gorm:"column:isActive;default:true"` // 是否激活
	Properties string    `gorm:"column:properties"`            // JSON 格式的属性
	UpdatedAt  time.Time `gorm:"column:updated_at"`            // 更新时间戳
}

func (Cart) TableName() string {
	return "cart_items"
}

type CartDao struct {
}

// 建立一个Dao 对象
var cartDao *CartDao
var cartOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewCartDaoInstance() *CartDao {
	cartOnce.Do(
		func() {
			cartDao = &CartDao{}
		})
	return cartDao
}

// 直接创建一个Cart
func (cartDao *CartDao) CreateCart(cart *Cart) error {
	// 如果没有重复，继续创建
	err := db.Create(cart).Error
	if err != nil {
		util.Logger.Error("create collect err: " + err.Error())
		return err
	}
	return nil
}

// 根据cartId和userId 查找购物车项
func (this *CartDao) FindCart(userId, cartId string) (*Cart, error) {
	var cart Cart
	err := db.Where("user_id = ? AND cart_id =  ? AND isActive=?", userId, cartId, 1).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Error(fmt.Sprintf("Cart item not found for userId: %s, cartId: %s, error: %v", userId, cartId, err))
			return nil, err
		}
		util.Logger.Error(fmt.Sprintf("Failed to find cart for userId: %s, cartId: %s, error: %v", userId, cartId, err))
		return nil, err
	}
	return &cart, nil
}

// 通过 user_id 获得 用户的cart_item
func (this *CartDao) GetAllCart(userId string) ([]Cart, error) {
	var carts []Cart
	// 查询条件：user_id 和 collect_status = 1
	err := db.Where("user_id = ? AND isActive = ?", userId, 1).Find(&carts).Error
	if err != nil {
		util.Logger.Error("find all collects by user_id and collect_status err: " + err.Error() + ", userId: " + userId)
		return nil, err
	}
	return carts, nil
}

// 修改某一个item的参数 SelectNum
// UpdateCartItem 更新指定用户购物车中商品的选择数量
func (this *CartDao) UpdateCartItem(userId string, cartID string, selectNum int) error {
	// 更新 select_num 字段
	err := db.Model(&Cart{}).
		Where("user_id = ? AND cart_id = ?", userId, cartID).
		Update("select_num", selectNum).Error
	if err != nil {
		util.Logger.Error(fmt.Sprintf("failed to update select_num for userId %s, cartID %s: %s", userId, cartID, err.Error()))
		return fmt.Errorf("failed to update select_num for userId %s, cartID %s: %v", userId, cartID, err)
	}

	return nil
}

// 修改某一个item的参数
func (this *CartDao) DeleteCartItem(userId string, cartID string) error {

	// 将 is_active 设置为 false，表示软删除
	err := db.Model(&Cart{}).
		Where("user_id = ? AND cart_id = ?", userId, cartID).
		Update("isActive", false).Error
	if err != nil {
		util.Logger.Error(fmt.Sprintf("failed to soft delete cart item for userId %s, cartID %s: %s", userId, cartID, err.Error()))
		return fmt.Errorf("failed to soft delete cart item for userId %s, cartID %s: %v", userId, cartID, err)
	}

	return nil
}
