package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type UpdateQuery struct {
	UserId   string `json:"uid" binding:"required"`
	GoodsNum string `json:"goods_num" binding:"required"`
}

// 为防止重名函数的干扰建立一个struct
type CollectController struct{}

// 获得收藏列表
func (col CollectController) GetCollectInfo(c *gin.Context) {
	uid := c.Param("uid")
	collects, err := service.FindALLCollect(uid)
	if err != nil {
		ErrorJson(c, "no found uid")
		return
	}

	SuccessJson(c, "成功获得用户信息", collects, int64(len(collects)))
}

// 创造 collect 信息
func (col CollectController) CreateCollectInfo(c *gin.Context) {
	var collectQuery service.CollectService

	// 绑定请求数据到 CollectQuery
	if err := c.ShouldBindJSON(&collectQuery); err != nil {
		RequestErrorJson(c, "请求参数错误："+err.Error())
		return
	}
	// 查重逻辑
	existingCollect, err := service.FindCollect(collectQuery.UserId, collectQuery.GoodsNum)
	if existingCollect != nil {
		ErrorJson(c, "记录已存在")
		return
	}
	// 调用服务层逻辑创建 Collect
	err = service.CreateCollect(collectQuery)
	if err != nil {
		ErrorJson(c, "创建失败："+err.Error())
		return
	}
	// 返回成功响应
	ChangeJson(c, "成功创建")
}

// 更新collect,软删除
func (col CollectController) UpdateCollectInfo(c *gin.Context) {
	var updateQuery UpdateQuery
	if err := c.ShouldBindJSON(&updateQuery); err != nil {
		RequestErrorJson(c, err.Error())
		return
	}
	// 创建activity
	if err := service.UpdateCollect(updateQuery.UserId, updateQuery.GoodsNum); err != nil {
		ErrorJson(c, "创建失败："+err.Error())
		return
	}
	ChangeJson(c, "成功修改用户信息")
}
