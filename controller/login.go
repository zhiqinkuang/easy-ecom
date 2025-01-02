package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiqinkuang/easy-ecom/service"
	"github.com/zhiqinkuang/easy-ecom/util"
)

type LoginController struct{}

func (login LoginController) GenerateCaptcha(c *gin.Context) {
	// 调用服务生成验证码
	masterImage, thumbImage, err := service.GenerateCaptcha()
	if err != nil {
		// 如果生成失败，返回错误响应
		ErrorJson(c, "生成验证码失败")
		return
	}

	// 返回成功响应
	SuccessJson(c, "success", gin.H{
		"masterImage": masterImage,
		"thumbImage":  thumbImage,
	}, 1)
}

// 发送短信验证码
func SendCode(c *gin.Context) {
	var req struct {
		Phone string `json:"phone" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Logger.Error("解析请求体失败: " + err.Error())
		RequestErrorJson(c, "请求体格式错误")
		return
	}
	code := service.Code()
	err := service.SendCode(req.Phone, code)
	if err != nil {
		util.Logger.Error("发送验证码失败: " + err.Error())
		ErrorJson(c, "发送验证码失败")
		return
	}

	SuccessJson(c, "验证码已发送", nil, 0)
}

// 注册用户
//func (login BannerController) Register(c *gin.Context) {
//	var req struct {
//		Phone     string `json:"phone" binding:"required"`
//		Code      string `json:"code" binding:"required"`
//		CaptchaID string `json:"captcha_id" binding:"required"`
//		Captcha   string `json:"captcha" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		util.Logger.Error("解析请求体失败: " + err.Error())
//		RequestErrorJson(c, "请求体格式错误")
//		return
//	}
//
//	// 校验图片验证码
//	if !service.VerifyCaptcha(req.CaptchaID, req.Captcha) {
//		util.Logger.Error("图片验证码错误")
//		ErrorJson(c, "图片验证码错误")
//		return
//	}
//	// 校验短信验证码
//
//	if !service.VerifySMSCode(req.Phone, req.Code) {
//		util.Logger.Error("验证码验证失败, 手机号: " + req.Phone)
//		ErrorJson(c, "验证码错误")
//		return
//	}
//
//	client := &repository.ClientLogin{
//		Username: req.Phone,
//		UID:      uuid.New().String(),
//		Status:   0,
//	}
//
//	err = service.CreateClient(client)
//	if err != nil {
//		util.Logger.Error("注册失败: " + err.Error())
//		ErrorJson(c, "用户注册失败")
//		return
//	}
//
//	SuccessJson(c, "用户注册成功", nil, 0)
//}

// 根据
// 登录用户
//func Login(c *gin.Context) {
//	var req struct {
//		Phone     string `json:"phone" binding:"required"`
//		Code      string `json:"code" binding:"required"`
//		CaptchaID string `json:"captcha_id" binding:"required"`
//		Captcha   string `json:"captcha" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		util.Logger.Error("解析请求体失败: " + err.Error())
//		RequestErrorJson(c, "请求体格式错误")
//		return
//	}
//
//	if !service.VerifyCaptcha(req.CaptchaID, req.Captcha) {
//		util.Logger.Error("图片验证码错误")
//		ErrorJson(c, "图片验证码错误")
//		return
//	}
//
//	if !service.VerifySMSCode(req.Phone, req.Code) {
//		util.Logger.Error("验证码验证失败, 手机号: " + req.Phone)
//		ErrorJson(c, "验证码错误")
//		return
//	}
//
//	client, err := service.GetClientByUserName(req.Phone)
//	if err != nil {
//		util.Logger.Error("登录失败, 手机号: " + req.Phone + ", 错误: " + err.Error())
//		ErrorJson(c, "用户不存在")
//		return
//	}
//
//	SuccessJson(c, "登录成功", map[string]interface{}{"uid": client.UID, "phone": client.Username}, 1)
//}
