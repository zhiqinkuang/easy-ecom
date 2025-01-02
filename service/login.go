package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/bindata/chars"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha/v2/click"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"github.com/zhiqinkuang/easy-ecom/util"
	_ "image/jpeg" // To support decoding jpeg images
	_ "image/png"  // To support decoding png images
	"log"
	"math/rand"
	"strconv"
	"time"
)

// GetClientByUID 根据 UID 获取单个用户
func GetClientByUID(uid string) (*repository.ClientLogin, error) {
	client, err := repository.NewClientLoginDaoInstance().GetClientByUID(uid)
	if err != nil {
		util.Logger.Error("获取用户失败, UID: " + uid + ", 错误: " + err.Error())
		return nil, err
	}
	return client, nil
}

// GetClientByUID 根据 UID 获取单个用户
func GetClientByUserName(username string) (*repository.ClientLogin, error) {
	client, err := repository.NewClientLoginDaoInstance().GetClientByUserName(username)
	if err != nil {
		util.Logger.Error("获取用户失败, UserName: " + username + ", 错误: " + err.Error())
		return nil, err
	}
	return client, nil
}

// CreateClient 创建一个新的用户
func CreateClient(client *repository.ClientLogin) error {

	// 创建新的用户
	err := repository.NewClientLoginDaoInstance().CreateClient(client)
	if err != nil {
		util.Logger.Error("创建用户失败, UID: " + client.UID + ", 错误: " + err.Error())
		return err
	}

	util.Logger.Info("用户创建成功, UID: " + client.UID)
	return nil
}

// SendMsg 向手机发送验证码
func SendCode(tel string, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accesskeyId>", "<accessSecret>")
	if err != nil {
		util.Logger.Error("Failed to create SMS client:" + err.Error())
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel             // 手机号变量值
	request.SignName = "凌睿工作室"             // 签名
	request.TemplateCode = "SMS_19586XXXX" // 模板编码
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		util.Logger.Error("Failed to Send SMS " + err.Error())
		return err
	}

	fmt.Println("Response Code:", response.Code)
	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		util.Logger.Error("Failed to Send SMS " + err.Error())
		return err
	}

	return nil
}

// Code 生成随机验证码
func Code() string {
	// 创建新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成 4 位随机数，范围是 [1000, 9999]
	code := r.Intn(9000) + 1000

	// 转换为字符串并返回
	return strconv.Itoa(code)
}

// 验证短信码
func Validation(validation string, userId string) bool {
	getcode, err := repository.Get(userId)
	if err != nil && getcode == validation {
		return true
	}
	return false
}

// 生成验证码并返回两个 Base64 字符串
func GenerateCaptcha() (string, string, error) {
	var textCapt click.Captcha
	builder := click.NewBuilder()

	// 获取字体资源
	fonts, err := fzshengsksjw.GetFont()
	if err != nil {
		log.Fatalln(err)
	}

	// 获取背景图片
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	// 设置验证码资源
	builder.SetResources(
		click.WithChars(chars.GetChineseChars()),
		click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(imgs),
	)

	textCapt = builder.Make()

	// 生成验证码
	captData, err := textCapt.Generate()
	if err != nil {
		return "", "", err
	}

	// 获取主图和缩略图，这些返回的是 JPEGImageData 类型
	masterImageData, err1 := captData.GetMasterImage().ToBase64Data()
	if err1 != nil {
		return "", "", fmt.Errorf("failed to decode master image: %v", err1)
	}
	thumbImageData, err2 := captData.GetThumbImage().ToBase64Data()
	if err2 != nil {
		return "", "", fmt.Errorf("failed to decode master image: %v", err2)
	}

	return masterImageData, thumbImageData, nil
}
