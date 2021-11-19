package response

import "gserver/internal/store/ent"

type RegisterSuccess struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type LoginSuccess struct {
	Token       string `json:"token"`
	UserId      int    `json:"user_id"`
	BirthdayStr string `json:"birthday_str"`
	ent.User
}

//短信接口返回
type SmsPhone struct {
	SmsCode string `json:"sms_code"`
	Phone   string `json:"phone"`
}

//个人中心资料
type PersonalDataItems struct {
	City []struct{}
	Itme []struct{}
}

/********题库**********/
type QuestionClass struct {
	ClassId int `json:"class_id"`
	/*	CourseId    int           `json:"course_id"`
	 *//*RecodeNum   int           `json:"recode_num"`*/ //班级刷题记录
	Name        string `json:"name"`
	CourseCount int    `json:"course_count"` //课程数量
	/*	Img         string        `json:"img"`
	 */ClassCourse []ClassCourse `json:"course_list"`
}

//班级课程
type ClassCourse struct {
	QuestionBankId   int    `json:"question_bank_id"` //
	QuestionNum      int    `json:"question_num"`
	CourseID         int    `json:"course_id"`
	QuestionBankName string `json:"question_bank_name"`
}

type QuestionBank struct {
	BankId           int    `json:"bank_id"`
	QuestionBankName string `json:"question_bank_name"`
	QuestionNum      int    `json:"question_num"` //题目数
}

type CateSmallCourse struct {
	Id        int    `json:"id"`
	SmallName string `json:"small_name"`
}

type QuestionBankType struct {
	QuestionBank
	Accuracy       float64 `json:"accuracy"`       //刷题正确率
	FinishRate     float64 `json:"finish_rate"`    //完成率
	ForecastScore  float64 `json:"forecast_score"` //预测分数
	UserStudyCount int     `json:"user_study_count"`
	TkExamTypeList []TkExamType
}

type TkExamType struct {
	Type          int    `json:"type"`
	QuestionCount int    `json:"question_count"`
	TotalCount    int    `json:"total_count"`
	UserCount     int    `json:"user_count"`
	Time          string `json:"time"`
}

type ExamList struct {
	MoNi    TkExamType
	KaoDian TkExamType
	NiNian  TkExamType
}
