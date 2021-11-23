package response

import (
	"time"
	"tkserver/internal/store/ent"
	"tkserver/pkg/htmltojson"
)

//题库试卷类型接口
type QuestionBankExam struct {
	Id                 int    `json:"exam_id"`
	BankId             int    `json:"bank_id" form:"bank_id"`
	TypeId             int    `json:"exam_question_type" form:"exam_question_type"`
	ExamName           string `json:"name" form:"name`
	ExamUserStudyCount int    `json:"exam_user_study_count" form:"exam_user_study_count`
	PracticeExamAt     int    `json:"practice_exam_at" form:"practice_exam_at"`
	/*	StarAt             string  `json:"star_at"`
		EndAt              string  `json:"end_at"`*/
	Time                 string  `json:"time"`
	Score                int     `json:"score"`
	QuestionCount        int     `json:"question_count"` //题目数量
	UserQuestionCount    int     `json:"user_question_count"`
	PassScore            int     `json:"pass_score"`             //通过分数
	ExamType             int     `json:"exam_type"`              //固定卷随机卷
	Accuracy             float64 `json:"accuracy"`               //正确率
	Difficulty           string  `json:"difficulty"`             //难易程度
	Desc                 string  `json:"desc"`                   //描述
	OrderStatus          int     `json:"order_status"`           //是否预约
	IsDo                 int     `json:"is_do"`                  //是否做了
	SimulationExamStatus int     `json:"simulation_exam_status"` //模拟考试状态
}

type GetSimulationExamStatus struct {
	SimulationExamStatus int `json:"simulation_exam_status"` //模拟考试状态
	SurplusTime          int `json:"surplus_time"`
}

//题库首页
type QuestionIndex struct {
	ClassId  int `json:"class_id" form :"class_id"`
	CourseId int `json:"course_id" form:"course_id"`
}

type QuestionMajorCheck struct {
	LevelList        []LevelList
	MajorList        []MajorList
	ItemCategoryList []ItemCategoryList
}

//level
type LevelList struct {
	LeveId   int    `json:"leve_id" form:"leve_id"`
	LeveName string `json:"leve_name" form:"leve_name"`
}

type MajorList struct {
	MajorId   int    `json:"major_id" form:"major_id"`
	MajorName string `json:"major_name" form:"major_name"`
}

type ItemCategoryList struct {
	ItemId   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

type BankListInfo struct {
	QuestionBankName string `json:"question_bank_name"`
	QuestionBankId   int    `json:"question_bank_id"`
}

type GroupCardList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	CodeUrl string `json:"code_url"`
}

//课程下题库接口
type CourseQuestionBank struct {
	QuestionId int `json:"question_id" form:"question_id"`
}

//章
type QuestionBankChapter struct {
	ChapterId   int    `json:"chapter_id"`
	ChapTerName string `json:"chap_ter_name"`
	SecList     []QuestionBankSection
}

//节
type QuestionBankSection struct {
	SecId             int    `json:"sec_id"`
	SecName           string `json:"sec_name"`
	QuestionCount     int    `json:"question_count"`
	UserQuestionCount int    `json:"user_question_count"`
	IsDo              int    `json:"is_do"`
	QuestionIds       []int  `json:"question_ids"`
}

//ID获取试卷详情
type GetTkExamPaperByIdResponse struct {
	/*TkPaperExamDetail
	DifficultyDescString string               `json:"difficulty_desc_string"`*/
	Partitions []ExamPaperPartition `json:"partitions"`
}

type ExamPaperPartition struct {
	ent.TkExamPaperPartition
	PartitionScore      int `json:"partition_score"`
	PartitionCount      int `json:"partition_count"`
	PartitionTotalScore int `json:"partition_total_score"`
	/*	FixedQuestions []TkQuestionByIdDetail `json:"fixed_questions"` //固定卷题目
	 */QuestionIDs []QuestionIdsCollection `json:"question_ids"`
	/*	PartitionTypeQuestions []ExamPartitionScore   `json:"partition_type_questions"`
	 */
}

type QuestionIdsCollection struct {
	Id                int                     `json:"id"`
	IsCollection      int                     `json:"is_collection"`
	Type              int                     `json:"type"`
	MaterialsQuestion []QuestionIdsCollection `json:"material_question"` //子题目
}

type ExamPartitionScore struct {
	QuestionType  int `json:"question_type"`
	QuestionScore int `json:"question_score"`
	QuestionCount int `json:"question_count"`
}

type TkPaperExamDetail struct {
	ent.TkExamPaper
	QuestionBankName string `json:"question_bank_name"`
	CreatedAdminName string `json:"created_admin_name"`
}

//id获取题目详情
type TkQuestionByIdDetail struct {
	TkQuestionDetail
	Materials []TkQuestionDetail `json:"material_question"`
}

type TkQuestionDetail struct {
	SectionId        int    `json:"section_id"` //小节id
	ChapterId        int    `json:"chapter_id"` //章id
	QuestionBankName string `json:"question_bank_name"`
	TkListDetail
}

type TkSecList struct {
	Id           int         `json:"id"`
	Type         int         `json:"type"`
	IsCollection int         `json:"is_collection"`
	Materials    []TkSecList `json:"material_question"`
}

type TkListDetail struct {
	Id               int                         `json:"id"`
	Name             string                      `json:"name"`
	NameIos          []htmltojson.Node           `json:"name_ios"`
	Type             int                         `json:"type"`
	DescIos          []htmltojson.Node           `json:"desc_ios"`
	Desc             string                      `json:"desc"`
	IsCollection     int                         `json:"is_collection"`
	Score            int                         `json:"score"`
	Difficulty       string                      `json:"difficulty"`
	CreatedAt        time.Time                   `json:"created_at"`
	CreatedAdminId   int                         `json:"created_admin_id"`
	CreatedAdminName string                      `json:"created_admin_name"`
	QuestionBankId   int                         `json:"question_bank_id"`
	QuestionBankName string                      `json:"question_bank_name"`
	KnowledgePoints  string                      `json:"knowledge_points"`
	Options          ent.TkQuestionAnswerOptions `json:"options"`
	OptionsIos       []AnswerListIos
}

type AnswerListIos struct {
	QuestionID    int               `json:"question_id"`
	IsRight       int               `json:"is_right"`
	Content       []htmltojson.Node `json:"content"`
	OptionNameIos []htmltojson.Node `json:"option_name_ios"`
	OptionName    string            `json:"option_name"`
}

type TkQuestionUserByIdDetail struct {
	TkUserListDetail
	Materials []TkQuestionDetail `json:"material_question"`
}

//
type TkUserListDetail struct {
	Id               int                         `json:"id"`
	Name             string                      `json:"name"`
	Type             int                         `json:"type"`
	Desc             string                      `json:"desc"`
	Score            int                         `json:"score"`
	UserAnswer       string                      `json:"user_answer"`
	Difficulty       string                      `json:"difficulty"`
	CreatedAt        time.Time                   `json:"created_at"`
	CreatedAdminId   int                         `json:"created_admin_id"`
	CreatedAdminName string                      `json:"created_admin_name"`
	QuestionBankId   int                         `json:"question_bank_id"`
	QuestionBankName string                      `json:"question_bank_name"`
	KnowledgePoints  string                      `json:"knowledge_points"`
	Options          ent.TkQuestionAnswerOptions `json:"options"`
}

//章节练习
type TkCsQuestionsList struct {
	ChapterId *int `json:"chapter_id" form:"chapter_id"`
	SectionId *int `json:"section_id" form:"section_id"`
	PageInfo
}

//章节下题目详情
type TkCsQuestionPageList struct {
	TkCsQuestionDetail `json:"list"`
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

type PageResponse struct {
	Total int `json:"total"`
}

type PageInfo struct {
	Page     *int `json:"current"  form:"current"`
	PageSize *int `json:"page_size"  form:"page_size"`
}

//答题卡
type UserAnswerSheet struct {
	Score              int     `json:"score"`
	Accuracy           float64 `json:"accuracy"`
	Duration           int     `json:"duration"`
	QuestionTotalCount int     `json:"question_total_count"`
	RightCount         int     `json:"right_count"`
	WrongCount         int     `json:"wrong_count"`
	NoCount            int     `json:"no_count"`
	Time               string  `json:"time"`
	/*	NextSecId int `json:"next_sec_id"`
	 */QuestionList []UserQuestionTypeList `json:"question_list"`
}

type UserQuestionTypeList struct {
	Id            int    `json:"id"`
	QuestionType  int    `json:"question_type"`
	IsRight       int    `json:"is_right"`
	IsDo          int    `json:"is_do"`
	UserAnswer    string `json:"user_answer"`
	Pid           int    `json:"supEid"`
	QuestionScore int    `json:"question_score"`
}

//错题本首页
type UserWrongQuestionList struct {
	ExamListType []UserTypeQuestion `json:"exam_list_type"`
}

type UserTypeQuestion struct {
	Count            int        `json:"count"`
	QuestionIds      []Question `json:"question_ids"`
	ExamQuestionType int        `json:"exam_question_type" ` //试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷
}

type Question struct {
	Id   int `json:"id"`
	Type int `json:"type"`
}

type QuestionDetail struct {
	TkListDetail
	Materials []TkListDetail `json:"materials"`
}

//随机卷筛选
type QuestionBankListSuiJi struct {
	QuestionId   int `json:"question_id"`
	QuestionType int `json:"question_type"`
}

//做题记录
type UserMakeQuestion struct {
	Name              string  `json:"name"`
	ExamId            int     `json:"exam_id"`
	ExamQuestionType  int     `json:"exam_question_type"`
	Chapter           int     `json:"chapter"`
	SecId             int     `json:"sec_id"`
	QuestionCount     int     `json:"question_count"`
	UserQuestionCount int     `json:"user_question_count"`
	IsDo              int     `json:"is_do"`
	Accuracy          float64 `json:"accuracy"`
	ExamType          int     `json:"exam_type"`
	PracticeExamAt    int     `json:"practice_exam_at" form:"practice_exam_at"`
}

//收藏首页
type UserCollectionList struct {
	TotalCount   int                          `json:"total_count"`
	ExamListType []UserCollectionTypeQuestion `json:"exam_list_type"`
}

type UserCollectionTypeQuestion struct {
	Count       int   `json:"count"`
	QuestionIds []int `json:"question_ids"`
	ExamType    int   `json:"exam_type" ` //试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷
}
