package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/router"
	"github.com/zhiqinkuang/easy-ecom/util"
	"os"
)

func main() {
	//使用初始化程序
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	//获得gin引擎
	r := router.Router()
	r.Use(gin.Logger())
	//处理GET请求
	err := r.Run(":9000")
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
