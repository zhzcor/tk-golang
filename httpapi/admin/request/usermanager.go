package request

import "time"

//用户问答列表
type UserAskAnswerPageList struct {
	TeacherId    *int `json:"teacher_id" form:"teacher_id" binding:"required"`
	AnswerStatus *int `json:"answer_status" form:"answer_status"`
	PageInfo
}

type ReplyUserAskAnswer struct {
	IdPtrOnly
	AnswerDesc    *string `json:"answer_desc" form:"answer_desc" binding:"required"`
	AttachmentIds []int   `json:"attachment_ids" form:"attachment_ids"`
}

//用户评论列表
type UserCourseAppraisePageList struct {
	Username *string    `json:"username" form:"username"`
	StartAt  *time.Time `json:"start_at" form:"start_at"`
	EndAt    *time.Time `json:"end_at" form:"end_at"`
	PageInfo
}

//课程评论回复
type ReplyUserCourseAppraise struct {
	IdPtrOnly
	TeacherReply *string `json:"teacher_reply" form:"teacher_reply" binding:"required"`
}
