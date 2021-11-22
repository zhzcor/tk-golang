package response

import "tkserver/internal/store/ent"

type CityListSuccess struct {
	List ent.Cities   `json:"list"`
	Page PageResponse `json:"page"`
}

type LevelListSuccess struct {
	List ent.Levels   `json:"list"`
	Page PageResponse `json:"page"`
}

type ItemCategoryListSuccess struct {
	List []ItemDetail `json:"list"`
	Page PageResponse `json:"page"`
}

type ItemDetail struct {
	ent.ItemCategory
	Children ent.ItemCategories
}

type MajorListSuccess struct {
	List ent.Majors   `json:"list"`
	Page PageResponse `json:"page"`
}

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

type AppVersionPageListSuccess struct {
	List ent.AppVersions `json:"list"`
	Page PageResponse    `json:"page"`
}

//登录日志
type GetLoginLogPageListResponse struct {
	List ent.AdminLoginLogs `json:"list"`
	Page PageResponse       `json:"page"`
}

type GetAgreementResponse struct {
	Type   int    `json:"type"`
	Detail string `json:"detail"`
}

type IndexStatistic struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}
