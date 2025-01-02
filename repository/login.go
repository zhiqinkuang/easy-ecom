package repository

import (
	"github.com/zhiqinkuang/easy-ecom/util"
	"sync"
	"time"
)

type ClientLogin struct {
	ID        int64     `gorm:"column:id;primaryKey"`                                                    // 用户ID
	UID       string    `gorm:"column:uid;unique;not null"`                                              // 唯一的用户 ID（使用 UUID）
	Username  string    `gorm:"column:username;unique;not null"`                                         // 用户名（手机号）
	Status    int8      `gorm:"column:status;default:0;not null"`                                        // 用户状态（0: 未激活，1: 正常，2: 封禁）
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

// TableName 设置表名
func (ClientLogin) TableName() string {
	return "clientlogin"
}

// BeforeCreate 在插入记录之前生成唯一的 UID

// 定义一个数据访问对象 (DAO) 来管理 `ClientLogin` 的 CRUD 操作
type ClientLoginDao struct {
}

// 创建一个单例对象
var clientLoginDao *ClientLoginDao
var clientLoginOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewClientLoginDaoInstance() *ClientLoginDao {
	clientLoginOnce.Do(
		func() {
			clientLoginDao = &ClientLoginDao{}
		})
	return clientLoginDao
}

// 根据 UID 获取单个用户
func (clientLoginDao *ClientLoginDao) GetClientByUID(uid string) (*ClientLogin, error) {
	var client ClientLogin
	err := db.Where("uid = ?", uid).First(&client).Error
	if err != nil {
		util.Logger.Error("find client by UID err: " + err.Error())
		return nil, err
	}
	return &client, nil
}

//func (clientLoginDao *ClientLoginDao) BeforeCreate() (err error) {
//	clientLoginDao.UID = uuid.New().String() // 使用 UUID 生成 UID
//	return nil
//}

// 根据 UID 获取单个用户
func (clientLoginDao *ClientLoginDao) GetClientByUserName(username string) (*ClientLogin, error) {
	var client ClientLogin
	err := db.Where("username = ?", username).First(&client).Error
	if err != nil {
		util.Logger.Error("find client by username err: " + err.Error())
		return nil, err
	}
	return &client, nil
}

// 创建一个新的用户
func (clientLoginDao *ClientLoginDao) CreateClient(client *ClientLogin) error {
	if err := db.Create(client).Error; err != nil {
		util.Logger.Error("create client err: " + err.Error())
		return err
	}
	return nil
}
