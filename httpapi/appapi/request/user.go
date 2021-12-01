package request

import "time"

type Register struct {
	Phone string `json:"phone"  form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type OldPwd struct {
	Password string `json:"password" form:"password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type Login struct {
	UserName string `json:"username"  form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type SmsLogin struct {
	Phone string `json:"phone"  form:"phone" binding:"required"`
	SmsCode string `json:"sms_code" form:"sms_code"`
}

type CheckCode struct {
	SmsCode string `json:"sms_code" form:"sms_code"`
	Phone string `json:"phone"  form:"phone" binding:"required"`
}

type PhoneSms struct {
	Phone string `json:"phone"  form:"phone" binding:"required"`
}

//设置密码
type SetPwd struct {
	SmsCode string `json:"sms_code" form:"sms_code"  binding:"required"`
	Phone string `json:"phone"  form:"phone" binding:"required"`
	Password   string  `json:"password" form:"password"`
}
//修改用户信息
type UpdateUserInfo struct {
	Username string `json:"username" form:"username" `
	Sex int `json:"sex" form:"sex" `
	Birthday string `json:"birthday" form:"birthday" `
	BirthdayTime time.Time `json:"birthday_time" form:"birthday_time"`
	Avatar  string `json:"avatar" form:"avatar"`
	FromItemCategoryID int `json:"from_item_category_id" form:"from_item_category_id"`
	FromCityID int `json:"from_city_id" form:"from_city_id"`
	SignRemark string `json:"sign_remark" form:"sign_remark"`
}

type UpdatePhone struct {
	SmsCode string `json:"sms_code" form:"sms_code"`
	OldSmsCode string `json:"old_sms_code" form:"old_sms_code"`
	NewPhone string `json:"new_phone" form:"new_phone"`
	Phone string `json:"phone" form:"phone" binding:"required"`
}

type WechatAppletLogin struct {
	Code string `json:"code" form:"code" binding:"required"`
	EncryptedData string `json:"encrypted_data" form:"encrypted_data" binding:"required"`
	Iv string `json:"iv" form:"iv" binding:"required"`
	CloudID int `json:"cloud_id" form:"cloud_id"`
}


