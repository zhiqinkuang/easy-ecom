package repository

import (
	"github.com/zhiqinkuang/easy-ecom/util"
	"sync"
	"time"
)

type Banner struct {
	ID        int64     `gorm:"column:id;primaryKey"`                                                    // 主键 ID
	UID       string    `gorm:"column:uid;unique;not null"`                                              // 唯一的用户 ID
	URL       string    `gorm:"column:url;not null"`                                                     // Banner 的 URL
	FileName  string    `gorm:"column:file_name;not null"`                                               // 文件名
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`                             // 创建时间戳
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间戳
	Status    bool      `gorm:"column:status;default:true"`                                              // Banner 状态，是否启用
}

func (Banner) TableName() string {
	return "banner"
}

type BannerDao struct {
}

// 建立一个Dao 对象
var bannerDao *BannerDao
var bannerOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewBannerDaoInstance() *BannerDao {
	bannerOnce.Do(
		func() {
			bannerDao = &BannerDao{}
		})
	return bannerDao
}

// 获取所有 Banner
func (bannerDao *BannerDao) GetAllBanners() ([]Banner, error) {
	var banners []Banner
	err := db.Where("status = ?", true).Find(&banners).Error
	if err != nil {
		util.Logger.Error("find all banners err: " + err.Error())
		return nil, err
	}
	return banners, nil
}
