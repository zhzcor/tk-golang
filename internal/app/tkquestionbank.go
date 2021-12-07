package app

import (
	"context"
	"encoding/xml"
	sql2 "entgo.io/ent/dialect/sql"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"
	"tkserver/httpapi/admin/request"
	request2 "tkserver/httpapi/appapi/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/tkchapter"
	"tkserver/internal/store/ent/tkexampartitionquestionlink"
	"tkserver/internal/store/ent/tkexamquestiontype"
	"tkserver/internal/store/ent/tkknowledgepoint"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/tkquestionbankcity"
	"tkserver/internal/store/ent/tkquestionbankmajor"
	"tkserver/internal/store/ent/tkquestionerrorfeedback"
	"tkserver/internal/store/ent/tkquestionsection"
	"tkserver/internal/store/ent/tksection"
	"tkserver/internal/store/ent/tkuserquestionbankrecord"
	"tkserver/pkg/docconv"
	"tkserver/pkg/oss"
	app "tkserver/pkg/requestparam"
)

type TkQuestionBank struct {
}

const (
	SingleSelect     int = 1 //单选题
	MultiSelect      int = 2
	JudgeQuestion    int = 3
	ShortAnswer      int = 4
	MaterialQuestion int = 5
)

//编辑题库
func (a TkQuestionBank) UpdateTkQuestionBank(ctx context.Context, id, itemId int, name string,
	levelId int, cityIds, majorIds []int, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestionBank.
		Query().SoftDelete().
		Where(tkquestionbank.Name(name)).
		Where(tkquestionbank.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	bc := BasicConfig{}
	if errorno := bc.ItemIdExist(ctx, itemId); errorno != nil {
		return 0, errorno
	}
	if errorno := a.TkQuestionBankIdExist(ctx, id); errorno != nil {
		return 0, errorno
	}
	questionBank, err := s.TkQuestionBank.UpdateOneID(id).
		SetName(name).
		SetItemCategoryID(itemId).
		SetLevelID(levelId).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	//删除地区、专业
	_, err = s.TkQuestionBankCity.Delete().Where(tkquestionbankcity.QuestionBankID(id)).Exec(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	_, err = s.TkQuestionBankMajor.Delete().Where(tkquestionbankmajor.QuestionBankID(id)).Exec(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	//添加地区、专业
	bulkCity := make([]*ent.TkQuestionBankCityCreate, len(cityIds))
	for i, v := range cityIds {
		bulkCity[i] = s.TkQuestionBankCity.Create().
			SetQuestionBankID(questionBank.ID).
			SetCityID(v)
	}
	_, err = s.TkQuestionBankCity.CreateBulk(bulkCity...).Save(ctx)
	if err != nil {
		return 0, err
	}

	bulkMajor := make([]*ent.TkQuestionBankMajorCreate, len(majorIds))
	for i, v := range majorIds {
		bulkMajor[i] = s.TkQuestionBankMajor.Create().
			SetQuestionBankID(questionBank.ID).
			SetMajorID(v)
	}
	_, err = s.TkQuestionBankMajor.CreateBulk(bulkMajor...).Save(ctx)
	if err != nil {
		return 0, err
	}

	return questionBank.ID, nil
}

//添加题库
func (a TkQuestionBank) AddTkQuestionBank(ctx context.Context, itemId int, name string,
	adminId int, levelId int, cityIds, majorIds []int, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestionBank.
		Query().SoftDelete().
		Where(tkquestionbank.Name(name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	bc := BasicConfig{}
	if errorno := bc.ItemIdExist(ctx, itemId); errorno != nil {
		return 0, errorno
	}
	tkAdd := s.TkQuestionBank.Create().
		SetName(name).
		SetItemCategoryID(itemId).
		SetLevelID(levelId).
		SetSortOrder(sortOrder)
	if adminId > 0 {
		tkAdd = tkAdd.SetCreatedAdminID(adminId)
	}
	questionBank, err := tkAdd.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	bulkCity := make([]*ent.TkQuestionBankCityCreate, len(cityIds))
	for i, v := range cityIds {
		bulkCity[i] = s.TkQuestionBankCity.Create().
			SetQuestionBankID(questionBank.ID).
			SetCityID(v)
	}
	_, err = s.TkQuestionBankCity.CreateBulk(bulkCity...).Save(ctx)
	if err != nil {
		return 0, err
	}

	bulkMajor := make([]*ent.TkQuestionBankMajorCreate, len(majorIds))
	for i, v := range majorIds {
		bulkMajor[i] = s.TkQuestionBankMajor.Create().
			SetQuestionBankID(questionBank.ID).
			SetMajorID(v)
	}
	_, err = s.TkQuestionBankMajor.CreateBulk(bulkMajor...).Save(ctx)
	if err != nil {
		return 0, err
	}

	return questionBank.ID, nil
}

//删除题库
func (a TkQuestionBank) DelTkQuestionBank(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkQuestionBank.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置题库状态
func (a TkQuestionBank) SetTkQuestionBankStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkQuestionBank.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断题库是否存在
func (a TkQuestionBank) TkQuestionBankIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestionBank.
		Query().SoftDelete().
		Where(tkquestionbank.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题库不存在",
		})
	}
	return nil
}

//编辑章
func (a TkQuestionBank) UpdateTkQuestionBankChapter(ctx context.Context, id int, name string, questionBankId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkChapter.
		Query().SoftDelete().
		Where(tkchapter.Name(name)).
		Where(tkchapter.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return 0, errorno
	}
	//if errorno := a.TkQuestionBankChapterIdExist(ctx, id); errorno != nil {
	//	return 0, errorno
	//}

	chapter, err := s.TkChapter.UpdateOneID(id).
		SetName(name).
		SetQuestionBankID(questionBankId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return chapter.ID, nil
}

//添加章
func (a TkQuestionBank) AddTkQuestionBankChapter(ctx context.Context, name string, questionBankId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkChapter.
		Query().SoftDelete().
		Where(tkchapter.Name(name)).
		Where(tkchapter.QuestionBankID(questionBankId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return 0, errorno
	}
	chapter, err := s.TkChapter.Create().
		SetName(name).
		SetQuestionBankID(questionBankId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return chapter.ID, nil
}

//删除章
func (a TkQuestionBank) DelTkQuestionBankChapter(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankChapterIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkChapter.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断题库章是否存在
func (a TkQuestionBank) TkQuestionBankChapterIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkChapter.
		Query().SoftDelete().
		Where(tkchapter.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题库章不存在",
		})
	}
	return nil
}

//编辑节
func (a TkQuestionBank) UpdateTkQuestionBankSection(ctx context.Context, id int, name string, chapterId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkSection.
		Query().SoftDelete().
		Where(tksection.Name(name)).
		Where(tksection.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	//if errorno := a.TkQuestionBankSectionIdExist(ctx, id); errorno != nil {
	//	return 0, errorno
	//}
	if errorno := a.TkQuestionBankChapterIdExist(ctx, chapterId); errorno != nil {
		return 0, errorno
	}
	section, err := s.TkSection.UpdateOneID(id).
		SetName(name).
		SetChapterID(chapterId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//添加节
func (a TkQuestionBank) AddTkQuestionBankSection(ctx context.Context, name string, chapterId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkSection.
		Query().SoftDelete().
		Where(tksection.Name(name)).
		Where(tksection.ChapterID(chapterId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := a.TkQuestionBankChapterIdExist(ctx, chapterId); errorno != nil {
		return 0, errorno
	}
	section, err := s.TkSection.Create().
		SetName(name).
		SetChapterID(chapterId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//删除节
func (a TkQuestionBank) DelTkQuestionBankSection(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankSectionIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkSection.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断题库节是否存在
func (a TkQuestionBank) TkQuestionBankSectionIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkSection.
		Query().SoftDelete().
		Where(tksection.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题库节不存在",
		})
	}
	return nil
}

//编辑知识点
func (a TkQuestionBank) UpdateTkKnowledgePoint(ctx context.Context, id int, name string, questionBankId int) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkKnowledgePointIdExist(ctx, id); errorno != nil {
		return 0, errorno
	}
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return 0, errorno
	}
	fined, err := s.TkKnowledgePoint.
		Query().SoftDelete().
		Where(tkknowledgepoint.Name(name)).
		Where(tkknowledgepoint.IDNEQ(id)).
		Where(tkknowledgepoint.QuestionBankID(questionBankId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	section, err := s.TkKnowledgePoint.UpdateOneID(id).
		SetName(name).
		SetQuestionBankID(questionBankId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//添加知识点
func (a TkQuestionBank) AddTkKnowledgePoint(ctx context.Context, name string, questionBankId int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.TkKnowledgePoint.
		Query().SoftDelete().
		Where(tkknowledgepoint.Name(name)).
		Where(tkknowledgepoint.QuestionBankID(questionBankId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return 0, errorno
	}
	section, err := s.TkKnowledgePoint.Create().
		SetName(name).
		SetQuestionBankID(questionBankId).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return section.ID, nil
}

//删除知识点
func (a TkQuestionBank) DelTkKnowledgePoint(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkKnowledgePointIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkKnowledgePoint.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断知识点是否存在
func (a TkQuestionBank) TkKnowledgePointIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkKnowledgePoint.
		Query().SoftDelete().
		Where(tkknowledgepoint.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题库知识点不存在",
		})
	}
	return nil
}

//添加题目
func (a TkQuestionBank) AddTkQuestion(ctx context.Context, req request.SetTkQuestion) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, req.QuestionBankId); errorno != nil {
		return 0, errorno
	}
	knowledgePoints, err := s.TkKnowledgePoint.Query().
		Where(tkknowledgepoint.IDIn(req.KnowledgePoint...)).
		All(ctx)
	if err != nil {
		return 0, err
	}
	question, err := s.TkQuestion.Create().
		AddKnowledgePoints(knowledgePoints...).
		SetQuestionBankID(req.QuestionBankId).
		SetType(uint8(req.Type)).
		SetDifficulty(uint8(req.Difficulty)).
		SetDesc(req.Desc).
		SetName(req.QuestionStem).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	//知识点题目数+1
	_, err = s.TkKnowledgePoint.Update().Where(tkknowledgepoint.IDIn(req.KnowledgePoint...)).AddQuestionCount(1).Save(ctx)
	if err != nil {
		return 0, err
	}
	//题库题目数+1
	_, err = s.TkQuestionBank.UpdateOneID(req.QuestionBankId).AddQuestionCount(1).Save(ctx)
	if err != nil {
		return 0, err
	}

	switch req.Type {
	case SingleSelect, MultiSelect, JudgeQuestion, ShortAnswer:
		bulk := make([]*ent.TkQuestionAnswerOptionCreate, len(req.Options))
		for i, v := range req.Options {
			bulk[i] = s.TkQuestionAnswerOption.Create().
				SetQuestionID(question.ID).
				SetOptionName(v.OptionName).
				SetContent(v.OptionContent).
				SetIsRight(uint8(v.IsRight))
		}
		_, err := s.TkQuestionAnswerOption.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
		break
	case MaterialQuestion: //材料题
		if app.IsNil(req.MaterialQuestion) || len(req.MaterialQuestion) == 0 {
			return 0, errorno.NewErr(errorno.MaterialQuestionNotFound)
		}
		bulkQuestions := make([]*ent.TkQuestionCreate, len(req.MaterialQuestion))
		for i, v := range req.MaterialQuestion {
			//知识点
			knowledgePoints, err := s.TkKnowledgePoint.Query().
				Where(tkknowledgepoint.IDIn(v.KnowledgePoint...)).
				All(ctx)
			if err != nil {
				return 0, err
			}
			//选项
			materialOptions := make([]*ent.TkQuestionAnswerOptionCreate, len(v.Options))
			if app.IsNil(v.Options) || len(v.Options) == 0 {
				return 0, errorno.NewErr(errorno.MaterialQuestionNotFound)
			}
			for mi, mv := range v.Options {
				materialOptions[mi] = s.TkQuestionAnswerOption.Create().
					SetOptionName(mv.OptionName).
					SetContent(mv.OptionContent).
					SetIsRight(uint8(mv.IsRight))
			}
			newMaterialOptions, err := s.TkQuestionAnswerOption.
				CreateBulk(materialOptions...).Save(ctx)
			if err != nil {
				return 0, err
			}
			//题目
			bulkQuestions[i] = s.TkQuestion.Create().
				AddKnowledgePoints(knowledgePoints...).
				AddAnswerOptions(newMaterialOptions...).
				SetPid(question.ID).
				SetQuestionBankID(req.QuestionBankId).
				SetType(uint8(v.Type)).
				SetDifficulty(uint8(v.Difficulty)).
				SetDesc(v.Desc).
				SetName(v.QuestionStem)
		}
		_, err := s.TkQuestion.CreateBulk(bulkQuestions...).Save(ctx)
		if err != nil {
			return 0, err
		}

		break
	}
	return question.ID, nil
}

//编辑题目
func (a TkQuestionBank) UpdateTkQuestion(ctx context.Context, req request.SetTkQuestion) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, req.QuestionBankId); errorno != nil {
		return 0, errorno
	}
	if errorno := a.TkQuestionIdExist(ctx, *req.Id); errorno != nil {
		return 0, errorno
	}
	knowledgePoints, err := s.TkKnowledgePoint.Query().SoftDelete().
		Where(tkknowledgepoint.IDIn(req.KnowledgePoint...)).
		All(ctx)
	if err != nil {
		return 0, err
	}
	//知识点题目数-1
	err = a.CalKnowledgePointsQuestionCount(ctx, req.QuestionBankId, -1)
	if err != nil {
		return 0, err
	}
	question, err := s.TkQuestion.UpdateOneID(*req.Id).
		ClearKnowledgePoints().
		ClearAnswerOptions().
		AddKnowledgePoints(knowledgePoints...).
		SetQuestionBankID(req.QuestionBankId).
		SetType(uint8(req.Type)).
		SetDifficulty(uint8(req.Difficulty)).
		SetDesc(req.Desc).
		SetName(req.QuestionStem).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	//知识点关联题目数 +1
	err = a.CalKnowledgePointsQuestionCount(ctx, req.QuestionBankId, 1)
	if err != nil {
		return 0, err
	}
	switch req.Type {
	case SingleSelect, MultiSelect, JudgeQuestion, ShortAnswer:
		//删除旧选项
		//s.TkQuestionAnswerOption.Update().
		//	Where(tkquestionansweroption.QuestionID(*req.Id)).
		//	SoftDelete().Save(ctx)
		bulk := make([]*ent.TkQuestionAnswerOptionCreate, len(req.Options))
		for i, v := range req.Options {
			bulk[i] = s.TkQuestionAnswerOption.Create().
				SetQuestionID(question.ID).
				SetOptionName(v.OptionName).
				SetContent(v.OptionContent).
				SetIsRight(uint8(v.IsRight))
		}
		_, err := s.TkQuestionAnswerOption.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
		break
	case MaterialQuestion: //材料题
		if app.IsNil(req.MaterialQuestion) || len(req.MaterialQuestion) == 0 {
			return 0, errorno.NewErr(errorno.MaterialQuestionNotFound)
		}
		//删除子题目
		_, err := s.TkQuestion.Delete().Where(tkquestion.Pid(*req.Id)).Exec(ctx)
		if err != nil {
			return 0, err
		}
		bulkQuestions := make([]*ent.TkQuestionCreate, len(req.MaterialQuestion))
		for i, v := range req.MaterialQuestion {
			//知识点
			knowledgePoints, err := s.TkKnowledgePoint.Query().
				Where(tkknowledgepoint.IDIn(v.KnowledgePoint...)).
				All(ctx)
			if err != nil {
				return 0, err
			}
			//选项
			materialOptions := make([]*ent.TkQuestionAnswerOptionCreate, len(v.Options))
			if app.IsNil(v.Options) || len(v.Options) == 0 {
				return 0, errorno.NewErr(errorno.MaterialQuestionNotFound)
			}
			for mi, mv := range v.Options {
				materialOptions[mi] = s.TkQuestionAnswerOption.Create().
					SetOptionName(mv.OptionName).
					SetContent(mv.OptionContent).
					SetIsRight(uint8(mv.IsRight))
			}
			newMaterialOptions, err := s.TkQuestionAnswerOption.
				CreateBulk(materialOptions...).Save(ctx)
			if err != nil {
				return 0, err
			}
			//题目
			bulkQuestions[i] = s.TkQuestion.Create().
				AddKnowledgePoints(knowledgePoints...).
				AddAnswerOptions(newMaterialOptions...).
				SetPid(question.ID).
				SetQuestionBankID(v.QuestionBankId).
				SetType(uint8(v.Type)).
				SetDifficulty(uint8(v.Difficulty)).
				SetDesc(v.Desc).
				SetName(v.QuestionStem)
		}
		_, err = s.TkQuestion.CreateBulk(bulkQuestions...).Save(ctx)
		if err != nil {
			return 0, err
		}

		break
	}
	return question.ID, nil
}

//删除题目
func (a TkQuestionBank) DelTkQuestion(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionIdExist(ctx, id); errorno != nil {
		return errorno
	}
	if errorno := a.TkQuestionIsUsedBySection(ctx, id); errorno != nil {
		return errorno
	}
	if errorno := a.TkQuestionIsUsedByExam(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkQuestion.UpdateOneID(id).
		ClearKnowledgePoints().
		ClearAnswerOptions().
		SoftDelete().
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	tkQuestion, err := s.TkQuestion.Query().Where(tkquestion.ID(id)).First(ctx)
	if err != nil {
		return err
	}
	s.TkQuestionBank.UpdateOneID(tkQuestion.QuestionBankID).AddQuestionCount(-1).Save(ctx)

	return nil
}

//id判断题目是否存在
func (a TkQuestionBank) TkQuestionIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestion.
		Query().SoftDelete().
		Where(tkquestion.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题目不存在",
		})
	}
	return nil
}

//id判断题目是否被章节使用
func (a TkQuestionBank) TkQuestionIsUsedBySection(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestionSection.Query().SoftDelete().
		Where(tkquestionsection.QuestionID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题目已添加至章节中，无法删除",
		})
	}
	return nil
}

//id判断题目是否被试卷使用
func (a TkQuestionBank) TkQuestionIsUsedByExam(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkExamPartitionQuestionLink.Query().SoftDelete().
		Where(tkexampartitionquestionlink.QuestionID(id)).
		Where(tkexampartitionquestionlink.ExamPaperPartitionIDNotNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "题目已添加至试卷中，无法删除",
		})
	}
	return nil
}

//加 减知识点题目数量
func (a TkQuestionBank) CalKnowledgePointsQuestionCount(ctx context.Context, questionBankId int, num int) error {
	s := store.WithContext(ctx)
	_, err := s.TkKnowledgePoint.Update().
		Where(tkknowledgepoint.QuestionBankID(questionBankId)).
		AddQuestionCount(num).Save(ctx)
	return err
}

//加 减题库题目数量
func (a TkQuestionBank) CalQuestionBankCount(ctx context.Context, questionBankId int, num int) error {
	s := store.WithContext(ctx)
	_, err := s.TkQuestionBank.UpdateOneID(questionBankId).
		AddQuestionCount(num).Save(ctx)
	return err
}

//加 减节题目数量
func (a TkQuestionBank) CalTkSectionQuestionCount(ctx context.Context, sectionId int, num int) error {
	s := store.WithContext(ctx)
	_, err := s.TkSection.UpdateOneID(sectionId).
		AddQuestionCount(num).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//加 减章题目数量
func (a TkQuestionBank) CalTkChapterQuestionCount(ctx context.Context, ChapterId int, num int) error {
	s := store.WithContext(ctx)
	_, err := s.TkChapter.UpdateOneID(ChapterId).
		AddQuestionCount(num).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//加 减题库分类题目数量
func (a TkQuestionBank) CalTkQuestionBankTypeCount(ctx context.Context, QuestionBankId, examType, num int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkExamQuestionType.
		Query().SoftDelete().
		Where(tkexamquestiontype.QuestionBankID(QuestionBankId)).
		Where(tkexamquestiontype.ExamQuestionType(uint8(examType))).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		_, err := s.TkExamQuestionType.Create().
			SetQuestionBankID(QuestionBankId).
			SetExamQuestionType(uint8(examType)).
			SetQuestionCount(num).
			Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
	} else {
		_, err := s.TkExamQuestionType.Update().
			Where(tkexamquestiontype.QuestionBankID(QuestionBankId)).
			Where(tkexamquestiontype.ExamQuestionType(uint8(examType))).
			AddQuestionCount(num).Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
	}

	return err
}

//添加题目至章节中
func (a TkQuestionBank) SetQuestionToSection(ctx context.Context, sectionId int, questionIds []int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankSectionIdExist(ctx, sectionId); errorno != nil {
		return errorno
	}
	lenCount := len(questionIds)
	bulk := make([]*ent.TkQuestionSectionCreate, lenCount)
	for i, v := range questionIds {
		bulk[i] = s.TkQuestionSection.Create().
			SetQuestionID(v).
			SetSectionID(sectionId)
	}

	result, err := s.TkQuestionSection.CreateBulk(bulk...).Save(ctx)
	savedCount := len(result)
	if err != nil {
		return err
	}
	sectionInfo, err := s.TkSection.Query().Where(tksection.ID(sectionId)).First(ctx)
	if err != nil {
		return err
	}
	err = a.CalTkSectionQuestionCount(ctx, sectionId, savedCount)
	if err != nil {
		return err
	}
	err = a.CalTkChapterQuestionCount(ctx, sectionInfo.ChapterID, savedCount)
	if err != nil {
		return err
	}

	//加 减题库分类题目数量
	chapterInfo, err := s.TkChapter.Query().Where(tkchapter.ID(sectionInfo.ChapterID)).First(ctx)
	if err != nil {
		return err
	}
	err = a.CalTkQuestionBankTypeCount(ctx, chapterInfo.QuestionBankID, 2, savedCount)
	if err != nil {
		return err
	}
	return nil
}

//移除章节中题目
func (a TkQuestionBank) RemoveQuestionToSection(ctx context.Context, sectionId int, questionIds []int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankSectionIdExist(ctx, sectionId); errorno != nil {
		return errorno
	}
	count, err := s.TkQuestionSection.Delete().
		Where(tkquestionsection.SectionID(sectionId)).
		Where(tkquestionsection.QuestionIDIn(questionIds...)).
		Exec(ctx)

	if err != nil {
		return err
	}
	sectionInfo, err := s.TkSection.Query().Where(tksection.ID(sectionId)).First(ctx)
	if err != nil {
		return err
	}
	err = a.CalTkSectionQuestionCount(ctx, sectionId, -count)
	if err != nil {
		return err
	}
	err = a.CalTkChapterQuestionCount(ctx, sectionInfo.ChapterID, -count)
	if err != nil {
		return err
	}

	chapterInfo, err := s.TkChapter.Query().Where(tkchapter.ID(sectionInfo.ChapterID)).First(ctx)
	if err != nil {
		return err
	}
	err = a.CalTkQuestionBankTypeCount(ctx, chapterInfo.QuestionBankID, 2, -count)
	if err != nil {
		return err
	}
	return nil
}

//错题反馈处理
func (a TkQuestionBank) DealTkErrorFeedback(ctx context.Context, id int, status int, dealRemark string, operatorAdminId int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkErrorFeedbackIdExist(ctx, id); errorno != nil {
		return errorno
	}
	fbAdd := s.TkQuestionErrorFeedback.UpdateOneID(id).
		SetDealRemark(dealRemark).
		SetStatus(uint8(status))
	if operatorAdminId > 0 {
		fbAdd = fbAdd.SetOperatorAdminID(operatorAdminId)
	}
	_, err := fbAdd.Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	return nil
}

//id判断错题反馈是否存在
func (a TkQuestionBank) TkErrorFeedbackIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkQuestionErrorFeedback.
		Query().SoftDelete().
		Where(tkquestionerrorfeedback.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "错题反馈记录不存在",
		})
	}
	return nil
}

//随机选取不同类型题目id
func (a TkQuestionBank) MicsSlice(origin []response.QuestionIdsCollection, count int) []response.QuestionIdsCollection {
	tmpOrigin := make([]response.QuestionIdsCollection, len(origin))
	copy(tmpOrigin, origin)
	//一定要seed
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(tmpOrigin), func(i int, j int) {
		tmpOrigin[i], tmpOrigin[j] = tmpOrigin[j], tmpOrigin[i]
	})

	result := make([]response.QuestionIdsCollection, 0, count)
	for index, value := range tmpOrigin {
		if index == count {
			break
		}
		result = append(result, value)
	}
	return result
}

//获取子部分描述
func (a TkQuestionBank) GetExamPartiTypeScore(typeId int, r *ent.TkExamPaperPartitionScore) (int, int, int) {
	switch typeId {
	case SingleSelect:
		return int(r.SingeSelect), int(r.SingeSelectCount), int(r.SingeSelect * r.SingeSelectCount)
	case MultiSelect:
		return int(r.MultiSelect), int(r.MultiSelectCount), int(r.MultiSelect * r.MultiSelectCount)

	case JudgeQuestion:
		return int(r.JudgeQuestion), int(r.JudgeQuestionCount), int(r.JudgeQuestion * r.JudgeQuestionCount)

	case ShortAnswer:
		return int(r.ShorterAnswer), int(r.ShorterAnswerCount), int(r.ShorterAnswer * r.ShorterAnswerCount)

	case MaterialQuestion:
		return int(r.MaterialQuestion), int(r.MaterialQuestionCount), int(r.MaterialQuestion * r.MaterialQuestionCount)
	default:
		return 0, 0, 0
	}
}

//随机描述
func (a TkQuestionBank) GetExamPartiTypeDesc(typeId int, score int, qCount int) string {
	s := strconv.Itoa(score)
	s1 := strconv.Itoa(qCount)
	s2 := strconv.Itoa(qCount * score)
	switch typeId {
	case SingleSelect:
		return "单选题：本题型共" + s1 + "小题，每小题" + s + "分，共" + s2 + "分。在每小题列出的四个备选项中只有一个是最符合题目要求的，请将其代码填写在题后的括号内。错选、多选或未选均无分。"
	case MultiSelect:
		return "多选题：本题型共" + s1 + "小题，每小题" + s + "分，共" + s2 + "分。在每小题列出的四个备选项中有多个符合题目要求的，请将其代码填写在题后的括号内。错选、多选或未选均无分。"

	case JudgeQuestion:
		return "判断题：本题型共" + s1 + "小题，每小题" + s + "分，共" + s2 + "分。在每小题列出 的两个备选项中只有一个是正确选项。"

	case ShortAnswer:
		return "简答题：共" + s1 + "小题，每题" + s + "分，共" + s2 + "分。"

	case MaterialQuestion:
		return "材料题：共" + s1 + "大题，每题" + s + "分，共" + s2 + "分。"
	default:
		return ""
	}
}

//固定
//描述
func (a TkQuestionBank) GetGuDingExamPartiTypeDesc(typeId int, score int, qCount int) string {
	s1 := strconv.Itoa(qCount)
	s2 := strconv.Itoa(score)
	switch typeId {
	case SingleSelect:
		return "单选题：本题型共" + s1 + "小题，共" + s2 + "分。在每小题列出的四个备选项中只有一个是最符合题目要求的，请将其代码填写在题后的括号内。错选、多选或未选均无分。"
	case MultiSelect:
		return "多选题：本题型共" + s1 + "小题，共" + s2 + "分。在每小题列出的四个备选项中有多个符合题目要求的，请将其代码填写在题后的括号内。错选、多选或未选均无分。"

	case JudgeQuestion:
		return "判断题：本题型共" + s1 + "小题，共" + s2 + "分。在每小题列出 的两个备选项中只有一个是正确选项。"

	case ShortAnswer:
		return "简答题：共" + s1 + "小题，共" + s2 + "分。"

	case MaterialQuestion:
		return "材料题：共" + s1 + "大题，共" + s2 + "分。"
	default:
		return ""
	}
}

//上传题目
func (a TkQuestionBank) UploadQuestion(ctx context.Context, f multipart.File, taskId, questionBankId int) error {
	//s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return errorno
	}
	doc, err := GetQuestionRes(f)
	if err != nil {
		return err
	}

	task := ImportTask{}
	err = task.UpdateImportTask(ctx, taskId, 2, "", len(doc))
	if err != nil {
		return err
	}

	for _, v := range doc {
		v.QuestionBankId = questionBankId
		_, err := a.AddTkQuestion(ctx, v)
		if err != nil {
			return err
		}
	}
	err = task.UpdateImportTask(ctx, taskId, 3, "", len(doc))
	if err != nil {
		return err
	}
	fmt.Println(doc)
	return nil

}

func GetQuestionRes(f multipart.File) ([]request.SetTkQuestion, error) {
	res, media, err := docconv.ConvertDocx2(f)
	if err != nil {
		return nil, err
	}

	// 用 <p></p> 切割
	stages := strings.Split(res, "<p></p>")

	type Node struct {
		Items      []string
		childItems []Node
	}

	var doc []request.SetTkQuestion
	var isQuestionArray bool // 是否是材料题
	var rawItems []*Node
	jumpStage := -1
	for si, stage := range stages {
		// 解析html
		item := &Node{}
		var questionListStartKeyWords = []string{
			"【材料题开始】",
			"[材料题开始]",
		}
		var questionListEndKeyWords = []string{
			"【材料题结束】",
			"[材料题结束]",
		}

		dec := xml.NewDecoder(strings.NewReader(stage))
		for {
			tt, err := dec.Token()
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err

			}
			switch v := tt.(type) {
			case xml.StartElement:
				if v.Name.Local == "file" {
					// 文件处理
					if m, o := media["media"][v.Attr[0].Value]; o {
						url, err := oss.Base64LoadUrl(m, "tk_question")
						if err != nil {
							return nil, err
						}
						// 默认都是图
						// 可选上传oss
						item.Items = append(item.Items, `<img src="`+url+`" />`)
					}

				}
			case xml.CharData:
				if strings.TrimSpace(string(v)) != "" {
					item.Items = append(item.Items, string(v))
				}
			}
		}
		//fmt.Println(len(rawItems), rawItems, rawItems[0])
		if len(item.Items) == 0 {
			continue
		}

		for _, word := range questionListStartKeyWords {
			if strings.Index(item.Items[0], word) != -1 {
				isQuestionArray = true
				jumpStage = si
				// 新增一个节点过去
				titleItems := ""
				for ti, tv := range item.Items {
					if ti > 0 {
						titleItems = titleItems + "<p>" + tv + "</p>"
					}
				}
				// 新增一个节点过去
				rawItems = append(rawItems, &Node{
					Items:      []string{titleItems},
					childItems: nil,
				})
			}
		}

		for _, word := range questionListEndKeyWords {
			if strings.Index(item.Items[len(item.Items)-1], word) != -1 {
				isQuestionArray = false
				jumpStage = si
			}
		}

		if isQuestionArray {
			// 直接怼最后一个rawItems的childnode
			if rawItems != nil {
				parentItem := rawItems[len(rawItems)-1]
				if si != jumpStage {
					parentItem.childItems = append(parentItem.childItems, *item)
				}
			}
		} else {
			if jumpStage == si {
				if rawItems != nil {
					parentItem := rawItems[len(rawItems)-1]
					item.Items = item.Items[:len(item.Items)-1]
					parentItem.childItems = append(parentItem.childItems, *item)
				}
			} else {
				rawItems = append(rawItems, item)
			}
		}
	}
	for _, item := range rawItems {
		if item.childItems == nil {
			quest, err := ParseQuestion(item.Items)
			if err != nil {
				return nil, err
			}
			doc = append(doc, *quest)
		} else {
			parentQuest := request.SetTkQuestion{Type: 5, QuestionStem: item.Items[0]}
			for _, childItem := range item.childItems {
				subQuest, err := ParseQuestion(childItem.Items)
				if err != nil {
					return nil, err
				}
				parentQuest.MaterialQuestion = append(parentQuest.MaterialQuestion, *subQuest)
			}
			doc = append(doc, parentQuest)
		}

	}
	//d, _ := json.Marshal(doc)
	//fmt.Println(string(d))
	return doc, err
}

func ParseQuestion(rawItems []string) (*request.SetTkQuestion, error) {
	// 题型判断
	// 填空题
	packfill := regexp.MustCompile(`\[\[(.*?)\]\]`)

	// 找到答案
	answerKeyWords := []string{
		"答案：",
		"答案:",
		"参考答案：",
		"参考答案:",
		"正确答案：",
		"正确答案:",
		"【答案】",
		"[答案]",
		"【正确答案】",
		"[正确答案]",
		"【参考答案】",
		"[参考答案]",
	}
	// 判断题依据条件
	judgeKeyWords := []string{
		"(正确)",
		"(错误)",
		"（正确）",
		"（错误）",
	}
	// 解析 依据条件
	descKeyWords := []string{
		"【解析】",
		"[解析]",
		"【答案解析】",
		"[答案解析]",
	}
	//材料题结束
	questionListEndKeyWords := []string{
		"【材料题结束】",
		"[材料题结束]",
	}
	questionItem := request.SetTkQuestion{}
	var rightOptions []string
	/********* 选择，判断题检测 ********/
	jumpLine := 100
	beforeContent := ""
	answerWord := ""
	descWord := ""
	for ri, item := range rawItems {
		var hasRule bool
		ii := packfill.FindAllString(item, -1)
		if len(ii) > 0 {
			questionItem.Type = 6 //判断 填空题
		}

		//图片解析
		im, err := regexp.MatchString(`<img src="`, string(item))
		if err != nil {
			return nil, err
		}
		if im {
			switch beforeContent {
			case "":
				break
			case "answer":
				rightOptions = append(rightOptions, item)
				continue
			case "desc":
				questionItem.Desc = questionItem.Desc + string(item)
				continue
			default:
				for qi, qv := range questionItem.Options {
					if qv.OptionName == beforeContent {
						questionItem.Options[qi].OptionContent = qv.OptionContent + string(item)
						break
					}
				}
				continue
			}
		}

		// 是否是选项
		itemRaw := []rune(item)
		m, err := regexp.MatchString(`[a-zA-Z]`, string(itemRaw[0]))
		if err != nil {
			return nil, err
		}
		if m && ri > 0 && ri < jumpLine {
			// 如果是判断题
			var isRight = 1
			content := string(itemRaw[2:])
			for _, word := range judgeKeyWords {
				if len(itemRaw) > 4 {
					search := itemRaw[:len(itemRaw)-4]
					if strings.Index(item, string(search)) != -1 {
						questionItem.Type = 3
						answerSearch := itemRaw[len(itemRaw)-4:]
						if strings.Index(string(answerSearch), "正确") != -1 {
							isRight = 2
						}
						content = strings.ReplaceAll(content, word, "")
					}
				}
			}

			// 是选项，就添加进去
			optionItem := request.TkQuestionOption{
				OptionContent: content,
				OptionName:    string(itemRaw[0]),
				IsRight:       isRight,
			}
			beforeContent = string(itemRaw[0])
			questionItem.Options = append(questionItem.Options, optionItem)
			hasRule = true
		}

		// 答案判断
		for wi, word := range answerKeyWords {

			if strings.Index(item, word) != -1 {
				answerWord = word
				jumpLine = wi + 1

			}
		}

		// 解析判断
		for _, word := range descKeyWords {
			if strings.Index(item, word) != -1 {
				descWord = word
				answerWord = ""
			}
		}
		//材料题结束（防止材料题结束录入解析中）
		for _, word := range questionListEndKeyWords {
			if strings.Index(item, word) != -1 {
				descWord = ""
			}
		}

		//答案处理
		if answerWord != "" {
			// 找到答案了，检查是否是 多选题
			// 剔除掉标记
			answerRaw := strings.ReplaceAll(item, answerWord, "")
			answerOptions := []rune(answerRaw)
			beforeContent = "answer"

			for _, option := range answerOptions {
				rightOptions = append(rightOptions, string(option))
			}
			rightOptions = append(rightOptions, "</br>")
			hasRule = true
			//answerWord = ""
		}

		//解析处理
		if descWord != "" {
			// 解析
			descRaw := strings.ReplaceAll(item, descWord, "")
			questionItem.Desc += "<p>" + descRaw + "</p>"
			beforeContent = "desc"
			hasRule = true
		}

		// 其它东西直接扔题干
		if !hasRule {
			questionItem.QuestionStem += item
		}

	}
	// 题目类型判断
	if len(questionItem.Options) > 0 {
		if len(rightOptions) == 2 {
			questionItem.Type = 1
		} else if len(rightOptions) > 2 {
			questionItem.Type = 2
		}
	}
	// 处理正确答案
	var newOptions []request.TkQuestionOption
	for _, option := range questionItem.Options {
		for _, rightOption := range rightOptions {
			if option.OptionName == rightOption {
				option.IsRight = 2
			}
		}
		newOptions = append(newOptions, option)
	}
	questionItem.Options = newOptions

	if questionItem.Type == 0 {
		questionItem.Type = 4
		// 答案放选项里
		questionItem.Options = []request.TkQuestionOption{
			{
				OptionContent: strings.Join(rightOptions, ""),
				OptionName:    "",
				IsRight:       1,
			},
		}
	}
	return &questionItem, nil
}

//做题记录统计
func (a TkQuestionBank) SetUserQuestionBankRecord(ctx context.Context, questionBankId, userId int, questionList []request2.UserRecodeAdd) error {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, questionBankId); errorno != nil {
		return errorno
	}
	record, err := s.TkUserQuestionBankRecord.Query().
		Where(tkuserquestionbankrecord.UserID(userId)).
		Where(tkuserquestionbankrecord.QuestionBankID(questionBankId)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	recordId := 0
	if ent.IsNotFound(err) {
		record, err := s.TkUserQuestionBankRecord.Create().
			SetUserID(userId).
			SetQuestionBankID(questionBankId).
			Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
		recordId = record.ID
	} else {
		recordId = record.ID
	}

	answerCount := record.RecordCount
	rightCount := record.CorrectCount
	wrongCount := record.WrongCount
	for _, v := range questionList {
		questionSql := s.TkQuestion.
			UpdateOneID(v.QuestionId).
			AddAnswerCount(1)

		if v.IsRight == 1 {
			wrongCount += 1
		} else if v.IsRight == 2 {
			rightCount += 1
			questionSql = questionSql.AddRightCount(1)
		}
		if v.IsRecordCount == 1 && v.IsChildren != 1 {
			answerCount += 1
		}

		_, err := questionSql.Save(ctx)
		if err != nil {
			break
		}
	}

	totalCount := float64(rightCount + wrongCount)
	correctRate, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(rightCount)/totalCount), 64)
	if err != nil {
		return err
	}

	_, err = s.TkUserQuestionBankRecord.
		UpdateOneID(recordId).
		SetRecordCount(answerCount).
		SetCorrectCount(rightCount).
		SetWrongCount(wrongCount).
		SetCorrectRate(correctRate).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	return nil
}

//获取地区题库默认(排序)
func (a TkQuestionBank) GetCityBankIds(ctx context.Context, cityId int) ([]ent.TkQuestionBank, error) {
	var res []ent.TkQuestionBank
	s := store.WithContext(ctx)
	cityInfo := s.TkQuestionBankCity.Query().Where(tkquestionbankcity.CityID(cityId)).WithTkQuestionBank(func(query *ent.TkQuestionBankQuery) {
		query.SoftDelete().Where(tkquestionbank.Status(1))
	}).Order(func(s *sql2.Selector) {
		t := sql2.Table(tkquestionbank.Table)
		s.Join(t).On(s.C(tkquestionbankcity.FieldQuestionBankID), t.C(tkquestionbank.FieldID))
		s.OrderBy(t.C(sql2.Desc(tkquestionbank.FieldCreatedAt)))
	}).AllX(ctx)
	for _, v := range cityInfo {
		if v.Edges.TkQuestionBank != nil {
			res = append(res, *v.Edges.TkQuestionBank)
		}
	}

	return res, nil
}
