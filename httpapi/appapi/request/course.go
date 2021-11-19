package request

type OnlySmallId struct {
	SmallId int `json:"small_id" form:"small_id" binding:"required"`
}

type OnlyCalendar struct {
	Date string `json:"date" form:"date" `
}

type CourseCalendar struct {
	Date string `json:"date" form:"date" `
	Type int    `json:"type" form:"type" binding:"required"` //1是直播课程 2为试卷
}

type ClassCourseList struct {
	ClassId int `json:"class_id" form:"class_id" binding:"required"`
}

type CourseDetailsById struct {
	CourseId int `json:"course_id" form:"course_id" binding:"required"`
}

//小节课程详情
type CourseSmallDetail struct {
	OnlySmallId
}

//随堂练习
type CourseSmallPracticeInClass struct {
	OnlySmallId
}

//学员老师问答
type UserTeacherAsk struct {
	TeacherId int    `json:"teacher_id" form:"teacher_id" binding:"required"`
	AskDesc   string `json:"ask_desc" form:"ask_desc"`
	ImgIds    []int  `json:"img_ids" form:"img_ids[]"`
}

//用户评价
type AddUserCourseAppraise struct {
	SmallId              int     `json:"small_id" form:"small_id" binding:"required"`
	CourseId             int     `json:"course_id" form:"course_id" binding:"required"`
	TeachAttitudeScore   float64 `json:"teach_attitude_score" form:"teach_attitude_score" `
	TeachContentScore    float64 `json:"teach_content_score" form:"teach_content_score"`
	TeachAtmosphereScore float64 `json:"teach_atmosphere_score" form:"teach_atmosphere_score"`
	CompositeScore       float64 `json:"composite_score" form:"composite_score"`
	TeacherImpression    string  `json:"teacher_impression" form:"teacher_impression"`
	Desc                 string  `json:"desc" form:"desc"`
	Type                 int     `json:"type" form:"type" binding:"required"` //评价类型：1:普通课程，2:直播课程，3:直播公开课，4:录播公开课
}

type GetUserCourseAppraise struct {
	CourseId int `json:"course_id" form:"course_id" binding:"required"`
	SmallId  int `json:"small_id" form:"small_id" binding:"required"`
	PageInfo
}

type PageInfo struct {
	Page int `json:"current"  form:"current"`
	/*	PageSize int `json:"page_size"  form:"page_size" binding:"eq=10"`
	 */
}

//消息
type MessageAskType struct {
	MessType
	PageInfo
}

type MessType struct {
	Type int `json:"type" form:"type" binding:"required"`
}

//视频观看记录
type AddVideoRecord struct {
	SmallId   int    `json:"small_id" form:"small_id" binding:"required"`
	Name      string `json:"name" form:"name"`
	VideoTime int    `json:"video_time" form:"video_time"`
}
