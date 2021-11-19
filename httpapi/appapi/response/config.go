package response

type AuthOssResponse struct {
	Accessid     string `json:"accessid"`
	Callback     string `json:"callback"`
	Endpoint     string `json:"endpoint"`
	EndpointHost string `json:"endpoint_host"`
	Expire       int64  `json:"expire"`
	Host         string `json:"host"`
	Key          string `json:"key"`
	Policy       string `json:"policy"`
	Signature    string `json:"signature"`
	CallbackInfo struct {
		CallbackBody     string `json:"callback_body"`
		CallbackBodyType string `json:"callback_body_type"`
		CallbackUrl      string `json:"callback_url"`
	} `json:"callback_info"`
	Token struct {
		AccessKeyId     string `json:"access_key_id"`
		AccessKeySecret string `json:"access_key_secret"`
		Expiration      string `json:"expiration"`
		SecurityToken   string `json:"security_token"`
	} `json:"token"`
}

type GetAgreementResponse struct {
	Type   int    `json:"type"`
	Detail string `json:"detail"`
}

type AppVersion struct {
	Id            int    `json:"id" form:"id"`
	Name          string `json:"name" form:"name" `
	Sn            string `json:"sn" form:"sn" `
	Url           string `json:"url" form:"url" `
	Remark        string `json:"remark" form:"remark" `
	PhoneType     int    `json:"phone_type" form:"phone_type"`
	IsForceUpdate int    `json:"is_force_update" form:"is_force_update"`
}