package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
)

func GetGoodAtr(userID string) (repository.GoodsAttribute, error) {
	carts, err := repository.NewGoodsAtrInstance().GetCartAttributeByGoodsID(userID)
	if err != nil {
		util.Logger.Error("获取商品属性失败: " + err.Error())
		return repository.GoodsAttribute{}, err
	}
	return *carts, nil
}
