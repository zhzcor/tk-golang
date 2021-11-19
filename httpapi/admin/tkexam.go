package admin

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/tkexampaper"
	"gserver/internal/store/ent/tkexampaperpartition"
	"gserver/internal/store/ent/tkexampartitionquestionlink"
	"gserver/internal/store/ent/tkquestion"
	"gserver/internal/store/ent/tkquestionbank"
	"gserver/internal/store/ent/tkuserexamscorerecord"
	"gserver/internal/store/ent/tkuserquestionrecord"
	"gserver/internal/store/ent/tkusersimulationteachermark"
	app2 "gserver/pkg/requestparam"
)

//添加（编辑）试卷
func SetTkExamPaper(ctx *gin.Context) (interface{}, error) {
	var req request.SetTkExamPaperRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	cm := app.Common{}
	adminId := ctx.GetInt(app.GlobalAdminId)
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := tk.UpdateTkExamPaper(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑试卷", *req.Name)
		} else {
			_, err := tk.AddTkExamPaper(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加试卷", *req.Name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除试卷
func DelTkExamPaper(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkExamPaper(ctx, req.Id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//获取试卷列表（带分页）
func GetTkExamPapersByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetTkPaperExamListByPageRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetTkPaperExamListByPageResponse{
		List: []response.TkPaperExamDetail{},
	}
	var tkPaperExamDetail response.TkPaperExamDetail
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkExamPaper.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(tkexampaper.NameContains(*req.Name)).SoftDelete()
		}
		if !app2.IsNil(req.ExamType) {
			query = query.Where(tkexampaper.ExamType(uint8(*req.ExamType)))
		}
		if !app2.IsNil(req.ExamQuestionType) {
			query = query.Where(tkexampaper.ExamQuestionType(uint8(*req.ExamQuestionType)))
		}
		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(tkexampaper.HasQuestionBankWith(
				tkquestionbank.NameContains(*req.QuestionBankName)))
		}
		if !app2.IsNil(req.QuestionBankId) {
			query = query.Where(tkexampaper.HasQuestionBankWith(
				tkquestionbank.ID(*req.QuestionBankId)))
		}
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count

		list, err := query.WithAdmin(func(query *ent.AdminQuery) {
			query.SoftDelete().Select("id", "real_name")
		}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete().Select("id", "name")
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			tkPaperExamDetail.ID = v.ID
			tkPaperExamDetail.Name = v.Name
			tkPaperExamDetail.ExamType = v.ExamType
			tkPaperExamDetail.ExamQuestionType = v.ExamQuestionType
			tkPaperExamDetail.QuestionCount = v.QuestionCount
			tkPaperExamDetail.CreatedAt = v.CreatedAt
			tkPaperExamDetail.CreatedAdminID = v.CreatedAdminID
			tkPaperExamDetail.QuestionBankID = v.QuestionBankID
			tkPaperExamDetail.EnableStatus = v.EnableStatus
			tkPaperExamDetail.StartAt = v.StartAt
			tkPaperExamDetail.EndAt = v.EndAt
			tkPaperExamDetail.ExamStatus = tk.GetExamStatus(v.StartAt, v.EndAt)
			tkPaperExamDetail.QuestionBankName = ""
			tkPaperExamDetail.CreatedAdminName = ""
			if !app2.IsNil(v.Edges.QuestionBank) {
				tkPaperExamDetail.QuestionBankName = v.Edges.QuestionBank.Name
			}
			if !app2.IsNil(v.Edges.Admin) {
				tkPaperExamDetail.QuestionBankName = v.Edges.Admin.RealName
			}
			res.List = append(res.List, tkPaperExamDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//设置试卷状态
func SetExamPaperStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetExamPaperStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetExamPaperStatus(ctx, *req.Id, *req.EnableStatus)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//获取试卷列表（不带分页）
func GetTkExamPapersListAll(ctx *gin.Context) (interface{}, error) {
	var req request.GetTkPaperExamListByPageRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := ent.TkExamPapers{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkExamPaper.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(tkexampaper.NameContains(*req.Name)).SoftDelete()
		}
		if !app2.IsNil(req.ExamType) {
			query = query.Where(tkexampaper.ExamType(uint8(*req.ExamType)))
		}
		if !app2.IsNil(req.ExamQuestionType) {
			query = query.Where(tkexampaper.ExamQuestionType(uint8(*req.ExamQuestionType)))
		}
		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(tkexampaper.HasQuestionBankWith(
				tkquestionbank.NameContains(*req.QuestionBankName)))
		}
		if !app2.IsNil(req.QuestionBankId) {
			query = query.Where(tkexampaper.HasQuestionBankWith(
				tkquestionbank.ID(*req.QuestionBankId)))
		}
		list, err := query.
			Select("id", "name", "exam_type", "exam_question_type", "question_bank_id").
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//id获取试卷详情
func GetTkExamPaperById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetTkExamPaperByIdResponse{
		Partitions: []response.ExamPaperPartition{},
	}
	examPartitionScore := response.ExamPartitionScore{}
	tkQuestionDetail := response.TkQuestionByIdDetail{}
	tk := app.TkQuestionBank{}
	//var res interface{}

	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := tk.TkExamPaperIdExist(ctx, *req.Id)
		if err != nil {
			return err
		}

		query, err := s.TkExamPaper.Query().SoftDelete().Where(tkexampaper.ID(*req.Id)).
			WithExamPartitions(func(query *ent.TkExamPaperPartitionQuery) {
				query.SoftDelete().WithExamPartitionLinks(func(query *ent.TkExamPartitionQuestionLinkQuery) {
					query.SoftDelete().WithQuestion(func(query *ent.TkQuestionQuery) {
						query.SoftDelete().WithAnswerOptions().WithKnowledgePoints().
							WithChildren(func(query *ent.TkQuestionQuery) {
								query.SoftDelete().WithKnowledgePoints().WithAnswerOptions()
							})
					})
				}).WithExamPartitionScores()
			}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("name")
		}).First(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		res.ID = query.ID
		res.Name = query.Name
		res.Difficulty = query.Difficulty
		res.QuestionBankID = query.QuestionBankID
		res.ExamQuestionType = query.ExamQuestionType
		res.ExamType = query.ExamType
		res.Score = query.Score
		res.CreatedAt = query.CreatedAt
		res.Desc = query.Desc
		res.Desc = query.Desc
		res.PassScore = query.PassScore
		res.Duration = query.Duration
		res.DurationType = query.DurationType
		res.StartAt = query.StartAt
		res.EndAt = query.EndAt
		res.QuestionCount = query.QuestionCount
		res.QuestionBankName = ""
		if !app2.IsNil(query.Edges.QuestionBank) {
			res.QuestionBankName = query.Edges.QuestionBank.Name
		}
		if !app2.IsNil(query.Edges.ExamPartitions) {
			for _, v := range query.Edges.ExamPartitions {
				examPaperPartition := response.ExamPaperPartition{}
				examPaperPartition.ID = v.ID
				examPaperPartition.Name = v.Name
				examPaperPartition.Desc = v.Desc
				examPaperPartition.Type = v.Type
				examPaperPartition.Duration = v.Duration
				examPaperPartition.ExamPaperID = v.ExamPaperID
				examPaperPartition.QuestionCount = v.QuestionCount
				examPaperPartition.CreatedAt = v.CreatedAt
				if !app2.IsNil(v.Edges.ExamPartitionLinks) {
					for _, lv := range v.Edges.ExamPartitionLinks {
						if !app2.IsNil(lv.Edges.Question) {
							tkQuestionDetail.Id = lv.Edges.Question.ID
							tkQuestionDetail.Name = lv.Edges.Question.Name
							tkQuestionDetail.Desc = lv.Edges.Question.Desc
							tkQuestionDetail.Type = int(lv.Edges.Question.Type)
							tkQuestionDetail.Score = int(lv.QuestionScore)
							tkQuestionDetail.Difficulty = int(lv.Edges.Question.Difficulty)
							tkQuestionDetail.KnowledgePoints = ent.TkKnowledgePoints{}
							tkQuestionDetail.Options = ent.TkQuestionAnswerOptions{}
							if !app2.IsNil(lv.Edges.Question.Edges.KnowledgePoints) {
								tkQuestionDetail.KnowledgePoints = lv.Edges.Question.Edges.KnowledgePoints
							}
							if !app2.IsNil(lv.Edges.Question.Edges.AnswerOptions) {
								tkQuestionDetail.Options = lv.Edges.Question.Edges.AnswerOptions
							}
							//材料题
							materialDetail := response.TkQuestionDetail{}
							if !app2.IsNil(lv.Edges.Question.Edges.Children) {
								for _, cv := range lv.Edges.Question.Edges.Children {
									materialDetail.Id = cv.ID
									materialDetail.Name = cv.Name
									materialDetail.Desc = cv.Desc
									materialDetail.Type = int(cv.Type)
									materialDetail.Score = int(lv.QuestionScore)
									materialDetail.Difficulty = int(cv.Difficulty)
									materialDetail.KnowledgePoints = ent.TkKnowledgePoints{}
									materialDetail.Options = ent.TkQuestionAnswerOptions{}
									if !app2.IsNil(cv.Edges.KnowledgePoints) {
										materialDetail.KnowledgePoints = cv.Edges.KnowledgePoints
									}
									if !app2.IsNil(cv.Edges.AnswerOptions) {
										materialDetail.Options = cv.Edges.AnswerOptions
									}
									tkQuestionDetail.Materials = append(tkQuestionDetail.Materials, materialDetail)
								}
							}
							examPaperPartition.FixedQuestions = append(examPaperPartition.FixedQuestions, tkQuestionDetail)
						}
					}
					if !app2.IsNil(v.Edges.ExamPartitionScores) { //默认分数
						sv := v.Edges.ExamPartitionScores[0]
						examPartitionScore.QuestionType = app.SingleSelect
						examPartitionScore.QuestionScore = int(sv.SingeSelect)
						examPartitionScore.QuestionCount = int(sv.SingeSelectCount)
						examPaperPartition.PartitionTypeQuestions = append(examPaperPartition.PartitionTypeQuestions, examPartitionScore)
						examPartitionScore.QuestionType = app.MultiSelect
						examPartitionScore.QuestionScore = int(sv.MultiSelect)
						examPartitionScore.QuestionCount = int(sv.MultiSelectCount)
						examPaperPartition.PartitionTypeQuestions = append(examPaperPartition.PartitionTypeQuestions, examPartitionScore)
						examPartitionScore.QuestionType = app.ShortAnswer
						examPartitionScore.QuestionScore = int(sv.ShorterAnswer)
						examPartitionScore.QuestionCount = int(sv.ShorterAnswerCount)
						examPaperPartition.PartitionTypeQuestions = append(examPaperPartition.PartitionTypeQuestions, examPartitionScore)
						examPartitionScore.QuestionType = app.JudgeQuestion
						examPartitionScore.QuestionScore = int(sv.JudgeQuestion)
						examPartitionScore.QuestionCount = int(sv.JudgeQuestionCount)
						examPaperPartition.PartitionTypeQuestions = append(examPaperPartition.PartitionTypeQuestions, examPartitionScore)
						examPartitionScore.QuestionType = app.MaterialQuestion
						examPartitionScore.QuestionScore = int(sv.MaterialQuestion)
						examPartitionScore.QuestionCount = int(sv.MaterialQuestionCount)
						examPaperPartition.PartitionTypeQuestions = append(examPaperPartition.PartitionTypeQuestions, examPartitionScore)
					}

				}
				res.Partitions = append(res.Partitions, examPaperPartition)
				if res.ExamType == 2 {
					break
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//查看模拟考试明细
func GetUserSimulationExamPageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetUserSimulationExamPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetUserSimulationExamPageListResponse{
		List: []response.UserSimulationExamDetail{},
	}
	var userSimulationExamDetail response.UserSimulationExamDetail
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkUserExamScoreRecord.Query().SoftDelete()
		if !app2.IsNil(req.ExamPaperId) {
			query = query.Where(tkuserexamscorerecord.ExamPaperID(*req.ExamPaperId))
		}
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count

		list, err := query.WithTeacher(func(query *ent.TeacherQuery) {
			query.SoftDelete().Select("id", "name")
		}).WithUser(func(query *ent.UserQuery) {
			query.SoftDelete().Select("id", "username", "phone")
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			userSimulationExamDetail.ID = v.ID
			userSimulationExamDetail.ObjectiveQuestionScore = v.ObjectiveQuestionScore
			userSimulationExamDetail.SubjectiveQuestionScore = v.SubjectiveQuestionScore
			userSimulationExamDetail.Status = v.Status
			userSimulationExamDetail.TotalScore = v.TotalScore
			userSimulationExamDetail.Rank = v.Rank
			userSimulationExamDetail.UserId = v.UserID
			userSimulationExamDetail.OperationTeacherID = v.OperationTeacherID
			userSimulationExamDetail.OperatorTeacherName = ""
			userSimulationExamDetail.Username = ""
			userSimulationExamDetail.Phone = ""
			if !app2.IsNil(v.Edges.Teacher) {
				userSimulationExamDetail.OperatorTeacherName = v.Edges.Teacher.Name
			}
			if !app2.IsNil(v.Edges.User) {
				userSimulationExamDetail.Username = v.Edges.User.Username
				userSimulationExamDetail.Phone = v.Edges.User.Phone
			}
			res.List = append(res.List, userSimulationExamDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取用户模拟考试主观题记录
func GetUserSimulationSubjectiveQuestions(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := []response.GetUserSubjectiveQuestionsDetail{}
	var userSubjectiveQuestionsDetail response.GetUserSubjectiveQuestionsDetail
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		record, err := s.TkUserExamScoreRecord.Query().
			Where(tkuserexamscorerecord.ID(*req.Id)).
			First(ctx)
		if err != nil {
			return err
		}
		userId := record.UserID
		examPaperId := record.ExamPaperID
		userSubjectiveQuestionsDetail.UserExamRecordId = *req.Id

		ExamPartitionIds, err := s.TkExamPaperPartition.Query().SoftDelete().
			Where(tkexampaperpartition.ExamPaperID(examPaperId)).IDs(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		var questionIds []int
		partitionLinks, err := s.TkExamPartitionQuestionLink.Query().SoftDelete().
			Where(tkexampartitionquestionlink.ExamPaperPartitionIDIn(ExamPartitionIds...)).
			Select("question_id", "exam_paper_partition_id", "question_score").
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range partitionLinks {
			questionIds = append(questionIds, v.QuestionID)
		}
		//查出试卷下所有主观题
		questions, err := s.TkQuestion.Query().SoftDelete().
			Where(tkquestion.IDIn(questionIds...)).
			Where(tkquestion.TypeIn(uint8(app.ShortAnswer), uint8(app.MaterialQuestion))).
			Where(tkquestion.PidIsNil()).
			WithAnswerOptions().
			WithUserRecords(func(query *ent.TkUserQuestionRecordQuery) {
				query.SoftDelete().Where(tkuserquestionrecord.UserID(userId)).
					Where(tkuserquestionrecord.ExamPaperID(examPaperId))
			}).
			WithUserExamQuestions(func(query *ent.TkUserSimulationTeacherMarkQuery) {
				query.SoftDelete().Where(tkusersimulationteachermark.UserExamRecordID(*req.Id))
			}).
			WithChildren(func(query *ent.TkQuestionQuery) { //材料题子题目
				query.SoftDelete().WithAnswerOptions().
					WithUserRecords(func(query *ent.TkUserQuestionRecordQuery) {
						query.SoftDelete().Where(tkuserquestionrecord.UserID(userId)).
							Where(tkuserquestionrecord.ExamPaperID(examPaperId))
					}).
					WithUserExamQuestions(func(query *ent.TkUserSimulationTeacherMarkQuery) {
						query.SoftDelete().Where(tkusersimulationteachermark.UserExamRecordID(*req.Id))
					})
			}).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range questions {
			userSubjectiveQuestionsDetail.QuestionId = v.ID
			userSubjectiveQuestionsDetail.QuestionName = v.Name
			userSubjectiveQuestionsDetail.QuestionType = int(v.Type)
			userSubjectiveQuestionsDetail.QuestionDesc = v.Desc
			userSubjectiveQuestionsDetail.TeacherDesc = ""
			userSubjectiveQuestionsDetail.TeacherScore = 0
			userSubjectiveQuestionsDetail.UserAnswer = ""
			if !app2.IsNil(v.Edges.UserExamQuestions) {
				if len(v.Edges.UserExamQuestions) > 0 {
					teacherMarks := v.Edges.UserExamQuestions[len(v.Edges.UserExamQuestions)-1]
					userSubjectiveQuestionsDetail.TeacherDesc = teacherMarks.Desc
					userSubjectiveQuestionsDetail.TeacherScore = int(teacherMarks.Score)
				}
			}
			if !app2.IsNil(v.Edges.UserRecords) {
				if len(v.Edges.UserRecords) > 0 {
					userRecord := v.Edges.UserRecords[len(v.Edges.UserRecords)-1]
					userSubjectiveQuestionsDetail.UserAnswer = userRecord.Answer
				}
			}
			userSubjectiveQuestionsDetail.Options = ent.TkQuestionAnswerOptions{}
			if !app2.IsNil(v.Edges.AnswerOptions) {
				userSubjectiveQuestionsDetail.Options = v.Edges.AnswerOptions
			}
			userSubjectiveQuestionsDetail.Score = 0
			for _, pv := range partitionLinks {
				if pv.QuestionID == v.ID {
					userSubjectiveQuestionsDetail.Score = int(pv.QuestionScore)
				}
			}

			userSubjectiveQuestionsDetail.Materials = []response.GetUserSubjectiveQuestionsDetail{}
			if !app2.IsNil(v.Edges.Children) {
				for _, cv := range v.Edges.Children {
					materialQuestionsDetail := response.GetUserSubjectiveQuestionsDetail{}
					materialQuestionsDetail.QuestionId = cv.ID
					materialQuestionsDetail.QuestionName = cv.Name
					materialQuestionsDetail.QuestionType = int(cv.Type)
					materialQuestionsDetail.QuestionDesc = cv.Desc
					materialQuestionsDetail.TeacherDesc = ""
					materialQuestionsDetail.TeacherScore = 0
					materialQuestionsDetail.UserAnswer = ""
					if !app2.IsNil(v.Edges.UserExamQuestions) {
						if len(cv.Edges.UserExamQuestions) > 0 {
							teacherMarks := cv.Edges.UserExamQuestions[len(cv.Edges.UserExamQuestions)-1]
							materialQuestionsDetail.TeacherDesc = teacherMarks.Desc
							materialQuestionsDetail.TeacherScore = int(teacherMarks.Score)
						}
					}
					if !app2.IsNil(cv.Edges.UserRecords) {
						if len(cv.Edges.UserRecords) > 0 {
							userRecord := cv.Edges.UserRecords[len(cv.Edges.UserRecords)-1]
							materialQuestionsDetail.UserAnswer = userRecord.Answer
						}
					}
					materialQuestionsDetail.Options = ent.TkQuestionAnswerOptions{}
					if !app2.IsNil(v.Edges.AnswerOptions) {
						materialQuestionsDetail.Options = cv.Edges.AnswerOptions
					}
					materialQuestionsDetail.Score = 0
					for _, spv := range partitionLinks {
						if spv.QuestionID == cv.ID {
							materialQuestionsDetail.Score = int(spv.QuestionScore)
						}
					}
					userSubjectiveQuestionsDetail.Materials = append(userSubjectiveQuestionsDetail.Materials, materialQuestionsDetail)
				}

			}

			res = append(res, userSubjectiveQuestionsDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//主观题判分
func SetUserSubjectiveScore(ctx *gin.Context) (interface{}, error) {
	var req request.SetSubjectiveScoreRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	operatorAdminId := ctx.GetInt(app.GlobalAdminId)
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err := tk.SetUserSubjectiveScore(ctx, req, operatorAdminId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//模拟考试排名
func RankTkExamPaper(ctx *gin.Context) (interface{}, error) {
	var req request.GetUserSimulationExamPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.RankTkExamPaper(ctx, *req.ExamPaperId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
