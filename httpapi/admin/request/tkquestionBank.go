package request

import (
	"time"
)

//题库相关request

type SetTkBank struct {
	IdPtrNillable
	NamePtrOnly
	ItemId    *int  `json:"item_category_id" form:"item_category_id" binding:"required,gte=0"`
	LevelId   *int  `json:"level_id" form:"level_id" binding:"required,gte=0"`
	SortOrder int   `json:"sort_order" form:"sort_order"`
	CityIds   []int `json:"city_ids" form:"city_ids" binding:"required"`
	MajorIds  []int `json:"major_ids" form:"major_ids" binding:"required"`
}

type TkBankStatus struct {
	IdOnly
	Status int `json:"status" form:"status" binding:"required,min=1,max=3"`
}

//题库列表
type TkBankList struct {
	Name             *string `json:"name" form:"name"`
	ItemId           *int    `json:"item_category_id" form:"item_category_id"`
	CityId           *int    `json:"city_id" form:"city_id"`
	MajorId          *int    `json:"major_id" form:"major_id"`
	LevelId          *int    `json:"level_id" form:"level_id"`
	Status           *int    `json:"status" form:"status"`
	CreatedAdminName *string `json:"created_admin_name" form:"created_admin_name"`
	PageInfo
}

//添加章
type SetChapter struct {
	IdPtrNillable
	NamePtrOnly
	QuestionBankId *int `json:"question_bank_id" form:"question_bank_id" binding:"required,gte=0"`
}

//添加节
type SetSection struct {
	IdPtrNillable
	NamePtrOnly
	ChapterId *int `json:"chapter_id" form:"chapter_id" binding:"required,gte=0"`
}

//题库知识点
type SetKnowledgePoint struct {
	IdPtrNillable
	NamePtrOnly
	QuestionBankId *int `json:"question_bank_id" form:"question_bank_id" binding:"required,gte=0"`
}

//题库章节列表
type GetTkChapterSectionList struct {
	QuestionBankId *int `json:"question_bank_id" form:"question_bank_id" binding:"required,gte=0"`
	PageInfo
}

//题目添加
type SetTkQuestion struct {
	IdPtrNillable
	QuestionBankId   int                `json:"question_bank_id" form:"question_bank_id" binding:"required,gte=0"`
	Type             int                `json:"type" form:"type" binding:"required,min=1,max=5"`       //题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题
	QuestionStem     string             `json:"question_stem" form:"question_stem" binding:"required"` //题干(html)
	Options          []TkQuestionOption `json:"options" form:"options"`                                //选项
	KnowledgePoint   []int              `json:"knowledge_points" form:"knowledge_points"`
	Difficulty       int                `json:"difficulty" form:"difficulty" binding:required,min=1,max=5` //难易度，1：易，2：较易，3：较难，4：难，5：一般
	Desc             string             `json:"desc" form:"desc"`                                          //解析(html)
	MaterialQuestion []SetTkQuestion    `json:"material_question" form:"material_question"`                //材料题
}
type TkQuestionOption struct {
	OptionContent string `json:"option_content" form:"option_content"` //选项内容(html)
	OptionName    string `json:"option_name" form:"option_name"`       //选项名称 （string）
	IsRight       int    `json:"is_right" form:"is_right"`             //是否正确 1：否，2：是"
}

//题库知识点列表(带分页)
type TkKnowledgePointList struct {
	Name             *string `json:"name" form:"name"`
	QuestionBankId   *int    `json:"question_bank_id" form:"question_bank_id"`
	QuestionBankName *string `json:"question_bank_name" form:"question_bank_name"`
	PageInfo
}

//获取题目列表（带分页）
type TkQuestionsList struct {
	QuestionType   *int    `json:"question_type" form:"question_type"`
	Difficulty     *int    `json:"difficulty" form:"difficulty"`
	KnowledgeName  *string `json:"knowledge_name" form:"knowledge_name"`
	QuestionBankId *int    `json:"question_bank_id" form:"question_bank_id"`
	QuestionName   *string `json:"question_name" form:"question_name"`
	PageInfo
}

//获取章节下题目列表（带分页）
type TkCsQuestionsList struct {
	ChapterId *int `json:"chapter_id" form:"chapter_id"`
	SectionId *int `json:"section_id" form:"section_id"`
	PageInfo
}

//为章节添加(移除)题目
type SetQuestionToSection struct {
	SectionId   *int  `json:"section_id" form:"section_id" binding:"required,gte=0"`
	QuestionIds []int `json:"question_ids" form:"question_ids" binding:"required"`
}

//获取章节下已有题目id数组
type GetQuestionIdsInSection struct {
	SectionId int `json:"section_id" form:"section_id" binding:"required,gte=0"`
}

//获取题目纠错列表（带分页）
type GetQuestionErrorFeedbackListRequest struct {
	QuestionBankName *string    `json:"question_bank_name" form:"question_bank_name"`
	QuestionName     *string    `json:"question_name" form:"question_name"`
	ErrorStartAt     *time.Time `json:"error_start_at" form:"error_start_at"`
	ErrorEndAt       *time.Time `json:"error_end_at" form:"error_end_at"`
	Phone            *string    `json:"phone" form:"phone"`
	Status           *int       `json:"status" form:"status"`
	ErrorType        *int       `json:"error_type" form:"error_type"`
	PageInfo
}

//错题反馈处理
type DealTkErrorFeedbackRequest struct {
	IdPtrOnly
	Status     *int    `json:"status" form:"status" binding:"required"`
	DealRemark *string `json:"deal_remark" form:"deal_remark" binding:"required"`
}

//新增试卷
type SetTkExamPaperRequest struct {
	IdPtrNillable
	NamePtrOnly
	Desc             *string              `json:"desc" form:"desc" binding:"required"`
	ExamQuestionType *int                 `json:"exam_question_type" form:"exam_question_type" binding:"required"`
	ExamType         *int                 `json:"exam_type" form:"exam_type" binding:"required,min=1,max=2"`
	StartAt          *time.Time           `json:"start_at" form:"end_at"`
	EndAt            *time.Time           `json:"end_at" form:"end_at"`
	QuestionCount    *int                 `json:"question_count" form:"question_count" binding:"required"`
	Difficulty       *int                 `json:"difficulty" form:"difficulty" binding:"required"`
	Score            *int                 `json:"score" form:"score" binding:"required"`
	PassScore        *int                 `json:"pass_score" form:"pass_score" binding:"required"`
	Duration         int                  `json:"duration" form:"duration"`
	DurationType     int                  `json:"duration_type" form:"duration_type" binding:"required"`
	QuestionBankId   *int                 `json:"question_bank_id" form:"duration_type" binding:"required"`
	Partitions       []ExamPaperPartition `json:"partitions" form:"partitions"`
}

type ExamPaperPartition struct {
	Name                   string                    `json:"name" form:"name"`
	Desc                   string                    `json:"desc" form:"desc"`
	Type                   int                       `json:"type" form:"type"`
	Duration               int                       `json:"duration" form:"duration"`
	FixedQuestions         []FixedPartitionQuestions `json:"fixed_questions" form:"fixed_questions"`                   //固定卷题目
	PartitionTypeQuestions []PartitionTypeQuestion   `json:"partition_type_questions" form:"partition_type_questions"` //默认分数 + 随机卷题目数
}

type FixedPartitionQuestions struct {
	QuestionId *int `json:"question_id" form:"question_id" binding:"required"'`
	Score      *int `json:"score" form:"score" binding:required`
}

type PartitionTypeQuestion struct {
	QuestionType  int `json:"question_type" form:"question_type" binding:"required"`
	QuestionScore int `json:"question_score" form:"question_score" binding:"required"`
	QuestionCount int `json:"question_count" form:"question_count" binding:"required"`
}

//获取试卷列表（带分页）
type GetTkPaperExamListByPageRequest struct {
	Name             *string `json:"name" form:"name"`
	ExamType         *int    `json:"exam_type" form:"exam_type"`
	ExamQuestionType *int    `json:"exam_question_type" form:"exam_question_type"`
	QuestionBankName *string `json:"question_bank_name" form:"question_bank_name"`
	QuestionBankId   *int    `json:"question_bank_id" form:"question_bank_id"`
	PageInfo
}

//查看模拟考试明细
type GetUserSimulationExamPageListRequest struct {
	ExamPaperId *int `json:"exam_paper_id" form:"exam_paper_id" binding:"required"`
	PageInfo
}

//主观题判分
type SetSubjectiveScoreRequest struct {
	UserExamRecordId *int                    `json:"user_exam_record_id" form:"user_exam_record_id" binding:"required"`
	Score            *int                    `json:"score" form:"score" binding:required`
	Details          []SubjectiveScoreDetail `json:"details" form:"details" binding:"required"`
}

type SubjectiveScoreDetail struct {
	QuestionId *int    `json:"question_id" form:"question_id" binding:"required"`
	Score      *int    `json:"score" form:"score" binding:"required"`
	Desc       *string `json:"desc" form:"desc" binding:"required"`
}

//设置试卷状态
type SetExamPaperStatusRequest struct {
	Id           *int `json:"id"  form:"id" binding:"required"`
	EnableStatus *int `json:"enable_status" form:"enable_status" binding:"required,min=1,max=2"`
}

//导入题目
type ImportQuestionRequest struct {
	QuestionBankId *int `json:"question_bank_id" form:"question_bank_id" binding:"required"`
}

//获取做题记录统计列表
type GetQuestionBankStatisticListRequest struct {
	QuestionBankName *string `json:"question_bank_name" form:"question_bank_name"`
	CateId           *int    `json:"cate_id" form:"cate_id"`
	UserId           *int    `json:"user_id" form:"user_id" binding:"required"`
	PageInfo
}

//获取用户做题记录统计题目详情列表
type GetUserQuestionBankStatisticDetailListRequest struct {
	QuestionName   *string `json:"question_name" form:"question_name"`
	QuestionType   *int    `json:"question_type" form:"question_type"`
	UserId         *int    `json:"user_id" form:"user_id" binding:"required"`
	QuestionBankId *int    `json:"question_bank_id" form:"question_bank_id" binding:"required"`
	PageInfo
}

//学生数据
type GetStudentStatisticRequest struct {
	Username *string `json:"username"`
	Phone    *string `json:"phone"`
	PageInfo
}
