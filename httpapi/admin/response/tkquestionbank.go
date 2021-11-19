package response

import (
	"gserver/internal/store/ent"
	"time"
)

type TkQuestionBank struct {
	ent.TkQuestionBank
	ItemCategoryName       string `json:"item_category_name"`
	ParentItemCategoryName string `json:"parent_item_category_name"`
	ParentItemCategoryId   int    `json:"parent_item_category_id"`
	CreatedAdminName       string `json:"created_admin_name"`
}

type TkBankStatistic struct {
	Total        int `json:"total"`
	EnableCount  int `json:"enable_count"`
	DisableCount int `json:"disable_count"`
}

//题库列表
type TkBankListSuccess struct {
	List      []TkQuestionBank `json:"list"`
	Page      PageResponse     `json:"page"`
	Statistic TkBankStatistic  `json:"statistic"`
}

type TkChapterSection struct {
	ent.TkChapter
	Sections ent.TkSections `json:"sections"`
}

//章节列表
type TkChapterSectionListSuccess struct {
	List []TkChapterSection `json:"list"`
	Page PageResponse       `json:"page"`
}

//题库知识点列表（分页）
type TkKnowledgePointPageListSuccess struct {
	List []TkKnowledgePointDetail `json:"list"`
	Page PageResponse             `json:"page"`
}
type TkKnowledgePointDetail struct {
	Id              int                   `json:"id"`
	Name            string                `json:"name"`
	QuestionCount   int                   `json:"question_count"`
	KnowledgePoints ent.TkKnowledgePoints `json:"knowledge_points"`
}

//知识点列表（分页）
type TkKnowledgePoints struct {
	List []TkPointDetail `json:"list"`
	Page PageResponse    `json:"page"`
}
type TkPointDetail struct {
	ent.TkKnowledgePoint
	QuestionBankName string `json:"question_bank_name"`
}

//题目列表（分页）
type TkQuestionListSuccess struct {
	List []TkListDetail `json:"list"`
	Page PageResponse   `json:"page"`
}
type TkListDetail struct {
	Id               int                         `json:"id"`
	Name             string                      `json:"name"`
	Type             int                         `json:"type"`
	Score            int                         `json:"score"`
	Difficulty       int                         `json:"difficulty"`
	CreatedAt        time.Time                   `json:"created_at"`
	CreatedAdminId   int                         `json:"created_admin_id"`
	CreatedAdminName string                      `json:"created_admin_name"`
	QuestionBankId   int                         `json:"question_bank_id"`
	QuestionBankName string                      `json:"question_bank_name"`
	KnowledgePoints  ent.TkKnowledgePoints       `json:"knowledge_points"`
	Options          ent.TkQuestionAnswerOptions `json:"options"`
}

//id获取题目详情
type TkQuestionByIdDetail struct {
	TkQuestionDetail
	Materials []TkQuestionDetail `json:"material_question"`
}
type TkQuestionDetail struct {
	TkListDetail
	Desc             string `json:"desc"`
	QuestionBankName string `json:"question_bank_name"`
}

//章节下题目详情
type TkCsQuestionPageList struct {
	List []TkCsQuestionDetail `json:"list"`
	Page PageResponse         `json:"page"`
}
type TkCsQuestionDetail struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	ChapterId   int    `json:"chapter_id"`
	ChapterName string `json:"chapter_name"`
	SectionId   int    `json:"section_id"`
	SectionName string `json:"section_name"`
}

//获取章节下已有题目id数组
type GetQuestionIdsInSection struct {
	QuestionIds []int `json:"question_ids"`
}

//获取题目纠错列表（带分页）
type GetQuestionErrorFeedbackListResponse struct {
	List []ErrorFeedbackDetail `json:"list"`
	Page PageResponse          `json:"page"`
}

type ErrorFeedbackDetail struct {
	QuestionId       int    `json:"question_id"`
	QuestionName     string `json:"question_name"`
	Type             int    `json:"type"`
	QuestionBankId   int    `json:"question_bank_id"`
	QuestionBankName string `json:"question_bank_name"`
	ent.TkQuestionErrorFeedback
}

//获取试卷列表（带分页）
type GetTkPaperExamListByPageResponse struct {
	List []TkPaperExamDetail `json:"list"`
	Page PageResponse        `json:"page"`
}

type TkPaperExamDetail struct {
	ent.TkExamPaper
	QuestionBankName string `json:"question_bank_name"`
	CreatedAdminName string `json:"created_admin_name"`
	ExamStatus       int    `json:"exam_status"`
}

//ID获取试卷详情
type GetTkExamPaperByIdResponse struct {
	TkPaperExamDetail
	Partitions []ExamPaperPartition `json:"partitions"`
}

type ExamPaperPartition struct {
	ent.TkExamPaperPartition
	FixedQuestions         []TkQuestionByIdDetail `json:"fixed_questions"` //固定卷题目
	PartitionTypeQuestions []ExamPartitionScore   `json:"partition_type_questions"`
}

type ExamPartitionScore struct {
	QuestionType  int `json:"question_type"`
	QuestionScore int `json:"question_score"`
	QuestionCount int `json:"question_count"`
}

//获取模拟考试成绩
type GetUserSimulationExamPageListResponse struct {
	List []UserSimulationExamDetail `json:"list"`
	Page PageResponse               `json:"page"`
}

type UserSimulationExamDetail struct {
	ent.TkUserExamScoreRecord
	UserId              int    `json:"user_id"`
	Username            string `json:"username"`
	Phone               string `json:"phone"`
	OperatorTeacherName string `json:"operator_teacher_name"`
}

//获取用户模拟考试主观题记录
type GetUserSubjectiveQuestionsDetail struct {
	UserExamRecordId int                                `json:"user_exam_record_id"`
	QuestionId       int                                `json:"question_id"`
	QuestionName     string                             `json:"question_name"`
	QuestionType     int                                `json:"question_type"`
	QuestionDesc     string                             `json:"question_desc"`
	Score            int                                `json:"score"`
	TeacherDesc      string                             `json:"teacher_desc"`
	TeacherScore     int                                `json:"teacher_score"`
	UserAnswer       string                             `json:"user_answer"`
	Options          ent.TkQuestionAnswerOptions        `json:"options"`
	Materials        []GetUserSubjectiveQuestionsDetail `json:"material_question"`
}

//获取题库下各题型题目数量
type BankQuestionTypeCount struct {
	QuestionBankId   int                           `json:"question_bank_id"`
	QuestionBankName string                        `json:"question_bank_name"`
	TotalCount       int                           `json:"total_count"`
	Detail           []BankQuestionTypeCountDetail `json:"detail"`
}
type BankQuestionTypeCountDetail struct {
	QuestionType  int `json:"question_type"`
	QuestionCount int `json:"question_count"`
}

type SimpleQuestionBank struct {
	IdSuccess
	Name string `json:"name"`
}

//获取做题记录统计列表
type GetQuestionBankStatisticListResponse struct {
	List []QuestionBankStatisticDetail `json:"list"`
	Page PageResponse                  `json:"page"`
}

type QuestionBankStatisticDetail struct {
	UserId           int     `json:"user_id"`
	Username         string  `json:"username"`
	QuestionBankId   int     `json:"question_bank_id"`
	QuestionBankName string  `json:"question_bank_name"`
	CateId           int     `json:"cate_id"`
	CateName         string  `json:"cate_name"`
	QuestionCount    int     `json:"question_count"`
	RecordCount      int     `json:"record_count"`
	CorrectRate      float64 `json:"correct_rate"`
	FinishRate       float64 `json:"finish_rate"`
}

//获取做题记录统计列表
type GetUserQuestionBankStatisticDetailListResponse struct {
	List []UserQuestionBankStatisticDetail `json:"list"`
	Page PageResponse                      `json:"page"`
}

type UserQuestionBankStatisticDetail struct {
	Id               int                         `json:"id"`
	UserId           int                         `json:"user_id"`
	QuestionBankId   int                         `json:"question_id"`
	QuestionId       int                         `json:"question_id"`
	QuestionName     string                      `json:"question_name"`
	UserAnswer       string                      `json:"user_answer"`
	IsRight          int                         `json:"is_right"`
	QuestionType     int                         `json:"question_type"`
	ExamQuestionType int                         `json:"exam_question_type"`
	CorrectRate      float64                     `json:"correct_rate"`
	Options          ent.TkQuestionAnswerOptions `json:"options"`
}

//学生数据
type GetStudentStatisticResponse struct {
	List []GetStudentStatisticDetail `json:"list"`
	Page PageResponse                `json:"page"`
}

type GetStudentStatisticDetail struct {
	Username           string  `json:"username"`
	Phone              string  `json:"phone"`
	UserId             int     `json:"user_id"`
	CateId             int     `json:"item_id"`
	CateName           string  `json:"cate_name"`
	CityId             int     `json:"city_id"`
	CityName           string  `json:"city_name"`
	AttendanceRate     float64 `json:"attendance_rate"`
	CourseFinishRate   float64 `json:"course_finish_rate"`
	QuestionFinishRate float64 `json:"question_finish_rate"`
}
