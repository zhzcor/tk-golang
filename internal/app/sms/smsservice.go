package sms

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"math/rand"
	"regexp"
	"time"
	"tkserver/internal/cache"
	"tkserver/internal/config"
	"tkserver/internal/errorno"
)

type Sms struct {
	Client dysmsapi.Client
}

func NewSmsService() *Sms {
	aliyunConfig := config.ServerConfig.Aliyun
	acid := aliyunConfig.AccessKeyId
	scrKey := aliyunConfig.AccessKeySecret
	Endpoint := "dysmsapi.aliyuncs.com"
	time := 2000
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &acid,
		// 您的AccessKey Secret
		AccessKeySecret: &scrKey,

		Endpoint: &Endpoint,

		ReadTimeout: &time,
	}
	// 访问的域名
	client, _ := dysmsapi.NewClient(config)

	service := Sms{
		Client: *client,
	}
	return &service
}

//发送短信userType 0：用户登录 1 :新用户注册
func (*Sms) SendSmsCode(client *dysmsapi.Client, phone string, param interface{}, userType int) (bool, error) {
	request := &dysmsapi.SendSmsRequest{}
	// 该参数值为假设值，请您根据实际情况进行填写
	request.SetPhoneNumbers(phone)
	paramJson, _ := json.Marshal(param)

	request.SetTemplateParam(string(paramJson))
	if userType > 0 { //成功提示短信
		request.SetTemplateCode("SMS_218725888")
		request.SetSignName("盛世明德教育")
	} else {
		//登录
		request.SetTemplateCode("SMS_218565948")
		request.SetSignName("盛世明德教育")
	}

	response, err := client.SendSms(request)
	if err != nil {
		return false, err
	}

	if *response.Body.Code != "OK" {
		return false, errorno.NewErr(errorno.Errno{
			Code:    20008,
			Message: *response.Body.Message,
		})
	}

	return true, nil
}

// 识别手机号码
func (s *Sms) IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if result {
		return true
	}
	return false
}

//生成随机验证码
func (s *Sms) RandCode() string {
	// 6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return rndCode
}

//验证code
func (s *Sms) CheckCode(c string, phone interface{}) bool {

	lru := cache.MemoryCache
	smsKey := cache.SmsCacheKey + phone.(string)

	code, ok := lru.Get(smsKey)
	if ok && code == nil {
		return false
	}

	str, _ := code.(string)

	if str == c {
		return true
	}

	return false

}
