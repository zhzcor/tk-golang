package response

import (
	"gserver/internal/store/ent"
	"time"
)

type CourseCalendar struct {
	Date  string `json:"date" `
	Ok    int    `json:"ok"`
	Today int    `json:"today"`
}

//模拟考试试卷
type CourseExam struct {
	Date   string `json:"date" `
	IsExam int    `json:"is_course"`
	ExamId int    `json:"exam_id"`
}

//学习首页日历
type CalendarList struct {
	Date     string `json:"date" `
	IsCourse int    `json:"is_course"`
	IsExam   int    `json:"is_exam"`
	Today    int    `json:"today"`
}

//已购课程或班级
type UserCourseList struct {
	Name          string `json:"name"`
	ClassId       int    `json:"class_id"`
	CourseId      int    `json:"course_id"`
	Img           string `json:"img"`
	CourseCount   int    `json:"course_count"`
	QuestionCount int    `json:"question_count"`
}

//模拟考试
type CourseMoNiKao struct {
	Name     string `json:"name"`
	TimeAt   string `json:"time_at"`
	Score    int    `json:"score"`
	Duration int    `json:"duration"`
}

//获取指定日期的直播直播课程
type GetDateLiveSmallCourse struct {
	LiveName       string    `json:"live_name"`
	LiveFinishType int       `json:"live_finish_type"` //完成条件：1：视频观看到最后，2：视频观看时长，3：音频收听到最后，4：音频收听时长，5：进入过直播间，6：下载过资料"
	LiveSmallId    int       `json:"live_small_id"`
	CourseImg      string    `json:"course_img"`
	LiveStarAt     string    `json:"live_start_at"`
	LiveStarAt2    string    `json:"live_start_at2"`
	LiveStarAt3    time.Time `json:"live_start_at3"`
	LiveStatus     int       `json:"live_status"`
	TeacherName    string    `json:"teacher_name"`
	TeacherAvatar  string    `json:"teacher_avatar"`
}

//模拟考试
/*type ExamQuestionTypeList struct {
	Id               int    `json:"id"`
	ExamName         string `json:"exam_name"`
	ExamQuestionType int    `json:"exam_question_type"`
	ExamType         int    `json:"exam_type"`
	Time             string `json:"time"`
	Score            int    `json:"scored"`
	Duration         int    `json:"duration"`
	IsOrder          int    `json:"is_order"`
}*/

type ExamQuestionTypeList struct {
	Id                 int    `json:"exam_id"`
	CourseName         string `json:"course_name"`
	BankId             int    `json:"bank_id" form:"bank_id"`
	TypeId             int    `json:"exam_question_type" form:"exam_question_type"`
	ExamName           string `json:"name" form:"name`
	ExamUserStudyCount int    `json:"exam_user_study_count" form:"exam_user_study_count`
	PracticeExamAt     int    `json:"practice_exam_at" form:"practice_exam_at"`
	/*	StarAt             string  `json:"star_at"`
		EndAt              string  `json:"end_at"`*/
	Time                 string    `json:"time"`
	StartAt              string    `json:"start_at"`
	StartAt2             time.Time `json:"start_at2"`
	Score                int       `json:"score"`
	QuestionCount        int       `json:"question_count"` //题目数量
	UserQuestionCount    int       `json:"user_question_count"`
	PassScore            int       `json:"pass_score"`             //通过分数
	ExamType             int       `json:"exam_type"`              //固定卷随机卷
	Accuracy             float64   `json:"accuracy"`               //正确率
	Difficulty           string    `json:"difficulty"`             //难易程度
	Desc                 string    `json:"desc"`                   //描述
	OrderStatus          int       `json:"order_status"`           //是否预约
	IsDo                 int       `json:"is_do"`                  //是否做了
	IsOrder              int       `json:"is_order"`               //是否预约
	SimulationExamStatus int       `json:"simulation_exam_status"` //模拟考试状态
}

//班级课程
type ClassCourseList struct {
	ClassName        string       `json:"class_name"`
	ClassId          int          `json:"class_id"`
	ClassCourseCount int          `json:"class_course_count"`
	CourseList       []CourseList `json:"course_list"`
}

type CourseList struct {
	CourseName    string  `json:"course_name"`
	TeacherName   string  `json:"teacher_name"`
	TeacherAvatar string  `json:"teacher_avatar"`
	ImgCover      string  `json:"img_cover"`
	CourseId      int     `json:"course_id"`
	ChapterCount  int     `json:"chapter_count"`
	SmallCourseCount   int     `json:"small_course_count"`
	SecCount      int     `json:"sec_count"`
	Learned       float64 `json:"learned"`
}

//课程详情目录+课时
type GetCourseSmallCateList struct {
	CourseId      int               `json:"course_id"`
	CourseName    string            `json:"course_name"`
	SmallCategory []SmallCateDetail `json:"small_category"`
	Chapters      []ChapterDetail   `json:"chapters"`
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

//目录
type ChapterDetail struct {
	ChapterId          int               `json:"chapter_id"`
	ChapterName        string            `json:"chapter_name"`
	SmallCategoryCount int               `json:"small_category_count"`
	SmallCategory      []SmallCateDetail `json:"small_category"`
	Sections           []SectionDetail   `json:"sections"`
}

type ChapterDetail2 struct {
	Type       int              `json:"type"`
	SmallType  int              `json:"small_type"` //1 直播 2 回放 3 点播
	LiveStatus int              `json:"live_status"`
	FileUrl    string           `json:"file_url"`
	LiveAt     string           `json:"live_at"`
	Name       string           `json:"name"`
	SmallId    int              `json:"small_id"`
	Count      int              `json:"count"`
	List       []ChapterDetail2 `json:"list"`
}

type SectionDetail struct {
	SectionId     int               `json:"section_id"`
	SectionName   string            `json:"section_name"`
	SmallCategory []SmallCateDetail `json:"small_category"`
}

//小节课程
type SmallDetail struct {
	Token       string `json:"token"`
	DomainUrl   string `json:"domain_url"`
	TeacherName string `json:"teacher_name"`
	Avatar      string `json:"avatar"`
	VideoId     int    `json:"video_id"`
	StudentCode string `json:"student_code"`
	PartnerKey  string `json:"partner_key"`
	SmallType   int    `json:"small_type"`  //1 直播 2 回放 3 点播
	LiveStatus  int    `json:"live_status"` // 1.未开始，2:直播中，3：直播结束，4：回放
	FileType    int    `json:"file_type"`   //1：视频，2：音频，3：直播，4：资料下载
	Detail
}

//课程小节详情
type Detail struct {
	Id              int    `json:"id"`
	CourseType      int    `json:"course_type"`
	SmallName       string `json:"small_name"`
	LiveRoomID      int    `json:"live_room_id"`
	ChapterID       int    `json:"chapter_id"`
	SectionID       int    `json:"section_id"`
	CourseID        int    `json:"course_id"`
	AttachmentCount int    `json:"attachment_count"`
	LiveSmallStart  string `json:"live_small_start"`
	LiveSmallTime   int    `json:"live_small_time"`
	LiveSmallStatus int    `json:"live_small_status"`
	QuestionBankId  int    `json:"question_bank_id"`
}

//课程试卷
type SmallCourseExam struct {
	ExamList      []QuestionBankExam `json:"exam_list"`
	LianxiList    []int              `json:"lianxi_list"`
	WorkList      []QuestionBankExam `json:"work_list"`
	MaterialsList []SmallMaterials   `json:"materials_list"`
}

//课程资料
type SmallMaterials struct {
	ent.Attachment
}

//老师问答老师列表
type TeacherAnswerList struct {
	ClassTeacher []TeacherDetail `json:"class_teacher"`
	TeacherList  []TeacherDetail `json:"teacher_list"`
}

type TeacherList struct {
	TeacherDetail
	CourseArr []string `json:"course_arr"`
}

type TeacherDetail struct {
	TeacherId   int      `json:"teacher_id"`
	Avatar      string   `json:"avatar"`
	Name        string   `json:"name"`
	Nickname    string   `json:"nickname"`
	Detail      string   `json:"detail"`
	TeachingAge int      `json:"teaching_age"`
	Majors      []string `json:"majors"`
	Tags        []string `json:"tags"`
	CourseArr   []string `json:"course_arr"`
}

type CourseTag struct {
	CourseName string `json:"course_name"`
}

//评价列表
type UserCourseAppraiseList struct {
	Id                   int      `json:"id"`
	SmallId              int      `json:"small_id"`
	CourseId             int      `json:"course_id"`
	UserId               int      `json:"user_id"`
	UserName             string   `json:"user_name"`
	Avatar               string   `json:"avatar"`
	TeacherImpression    []string `json:"teacher_impression"`     //印象标签
	CompositeScore       float64  `json:"composite_score"`        //综合分
	TeachAttitudeScore   float64  `json:"teach_attitude_score"`   //授课态度评分
	TeachContentScore    float64  `json:"teach_content_score"`    //授课内容评分
	TeachAtmosphereScore float64  `json:"teach_atmosphere_score"` //授课氛围评分
	Desc                 string   `json:"desc"`
	CourseName           string   `json:"course_name"`
	Time                 string   `json:"time"`
	/*	Page                 PageResponse `json:"page"`
	 */
}

//公开课
type LiveOpenCourses struct {
	LiveSmallId int `json:"live_small_id"`
}

type LiveDetail struct {
	LiveSmallId   int       `json:"live_small_id"`
	LiveAt        string    `json:"live_at"`
	LiveTime      time.Time `json:"live_time"`
	LiveDate      string    `json:"live_date"`
	LiveDateAt    string    `json:"live_date_at"`
	LiveStatus    int       `json:"live_status"`
	LiveName      string    `json:"live_name"`
	LiveDesc      string    `json:"live_desc"`
	TeacherName   string    `json:"teacher_name"`
	CourseName    string    `json:"course_name"`
	CourseId      int       `json:"course_id"`
	TeacherAvatar string    `json:"teacher_avatar"`
	PeopleNum     int       `json:"people_num"`
	ImgCover      string    `json:"img_cover"`
	IsEnroll      int       `json:"is_enroll"`
}

//推荐精品课程
type RecommendCourse struct {
	CourseId      int    `json:"course_id"`
	TeacherName   string `json:"teacher_name"`
	TeacherAvatar string `json:"teacher_avatar"`
	LiveStatus    int    `json:"live_status"`
	ImgCover      string `json:"img_cover"`
	SmallCount    int    `json:"small_count"`
	CourseName    string `json:"course_name"`
	CourseDesc    string `json:"course_desc"`
	StudyCount    int    `json:"study_count"`
}

//app首页
type AppIndexBanner struct {
	Img  string `json:"img"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

//首页课程
type AppIndexCourse struct {
	LiveAt        string `json:"live_at"`
	LiveSmallId   int    `json:"live_small_id"`
	CourseName    string `json:"course_name"`
	ImgCover      string `json:"img_cover"`
	TeacherAvatar string `json:"teacher_avatar"`
	LiveStatus    int    `json:"live_status"` // 1.未开始，2:直播中，3：直播结束，4：回放
	TeacherName   string `json:"teacher_name"`
	StudyCount    int    `json:"study_count"`
}

type AppIndexRecodeCourse struct {
	CourseId      int    `json:"course_id"`
	SmallCount    int    `json:"small_count"`
	CourseName    string `json:"course_name"`
	ImgCover      string `json:"img_cover"`
	TeacherAvatar string `json:"teacher_avatar"`
	TeacherName   string `json:"teacher_name"`
}

type AskMessage struct {
	TeacherName   string   `json:"teacher_name"`
	TeacherAvatar string   `json:"teacher_avatar"`
	TimeAt        string   `json:"time_at"`
	ReplyDesc     string   `json:"reply_desc"`
	ReplyImg      []string `json:"reply_img"`
	AnswerDesc    string   `json:"answer_desc"`
	AnswerImg     string   `json:"answer_img"`
}

type StyMessage struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Desc string `json:"desc"`
}

//课程观看记录
type VideoRecord struct {
	SmallType     int    `json:"small_type"` //1 直播 2 回放 3 点播
	CourseId      int    `json:"course_id"`
	LiveStatus    int    `json:"live_status"`
	SmallName     string `json:"small_name"`
	SmallId       int    `json:"small_id"`
	ImgCover      string `json:"img_cover"`
	TeacherName   string `json:"teacher_name"`
	TeacherAvatar string `json:"teacher_avatar"`
	ViewAt        string `json:"view_at"`
}
