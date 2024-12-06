package controller

import "github.com/gin-gonic/gin"

// 返回json的格式

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type ChangeStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

type ErrorStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// 传入要返回的数据,返回固定格式的json
func SuccessJson(c *gin.Context, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  200,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

// 服务错误json格式数据
func ErrorJson(c *gin.Context, msg interface{}) {
	json := &ErrorStruct{
		Code: 500,
		Msg:  msg,
	}
	c.JSON(500, json)
}

// 返回请求错误
func RequestErrorJson(c *gin.Context, msg interface{}) {
	json := &ErrorStruct{
		Code: 400,
		Msg:  msg,
	}
	c.JSON(400, json)
}

// 修改成功返回数据
func ChangeJson(c *gin.Context, msg interface{}) {
	json := &ChangeStruct{
		Code: 200,
		Msg:  msg,
	}
	c.JSON(200, json)
}
