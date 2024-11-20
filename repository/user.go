package repository

import (
	"github.com/zhiqinkuang/easy-ecom/util"
	"sync"
)

// 用户信息表的module
type User struct {
	Id        int64  `gorm:"column:id"`
	UserId    string `gorm:"column:user_id"`
	Name      string `gorm:"column:name"`
	Phone     string `gorm:"column:phone"`
	Avatar    string `gorm:"column:avatar_url"`
	Address   string `gorm:"column:address"`
	MsgStatus int64  `gorm:"column:msg_status"`
}

func (User) TableName() string {
	return "user_msg"
}

type UserDao struct {
}

// 建立一个Dao 对象
var postDao *UserDao
var userOnce sync.Once

// 使用单例模式创建一个单例 DAO 对象
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			postDao = &UserDao{}
		})
	return postDao
}

// 通过 user_id 获得 user_msg
func getUserMsg(userId string) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userId).Find(&user).Error
	if err != nil {
		util.Logger.Error("find user by user_id err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

// 创建user_msg
func createUserMsg(user User) error {
	if err := db.Create(user).Error; err != nil {
		util.Logger.Error("insert user err:" + err.Error())
		return err
	}
	return nil
}

// 更新user_msg 记录
func updateUserMsg(user User) error {
	err := db.Model(&User{UserId: user.UserId}).Updates(map[string]interface{}{"name": user.Name, "avatar_url": user.Avatar, "phone": user.Phone, "address": user.Address}).Error
	if err != nil {
		util.Logger.Error("update user err:" + err.Error())
		return err
	}
	return nil
}

// 删除user_msg 记录
func deleteUserMsg(user User) error {
	err := db.Model(&User{UserId: user.UserId}).Updates(map[string]interface{}{"msg_status": 0}).Error
	if err != nil {
		util.Logger.Error("update user err:" + err.Error())
		return err
	}
	return nil
}
