package request

//oss
type AuthOssRequest struct {
	Action *string   `json:"action" form:"action" binding:"required"`
}

//获取协议详情
type GetAgreement struct {
	Type int `json:"type" form:"type" binding:"required"`//1：服务协议，2：隐私政策，3：ios充值协议，4：账户注销协议，5：App温馨提示，6：优惠卷使用规则

}

//app版本
type PhoneType struct {
	Type int `json:"type" form:"type" binding:"required"` //适用机型 1:ios 2:android
}


