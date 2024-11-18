package main

import (
	"github.com/zhiqinkuang/easy-ecom/router"
)

func main() {
	//获得gin引擎
	r := router.Router()
	//处理GET请求
	err := r.Run(":9000")
	if err != nil {
		return
	}
}
