package service

import (
	"errors"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
	"time"
)

// UserMsg 的结构体
type CollectService struct {
	GoodsNum      string  `json:"goods_num" binding:"required"` // 商品编号，必填
	GoodsName     string  `json:"goods_name"`                   // 商品名称
	GoodsPrice    float64 `json:"goods_price"`                  // 商品价格
	CollectStatus int64   `json:"collect_status"`               // 收藏状态
	GoodsImg      string  `json:"goods_img"`                    // 商品图片
	Tags          int64   `json:"tags"`                         // 商品标签
	UserId        string  `json:"user_id" binding:"required"`   // 用户 ID，必填
	GoodsDes      string  `json:"goods_des"`                    // 商品描述
}

// NewCollectService 创建服务对象
func NewCollectService(params CollectService) *CollectService {
	return &CollectService{
		GoodsNum:      params.GoodsNum,
		GoodsName:     params.GoodsName,
		GoodsPrice:    params.GoodsPrice,
		CollectStatus: params.CollectStatus,
		GoodsImg:      params.GoodsImg,
		Tags:          params.Tags,
		UserId:        params.UserId,
		GoodsDes:      params.GoodsDes,
	}
}

// CreateCollect 创建收藏记录
func CreateCollect(params CollectService) error {
	service := NewCollectService(params)
	return service.add()
}

// 更新字段,软删除
func UpdateCollect(UserId, GoodsNum string) error {
	return update(UserId, GoodsNum)
}

func FindCollect(userId, goods_num string) (*repository.Collect, error) {
	col, err := repository.NewCollectDaoInstance().GetCollect(userId, goods_num)
	if err != nil {
		return nil, err
	}
	return col, nil
}

func FindALLCollect(userId string) ([]repository.Collect, error) {
	col, err := repository.NewCollectDaoInstance().GetAllCollect(userId)
	if err != nil {
		return nil, err
	}
	return col, nil
}

func checkUpdate(GoodsNum, UserId string) (err error) {
	// 校验字段
	if GoodsNum == "" {
		err = errors.New("GoodsNum is required")
		util.Logger.Error("params Error " + err.Error())
		return err
	}
	if UserId == "" {
		err = errors.New("GoodsNum is required")
		util.Logger.Error("params Error " + err.Error())
		return err
	}

	return nil
}

func (c *CollectService) checkParam() (err error) {
	// 校验字段
	if c.GoodsNum == "" {
		err = errors.New("GoodsNum is required")
		util.Logger.Error("params Error " + err.Error())
		return err
	}
	if c.UserId == "" {
		err = errors.New("GoodsNum is required")
		util.Logger.Error("params Error " + err.Error())
		return err
	}

	return nil
}
func (this *CollectService) add() error {
	if err := this.checkParam(); err != nil {
		return err
	}
	collect := repository.Collect{
		GoodsNum:      this.GoodsNum,
		GoodsName:     this.GoodsName,
		GoodsPrice:    this.GoodsPrice,
		CollectStatus: this.CollectStatus,
		GoodsImg:      this.GoodsImg,
		Tags:          this.Tags,
		UserId:        this.UserId,
		GoodsDes:      this.GoodsDes,
		CreatedAt:     time.Now().Unix(),
	}
	err := repository.NewCollectDaoInstance().CreateCollect(&collect)
	if err != nil {
		return nil
	}
	return nil
}

func update(UserId, GoodsNum string) error {
	if err := checkUpdate(UserId, GoodsNum); err != nil {
		return err
	}
	err := repository.NewCollectDaoInstance().UpdateCollectStatusAndDeleteAt(UserId, GoodsNum, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
