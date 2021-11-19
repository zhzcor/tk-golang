package response

import "tkserver/internal/store/ent"

type AdPageListResponse struct {
	List []AdPageDetail `json:"list"`
	Page PageResponse   `json:"page"`
}
type AdPageDetail struct {
	ent.Advertise
	AttachmentName string `json:"attachment_name"`
}

type SharePosterPageListResponse struct {
	List []SharePosterDetail `json:"list"`
	Page PageResponse        `json:"page"`
}

type SharePosterDetail struct {
	ent.SharePoster
	ImgFileName string `json:"img_file_name"`
}

type MessagePageList struct {
	List []MessageDetail `json:"list"`
	Page PageResponse    `json:"page"`
}

type MessageDetail struct {
	ent.Message
	MessageTypeName  string `json:"message_type_name"`
	CreatedAdminName string `json:"created_admin_name"`
	AttachmentName   string `json:"attachment_name"`
}
