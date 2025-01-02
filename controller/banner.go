package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
)

type BannerController struct{}

// 用uid获得购物车列表
func (banner BannerController) GetBannerList(c *gin.Context) {
	banners, err := service.GetBanner()
	if err != nil {
		ErrorJson(c, "no banner")
		return
	}
	SuccessJson(c, "成功获得banner信息", banners, int64(len(banners)))
}
