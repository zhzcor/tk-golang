package app

import (
	"context"
	"database/sql"
	"time"
	"tkserver/httpapi/admin/request"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/tkexampaper"
	"tkserver/internal/store/ent/tkexampaperpartition"
	"tkserver/internal/store/ent/tkexampaperpartitionscore"
	"tkserver/internal/store/ent/tkexampartitionquestionlink"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/internal/store/ent/tkquestionsection"
	"tkserver/internal/store/ent/tkuserexamscorerecord"
	app "tkserver/pkg/requestparam"
)

type questionList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

//添加试卷
func (a TkQuestionBank) AddTkExamPaper(ctx context.Context, req request.SetTkExamPaperRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, *req.QuestionBankId); errorno != nil {
		return 0, errorno
	}
	partitionCount := len(req.Partitions)
	if partitionCount <= 0 {
		return 0, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "参数错误",
		})
	}
	_, examPartitions, err := a.setExamPartition(ctx, req)
	if err != nil {
		return 0, err
	}
	paperCreate := s.TkExamPaper.Create().AddExamPartitions(examPartitions...).
		SetQuestionCount(*req.QuestionCount).
		SetName(*req.Name).
		SetDifficulty(uint8(*req.Difficulty)).
		SetDesc(*req.Desc).
		SetDuration(req.Duration).
		SetDurationType(req.DurationType).
		SetQuestionBankID(*req.QuestionBankId).
		SetScore(*req.Score).
		SetScore(*req.Score).
		SetPassScore(*req.PassScore).
		SetExamQuestionType(uint8(*req.ExamQuestionType)).
		SetExamType(uint8(*req.ExamType))
	if !app.IsNil(req.StartAt) && !app.IsNil(req.EndAt) {
		paperCreate = paperCreate.SetStartAt(*req.StartAt).
			SetEndAt(*req.EndAt)
	}

	_, err = paperCreate.Save(ctx)
	if err != nil {
		return 0, err
	}

	err = a.CalTkQuestionBankTypeCount(ctx, *req.QuestionBankId, *req.ExamQuestionType, *req.QuestionCount)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

//编辑试卷
func (a TkQuestionBank) UpdateTkExamPaper(ctx context.Context, req request.SetTkExamPaperRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkQuestionBankIdExist(ctx, *req.QuestionBankId); errorno != nil {
		return 0, errorno
	}
	if errorno := a.TkExamPaperIdExist(ctx, *req.Id); errorno != nil {
		return 0, errorno
	}

	examPaperInfo, err := s.TkExamPaper.Query().Where(tkexampaper.ID(*req.Id)).First(ctx)
	if err != nil {
		return 0, err
	}
	partitionCount := len(req.Partitions)
	if partitionCount <= 0 {
		return 0, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "参数错误",
		})
	}
	partitionIds, err := s.TkExamPaperPartition.Query().
		Where(tkexampaperpartition.ExamPaperID(*req.Id)).
		IDs(ctx)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	//删除试卷关联信息
	if len(partitionIds) > 0 {
		_, err = s.TkExamPaperPartitionScore.Delete().
			Where(tkexampaperpartitionscore.ExamPaperPartitionIDIn(partitionIds...)).
			Exec(ctx)
		if err != nil {
			return 0, err
		}
		_, err = s.TkExamPartitionQuestionLink.Delete().
			Where(tkexampartitionquestionlink.ExamPaperPartitionIDIn(partitionIds...)).
			Exec(ctx)
		if err != nil {
			return 0, err
		}
	}

	_, err = s.TkExamPaperPartition.Delete().
		Where(tkexampaperpartition.ExamPaperID(*req.Id)).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	_, examPartitions, err := a.setExamPartition(ctx, req)
	if err != nil {
		return 0, err
	}
	paperCreate := s.TkExamPaper.UpdateOneID(*req.Id).
		AddExamPartitions(examPartitions...).
		SetQuestionCount(*req.QuestionCount).
		SetQuestionCount(*req.QuestionCount).
		SetDifficulty(uint8(*req.Difficulty)).
		SetName(*req.Name).
		SetDesc(*req.Desc).
		SetDuration(req.Duration).
		SetDurationType(req.DurationType).
		SetQuestionBankID(*req.QuestionBankId).
		SetScore(*req.Score).
		SetScore(*req.Score).
		SetPassScore(*req.PassScore).
		SetExamQuestionType(uint8(*req.ExamQuestionType)).
		SetExamType(uint8(*req.ExamType))
	if !app.IsNil(req.StartAt) && !app.IsNil(req.EndAt) {
		paperCreate = paperCreate.SetStartAt(*req.StartAt).
			SetEndAt(*req.EndAt)
	}

	_, err = paperCreate.Save(ctx)
	if err != nil {
		return 0, err
	}

	staticQuestionCount := *req.QuestionCount - examPaperInfo.QuestionCount
	err = a.CalTkQuestionBankTypeCount(ctx, *req.QuestionBankId, *req.ExamQuestionType, staticQuestionCount)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

//id判断试卷是否存在
func (a TkQuestionBank) TkExamPaperIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkExamPaper.
		Query().SoftDelete().
		Where(tkexampaper.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "试卷不存在",
		})
	}
	return nil
}

//删除id
func (a TkQuestionBank) DelUserExam(ctx context.Context, id int, uid int, typeId int) bool {

	s := store.WithContext(ctx)
	query := s.TkUserExamScoreRecord.Delete().Where(tkuserexamscorerecord.UserID(uid))
	if typeId == 1 {
		query.Where(tkuserexamscorerecord.ExamPaperID(id))
	} else {
		query.Where(tkuserexamscorerecord.SectionID(id))
	}
	_, err := query.Exec(ctx)
	if err != nil {
		return false
	}
	return true
}

//添加试卷子部分
func (a TkQuestionBank) setExamPartition(ctx context.Context, req request.SetTkExamPaperRequest) (int, []*ent.TkExamPaperPartition, error) {
	//试卷子部分
	s := store.WithContext(ctx)
	Ecount := 0
	partitionBulk := make([]*ent.TkExamPaperPartitionCreate, len(req.Partitions))
	partitionRandomBulk := make([]*ent.TkExamPaperPartitionCreate, 5)
	for i, v := range req.Partitions {
		PCount := 0 //试卷子部分题目数量
		//默认分数
		exaMPartitionScoreDetail := ExamPartitionScores{}
		randomCount := 0 //随机卷题目数量

		for _, sv := range v.PartitionTypeQuestions {
			switch sv.QuestionType {
			case SingleSelect:
				exaMPartitionScoreDetail.SelectCount = uint8(sv.QuestionCount)
				exaMPartitionScoreDetail.SelectScore = uint8(sv.QuestionScore)
				break
			case MultiSelect:
				exaMPartitionScoreDetail.MultiCount = uint8(sv.QuestionCount)
				exaMPartitionScoreDetail.MultiScore = uint8(sv.QuestionScore)
				break
			case JudgeQuestion:
				exaMPartitionScoreDetail.JudgeCount = uint8(sv.QuestionCount)
				exaMPartitionScoreDetail.JudgeScore = uint8(sv.QuestionScore)
				break
			case ShortAnswer:
				exaMPartitionScoreDetail.ShortCount = uint8(sv.QuestionCount)
				exaMPartitionScoreDetail.ShortScore = uint8(sv.QuestionScore)
				break
			case MaterialQuestion:
				exaMPartitionScoreDetail.MaterialCount = uint8(sv.QuestionCount)
				exaMPartitionScoreDetail.MaterialScore = uint8(sv.QuestionScore)
				break
			default:
				return 0, nil, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "题目类型不存在",
				})
			}
			randomCount += sv.QuestionCount
		}

		partitionTypeQuestionCount := len(v.PartitionTypeQuestions)
		switch *req.ExamType {
		case 1: //固定卷
			partitionBulkM := s.TkExamPaperPartition.Create().
				SetName(v.Name).
				SetDesc(v.Desc).
				SetDuration(v.Duration).
				SetType(v.Type)

			partitionQuestionCount := len(v.FixedQuestions)
			PCount = partitionQuestionCount
			if partitionQuestionCount <= 0 {
				return 0, nil, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "固定卷题目不能为空",
				})
			}
			partitionQuestionBulk := make([]*ent.TkExamPartitionQuestionLinkCreate, partitionQuestionCount)
			for linkI, linkV := range v.FixedQuestions {
				partitionQuestionBulk[linkI] = s.TkExamPartitionQuestionLink.
					Create().
					SetQuestionID(*linkV.QuestionId).
					SetQuestionScore(uint8(*linkV.Score))
			}
			partitionQuestionLinks, err := s.TkExamPartitionQuestionLink.
				CreateBulk(partitionQuestionBulk...).Save(ctx)
			if err != nil {
				return 0, nil, err
			}
			partitionBulkM = partitionBulkM.AddExamPartitionLinks(partitionQuestionLinks...)

			examPaperPartitionScore, err := a.addExamPaperPartitionScore(ctx, exaMPartitionScoreDetail)
			if err != nil {
				return 0, nil, err
			}

			partitionBulk[i] = partitionBulkM.
				AddExamPartitionScores(examPaperPartitionScore).
				SetQuestionCount(uint8(PCount))

			break
		case 2: //随机卷
			if partitionTypeQuestionCount <= 0 {
				return 0, nil, errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "随机卷默认分数不能为空",
				})
			}

			for s1i, s1v := range v.PartitionTypeQuestions {
				switch s1v.QuestionType {
				case SingleSelect:
					single1, err := a.setPartitionBulk(ctx, "单选题", SingleSelect,
						exaMPartitionScoreDetail.SelectCount, exaMPartitionScoreDetail)
					if err != nil {
						return 0, nil, err
					}
					partitionRandomBulk[s1i] = single1
					break
				case MultiSelect:
					multi1, err := a.setPartitionBulk(ctx, "多选题", MultiSelect,
						exaMPartitionScoreDetail.MultiCount, exaMPartitionScoreDetail)
					if err != nil {
						return 0, nil, err
					}
					partitionRandomBulk[s1i] = multi1
					break
				case JudgeQuestion:
					judge1, err := a.setPartitionBulk(ctx, "判断题", JudgeQuestion,
						exaMPartitionScoreDetail.JudgeCount, exaMPartitionScoreDetail)
					if err != nil {
						return 0, nil, err
					}
					partitionRandomBulk[s1i] = judge1
					break
				case ShortAnswer:
					short1, err := a.setPartitionBulk(ctx, "简答题", ShortAnswer,
						exaMPartitionScoreDetail.ShortCount, exaMPartitionScoreDetail)
					if err != nil {
						return 0, nil, err
					}
					partitionRandomBulk[s1i] = short1
					break
				case MaterialQuestion:
					materail1, err := a.setPartitionBulk(ctx, "材料题", MaterialQuestion,
						exaMPartitionScoreDetail.MaterialCount, exaMPartitionScoreDetail)
					if err != nil {
						return 0, nil, err
					}
					partitionRandomBulk[s1i] = materail1
					break
				}
			}
			break
		}

		if *req.ExamType == 2 { //子部分题目数量计算
			PCount = randomCount
		}
		Ecount += PCount

	}

	if *req.ExamType == 1 {
		examPartitions, err := s.TkExamPaperPartition.CreateBulk(partitionBulk...).Save(ctx)
		if err != nil {
			return 0, nil, err
		}
		return Ecount, examPartitions, nil
	} else {
		examPartitions, err := s.TkExamPaperPartition.CreateBulk(partitionRandomBulk...).Save(ctx)
		if err != nil {
			return 0, nil, err
		}
		return Ecount, examPartitions, nil
	}
}

func (a TkQuestionBank) setPartitionBulk(ctx context.Context, name string, questionType int, questionCount uint8,
	exaMPartitionScoreDetail ExamPartitionScores) (*ent.TkExamPaperPartitionCreate, error) {
	scoreDetail, err := a.addExamPaperPartitionScore(ctx, exaMPartitionScoreDetail)
	if err != nil {
		return nil, err
	}
	return store.WithContext(ctx).TkExamPaperPartition.Create().
		SetName(name).
		SetDesc("").
		SetDuration(0).
		SetType(questionType).
		SetQuestionCount(questionCount).
		AddExamPartitionScores(scoreDetail), nil
}

func (a TkQuestionBank) addExamPaperPartitionScore(ctx context.Context, exaMPartitionScoreDetail ExamPartitionScores) (*ent.TkExamPaperPartitionScore, error) {
	return store.WithContext(ctx).TkExamPaperPartitionScore.Create().
		SetSingeSelect(exaMPartitionScoreDetail.SelectScore).SetSingeSelectCount(exaMPartitionScoreDetail.SelectCount).
		SetMultiSelect(exaMPartitionScoreDetail.MaterialScore).SetMultiSelectCount(exaMPartitionScoreDetail.MultiCount).
		SetJudgeQuestion(exaMPartitionScoreDetail.JudgeScore).SetJudgeQuestionCount(exaMPartitionScoreDetail.JudgeCount).
		SetShorterAnswer(exaMPartitionScoreDetail.ShortScore).SetShorterAnswerCount(exaMPartitionScoreDetail.ShortCount).
		SetMaterialQuestion(exaMPartitionScoreDetail.MaterialScore).SetMaterialQuestionCount(exaMPartitionScoreDetail.MaterialCount).
		Save(ctx)
}

//子部分（随机卷）分数，题目数
type ExamPartitionScores struct {
	SelectScore   uint8
	SelectCount   uint8
	MultiScore    uint8
	MultiCount    uint8
	JudgeScore    uint8
	JudgeCount    uint8
	ShortScore    uint8
	ShortCount    uint8
	MaterialScore uint8
	MaterialCount uint8
}

//删除试卷
func (a TkQuestionBank) DelTkExamPaper(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkExamPaperIdExist(ctx, id); errorno != nil {
		return errorno
	}

	examPaperInfo, err := s.TkExamPaper.Query().Where(tkexampaper.ID(id)).First(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	_, err = s.TkExamPaperPartition.Update().SoftDelete().
		Where(tkexampaperpartition.ExamPaperID(id)).
		ClearExamPartitionLinks().
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	err = s.TkExamPaper.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	_, err = s.TkExamPaperPartition.Delete().Where(tkexampaperpartition.ExamPaperID(id)).Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	err = a.CalTkQuestionBankTypeCount(ctx, examPaperInfo.QuestionBankID, int(examPaperInfo.ExamQuestionType), -examPaperInfo.QuestionCount)
	if err != nil {
		return errorno.NewStoreErr(err)
	}

	return nil
}

//设置试卷状态
func (a TkQuestionBank) SetExamPaperStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkExamPaperIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.TkExamPaper.UpdateOneID(id).SetEnableStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//根据考试开始结束时间获取考试状态
func (a TkQuestionBank) GetExamStatus(startAt, endAt time.Time) int {
	now := time.Now()
	if now.After(endAt) {
		return 2
	}
	if now.After(startAt) {
		return 3
	}
	return 1
}

//主观题判分
func (a TkQuestionBank) SetUserSubjectiveScore(ctx context.Context, req request.SetSubjectiveScoreRequest, operatorAdminId int) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.TkUserExamRecordIdExist(ctx, *req.UserExamRecordId); errorno != nil {
		return 0, errorno
	}

	recordInfo, err := s.TkUserExamScoreRecord.Query().Where(tkuserexamscorerecord.ID(*req.UserExamRecordId)).First(ctx)
	if err != nil {
		return 0, err
	}

	_, err = s.TkUserExamScoreRecord.UpdateOneID(*req.UserExamRecordId).
		ClearUserExamDetails().
		SetStatus(2).
		SetSubjectiveQuestionScore(uint8(*req.Score)).
		SetTotalScore(recordInfo.ObjectiveQuestionScore + uint8(*req.Score)).
		SetOperationTeacherID(operatorAdminId).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	if len(req.Details) > 0 {
		markBulk := make([]*ent.TkUserSimulationTeacherMarkCreate, len(req.Details))
		for i, v := range req.Details {
			markBulk[i] = s.TkUserSimulationTeacherMark.
				Create().
				SetQuestionID(*v.QuestionId).
				SetScore(uint8(*v.Score)).
				SetDesc(*v.Desc).
				SetUserExamRecordID(*req.UserExamRecordId)
		}
		_, err := s.TkUserSimulationTeacherMark.CreateBulk(markBulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
	}
	return 0, nil
}

//id判断用户考试记录是否存在
func (a TkQuestionBank) TkUserExamRecordIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.TkUserExamScoreRecord.
		Query().SoftDelete().
		Where(tkuserexamscorerecord.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "用户模拟考试记录不存在",
		})
	}
	return nil
}

//查询题库类型试卷
func (a TkQuestionBank) GetBankExamTypeList(ctx context.Context, typeId int, BankId int) ([]*ent.TkExamPaper, error) {
	s := store.WithContext(ctx)
	moniList, err := s.TkExamPaper.Query().Where(tkexampaper.QuestionBankID(BankId)).
		Where(tkexampaper.ExamType(uint8(typeId))).All(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return moniList, nil
}

//查询类型试卷user成绩记录表或统计答题数
func (u TkQuestionBank) GetExamTypeListUserRecode(ctx context.Context, typeId int, BankId int, uid int) ([]*ent.TkExamPaper, error) {
	s := store.WithContext(ctx)
	moniList, err := s.TkExamPaper.Query().Where(tkexampaper.EnableStatus(1), tkexampaper.QuestionBankID(BankId)).
		Where(tkexampaper.ExamQuestionType(uint8(typeId))).WithUserExamPapers(func(query2 *ent.TkUserExamScoreRecordQuery) {
		query2.Where(tkuserexamscorerecord.UserID(uid))
	}).All(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return moniList, nil
}

//查询试卷或章节下题目
func (U TkQuestionBank) GetQuestionTypeSeEList(ctx context.Context, id int, typeId int) ([]*ent.TkQuestion, error) {

	data := []*ent.TkQuestion{}
	s := store.WithContext(ctx)
	if typeId == SectionLianXi {
		query := s.TkQuestionSection.Query().SoftDelete().
			Where(tkquestionsection.SectionID(id))
		////列表
		list, err := query.
			SoftDelete().WithSectionQuestion(func(query2 *ent.TkQuestionQuery) {
			query2.SoftDelete().Where(tkquestion.PidIsNil()).WithChildren()
		}).All(ctx)

		if err != nil {
			return nil, err
		}

		for _, v := range list {
			p := &ent.TkQuestion{}
			p = v.Edges.SectionQuestion
			data = append(data, p)
		}

		return data, nil

	} else if typeId != 0 {
		emxa, err := s.TkExamPaper.Query().SoftDelete().Where(tkexampaper.ID(id)).
			WithExamPartitions(func(query *ent.TkExamPaperPartitionQuery) {
				query.SoftDelete().WithExamPartitionLinks(func(query1 *ent.TkExamPartitionQuestionLinkQuery) {
					query1.SoftDelete().WithQuestion(func(query2 *ent.TkQuestionQuery) {
						query2.Where(tkquestion.PidIsNil()).SoftDelete().WithChildren()
					})
				})
			}).First(ctx)

		if err != nil {
			return nil, err
		}

		for _, v := range emxa.Edges.ExamPartitions {
			for _, v1 := range v.Edges.ExamPartitionLinks {
				p := &ent.TkQuestion{}
				p = v1.Edges.Question
				data = append(data, p)
			}

		}

		return data, err

	}

	return nil, nil

}

//查询题库类型试卷
func (a TkQuestionBank) RankTkExamPaper(ctx context.Context, examPaperId int) error {
	s := store.WithContext(ctx)
	if errorno := a.TkExamPaperIdExist(ctx, examPaperId); errorno != nil {
		return errorno
	}

	list, err := s.TkUserExamScoreRecord.Query().SoftDelete().
		Where(tkuserexamscorerecord.ExamPaperID(examPaperId)).
		Select("id", "total_score").
		Order(ent.Desc("total_score")).
		All(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	r := 1
	for _, l := range list {
		_, err = s.TkUserExamScoreRecord.UpdateOneID(l.ID).SetRank(r).Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
		r++
	}
	return nil
}
