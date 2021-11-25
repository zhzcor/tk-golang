package request

import "time"

//添加(编辑)标签（地区项目专业）
type SetBasicTag struct {
	Id        int    `json:"id"  form:"id"`
	Name      string `json:"name"  form:"name" binding:"required"`
	Code      string `json:"code" form:"code"`
	SortOrder int    `json:"sort_order" form:"sort_order"`
	Desc      string `json:"desc" form:"desc"`
}

//添加(编辑)群名片
type SetGroupCard struct {
	Id           int    `json:"id"  form:"id"`
	Title        string `json:"title"  form:"title" binding:"required"`
	SubTitle     string `json:"sub_title" form:"sub_title" binding:"required"`
	SortOrder    int    `json:"sort_order" form:"sort_order"`
	Status       int    `json:"status" form:"status" binding:"required"`
	AttachmentId int    `json:"attachment_id" form:"attachment_id" binding:"required,gte=0"`
}

//type SetItemCategory struct {
//	SetBasicTag
//	//Pid int `json:"pid" form:"pid"`
//}

//基础标签分页查询
type BasicTagPageList struct {
	Name   string `json:"name"  form:"name"`
	Code   string `json:"code" form:"code"`
	Status *int   `json:"status" form:"status"`
	PageInfo
}

type BasicConfigStatus struct {
	IdOnly
	Status int `json:"status" form:"status" binding:"required,min=1,max=2"`
}

//oss
type AuthOssRequest struct {
	Action *string   `json:"action" form:"action" binding:"required"`
	Media  *OssMedia `json:"media" form:"media"`
}
type OssMedia struct {
	AdminId int `json:"admin_id" form:"admin_id"`
}

//app协议
type SetAgreement struct {
	//Name   *string `json:"name" form:"name" binding:"required"`
	Type   *int    `json:"type" form:"type" binding:"required,min=1,max=11"`
	Detail *string `json:"detail" form:"detail" binding:"required"`
}

//获取协议详情
type GetAgreement struct {
	Type *int `json:"type" form:"type" binding:"required,min=1,max=11"`
}

//添加（编辑）app版本
type SetAppVersion struct {
	Id            *int    `json:"id" form:"id"`
	Name          *string `json:"name" form:"name" binding:"required"`
	Sn            *string `json:"sn" form:"sn" binding:"required"`
	Url           *string `json:"url" form:"url" binding:"required"`
	Remark        *string `json:"remark" form:"remark" binding:"required"`
	PhoneType     *int    `json:"phone_type" form:"phone_type" binding:"required"`
	IsForceUpdate *int    `json:"is_force_update" form:"is_force_update" binding:"required"`
}

//app版本列表
type GetAppVersionPageList struct {
	Name      *string `json:"name" form:"name"`
	Sn        *string `json:"sn" form:"sn"`
	PhoneType *int    `json:"phone_type" form:"phone_type"`
	PageInfo
}

//登录日志
type GetLoginLogPageList struct {
	PageInfo
}

//登录日志
type GetOperateLogPageList struct {
	AdminName *string    `json:"admin_name" form:"admin"`
	StartAt   *time.Time `json:"start_at" form:"start_at"`
	EndAt     *time.Time `json:"end_at" form:"end_at"`
	PageInfo
}

//登录日志
type GetLoginUserPageList struct {
	Username *string    `json:"username" form:"username"`
	StartAt  *time.Time `json:"start_at" form:"start_at"`
	EndAt    *time.Time `json:"end_at" form:"end_at"`
	PageInfo
}

//oss回调
type OssCallBack struct {
	AdminId  int    `json:"admin_id" form:"admin_id"`
	UserId   int    `json:"user_id" form:"user_id"`
	Size     int    `json:"size" form:"size"`
	MimeType string `json:"mimeType" form:"mimeType"`
	Filename string `json:"filename" form:"filename"`
}
