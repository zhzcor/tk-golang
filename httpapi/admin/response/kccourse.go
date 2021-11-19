package response

import (
	"time"
	"tkserver/internal/store/ent"
)

//课程列表（分页）
type CoursePageListResponse struct {
	List      []CoursePageDetail `json:"list"`
	Page      PageResponse       `json:"page"`
	Statistic CourseStatistic    `json:"statistic"`
}
type CoursePageDetail struct {
	ItemCategoryName string `json:"item_category_name"`
	CreatedAdminName string `json:"created_admin_name"`
	ItemId           int    `json:"item_id"`
	ent.KcCourse
}

type CourseSimple struct {
	Id         int    `json:"id"`
	CourseName string `json:"course_name"`
}

type CourseStatistic struct {
	Total       int `json:"total"`
	Published   int `json:"published"`
	Unpublished int `json:"unpublished"`
	Closed      int `json:"closed"`
}

//课程老师列表
type CourseTeacherPageListResponse struct {
	List []CourseTeacherDetail `json:"list"`
	Page PageResponse          `json:"page"`
}

type CourseTeacherDetail struct {
	TeacherName string `json:"teacher_name"`
	Avatar      string `json:"avatar"`
	ent.KcCourseTeacher
}

//课程用户列表
type CourseUserPageListResponse struct {
	List []CourseUserDetail `json:"list"`
	Page PageResponse       `json:"page"`
}

type CourseUserDetail struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	ent.KcUserCourse
}

type UserDetail struct {
	ent.User
	ItemCategoryName       string `json:"item_category_name"`
	CityName               string `json:"city_name"`
	ParentItemCategoryId   int    `json:"parent_item_category_id"`
	ParentItemCategoryName string `json:"parent_item_category_name"`
}

type CourseQuestionBank struct {
	ItemCategoryName       string `json:"item_category_name"`
	ParentItemCategoryName string `json:"parent_item_category_name"`
	QuestionBankName       string `json:"question_bank_name"`
	QuestionBankId         int    `json:"question_bank_id"`
	CourseId               int    `json:"course_id"`
}

//获取课程下所有章节课时（无分页）
type GetCourseSmallCateList struct {
	CourseId      int               `json:"course_id"`
	CourseName    string            `json:"course_name"`
	SmallCategory []SmallCateDetail `json:"small_category"`
	Chapters      []ChapterDetail   `json:"chapters"`
	SmallCount    int               `json:"small_count"`
}

type SmallCateDetail struct {
	Id              int    `json:"id"`
	SmallName       string `json:"small_name"`
	Duration        int    `json:"duration"`
	Status          int    `json:"status"`
	AttachmentCount int    `json:"attachment_count"`
	QuestionCount   int    `json:"question_count"`
	HomeworkCount   int    `json:"homework_count"`
	ExamCount       int    `json:"exam_count"`
	Type            int    `json:"type"`
}

type ChapterDetail struct {
	ChapterId     int               `json:"chapter_id"`
	ChapterName   string            `json:"chapter_name"`
	SmallCategory []SmallCateDetail `json:"small_category"`
	Sections      []SectionDetail   `json:"sections"`
}

type SectionDetail struct {
	SectionId     int               `json:"section_id"`
	SectionName   string            `json:"section_name"`
	SmallCategory []SmallCateDetail `json:"small_category"`
}

type CourseDetail struct {
	IdSuccess
	CourseName        string     `json:"course_name"`
	CourseType        int        `json:"course_type"`
	Majors            ent.Majors `json:"majors"`
	ItemId            int        `json:"item_id"`
	ItemName          string     `json:"item_name"`
	CityId            int        `json:"city_id"`
	CityName          string     `json:"city_name"`
	CourseCoverImgId  int        `json:"course_cover_img_id"`
	CourseCoverImgUrl string     `json:"course_cover_img_url"`
	CourseDesc        string     `json:"course_desc"`
	PushStatus        int        `json:"push_status"`
	CoursePrice       float64    `json:"course_price"`
	CreatedAt         time.Time  `json:"created_at"`
	OpenLiveStart     time.Time  `json:"open_live_start"`
	OpenLiveDuration  int        `json:"open_live_duration"`
	PeopleNum         int        `json:"people_num"`
}

//课时试卷、作业列表（分页）
type SmallCourseExamPageListResponse struct {
	List []SmallCourseExamPageDetail `json:"list"`
	Page PageResponse                `json:"page"`
}

type SmallCourseExamPageDetail struct {
	ExamPaperName    string `json:"exam_paper_name"`
	ExamPaperId      int    `json:"exam_paper_id"`
	ExamPaperType    int    `json:"exam_paper_type"`
	QuestionBankId   int    `json:"question_bank_id"`
	QuestionBankName string `json:"question_bank_name"`
	IsSelected       int    `json:"is_selected"`
}

//课时练习列表（分页）
type SmallCourseQuestionPageListResponse struct {
	List []SmallCourseQuestionPageDetail `json:"list"`
	Page PageResponse                    `json:"page"`
}

type SmallCourseQuestionPageDetail struct {
	QuestionName     string `json:"question_name"`
	QuestionId       int    `json:"question_id"`
	QuestionType     int    `json:"question_type"`
	QuestionBankId   int    `json:"question_bank_id"`
	QuestionBankName string `json:"question_bank_name"`
	IsSelected       int    `json:"is_selected"`
}

//课时资料列表（分页）
type SmallCourseAttachmentPageListResponse struct {
	List []SmallCourseAttachmentPageDetail `json:"list"`
	Page PageResponse                      `json:"page"`
}

type SmallCourseAttachmentPageDetail struct {
	AttachmentId   int    `json:"attachment_id"`
	AttachmentName string `json:"attachment_name"`
	AttachmentUrl  string `json:"attachment_url"`
}

//获取课程已添加teacher id列表
type GetTeacherIdsInCourse struct {
	TeacherIds []int `json:"teacher_ids"`
}

//获取课程已添加user id列表
type GetUserIdsInCourse struct {
	UserIds []int `json:"User_ids"`
}

//课程伪直播列表（分页）
type GetFalseVideoPageListResponse struct {
	List ent.KcVideoUploadTasks `json:"list"`
	Page PageResponse           `json:"page"`
}

//课程回放列表（分页）
type GetLiveBackPageListResponse struct {
	List ent.KcCourseSmallCategories `json:"list"`
	Page PageResponse                `json:"page"`
}

//课程直播列表
type GetLiveSmallCatePageList struct {
	List []LiveSmallCateDetail `json:"list"`
	Page PageResponse          `json:"page"`
}

type LiveSmallCateDetail struct {
	IdSuccess
	SmallName      string    `json:"small_name"`
	LiveSmallStart time.Time `json:"live_small_start"`
	LiveSmallTime  int       `json:"live_small_time"`
	CourseId       int       `json:"course_id"`
	CourseName     string    `json:"course_name"`
	Status         int       `json:"status"`
	StudyCount     int       `json:"study_count"`
}

type BjyCallbackResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CourseSmallCategoryById struct {
	Id       int `json:"id"`
	CourseId int `json:"course_id"`
	//CourseType          int       `json:"course_type"`
	ChapterId           int       `json:"chapter_id"`
	SectionId           int       `json:"section_id"`
	SmallName           string    `json:"small_name"`
	Type                int       `json:"type"`
	LiveSmallTime       int       `json:"live_small_time"`
	FinishType          int       `json:"finish_type"`
	ViewingTime         int       `json:"viewing_time"`
	FalseVideoId        int       `json:"false_video_id"`
	FalseVideoName      string    `json:"false_video_name"`
	OrderVideoAttachId  int       `json:"order_video_attach_id"`
	OrderVideoName      string    `json:"order_video_name"`
	CoursewareAttachId  int       `json:"courseware_attach_id"`
	CoursewareAttachUrl string    `json:"courseware_attach_url"`
	CoursewareName      string    `json:"courseware_name"`
	LiveSmallStart      time.Time `json:"live_small_start"`
	LiveSmallRemark     string    `json:"live_small_remark"`
}
