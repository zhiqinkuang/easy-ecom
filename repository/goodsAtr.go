package repository

import (
	"gorm.io/gorm"
	"sync"
)

// 只包含所需的字段
// 'goodsID', 'stock_num', 'price', 'properties', 'image_url'
type GoodsAttribute struct {
	GoodsID    string  `gorm:"column:goodsId"`    // 商品唯一标识
	ImageURL   string  `gorm:"column:image_url"`  // 商品图片 URL
	Price      float64 `gorm:"column:price"`      // 商品价格
	StockNum   int     `gorm:"column:stock_num"`  // 库存数量
	Properties string  `gorm:"column:properties"` // JSON 类型字段，存储商品属性
}

func (GoodsAttribute) TableName() string {
	return "cartattribute"
}

type GoodsAtrDao struct {
}

// 建立一个 Dao 对象
var goodsAtrDao *GoodsAtrDao
var goodsAtrOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewGoodsAtrInstance() *GoodsAtrDao {
	goodsAtrOnce.Do(
		func() {
			goodsAtrDao = &GoodsAtrDao{}
		})
	return goodsAtrDao
}

// 根据 GoodsID 获取商品信息的方法
func (this GoodsAtrDao) GetCartAttributeByGoodsID(goodsId string) (*GoodsAttribute, error) {
	var cartAttr GoodsAttribute
	// 查询数据库
	if err := db.Select("goodsId, image_url, price, stock_num, properties").Where("goodsId = ? AND isActive = ?", goodsId, 1).First(&cartAttr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err // 未找到记录
		}
		return nil, err // 其他错误
	}
	return &cartAttr, nil
}
