package repository

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

// Goods 定义了 goods 表的结构
type Goods struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	GoodsID      string    `gorm:"unique;size:255;not null;column:goodsid" json:"goodsid"`
	Name         string    `gorm:"size:255;not null;column:name" json:"name"`
	ImageURL     string    `gorm:"size:255;column:image_url" json:"image_url"`
	Price        float64   `gorm:"type:decimal(10,2);not null;column:price" json:"price"`
	Quantity     int32     `gorm:"not null;column:quantity" json:"quantity"`
	InStock      bool      `gorm:"not null;column:instock" json:"instock"`
	CategoryID   string    `gorm:"size:50;not null;column:categoryid" json:"categoryid"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
	CategoryItem string    `gorm:"type:json;column:categoryItem" json:"category_item"`
}

// TableName 设置表名为 goods
func (Goods) TableName() string {
	return "goods"
}

// GoodsDao 定义了与 goods 表交互的方法
type GoodsDao struct {
}

// 单例模式实现 GoodsDao 对象

var goodsDao *GoodsDao
var goodsOnce sync.Once

// NewGoodsDao 单例初始化 GoodsDao
func NewGoodsDaoInstance() *GoodsDao {
	goodsOnce.Do(func() {
		goodsDao = &GoodsDao{}
	})
	return goodsDao
}

// GetGoodsByCategory 根据 CategoryID 获取商品信息
func (this *GoodsDao) GetGoodsByCategory(categoryID string) ([]Goods, error) {
	var goodsList []Goods

	// 查询数据库
	if err := db.Where("categoryid = ? AND instock = ?", categoryID, true).
		Find(&goodsList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 未找到记录时返回空数组
		}
		return nil, err // 其他错误
	}
	return goodsList, nil
}
