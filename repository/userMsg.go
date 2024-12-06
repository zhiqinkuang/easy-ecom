package repository

import (
	"github.com/zhiqinkuang/easy-ecom/util"
	"sync"
)

// 用户信息表的module
type UserMsg struct {
	Id        int64  `gorm:"column:id"`
	UserId    string `gorm:"column:user_id"`
	Name      string `gorm:"column:name"`
	Phone     string `gorm:"column:phone"`
	Avatar    string `gorm:"column:avatar_url"`
	Address   string `gorm:"column:address"`
	MsgStatus int64  `gorm:"column:msg_status"`
}

func (UserMsg) TableName() string {
	return "user_msg"
}

type UserMsgDao struct {
}

// 建立一个Dao 对象
var userMsgDao *UserMsgDao
var userMsgOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewUserMsgDaoInstance() *UserMsgDao {
	userMsgOnce.Do(
		func() {
			userMsgDao = &UserMsgDao{}
		})
	return userMsgDao
}

// 通过 user_id 获得 user_msg
func (this *UserMsgDao) GetUserMsg(userId string) (*UserMsg, error) {
	var user UserMsg
	err := db.Where("user_id = ?", userId).Find(&user).Error
	if err != nil {
		util.Logger.Error("find user by user_id err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

// 创建user_msg
func (this *UserMsgDao) CreateUMsg(user UserMsg) error {
	if err := db.Create(&user).Error; err != nil {
		util.Logger.Error("insert user err:" + err.Error())
		return err
	}
	return nil
}

// 更新user_msg 记录
func (this *UserMsgDao) UpdateUMsg(user UserMsg) error {
	// 动态生成更新字段
	updates := map[string]interface{}{}
	if user.Name != "" {
		updates["name"] = user.Name
	}
	if user.Avatar != "" {
		updates["avatar_url"] = user.Avatar
	}
	if user.Phone != "" {
		updates["phone"] = user.Phone
	}
	if user.Address != "" {
		updates["address"] = user.Address
	}

	// 直接更新，不使用事务
	if err := db.Model(&UserMsg{}).Where("user_id = ?", user.UserId).Updates(updates).Error; err != nil {
		util.Logger.Error("update user err: " + err.Error())
		return err
	}
	util.Logger.Info(user.UserId + "update userMsg success")
	return nil
}
