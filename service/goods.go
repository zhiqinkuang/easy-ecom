package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
)

func GetGoodByCategory(category string) ([]repository.Goods, error) {
	carts, err := repository.NewGoodsDaoInstance().GetGoodsByCategory(category)
	if err != nil {
		util.Logger.Error("根据category 获取商品属性失败: " + err.Error())
		return nil, err
	}
	return carts, nil
}
