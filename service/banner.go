package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
)

func GetBanner() ([]repository.Banner, error) {
	banners, err := repository.NewBannerDaoInstance().GetAllBanners()
	if err != nil {
		util.Logger.Error("获取购物车记录失败: " + err.Error())
		return nil, err
	}
	return banners, nil
}
