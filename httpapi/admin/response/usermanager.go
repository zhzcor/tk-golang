package response

import "gserver/internal/store/ent"

//学员老师问答列表
type AskAnswerPageListResponse struct {
	List []AskAnswerDetail `json:"list"`
	Page PageResponse      `json:"page"`
}
type AskAnswerDetail struct {
	ent.UserAskAnswer
	Username          string             `json:"username"`
	TeacherName       string             `json:"teacher_name"`
	UserAvatarUrl     string             `json:"user_avatar_url"`
	AskAttachments    []AttachmentDetail `json:"ask_attachments"`
	AnswerAttachments []AttachmentDetail `json:"answer_attachments"`
}

type AttachmentDetail struct {
	AttachmentId  int    `json:"attachment_id"`
	AttachmentUrl string `json:"attachment_url"`
}

//学员评论列表
type UserCourseAppraisePageListResponse struct {
	List []UserCourseAppraiseDetail `json:"list"`
	Page PageResponse               `json:"page"`
}
type UserCourseAppraiseDetail struct {
	ent.UserCourseAppraise
	Username      string `json:"username"`
	SmallName     string `json:"small_name"`
	CourseName    string `json:"course_name"`
	Phone         string `json:"phone"`
	UserAvatarUrl string `json:"user_avatar_url"`
}
