package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type UserInfoQuery struct {
	UserId  string `json:"uid" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Avatar  string `json:"avatar" binding:"required"`
	Address string `json:"address" binding:"required"`
}

// 为防止重名函数的干扰建立一个struct
type UserMsgController struct{}

func (u UserMsgController) GetUserInfo(c *gin.Context) {
	uid := c.Param("uid")
	umsg, err := service.QueryUserMsg(uid)
	if err != nil {
		ErrorJson(c, "no found uid")
	}
	res := UserInfoQuery{
		Name:    umsg.Name,
		Phone:   umsg.Phone,
		Avatar:  umsg.Avatar,
		Address: umsg.Address,
	}
	SuccessJson(c, "成功获得用户信息", res, 1)
}

func (u UserMsgController) CreateUserInfo(c *gin.Context) {
	var userQuery UserInfoQuery
	if err := c.ShouldBindJSON(&userQuery); err != nil {
		RequestErrorJson(c, err.Error())
	}
	// 创建activity
	if err := service.CreateUserMsg(userQuery.UserId, userQuery.Name, userQuery.Phone, userQuery.Avatar, userQuery.Address); err != nil {
		ErrorJson(c, err.Error())
	}
	ChangeJson(c, "成功创建")

}

func (u UserMsgController) UpdateUserInfo(c *gin.Context) {
	var userQuery UserInfoQuery
	if err := c.ShouldBindJSON(&userQuery); err != nil {
		RequestErrorJson(c, err.Error())
	}
	// 创建activity
	if err := service.UpdateUserMsg(userQuery.UserId, userQuery.Name, userQuery.Phone, userQuery.Avatar, userQuery.Address); err != nil {
		ErrorJson(c, err.Error())
	}
	ChangeJson(c, "成功修改用户信息")

}
