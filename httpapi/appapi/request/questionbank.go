package request

//通用单个题库ID
type OnlyQuestionBankId struct {
	QuestionBankId int `json:"question_bank_id" form:"question_bank_id" binding:"required"`
}

//题库试卷类型接口
type QuestionBankExamList struct {
	QuestionBankId int `json:"question_bank_id" form:"question_bank_id"`
	TypeId         int `json:"type_id" form:"type_id"`
}

//题库首页
type QuestionIndex struct {
	/*	ClassId  int `json:"class_id" form:"class_id"`
		CourseId int `json:"course_id" form:"course_id"`*/
	CityId int `json:"city_id" form:"city_id" binding:"required"`
}

//地区专业题库题库
type MajorsQuestionBank struct {
	LevelId        int `json:"level_id" form:"level_id"`                 //层次id
	MajorId        int `json:"major_id" form:"major_id"`                 //专业id
	ItemCategoryId int `json:"item_category_id" form:"item_category_id"` //项目id
	CityId         int `json:"city_id" form:"city_id" binding:"required"`                   //地区id
}

//题库接口
type CourseQuestionBank struct {
	QuestionId int `json:"question_id" form:"question_id"`
	CityId     int `json:"city_id" form:"city_id" binding:"required"`
}

//卷子下题目列表
type ExamQuestionList struct {
	Id                 int    `json:"id" form:"id" binding:"required"`
	LookReport         string `json:"look_report"`
	AnswerAgain        int    `json:"answer_again" form:"answer_again"`
	QuestionUserAction int    `json:"question_user_action" form:"question_user_action"` //标示用户是查看题目还是做题 1:做题 0 查看
}

//获取章节下题目列表
type TkCsQuestionsList struct {
	ChapterId          *int `json:"chapter_id" form:"chapter_id"`
	SectionId          *int `json:"section_id" form:"section_id"`
	AnswerAgain        int  `json:"answer_again" form:"answer_again"`
	QuestionBankId     int  `json:"question_bank_id" form:"question_bank_id"`
	QuestionUserAction int  `json:"question_user_action" form:"question_user_action"` //标示用户是查看题目还是做题 1:做题 0 查看
}

//题库答题统计
type UserRecodeList struct {
	ExamQuestionType  int             `json:"exam_question_type" form:"exam_question_type"` //试卷类型
	ExamType          int             `json:"exam_type" form:"exam_type"`                   //考试类型类型
	ExamId            int             `json:"exam_id" form:"exam_id"`                       //试卷id
	SecId             int             `json:"sec_id" form:"sec_id"`                         //小节Id
	Score             int             `json:"score" form:"score"`
	RightCount        int             `json:"right_count" form:"right_count"` //答对题数
	NoCount           int             `json:"no_count" form:"no_count"`       //未作答数
	WrongCount        int             `json:"wrong_count" form:"wrong_count"` //搭错题
	TotalCount        int             `json:"total_count" form:"total_count"` //总题数
	Duration          int             `json:"duration" form:"duration"`       //答题时长
	QuestionBankId    int             `json:"question_bank_id" form:"question_bank_id"`
	QuestionList      []UserRecodeAdd `json:"question_list" form:"question_list[]"`
	RandomQuestionIds []int           `json:"random_question_ids" form:"random_question_ids[]"`
}

type UserRecodeAdd struct {
	QuestionId    int    `json:"question_id" form:"question_id"`
	Answer        string `json:"answer" form:"answer" `
	IsRight       int    `json:"is_right" form:"is_right" `
	QuestionType  int    `json:"question_type" form:"question_type"`
	IsRecordCount int    `json:"is_record_count" form:"is_record_count"`
	IsChildren    int    `json:"is_children" form:"is_children"`
}

type UserAnswerSheet struct {
	ExamType         int `json:"exam_type" form:"exam_type"` //考试类型，1：固定卷，2：随机卷
	SecId            int `json:"sec_id" form:"sec_id"`       //小节Id
	ExamId           int `json:"exam_id" form:"exam_id"`
	ExamQuestionType int `json:"exam_question_type" form:"exam_question_type" binding:"required"` //试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷
}

//错题本列表
type UserWrongQuestionList struct {
	ExamQuestionType int `json:"exam_question_type" form:"exam_question_type" binding:"required"` //试卷类型
	OnlyQuestionBankId
}

//单个题目详情
type QuestionDetail struct {
	QuestionId int `json:"question_id" form:"question_id" binding:"required"`
	ExamId     int `json:"exam_id" form:"question_id"`
	SecId      int `json:"sec_id" form:"sec_id"`
	ChapterId  int `json:"chapter_id" form:"chapter_id"`
}

//题目ID组批量获取题目详情
type QuestionIdArrDetail struct {
	QuestionIds []int `json:"question_ids" form:"question_ids" binding:"required"`
	ExamId      int   `json:"exam_id" form:"question_id"`
	SecId       int   `json:"sec_id" form:"sec_id"`
	ChapterId   int   `json:"chapter_id" form:"chapter_id"`
}

//增加错题记录
type AddWrongQuestion struct {
	QuestionId   int    `json:"question_id" form:"question_id" binding:"required"`
	ExamType     int    `json:"exam_type" form:"exam_type"`
	QuestionType int    `json:"question_type" form:"question_type"`
	Answer       string `json:"answer" form:"answer"`
	OnlyQuestionBankId
}

//题目就错反馈
type AddFeedBack struct {
	QuestionId int    `json:"question_id" form:"question_id" binding:"required"`
	UserName   string `json:"user_name" form:"user_name" binding:"required"`
	Phone      string `json:"phone" form:"phone" binding:"required"`
	ErrorDesc  string `json:"error_desc" form:"error_desc"`
	ErrorType  int    `json:"error_type" form:"error_type"`
}

//收藏
type AddCollection struct {
	TypeId           int `json:"type_id" form:"type_id"`
	SecId            int `json:"sec_id" form:"sec_id"` //小节Id
	Id               int `json:"id" form:"id" binding:"required"`
	ExamId           int `json:"exam_id" form:"exam_id"`
	ExamQuestionType int `json:"exam_question_type" form:"exam_question_type" binding:"required"` //试卷题目分类，1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷
	OnlyQuestionBankId
}

type GetMakeUserQuestion struct {
	QuestionBankId int `json:"question_bank_id" form:"question_bank_id" binding:"required"`
}

//做题记录（试卷章节名称）
type AddMakeUserQuestionRecode struct {
	ExamId int `json:"exam_id" form:"exam_id"`
	SecId  int `json:"sec_id" form:"sec_id"` //小节Id
}

type QuestionIdOnly struct {
	QuestionId int `json:"question_id" form:"question_id" binding:"required"`
}

//移除错题记录
type RemoveWrongQuestion struct {
	QuestionIdOnly
	OnlyQuestionBankId
}

//模拟考试状态
type SimulationExamStatus struct {
	ExamId int `json:"exam_id" form:"exam_id" binding:"required"`
}
