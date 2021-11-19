package request

import "time"

//推广管理

//添加(编辑)广告
type SetAdvertise struct {
	Id           *int       `json:"id"  form:"id"`
	Name         *string    `json:"name"  form:"name" binding:"required"`
	Status       *int       `json:"status" form:"status" binding:"required"`
	Position     *int       `json:"position" form:"position" binding:"required"`
	StartAt      *time.Time `json:"start_at" form:"start_at" binding:"required"`
	EndAt        *time.Time `json:"end_at" form:"end_at" binding:"required"`
	AttachmentId *int       `json:"attachment_id" form:"attachment_id" binding:"required"`
	AdUrl        *string    `json:"ad_url" form:"ad_url" form:"required"`
	Remark       string     `json:"remark" form:"remark"`
	SortOrder    int        `json:"sort_order" form:"sort_order"`
}

//广告分页查询
type AdPageListRequest struct {
	Name     *string `json:"name"  form:"name"`
	Status   *int    `json:"status" form:"status"`
	Position *int    `json:"position" form:"position"`
	PageInfo
}

//添加(编辑)分享海报
type SetSharePostRequest struct {
	Id           *int    `json:"id"  form:"id"`
	Name         *string `json:"name"  form:"name" binding:"required"`
	Status       *int    `json:"status" form:"status" binding:"required"`
	SortOrder    *int    `json:"sort_order" form:"sort_order" binding:"required"`
	Url          *string `json:"url" form:"url"`
	Remark       *string `json:"remark" form:"remark"`
	AttachmentId *int    `json:"attachment_id" form:"attachment_id" binding:"required"`
}

//分享海报分页查询
type SharePosterPageListRequest struct {
	Status *int `json:"status" form:"status"`
	PageInfo
}

//设置消息公告
type SetMessageRequest struct {
	IdPtrNillable
	NamePtrOnly
	MessageTypeId *int       `json:"message_type_id" form:"message_type_id" binding:"required"`
	AttachmentId  *int       `json:"attachment_id" form:"attachment_id"`
	Detail        *string    `json:"detail" form:"detail" binding:"required"`
	PublishType   *int       `json:"publish_type" form:"publish_type" binding:"required"`
	AutoPublishAt *time.Time `json:"auto_publish_at" form:"auto_publish_at"`
	CourseId      *int       `json:"course_id" form:"course_id"`
	ClassId       *int       `json:"class_id" form:"class_id"`
}

//消息公告分页查询
type MessagePageListRequest struct {
	Status        *int    `json:"status" form:"status"`
	Name          *string `json:"name" form:"name"`
	MessageTypeId *int    `json:"message_type_id" form:"message_type_id"`
	CourseId      *int    `json:"course_id" form:"course_id"`
	ClassId       *int    `json:"class_id" form:"class_id"`
	PageInfo
}
