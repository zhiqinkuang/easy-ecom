package repository

import (
	"github.com/zhiqinkuang/easy-ecom/util"
	"sync"
)

type Collect struct {
	Id            int64   `gorm:"column:id"`             // 主键 ID
	GoodsNum      string  `gorm:"column:goods_num"`      // 商品编号
	GoodsName     string  `gorm:"column:goods_name"`     // 商品名称
	GoodsPrice    float64 `gorm:"column:goods_price"`    // 商品价格
	CollectStatus int64   `gorm:"column:collect_status"` // 收藏状态
	GoodsImg      string  `gorm:"column:goods_img"`      // 商品图片
	Tags          int64   `gorm:"column:tags"`           // 标签
	UserId        string  `gorm:"column:user_id"`        // 用户 ID
	GoodsDes      string  `gorm:"column:goods_des"`      // 商品描述
	CreatedAt     int64   `gorm:"column:created_at"`     // 创建时间（UNIX 时间戳）
	DeleteAt      int64   `gorm:"column:delete_at"`      // 删除时间（UNIX 时间戳）
}

func (Collect) TableName() string {
	return "collect"
}

type CollectDao struct {
}

// 建立一个Dao 对象
var collectDao *CollectDao
var collectOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewCollectDaoInstance() *CollectDao {
	collectOnce.Do(
		func() {
			collectDao = &CollectDao{}
		})
	return collectDao
}

// 通过 user_id 获得 collect
func (this *CollectDao) GetCollect(userId, goodsNum string) (*Collect, error) {
	var collect Collect
	// 添加多个查询条件：user_id、collect_status 和 goods_num
	err := db.Where("user_id = ? AND collect_status = ? AND goods_num = ?", userId, 1, goodsNum).First(&collect).Error
	if err != nil {
		util.Logger.Error("find collect by user_id, collect_status, and goods_num err: " + err.Error())
		return nil, err
	}
	return &collect, nil
}

// 添加collect 数据
func (this *CollectDao) CreateCollect(collect *Collect) error {
	// 如果没有重复，继续创建
	err := db.Create(&collect).Error
	if err != nil {
		util.Logger.Error("create collect err: " + err.Error())
		return err
	}
	return nil
}

// 修改collect数据，软删除
func (this *CollectDao) UpdateCollectStatusAndDeleteAt(userId, goodsNum string, newDeleteAt int64) error {
	// 使用 Updates 更新指定字段
	err := db.Model(&Collect{}).Where("user_id = ? AND goods_num = ?", userId, goodsNum).
		Updates(map[string]interface{}{
			"collect_status": 0,
			"delete_at":      newDeleteAt,
		}).Error
	if err != nil {
		util.Logger.Error("update collect status and delete_at err: " + err.Error())
		return err
	}
	return nil
}

// 通过user_id 查找collect数据
func (this *CollectDao) GetAllCollect(userId string) ([]Collect, error) {
	var collects []Collect
	// 查询条件：user_id 和 collect_status = 1
	err := db.Where("user_id = ? AND collect_status = ?", userId, 1).Find(&collects).Error
	if err != nil {
		util.Logger.Error("find all collects by user_id and collect_status err: " + err.Error() + ", userId: " + userId)
		return nil, err
	}
	return collects, nil
}
