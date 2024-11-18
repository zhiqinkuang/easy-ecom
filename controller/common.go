package controller

import "github.com/gin-gonic/gin"

// 返回json的格式

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type ErrorStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// 传入要返回的数据,返回固定格式的json
func SuccessJson(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

// 错误json格式数据
func ErrorJson(c *gin.Context, code int, msg interface{}) {
	json := &ErrorStruct{
		Code: code,
		Msg:  msg,
	}
	c.JSON(400, json)
}
