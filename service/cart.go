package service

import (
	"errors"
	"fmt"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
	"time"
)

type CartService struct {
	CartID     string  `json:"cart_id" binding:"required"` // 购物车ID，必填
	UserID     string  `json:"user_id" binding:"required"` // 用户ID，必填
	ProductID  string  `json:"product_id"`                 // 商品ID
	SelectNum  int     `json:"select_num"`                 // 选择数量
	Price      float64 `json:"price"`                      // 商品价格
	ImageURL   string  `json:"image_url"`                  // 商品图片URL
	IsActive   bool    `json:"is_active"`                  // 是否激活
	Properties string  `json:"properties"`                 // 商品属性
}

// NewCartService 创建购物车服务对象
func NewCartService(params CartService) *CartService {
	return &params
}

// CreateCartItem 创建购物车记录
func CreateCartItem(params CartService) error {
	service := NewCartService(params)
	return service.add()
}

// 验证 userID 和 cartID 是否存在
func (s *CartService) checkParam() error {
	if s.UserID == "" {
		util.Logger.Error("userID不能为空")
		return errors.New("userID不能为空")
	}
	if s.CartID == "" {
		util.Logger.Error("cartID不能为空")
		return errors.New("cartID不能为空")
	}
	return nil
}

// 验证参数的合法性
func checkID(userID, cartID string) error {
	if userID == "" {
		util.Logger.Error("userID不能为空")
		return errors.New("userID不能为空")
	}
	if cartID == "" {
		util.Logger.Error("cartID不能为空")
		return errors.New("cartID不能为空")
	}
	return nil
}

// 添加购物车记录
func (s *CartService) add() error {
	if err := s.checkParam(); err != nil {
		return err
	}

	cart := repository.Cart{
		CartID:     s.CartID,
		UserID:     s.UserID,
		ProductID:  s.ProductID,
		SelectNum:  s.SelectNum,
		Price:      s.Price,
		ImageURL:   s.ImageURL,
		IsActive:   s.IsActive,
		Properties: s.Properties,
		UpdatedAt:  time.Now(),
	}

	if err := repository.NewCartDaoInstance().CreateCart(&cart); err != nil {
		util.Logger.Error("创建购物车记录失败: " + err.Error())
		return err
	}
	return nil
}

// 更新购物车记录
func UpdateCartItem(userID, cartID string, selectNum int) error {
	if err := checkID(userID, cartID); err != nil {
		return err
	}
	if err := repository.NewCartDaoInstance().UpdateCartItem(userID, cartID, selectNum); err != nil {
		util.Logger.Error("更新购物车记录失败: " + err.Error())
		return err
	}
	return nil
}

// 删除购物车记录
func DeleteCart(userID, cartID string) error {
	if err := checkID(userID, cartID); err != nil {
		return err
	}
	if err := repository.NewCartDaoInstance().DeleteCartItem(userID, cartID); err != nil {
		util.Logger.Error("删除购物车记录失败: " + err.Error())
		return err
	}
	return nil
}

// 获取用户所有购物车记录
func GetCart(userID string) ([]repository.Cart, error) {
	carts, err := repository.NewCartDaoInstance().GetAllCart(userID)
	if err != nil {
		util.Logger.Error("获取购物车记录失败: " + err.Error())
		return nil, err
	}
	return carts, nil
}

// GetCartItem retrieves a single cart record based on userID and cartID
func GetCartItem(userID, cartID string) (repository.Cart, error) {
	// Validate the input IDs
	if err := checkID(userID, cartID); err != nil {
		return repository.Cart{}, fmt.Errorf("invalid ID: %w", err)
	}

	// Use a DAO instance to retrieve the cart
	cartDao := repository.NewCartDaoInstance()
	cart, err := cartDao.FindCart(userID, cartID)
	if err != nil {
		return repository.Cart{}, fmt.Errorf("failed to find cart: %w", err)
	}

	return *cart, nil
}
