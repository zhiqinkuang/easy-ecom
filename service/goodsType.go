package service

import (
	"errors"
	"github.com/zhiqinkuang/easy-ecom/repository"
)

// GetCategoriesByParentAndActive 根据 ParentCategoryID 和 IsActive 获取类别信息
func GetCategoriesByParent(parentID string) ([]repository.GoodsCategory, error) {
	if parentID == "" {
		return nil, errors.New("ParentCategoryID cannot be empty")
	}
	typeList, err := repository.NewGoodsCategoryInstance().GetCategoriesByParent(parentID)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}

// GetCategoriesByActiveAndLevel 根据 IsActive 和 CategoryLevel 获取类别信息
func GetCategoriesByActiveAndLevel(level int) ([]repository.GoodsCategory, error) {
	if level <= 0 && level > 3 {
		return nil, errors.New("CategoryLevel must be greater than 0 Less than 4")
	}
	typeList, err := repository.NewGoodsCategoryInstance().GetCategoriesByLevel(level)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}
