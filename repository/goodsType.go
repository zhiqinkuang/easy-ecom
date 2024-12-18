package repository

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type GoodsCategory struct {
	ID               int       `gorm:"column:ID;primaryKey;autoIncrement"`              // 主键，自增
	CategoryID       string    `gorm:"column:CategoryID;unique;not null;size:50"`       // 类别 ID，唯一
	CategoryName     string    `gorm:"column:CategoryName;not null;size:255"`           // 类别名称
	ParentCategoryID string    `gorm:"column:ParentCategoryID;size:50"`                 // 父类别 ID，可为空
	ImageURL         string    `gorm:"column:ImageURL;size:255"`                        // 类别图片 URL，可为空
	IsActive         bool      `gorm:"column:IsActive;default:true"`                    // 是否激活，默认值为 1（true）
	CategoryLevel    int       `gorm:"column:CategoryLevel;not null"`                   // 类别等级
	UpdateTime       time.Time `gorm:"column:UpdateTime;type:timestamp;autoUpdateTime"` // 更新时间，自动更新
}

// TableName 设置表名为 goodscategory
func (GoodsCategory) TableName() string {
	return "goodscategory"
}

type GoodsCategoryDao struct {
}

var goodsCategoryDao *GoodsCategoryDao
var goodsCategoryOnce sync.Once

// NewGoodsCategoryInstance 创建 GoodsCategoryDao 的单例实例
func NewGoodsCategoryInstance() *GoodsCategoryDao {
	goodsCategoryOnce.Do(
		func() {
			goodsCategoryDao = &GoodsCategoryDao{}
		})
	return goodsCategoryDao
}

// GetCategoryByID 根据 CategoryID 获取类别信息
func (dao GoodsCategoryDao) GetCategoryByID(categoryID string) (*GoodsCategory, error) {
	var category GoodsCategory
	if err := db.Select("ID  ,CategoryID, CategoryName, ParentCategoryID, ImageURL, IsActive, CategoryLevel, UpdateTime").Where("CategoryID = ? AND IsActive = ? ", categoryID, true).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err // 未找到记录
		}
		return nil, err // 其他错误
	}
	return &category, nil
}

// GetCategoriesByActiveAndLevel 根据 IsActive 和 CategoryLevel 获取类别信息
func (dao GoodsCategoryDao) GetCategoriesByLevel(level int) ([]GoodsCategory, error) {
	var categories []GoodsCategory
	if err := db.Select("ID  ,CategoryID, CategoryName, ParentCategoryID, ImageURL, IsActive, CategoryLevel, UpdateTime").Where("IsActive = ? AND CategoryLevel = ?", true, level).Find(&categories).Error; err != nil {
		return nil, err // 查询错误
	}
	return categories, nil
}

// GetCategoriesByParentAndActive 根据 ParentCategoryID
func (dao GoodsCategoryDao) GetCategoriesByParent(parentID string) ([]GoodsCategory, error) {
	var categories []GoodsCategory
	if err := db.Select("ID ,CategoryID, CategoryName, ParentCategoryID, ImageURL, IsActive, CategoryLevel, UpdateTime").Where("ParentCategoryID = ? AND IsActive = ?", parentID, true).Find(&categories).Error; err != nil {
		return nil, err // 查询错误
	}
	return categories, nil
}
