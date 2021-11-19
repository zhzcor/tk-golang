package app

import (
	"context"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/tkuserexamscorerecord"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"
)

type UserRecode struct {
}

func NewUserRecode() *UserRecode {
	return &UserRecode{}
}

//1：模拟考试，2：章节练习，3：历年真题，4：通关必做300题，5：考前密押卷
const (
	MoNiKaoShi    = 1
	SectionLianXi = 2
	NinianZhenTI  = 3
	T300          = 4
	MiYa          = 5
)

//根据题库Ids UserId查询题目数
func (u *UserRecode) GetUserRecodeByIds(ctx context.Context, qIds []int, uid int) (map[int]*ent.TkUserQuestionBankRecord, error) {
	s := store.WithContext(ctx)
	list, err := s.TkUserQuestionBankRecord.Query().Where(tkuserquestionbankrecord.QuestionBankIDIn(qIds...)).
		Where(tkuserquestionbankrecord.UserID(uid)).All(ctx)
	if err != nil {
		return nil, err
	}
	info := make(map[int]*ent.TkUserQuestionBankRecord)
	for _, v := range list {
		info[v.QuestionBankID] = v
	}
	return info, nil

}

func (u *UserRecode) GetUserChapterRecode(ctx context.Context, id int, ids []int) (map[int]int, error) {
	if len(ids) > 0 && id > 0 {
		s := store.WithContext(ctx)
		list, err := s.TkUserExamScoreRecord.Query().
			Where(tkuserexamscorerecord.UserID(id), tkuserexamscorerecord.SectionIDIn(ids...)).All(ctx)
		if err != nil {
			return nil, err
		}
		var data = map[int]int{}

		for _, v := range list {
			data[v.SectionID] = v.TotalCount - v.NoAnswerCount
		}

		return data, err
	}

	return nil, nil
}

//是否有试卷或小节成绩
func (u *UserRecode) IsUserExamScoreRecode(ctx context.Context, id int, examTypeId int, uid int) (bool, error) {
	s := store.WithContext(ctx)
	query := s.TkUserExamScoreRecord.Query().Where(tkuserexamscorerecord.UserID(uid))

	if examTypeId > 0 {
		if examTypeId != SectionLianXi {
			query.Where(tkuserexamscorerecord.ExamPaperID(id))
		} else {
			query.Where(tkuserexamscorerecord.SectionID(id))
		}

		ok, err := query.Exist(ctx)

		return ok, err

	}

	return false, nil

}
