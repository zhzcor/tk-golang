package request

import "time"

//添加(编辑)课程
type SetCourse struct {
	Id         *int    `json:"id"  form:"id"`
	CourseName *string `json:"course_name"  form:"course_name" binding:"required"`
	//Status            *int       `json:"status" form:"status" binding:"required"`
	CourseType        *int     `json:"course_type" form:"course_type" binding:"required"`
	MajorIds          []int    `json:"major_ids" form:"major_ids"`
	ItemId            *int     `json:"item_id" form:"item_id"`
	CityId            *int     `json:"city_id" form:"city_id"`
	CoverAttachmentId *int     `json:"cover_attachment_id" form:"cover_attachment_id"`
	CourseDesc        *string  `json:"course_desc" form:"course_desc"`
	CoursePrice       *float64 `json:"course_price" form:"course_price"`
	PeriodType        *int     `json:"periods_type" form:"periods_type"`
	//StartDate         *time.Time `json:"start_date" form:"start_date"`
	//EndDate           *time.Time `json:"end_date" form:"end_date"`
	//ClosingDate       *time.Time `json:"closing_date" form:"closing_date"`
	//DaysValidity      *int       `json:"days_validity" form:"days_validity"`
	OpenLiveStart    *time.Time `json:"open_live_start" form:"open_live_start"`
	OpenLiveDuration *int       `json:"open_live_duration" form:"open_live_duration"`
}

//课程列表（分页）
type CoursePageListRequest struct {
	CourseName       *string `json:"course_name" form:"course_name"`
	CreatedAdminName *string `json:"created_admin_name" form:"created_admin_name"`
	PushStatus       *int    `json:"push_status" form:"push_status"`
	CourseType       []uint8 `json:"course_type" form:"course_type"`
	CateId           *int    `json:"cate_id" form:"cate_id"`
	PageInfo
}

type CourseIdRequest struct {
	CourseId *int `json:"course_id" form:"course_id" binding:"required"`
	PageInfo
}

//设置课程状态
type SetCourseStatus struct {
	Id         *int `json:"id"  form:"id" binding:"required"`
	PushStatus *int `json:"push_status" form:"push_status" binding:"required,min=1,max=3"`
}

//设置课程老师
type SetCourseTeacher struct {
	CourseId   *int `json:"course_id"  form:"course_id" binding:"required"`
	TeacherId  *int `json:"teacher_id" form:"teacher_id" binding:"required"`
	ShowStatus *int `json:"show_status" form:"show_status"`
	SortOrder  *int `json:"sortOrder" form:"sort_order"`
}

//设置课程学员
type SetCourseUser struct {
	CourseId *int `json:"course_id"  form:"course_id" binding:"required"`
	UserId   *int `json:"user_id" form:"user_id" binding:"required"`
}

//修改学员有效期
type UpdateCourseUserValidity struct {
	IdPtrOnly
	PeriodType *int `json:"period_type" form:"period_type" binding:"required"`
}

type GetCourseTeacherPageList struct {
	IdPtrOnly
	PageInfo
}

//课程用户列表
type GetCourseUserPageList struct {
	IdPtrOnly
	PageInfo
	Username *string `json:"username" form:"username"`
	Phone    *string `json:"phone" form:"phone"`
}

//设置学员有效期
type SetCourseUserValidity struct {
	UserId      *int       `json:"user_id" form:"user_id" binding:"required"`
	CourseId    *int       `json:"course_id" form:"course_id" binding:"required"`
	ClosingDate *time.Time `json:"closing_date" form:"closing_date" form:"required"`
}

//课程添加题库
type SetCourseQuestionBank struct {
	QuestionBankId *int `json:"question_bank_id" form:"question_bank_id"`
	CourseId       *int `json:"course_id" form:"course_id" binding:"required"`
}

//添加章
type SetCourseChapter struct {
	IdPtrNillable
	Title    *string `json:"title" form:"title" binding:"required"`
	CourseId *int    `json:"course_id" form:"course_id" binding:"required,gte=0"`
}

//添加节
type SetCourseSection struct {
	IdPtrNillable
	Title     *string `json:"title" form:"title" binding:"required"`
	ChapterId *int    `json:"chapter_id" form:"chapter_id" binding:"required,gte=0"`
}

//添加（编辑）课时
type SetCourseSmallCategory struct {
	Id                 *int       `json:"id" form:"id"`
	CourseId           *int       `json:"course_id" form:"course_id" binding:"required"`
	CourseType         *int       `json:"course_type" form:"course_type" binding:"required"`
	ChapterId          *int       `json:"chapter_id" form:"chapter_id"`
	SectionId          *int       `json:"section_id" form:"section_id"`
	SmallName          *string    `json:"small_name" form:"small_name" binding:"required"`
	Type               *int       `json:"type" form:"type" binding:"required"`
	LiveSmallTime      *int       `json:"live_small_time" form:"live_small_time"`
	FinishType         int        `json:"finish_type" form:"finish_type"`
	ViewingTime        *int       `json:"viewing_time" form:"viewing_time"`
	FalseVideoId       *int       `json:"false_video_id" form:"false_video_id"`
	OrderVideoAttachId *int       `json:"order_video_attach_id" form:"order_video_attach_id"`
	OrderVideoName     string     `json:"order_video_name" form:"order_video_name"`
	CoursewareAttachId *int       `json:"courseware_attach_id" form:"courseware_attach_id"`
	CoursewareName     *string    `json:"courseware_name" form:"courseware_name"`
	LiveSmallStart     *time.Time `json:"live_small_start" form:"live_small_start"`
	LiveSmallRemark    *string    `json:"live_small_remark" form:"live_small_remark"`
}

//添加课时（试卷）
type SetSmallCateExamPagerRequest struct {
	SmallCategoryId *int   `json:"small_category_id" form:"small_category_id" binding:"required"`
	Type            *int   `json:"type" form:"type" binding:"required,min=1,max=2"`
	ExamPaperIds    *[]int `json:"exam_paper_ids" form:"exam_paper_ids" binding:"required"`
}

//添加课时练习
type SetSmallCateQuestionRequest struct {
	SmallCategoryId *int   `json:"small_category_id" form:"small_category_id" binding:"required"`
	QuestionIds     *[]int `json:"question_ids" form:"question_ids" binding:"required"`
}

//添加课时附件
type SetSmallCateAttachmentRequest struct {
	SmallCategoryId *int                  `json:"small_category_id" form:"small_category_id" binding:"required"`
	Attachments     []SmallCateAttachment `json:"Attachments" form:"Attachments"`
}

type SmallCateAttachment struct {
	AttachmentId   *int    `json:"attachment_id" form:"attachment_id" binding:"required"`
	AttachmentName *string `json:"attachment_name" form:"attachment_name" binding:"required"`
}

//移除课时（试卷）
type RemoveSmallCateExamPagerRequest struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	Type            *int `json:"type" form:"type" binding:"required,min=1,max=2"`
	ExamPaperId     *int `json:"exam_paper_id" form:"exam_paper_id" binding:"required"`
}

//移除课时练习
type RemoveSmallCateQuestionRequest struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	QuestionId      *int `json:"question_id" form:"question_id" binding:"required"`
}

//移除课时附件
type RemoveSmallCateAttachmentRequest struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	AttachmentId    *int `json:"attachment_id" form:"attachment_id" binding:"required"`
}

//设置课程状态
type SetCourseTeacherStatus struct {
	IdPtrOnly
	ShowStatus *int `json:"show_status" form:"show_status" binding:"required,min=1,max=2"`
}

//课时试卷、作业列表（分页）
type SmallCourseExamPageList struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	Type            *int `json:"type" form:"type" binding:"required,min=1,max=2"`
	QuestionBankId  *int `json:"question_bank_id" form:"question_bank_id"`
	ExamPaperType   *int `json:"exam_paper_type" form:"exam_paper_type"`
	PageInfo
}

//已选中 课时试卷、作业列表（分页）
type SelectedSmallCourseExamPageList struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	Type            *int `json:"type" form:"type" binding:"required,min=1,max=2"`
	PageInfo
}

//课时练习列表（分页）
type SmallCourseQuestionPageList struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	QuestionBankId  *int `json:"question_bank_id" form:"question_bank_id"`
	QuestionType    *int `json:"question_type" form:"question_type"`
	PageInfo
}

//已选中 课时资料列表（分页）
type SelectedSmallCourseAttachmentPageList struct {
	SmallCategoryId *int `json:"small_category_id" form:"small_category_id" binding:"required"`
	PageInfo
}

//上传伪直播视频
type SetFalseVideoRequest struct {
	IdPtrNillable
	AttachmentId *int    `json:"attachment_id" form:"attachment_id" binding:"required"`
	CourseId     *int    `json:"course_id" form:"course_id" binding:"required"`
	Title        *string `json:"title" form:"title" binding:"required"`
	VideoName    *string `json:"video_name" form:"video_name" binding:"required"`
	Length       *int    `json:"length" form:"length"`
}

//课程伪直播列表
type GetFalseVideoPageList struct {
	CourseIdRequest
	Title  *string `json:"title" form:"title"`
	Status *int    `json:"status" form:"status"`
	PageInfo
}

type ReplaceLiveBack struct {
	IdPtrOnly
	AttachmentId *int   `json:"attachment_id" form:"attachment_id" binding:"required"`
	VideoName    string `json:"video_name" form:"video_name" binding:"required"`
}

//课程直播列表
type GetLiveSmallCatePageList struct {
	LiveStartAt *time.Time `json:"live_start_at" form:"live_start_at"`
	LiveEndAt   *time.Time `json:"live_end_at" form:"live_end_at"`
	SmallName   *string    `json:"small_name" form:"small_name"`
	Status      []int8     `json:"status" form:"status"`
	PageInfo
}

//上下课回调
type ClassCallbackRequest struct {
	RoomId    int    `json:"room_id" form:"room_id" binding:"required"`
	Op        string `json:"op" form:"op" binding:"required"`
	OpTime    string `json:"op_time" form:"op_time"`
	Qid       string `json:"qid" form:"qid"`
	Timestamp int    `json:"timestamp" form:"timestamp" binding:"required"`
	Sign      string `json:"sign" form:"sign" binding:"required"`
}

//点播，回放处理
type VideoCallbackRequest struct {
	RoomId     *int   `json:"room_id" form:"room_id"`
	VideoId    int    `json:"video_id" form:"video_id"`
	Status     int    `json:"status" form:"status"`
	PrefaceUrl string `json:"preface_url" form:"preface_url"`
	TotalSize  int    `json:"total_size" form:"total_size"`
	Length     int    `json:"length" form:"length"`
	Qid        string `json:"qid" form:"qid"`
	Timestamp  int    `json:"timestamp" form:"timestamp" binding:"required"`
	Sign       string `json:"sign" form:"sign" binding:"required"`
}
