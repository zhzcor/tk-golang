package admin

import (
	"context"
	"database/sql"
	sql2 "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/admin"
	"gserver/internal/store/ent/itemcategory"
	"gserver/internal/store/ent/tkchapter"
	"gserver/internal/store/ent/tkknowledgepoint"
	"gserver/internal/store/ent/tkquestion"
	"gserver/internal/store/ent/tkquestionbank"
	"gserver/internal/store/ent/tkquestionerrorfeedback"
	"gserver/internal/store/ent/tkquestionsection"
	"gserver/internal/store/ent/tksection"
	"gserver/internal/store/ent/tkuserquestionbankrecord"
	"gserver/internal/store/ent/tkuserquestionrecord"
	"gserver/internal/store/ent/user"
	app2 "gserver/pkg/requestparam"
	"os"
	"time"
)

//添加（编辑）题库
func SetTkQuestionBank(ctx *gin.Context) (interface{}, error) {
	var req request.SetTkBank
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
			_, err := tk.UpdateTkQuestionBank(ctx, *req.Id, *req.ItemId, *req.Name)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑题库", *req.Name)
		} else {
			_, err := tk.AddTkQuestionBank(ctx, *req.ItemId, *req.Name, adminId)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加题库", *req.Name)

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除题库
func DelTkQuestionBank(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkQuestionBank(ctx, req.Id)
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

//设置题库状态
func SetTkQuestionBankStatus(ctx *gin.Context) (interface{}, error) {
	var req request.TkBankStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.SetTkQuestionBankStatus(ctx, req.Id, req.Status)
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

//获取题库列表
func GetTkQuestionBankByPage(ctx *gin.Context) (interface{}, error) {
	var req request.TkBankList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkBankListSuccess{
		List: []response.TkQuestionBank{},
	}
	var tkBank response.TkQuestionBank
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkQuestionBank.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(tkquestionbank.NameContains(*req.Name)).SoftDelete()
		}
		if !app2.IsNil(req.CreatedAdminName) {
			query = query.Where(tkquestionbank.HasAdminWith(admin.RealNameContains(*req.CreatedAdminName)))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(tkquestionbank.Status(uint8(*req.Status)))
		}
		if !app2.IsNil(req.ItemId) {
			items, err := s.ItemCategory.Query().
				Where(itemcategory.Or(itemcategory.Pid(*req.ItemId), itemcategory.ID(*req.ItemId))).
				IDs(ctx)
			if err != nil {
				if err == sql.ErrNoRows {
					return nil
				}
				return err
			}
			query = query.Where(tkquestionbank.ItemCategoryIDIn(items...))
		}
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		enableCount, err := query.Clone().Where(tkquestionbank.Status(1)).Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		res.Statistic.Total = count
		res.Statistic.EnableCount = enableCount
		res.Statistic.DisableCount = count - enableCount

		list, err := query.WithAdmin().
			WithItemCategory(func(query *ent.ItemCategoryQuery) {
				query.WithParent()
			}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			tkBank.ID = v.ID
			tkBank.Status = v.Status
			tkBank.Name = v.Name
			tkBank.CreatedAdminID = v.CreatedAdminID
			tkBank.CreatedAt = v.CreatedAt
			tkBank.QuestionCount = v.QuestionCount
			tkBank.ItemCategoryID = v.ItemCategoryID
			tkBank.ItemCategoryName = v.Edges.ItemCategory.Name
			tkBank.CreatedAdminName = ""
			tkBank.ParentItemCategoryName = ""
			tkBank.ParentItemCategoryId = 0
			if !app2.IsNil(v.Edges.Admin) {
				tkBank.CreatedAdminName = v.Edges.Admin.RealName
			}
			if !app2.IsNil(v.Edges.ItemCategory.Edges.Parent) {
				tkBank.ParentItemCategoryName = v.Edges.ItemCategory.Edges.Parent.Name
				tkBank.ParentItemCategoryId = v.Edges.ItemCategory.Edges.Parent.ID
			}
			res.List = append(res.List, tkBank)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//添加(编辑)题库 章
func SetTkQuestionBankChapter(ctx *gin.Context) (interface{}, error) {
	var req request.SetChapter
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := tk.UpdateTkQuestionBankChapter(ctx, *req.Id, *req.Name, *req.QuestionBankId)
			if err != nil {
				return err
			}
		} else {
			_, err := tk.AddTkQuestionBankChapter(ctx, *req.Name, *req.QuestionBankId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//删除章
func DelTkQuestionBankChapter(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkQuestionBankChapter(ctx, req.Id)
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

//添加(编辑)题库 节
func SetTkQuestionBankSection(ctx *gin.Context) (interface{}, error) {
	var req request.SetSection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := tk.UpdateTkQuestionBankSection(ctx, *req.Id, *req.Name, *req.ChapterId)
			if err != nil {
				return err
			}
		} else {
			_, err := tk.AddTkQuestionBankSection(ctx, *req.Name, *req.ChapterId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//删除节
func DelTkQuestionBankSection(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkQuestionBankSection(ctx, req.Id)
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

//获取章节列表(带分页)
func GetTkQuestionBankChapterSectionByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetTkChapterSectionList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkChapterSectionListSuccess{
		List: []response.TkChapterSection{},
	}
	tkChapterSection := response.TkChapterSection{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkChapter.Query().SoftDelete().
			Where(tkchapter.QuestionBankID(*req.QuestionBankId))
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count

		list, err := query.Select("id", "name", "question_count", "question_bank_id").
			WithSections(func(query *ent.TkSectionQuery) {
				query.SoftDelete().
					Select("id", "name", "question_count", "chapter_id")
			}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			tkChapterSection.ID = v.ID
			tkChapterSection.Name = v.Name
			tkChapterSection.QuestionBankID = v.QuestionBankID
			tkChapterSection.QuestionCount = v.QuestionCount
			if !app2.IsNil(v.Edges.Sections) {
				tkChapterSection.Sections = v.Edges.Sections
			}
			res.List = append(res.List, tkChapterSection)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//添加(编辑)知识点
func SetTkKnowledgePoint(ctx *gin.Context) (interface{}, error) {
	var req request.SetKnowledgePoint
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := tk.UpdateTkKnowledgePoint(ctx, *req.Id, *req.Name, *req.QuestionBankId)
			if err != nil {
				return err
			}
		} else {
			_, err := tk.AddTkKnowledgePoint(ctx, *req.Name, *req.QuestionBankId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//删除知识点
func DelTkKnowledgePoint(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkKnowledgePoint(ctx, req.Id)
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

//获取题库知识点列表（分页）
func GetTkKnowledgePointsByPage(ctx *gin.Context) (interface{}, error) {
	var req request.TkKnowledgePointList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkKnowledgePointPageListSuccess{
		List: []response.TkKnowledgePointDetail{},
	}
	var tkKnowledgeDetail response.TkKnowledgePointDetail
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkQuestionBank.Query().SoftDelete()

		if !app2.IsNil(req.Name) {
			query = query.Where(tkquestionbank.HasKnowledgePointsWith(
				tkknowledgepoint.NameContains(*req.Name),
				tkknowledgepoint.DeletedAtIsNil()))
		}
		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(tkquestionbank.NameContains(*req.QuestionBankName))
		}
		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.WithKnowledgePoints(func(query *ent.TkKnowledgePointQuery) {
			query.SoftDelete()
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			tkKnowledgeDetail.Id = v.ID
			tkKnowledgeDetail.Name = v.Name
			tkKnowledgeDetail.QuestionCount = v.QuestionCount
			if !app2.IsNil(v.Edges.KnowledgePoints) {
				tkKnowledgeDetail.KnowledgePoints = v.Edges.KnowledgePoints
			}
			res.List = append(res.List, tkKnowledgeDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取知识点列表（分页）
func GetTkKnowledgePoints(ctx *gin.Context) (interface{}, error) {
	var req request.TkKnowledgePointList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkKnowledgePoints{
		List: []response.TkPointDetail{},
	}
	pointDetail := response.TkPointDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkKnowledgePoint.Query().SoftDelete()

		if !app2.IsNil(req.Name) {
			query = query.Where(tkknowledgepoint.NameContains(*req.Name))
		}
		if !app2.IsNil(req.QuestionBankId) {
			query = query.Where(tkknowledgepoint.QuestionBankID(*req.QuestionBankId))
		}
		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(tkknowledgepoint.HasQuestionBankWith(
				tkquestionbank.NameContains(*req.QuestionBankName)))
		}
		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.WithQuestionBank().ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			pointDetail.ID = v.ID
			pointDetail.Name = v.Name
			pointDetail.QuestionBankID = v.QuestionBankID
			pointDetail.QuestionCount = v.QuestionCount
			pointDetail.QuestionBankName = ""
			if !app2.IsNil(v.Edges.QuestionBank) {
				pointDetail.QuestionBankName = v.Edges.QuestionBank.Name
			}
			res.List = append(res.List, pointDetail)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//添加（编辑）题目
func SetTkQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.SetTkQuestion
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
			_, err := tk.UpdateTkQuestion(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑题目", req.QuestionStem)
		} else {
			_, err := tk.AddTkQuestion(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加题目", req.QuestionStem)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除题目
func DelTkQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = tk.DelTkQuestion(ctx, req.Id)
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

//获取题目列表（分页）
func GetTkQuestionsByPage(ctx *gin.Context) (interface{}, error) {
	var req request.TkQuestionsList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkQuestionListSuccess{
		List: []response.TkListDetail{},
	}
	TkListDetail := response.TkListDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query := s.TkQuestion.Query().SoftDelete().Where(tkquestion.PidIsNil())

		if !app2.IsNil(req.KnowledgeName) { //知识点
			query = query.Where(tkquestion.HasKnowledgePointsWith(
				tkknowledgepoint.NameContains(*req.KnowledgeName),
				tkknowledgepoint.DeletedAtIsNil()))
		}
		if !app2.IsNil(req.QuestionName) { //题目名称
			query = query.Where(tkquestion.NameContains(*req.QuestionName))
		}
		if !app2.IsNil(req.QuestionType) { //题目类型
			query = query.Where(tkquestion.Type(uint8(*req.QuestionType)))
		}
		if !app2.IsNil(req.Difficulty) { //难易度
			query = query.Where(tkquestion.Difficulty(uint8(*req.Difficulty)))
		}
		if !app2.IsNil(req.QuestionBankId) { //知识点
			query = query.Where(tkquestion.HasQuestionBankWith(
				tkquestionbank.ID(*req.QuestionBankId),
				tkquestionbank.DeletedAtIsNil()))
		}

		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.WithKnowledgePoints(func(query *ent.TkKnowledgePointQuery) {
			query.SoftDelete().Select("name", "id")
		}).WithAdmin(func(query *ent.AdminQuery) {
			query.Select("real_name")
		}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("name")
		}).Order(ent.Desc("id")).ForPage(req.Page, req.PageSize).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			TkListDetail.Id = v.ID
			TkListDetail.Name = v.Name
			TkListDetail.QuestionBankId = v.QuestionBankID
			TkListDetail.Difficulty = int(v.Difficulty)
			TkListDetail.Type = int(v.Type)
			TkListDetail.CreatedAt = *v.CreatedAt
			TkListDetail.QuestionBankName = ""
			if !app2.IsNil(v.Edges.QuestionBank) {
				TkListDetail.QuestionBankName = v.Edges.QuestionBank.Name
			}
			TkListDetail.CreatedAdminName = ""
			if !app2.IsNil(v.Edges.Admin) {
				TkListDetail.CreatedAdminName = v.Edges.Admin.RealName
			}
			TkListDetail.KnowledgePoints = ent.TkKnowledgePoints{}
			if !app2.IsNil(v.Edges.KnowledgePoints) {
				TkListDetail.KnowledgePoints = v.Edges.KnowledgePoints
			}
			res.List = append(res.List, TkListDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//id获取题目详情
func GetTkQuestionsById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkQuestionByIdDetail{}
	tkIdDetail := response.TkQuestionDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		query, err := s.TkQuestion.Query().SoftDelete().Where(tkquestion.ID(*req.Id)).
			WithKnowledgePoints(func(query *ent.TkKnowledgePointQuery) {
				query.SoftDelete().Select("name", "id")
			}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.Select("name")
		}).WithAnswerOptions(func(query *ent.TkQuestionAnswerOptionQuery) {
			query.SoftDelete()
		}).First(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		res.Id = query.ID
		res.Name = query.Name
		res.QuestionBankId = query.QuestionBankID
		res.Difficulty = int(query.Difficulty)
		res.Type = int(query.Type)
		res.CreatedAt = *query.CreatedAt
		res.Desc = query.Desc
		res.KnowledgePoints = ent.TkKnowledgePoints{}
		if !app2.IsNil(query.Edges.KnowledgePoints) {
			res.KnowledgePoints = query.Edges.KnowledgePoints
		}
		res.Options = ent.TkQuestionAnswerOptions{}
		if !app2.IsNil(query.Edges.AnswerOptions) {
			res.Options = query.Edges.AnswerOptions
		}
		res.QuestionBankName = ""
		if !app2.IsNil(query.Edges.QuestionBank) {
			res.QuestionBankName = query.Edges.QuestionBank.Name
		}

		if query.Type == uint8(app.MaterialQuestion) { //材料题
			materialQuestions, err := s.TkQuestion.Query().SoftDelete().
				Where(tkquestion.Pid(*req.Id)).
				WithKnowledgePoints(func(query *ent.TkKnowledgePointQuery) {
					query.SoftDelete().Select("name", "id")
				}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
				query.Select("name")
			}).WithAnswerOptions(func(query *ent.TkQuestionAnswerOptionQuery) {
				query.SoftDelete()
			}).All(ctx)

			if err != nil {
				if err == sql.ErrNoRows {
					res.Materials = []response.TkQuestionDetail{}
					return nil
				}
				return err
			}

			for _, mv := range materialQuestions {
				tkIdDetail.Id = mv.ID
				tkIdDetail.Name = mv.Name
				tkIdDetail.QuestionBankId = mv.QuestionBankID
				tkIdDetail.Difficulty = int(mv.Difficulty)
				tkIdDetail.Type = int(mv.Type)
				tkIdDetail.CreatedAt = *mv.CreatedAt
				tkIdDetail.Desc = mv.Desc
				tkIdDetail.KnowledgePoints = ent.TkKnowledgePoints{}
				if !app2.IsNil(mv.Edges.KnowledgePoints) {
					tkIdDetail.KnowledgePoints = mv.Edges.KnowledgePoints
				}
				tkIdDetail.Options = ent.TkQuestionAnswerOptions{}
				if !app2.IsNil(mv.Edges.AnswerOptions) {
					tkIdDetail.Options = mv.Edges.AnswerOptions
				}
				tkIdDetail.QuestionBankName = ""
				if !app2.IsNil(mv.Edges.QuestionBank) {
					tkIdDetail.QuestionBankName = mv.Edges.QuestionBank.Name
				}
				res.Materials = append(res.Materials, tkIdDetail)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取章节下题目列表（带分页）
func GetTkCsQuestionList(ctx *gin.Context) (interface{}, error) {
	var req request.TkCsQuestionsList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.TkCsQuestionPageList{
		List: []response.TkCsQuestionDetail{},
	}
	tkCsQuestionDetail := response.TkCsQuestionDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		var sectionIds []int
		if !app2.IsNil(req.SectionId) {
			sectionIds = append(sectionIds, *req.SectionId)
		}
		if !app2.IsNil(req.ChapterId) {
			chapterSectionIds, err := s.TkSection.Query().SoftDelete().
				Where(tksection.ChapterID(*req.ChapterId)).
				IDs(ctx)
			if err != nil {
				return err
			}
			sectionIds = append(sectionIds, chapterSectionIds...)
		}

		query := s.TkQuestionSection.Query().SoftDelete().
			Where(tkquestionsection.SectionIDIn(sectionIds...))
		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		////列表
		list, err := query.ForPage(req.Page, req.PageSize).
			WithSectionQuestion(func(query *ent.TkQuestionQuery) {
				query.SoftDelete().Select("id", "name", "type")
			}).
			WithQuestionSection(func(query *ent.TkSectionQuery) {
				query.SoftDelete().Select("id", "name", "chapter_id").WithChapter()
			}).Order(func(s *sql2.Selector) {
			// 连接用户表，以通过 owner-name 和 pet-name 进行排序.
			t := sql2.Table(tkquestion.Table)
			s.Join(t).On(s.C(tkquestionsection.SectionQuestionColumn), t.C(tkquestion.FieldID))
			s.OrderBy(t.C(tkquestion.FieldType))
		}).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			if !app2.IsNil(v.Edges.SectionQuestion) {
				tkCsQuestionDetail.Id = v.Edges.SectionQuestion.ID
				tkCsQuestionDetail.Name = v.Edges.SectionQuestion.Name
				tkCsQuestionDetail.Type = int(v.Edges.SectionQuestion.Type)
				tkCsQuestionDetail.SectionId = v.SectionID
				tkCsQuestionDetail.SectionName = ""
				if !app2.IsNil(v.Edges.QuestionSection) {
					tkCsQuestionDetail.SectionName = v.Edges.QuestionSection.Name
					tkCsQuestionDetail.ChapterId = v.Edges.QuestionSection.ChapterID
					tkCsQuestionDetail.ChapterName = ""
					if !app2.IsNil(v.Edges.QuestionSection.Edges.Chapter) {
						tkCsQuestionDetail.ChapterName = v.Edges.QuestionSection.Edges.Chapter.Name
					}
				}
				res.List = append(res.List, tkCsQuestionDetail)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//为章节添加题目
func SetQuestionToSection(ctx *gin.Context) (interface{}, error) {
	var req request.SetQuestionToSection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := tk.SetQuestionToSection(ctx, *req.SectionId, req.QuestionIds)
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

//移除章节下题目
func RemoveQuestionToSection(ctx *gin.Context) (interface{}, error) {
	var req request.SetQuestionToSection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := tk.RemoveQuestionToSection(ctx, *req.SectionId, req.QuestionIds)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//获取章节下已有题目id数组
func GetQuestionIdsInSection(ctx *gin.Context) (interface{}, error) {
	var req request.GetQuestionIdsInSection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetQuestionIdsInSection{
		QuestionIds: []int{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := s.TkQuestionSection.Query().SoftDelete().
			Where(tkquestionsection.SectionID(req.SectionId)).
			Where(tkquestionsection.QuestionIDNotNil()).
			Select("question_id").Scan(ctx, &res.QuestionIds)

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取题目纠错列表
func GetQuestionErrorFeedbackList(ctx *gin.Context) (interface{}, error) {
	var req request.GetQuestionErrorFeedbackListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetQuestionErrorFeedbackListResponse{
		List: []response.ErrorFeedbackDetail{},
	}
	errorFeedbackDetail := response.ErrorFeedbackDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.TkQuestionErrorFeedback.Query().SoftDelete()

		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(
				tkquestionerrorfeedback.HasQuestionWith(
					tkquestion.HasQuestionBankWith(
						tkquestionbank.NameContains(*req.QuestionBankName))))
		}
		if !app2.IsNil(req.QuestionName) {
			query = query.Where(
				tkquestionerrorfeedback.HasQuestionWith(
					tkquestion.NameContains(*req.QuestionName)))
		}
		if !app2.IsNil(req.Phone) {
			query = query.Where(tkquestionerrorfeedback.Phone(*req.Phone))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(tkquestionerrorfeedback.Status(uint8(*req.Status)))
		}
		if !app2.IsNil(req.ErrorType) {
			query = query.Where(tkquestionerrorfeedback.ErrorType(uint8(*req.ErrorType)))
		}
		if !app2.IsNil(req.ErrorStartAt) && !app2.IsNil(req.ErrorEndAt) {
			query = query.Where(tkquestionerrorfeedback.CreatedAtIn(*req.ErrorStartAt, *req.ErrorEndAt))
		}

		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		////列表
		list, err := query.ForPage(req.Page, req.PageSize).
			WithAdmin().
			WithQuestion(func(query *ent.TkQuestionQuery) {
				query.SoftDelete().WithQuestionBank()
			}).Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			if !app2.IsNil(v.Edges.Question) {
				errorFeedbackDetail.ID = v.ID
				errorFeedbackDetail.QuestionId = v.Edges.Question.ID
				errorFeedbackDetail.QuestionName = v.Edges.Question.Name
				errorFeedbackDetail.Type = int(v.Edges.Question.Type)
				errorFeedbackDetail.QuestionBankName = ""
				errorFeedbackDetail.QuestionBankId = 0
				if !app2.IsNil(v.Edges.Question.Edges.QuestionBank) {
					errorFeedbackDetail.QuestionBankName = v.Edges.Question.Edges.QuestionBank.Name
					errorFeedbackDetail.QuestionBankId = v.Edges.Question.Edges.QuestionBank.ID
				}
				errorFeedbackDetail.CreatedAt = v.CreatedAt
				errorFeedbackDetail.Phone = v.Phone
				errorFeedbackDetail.Username = v.Username
				errorFeedbackDetail.Status = v.Status
				errorFeedbackDetail.ErrorType = v.ErrorType
				errorFeedbackDetail.DealRemark = v.DealRemark
				errorFeedbackDetail.ErrorDesc = v.ErrorDesc
				res.List = append(res.List, errorFeedbackDetail)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//错题反馈处理
func DealTkErrorFeedback(ctx *gin.Context) (interface{}, error) {
	var req request.DealTkErrorFeedbackRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	tk := app.TkQuestionBank{}
	operatorAdminId := ctx.GetInt(app.GlobalAdminId)
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := tk.DealTkErrorFeedback(ctx, *req.Id, *req.Status, *req.DealRemark, operatorAdminId)
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

//获取题库下各题型题目数量
func GetTkBankQuestionTypeCount(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.BankQuestionTypeCount{
		Detail: []response.BankQuestionTypeCountDetail{},
	}
	bankQuestionTypeCountDetail := response.BankQuestionTypeCountDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		questionbank, err := s.TkQuestionBank.Query().Where(tkquestionbank.ID(req.Id)).First(ctx)
		if err != nil {
			return err
		}
		res.QuestionBankId = questionbank.ID
		res.QuestionBankName = questionbank.Name

		query := s.TkQuestion.Query().SoftDelete().
			Where(tkquestion.PidIsNil()).
			Where(tkquestion.QuestionBankID(req.Id)).
			Select("type")

		//单选
		selectCount, err := query.Clone().Where(tkquestion.Type(uint8(app.SingleSelect))).Count(ctx)
		if err != nil {
			return err
		}
		bankQuestionTypeCountDetail.QuestionType = app.SingleSelect
		bankQuestionTypeCountDetail.QuestionCount = selectCount
		res.Detail = append(res.Detail, bankQuestionTypeCountDetail)

		//多选
		multiCount, err := query.Clone().Where(tkquestion.Type(uint8(app.MultiSelect))).Count(ctx)
		if err != nil {
			return err
		}
		bankQuestionTypeCountDetail.QuestionType = app.MultiSelect
		bankQuestionTypeCountDetail.QuestionCount = multiCount
		res.Detail = append(res.Detail, bankQuestionTypeCountDetail)

		//多选
		shortCount, err := query.Clone().Where(tkquestion.Type(uint8(app.ShortAnswer))).Count(ctx)
		if err != nil {
			return err
		}
		bankQuestionTypeCountDetail.QuestionType = app.ShortAnswer
		bankQuestionTypeCountDetail.QuestionCount = shortCount
		res.Detail = append(res.Detail, bankQuestionTypeCountDetail)

		//判断
		judgeCount, err := query.Clone().Where(tkquestion.Type(uint8(app.JudgeQuestion))).Count(ctx)
		if err != nil {
			return err
		}
		bankQuestionTypeCountDetail.QuestionType = app.JudgeQuestion
		bankQuestionTypeCountDetail.QuestionCount = judgeCount
		res.Detail = append(res.Detail, bankQuestionTypeCountDetail)

		//材料
		materialCount, err := query.Clone().Where(tkquestion.Type(uint8(app.MaterialQuestion))).Count(ctx)
		if err != nil {
			return err
		}
		bankQuestionTypeCountDetail.QuestionType = app.MaterialQuestion
		bankQuestionTypeCountDetail.QuestionCount = materialCount
		res.Detail = append(res.Detail, bankQuestionTypeCountDetail)

		totalCount, err := query.Clone().Count(ctx)
		if err != nil {
			return err
		}
		res.TotalCount = totalCount
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取题库id name列表
func GetTkSimpleQuestionBank(ctx *gin.Context) (interface{}, error) {
	res := []response.SimpleQuestionBank{}
	err := store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		list, err := s.TkQuestionBank.Query().
			SoftDelete().
			Select("id", "name").
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SimpleQuestionBank{}
			detail.Id = v.ID
			detail.Name = v.Name
			res = append(res, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//上传题目
func UploadQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.ImportQuestionRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}
	fileDetail, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileDetail.Close()
	//获取文件名
	dirname := "static/storage/import/questions/"
	fileName := fmt.Sprintf("%s%d%s", dirname, time.Now().UnixNano(), file.Filename)
	fmt.Println("文件名：", fileName)

	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		return nil, err
	}

	if err := ctx.SaveUploadedFile(file, fileName); err != nil {
		return nil, err
	}

	tk := app.TkQuestionBank{}
	it := app.ImportTask{}
	taskId, err := it.AddImportTask(ctx, 0, "导入题目", fileName)
	if err != nil {
		return nil, err
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := tk.UploadQuestion(ctx, fileDetail, taskId, *req.QuestionBankId)
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

//获取做题记录统计列表
func GetQuestionBankStatisticList(ctx *gin.Context) (interface{}, error) {
	var req request.GetQuestionBankStatisticListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetQuestionBankStatisticListResponse{
		List: []response.QuestionBankStatisticDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.TkUserQuestionBankRecord.Query().SoftDelete().
			Where(tkuserquestionbankrecord.UserID(*req.UserId))

		if !app2.IsNil(req.QuestionBankName) {
			query = query.Where(tkuserquestionbankrecord.HasQuestionBankWith(
				tkquestionbank.NameContains(*req.QuestionBankName)))
		}
		if !app2.IsNil(req.CateId) {
			query = query.Where(tkuserquestionbankrecord.HasQuestionBankWith(
				tkquestionbank.ItemCategoryID(*req.CateId)))
		}

		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.ForPage(req.Page, req.PageSize).
			WithUser(func(query *ent.UserQuery) {
				query.Select("id", "username")
			}).
			WithQuestionBank(
				func(query *ent.TkQuestionBankQuery) {
					query.WithItemCategory(
						func(query *ent.ItemCategoryQuery) {
							query.WithParent()
						})
				}).
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			detail := response.QuestionBankStatisticDetail{}
			detail.RecordCount = v.RecordCount
			detail.CorrectRate = v.CorrectRate
			detail.UserId = v.UserID
			detail.QuestionBankId = v.QuestionBankID
			detail.CateId = 0
			detail.Username = ""
			detail.QuestionBankName = ""
			detail.CateName = ""
			if !app2.IsNil(v.Edges.User) {
				detail.Username = v.Edges.User.Username
			}
			if !app2.IsNil(v.Edges.QuestionBank) {
				detail.QuestionBankName = v.Edges.QuestionBank.Name
				detail.CateId = v.Edges.QuestionBank.ItemCategoryID
				detail.QuestionCount = v.Edges.QuestionBank.QuestionCount
				if detail.QuestionCount > 0 {
					finishRate, err := app2.DivRate(v.RecordCount, detail.QuestionCount)

					if err != nil {
						return err
					}
					if detail.FinishRate > 1 {
						detail.FinishRate = 1
					} else {
						detail.FinishRate = finishRate
					}
				}

				if !app2.IsNil(v.Edges.QuestionBank.Edges.ItemCategory) {
					detail.CateName = v.Edges.QuestionBank.Edges.ItemCategory.Name
					if !app2.IsNil(v.Edges.QuestionBank.Edges.ItemCategory.Edges.Parent) {
						detail.CateName = v.Edges.QuestionBank.Edges.ItemCategory.Edges.Parent.Name + "-" + detail.CateName
					}
				}
			}
			res.List = append(res.List, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取用户做题记录统计详细列表
func GetUserQuestionBankStatisticDetailList(ctx *gin.Context) (interface{}, error) {
	var req request.GetUserQuestionBankStatisticDetailListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetUserQuestionBankStatisticDetailListResponse{
		List: []response.UserQuestionBankStatisticDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.TkUserQuestionRecord.Query().SoftDelete().
			Where(tkuserquestionrecord.UserID(*req.UserId)).
			Where(tkuserquestionrecord.QuestionBankID(*req.QuestionBankId))

		if !app2.IsNil(req.QuestionName) {
			query = query.Where(tkuserquestionrecord.HasQuestionWith(
				tkquestion.NameContains(*req.QuestionName)))
		}
		if !app2.IsNil(req.QuestionType) {
			query = query.Where(tkuserquestionrecord.HasQuestionWith(
				tkquestion.Type(uint8(*req.QuestionType))))
		}

		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.ForPage(req.Page, req.PageSize).
			WithQuestion(func(query *ent.TkQuestionQuery) {
				query.WithAnswerOptions()
			}).
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			detail := response.UserQuestionBankStatisticDetail{}
			detail.Id = v.ID
			detail.UserId = v.UserID
			detail.QuestionBankId = v.QuestionBankID
			detail.QuestionId = v.QuestionID
			detail.QuestionName = ""
			detail.UserAnswer = v.Answer
			detail.IsRight = int(v.IsRight)
			detail.QuestionType = 0
			detail.CorrectRate = 0
			detail.Options = ent.TkQuestionAnswerOptions{}
			if !app2.IsNil(v.Edges.Question) {
				detail.QuestionName = v.Edges.Question.Name
				detail.QuestionType = int(v.Edges.Question.Type)
				if v.Edges.Question.AnswerCount > 0 {
					detail.CorrectRate, err = app2.DivRate(v.Edges.Question.RightCount, v.Edges.Question.AnswerCount)
					if err != nil {
						return err
					}
				}
				if !app2.IsNil(v.Edges.Question.Edges.AnswerOptions) && len(v.Edges.Question.Edges.AnswerOptions) > 0 {
					detail.Options = v.Edges.Question.Edges.AnswerOptions
				}
			}

			res.List = append(res.List, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//学生数据
func GetStudentStatisticList(ctx *gin.Context) (interface{}, error) {
	var req request.GetStudentStatisticRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetStudentStatisticResponse{
		List: []response.GetStudentStatisticDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.User.Query().SoftDelete()

		if !app2.IsNil(req.Phone) {
			query = query.Where(user.PhoneContains(*req.Phone))
		}
		if !app2.IsNil(req.Username) {
			query = query.Where(user.UsernameContains(*req.Username))
		}

		//分页
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		//列表
		list, err := query.
			WithCity().WithCate().
			ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			detail := response.GetStudentStatisticDetail{}
			detail.Username = v.Username
			detail.UserId = v.ID
			detail.Phone = v.Phone
			detail.CityId = v.FromCityID
			detail.CateId = v.FromItemCategoryID
			detail.CityName = ""
			detail.CateName = ""
			detail.AttendanceRate = 0
			detail.CourseFinishRate = 0
			detail.QuestionFinishRate = 0

			if !app2.IsNil(v.Edges.City) {
				detail.CityName = v.Edges.City.Name
			}
			if !app2.IsNil(v.Edges.Cate) {
				detail.CateName = v.Edges.Cate.Name
			}

			res.List = append(res.List, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
