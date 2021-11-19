package admin

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tkserver/httpapi/admin/request"
	"tkserver/httpapi/admin/response"
	"tkserver/internal/app"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kcsmallcategoryattachment"
	"tkserver/internal/store/ent/kcsmallcategoryexampaper"
	"tkserver/internal/store/ent/kcsmallcategoryquestion"
	"tkserver/internal/store/ent/tkexampaper"
	"tkserver/internal/store/ent/tkquestion"
	"tkserver/pkg/log"
	app2 "tkserver/pkg/requestparam"
)

//添加(编辑)课时
func SetKcCourseSmallCategory(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseSmallCategory
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	cm := app.Common{}
	adminId := ctx.GetInt(app.GlobalAdminId)
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateKcCourseSmallCategory(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑课程", *req.SmallName)
		} else {
			_, err := c.AddKcCourseSmallCategory(ctx, req)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加课程", *req.SmallName)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除课时
func DelKcCourseSmallCategory(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelKcSmallCategory(ctx, req.Id)
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

//添加课时试卷
func SetKcSmallCategoryExamPaper(ctx *gin.Context) (interface{}, error) {
	var req request.SetSmallCateExamPagerRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetSmallCategoryExamPaper(ctx, req)
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

//添加课时题目
func SetKcSmallCategoryQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.SetSmallCateQuestionRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetSmallCategoryQuestion(ctx, req)
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

//添加课时附件
func SetKcSmallCategoryAttachment(ctx *gin.Context) (interface{}, error) {
	var req request.SetSmallCateAttachmentRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetSmallCategoryAttachment(ctx, req)
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

//移除课时试卷
func RemoveKcSmallCategoryExamPaper(ctx *gin.Context) (interface{}, error) {
	var req request.RemoveSmallCateExamPagerRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.RemoKcSmallCategoryExamPaper(ctx, req)
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

//移除课时题目
func RemoveKcSmallCategoryQuestion(ctx *gin.Context) (interface{}, error) {
	var req request.RemoveSmallCateQuestionRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.RemoveSmallCategoryQuestion(ctx, req)
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

//移除课时附件
func RemoveKcSmallCategoryAttachment(ctx *gin.Context) (interface{}, error) {
	var req request.RemoveSmallCateAttachmentRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.RemoveSmallCategoryAttachment(ctx, req)
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

//获取课程下所有章节课时（无分页）
func GetCourseSmallCateList(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetCourseSmallCateList{}
	s := store.WithContext(ctx)
	result, err := s.KcCourse.Query().SoftDelete().
		Where(kccourse.ID(*req.Id)).WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
		query.SoftDelete().Where(kccoursesmallcategory.ChapterIDIsNil(), kccoursesmallcategory.SectionIDIsNil()).
			Select("id", "small_name", "live_small_time", "push_status", "type",
				"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count")
	}).WithCourseChapters(func(query *ent.KcCourseChapterQuery) {
		query.WithCourseSmallChapters(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Select("id", "small_name", "live_small_time", "push_status", "type",
				"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count")
		}).WithChapterSections(func(query *ent.KcCourseSectionQuery) {
			query.WithCourseSmallSections(func(query *ent.KcCourseSmallCategoryQuery) {
				query.SoftDelete().Select("id", "small_name", "live_small_time", "push_status", "type",
					"course_id", "section_id", "chapter_id", "attachment_count", "question_count", "exam_count", "homework_count")
			})
		})
	}).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return res, nil
		}
		return nil, err
	}

	smallCateCount, err := s.KcCourseSmallCategory.Query().SoftDelete().
		Where(kccoursesmallcategory.CourseID(*req.Id)).Count(ctx)
	if err != nil {
		return nil, errorno.NewStoreErr(err)
	}

	res.CourseId = result.ID
	res.CourseName = result.CourseName
	res.SmallCount = smallCateCount
	res.SmallCategory = []response.SmallCateDetail{}
	res.Chapters = []response.ChapterDetail{}
	if !app2.IsNil(result.Edges.CourseSmallCategorys) {
		for _, v := range result.Edges.CourseSmallCategorys { //课程课时
			detail := response.SmallCateDetail{}
			detail.Id = v.ID
			detail.SmallName = v.SmallName
			detail.Duration = v.LiveSmallTime
			detail.Status = int(v.PushStatus)
			detail.AttachmentCount = v.AttachmentCount
			detail.QuestionCount = v.QuestionCount
			detail.ExamCount = v.ExamCount
			detail.HomeworkCount = v.HomeworkCount
			detail.Type = int(v.Type)
			res.SmallCategory = append(res.SmallCategory, detail)
		}
	}

	if !app2.IsNil(result.Edges.CourseChapters) {
		for _, cv := range result.Edges.CourseChapters { //章
			chapterDetail := response.ChapterDetail{}
			chapterDetail.ChapterId = cv.ID
			chapterDetail.ChapterName = cv.Title
			chapterDetail.SmallCategory = []response.SmallCateDetail{}
			chapterDetail.Sections = []response.SectionDetail{}
			if !app2.IsNil(cv.Edges.CourseSmallChapters) { //章课时
				for _, cs := range cv.Edges.CourseSmallChapters {
					cSDetail := response.SmallCateDetail{}
					cSDetail.Id = cs.ID
					cSDetail.SmallName = cs.SmallName
					cSDetail.Duration = cs.LiveSmallTime
					cSDetail.Status = int(cs.PushStatus)
					cSDetail.AttachmentCount = cs.AttachmentCount
					cSDetail.QuestionCount = cs.QuestionCount
					cSDetail.ExamCount = cs.ExamCount
					cSDetail.HomeworkCount = cs.HomeworkCount
					chapterDetail.SmallCategory = append(chapterDetail.SmallCategory, cSDetail)
				}
			}

			if !app2.IsNil(cv.Edges.ChapterSections) {
				for _, sc := range cv.Edges.ChapterSections { //节
					sectionDetail := response.SectionDetail{}
					sectionDetail.SectionId = sc.ID
					sectionDetail.SectionName = sc.Title
					sectionDetail.SmallCategory = []response.SmallCateDetail{}
					if !app2.IsNil(sc.Edges.CourseSmallSections) {
						for _, s := range sc.Edges.CourseSmallSections { //节课时
							sDetail := response.SmallCateDetail{}
							sDetail.Id = s.ID
							sDetail.SmallName = s.SmallName
							sDetail.Duration = s.LiveSmallTime
							sDetail.Status = int(s.PushStatus)
							sDetail.AttachmentCount = s.AttachmentCount
							sDetail.QuestionCount = s.QuestionCount
							sDetail.ExamCount = s.ExamCount
							sDetail.HomeworkCount = s.HomeworkCount
							sectionDetail.SmallCategory = append(sectionDetail.SmallCategory, sDetail)
						}
					}
					chapterDetail.Sections = append(chapterDetail.Sections, sectionDetail)
				}
			}
			res.Chapters = append(res.Chapters, chapterDetail)
		}
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

//id获取课程信息
func GetCourseDetailById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.CourseDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		course, err := store.WithContext(ctx).KcCourse.
			Query().SoftDelete().
			Where(kccourse.ID(*req.Id)).
			WithItem().
			WithCity().
			WithMajor().
			WithAttachment().WithCourseSmallCategorys(func(query *ent.KcCourseSmallCategoryQuery) {
			query.SoftDelete().Select("id", "course_id", "small_name", "live_small_start", "live_small_time")
		}).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		res.Id = course.ID
		res.CourseName = course.CourseName
		res.CoursePrice = course.CoursePrice
		res.PushStatus = int(course.PushStatus)
		res.CourseType = int(course.CourseType)
		res.CourseDesc = course.CourseDesc
		res.CreatedAt = *course.CreatedAt
		res.Majors = ent.Majors{}
		res.CityName = ""
		res.CityId = course.CityID
		res.ItemName = ""
		res.ItemId = course.CateID
		res.CourseCoverImgId = course.CourseCoverImgID
		res.CourseCoverImgUrl = ""
		res.PeopleNum = course.PeopleNum
		if !app2.IsNil(course.Edges.Major) {
			res.Majors = course.Edges.Major
		}
		if !app2.IsNil(course.Edges.City) {
			res.CityName = course.Edges.City.Name
		}
		if !app2.IsNil(course.Edges.Item) {
			res.ItemName = course.Edges.Item.Name
		}
		if !app2.IsNil(course.Edges.City) {
			res.CityName = course.Edges.City.Name
		}
		if !app2.IsNil(course.Edges.Attachment) {
			res.CourseCoverImgUrl = app2.GetOssHost() + course.Edges.Attachment.Filename
		}
		res.OpenLiveDuration = 0
		if course.CourseType == 3 && !app2.IsNil(course.Edges.CourseSmallCategorys) && len(course.Edges.CourseSmallCategorys) > 0 { //直播公开课
			smallCourse := course.Edges.CourseSmallCategorys[len(course.Edges.CourseSmallCategorys)-1]
			res.OpenLiveDuration = smallCourse.LiveSmallTime
			res.OpenLiveStart = smallCourse.LiveSmallStart
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//课时试卷、作业列表（分页）
func GetSmallCourseExamPageList(ctx *gin.Context) (interface{}, error) {
	var req request.SmallCourseExamPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SmallCourseExamPageListResponse{
		List: []response.SmallCourseExamPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).TkExamPaper.Query().SoftDelete()

		if !app2.IsNil(req.QuestionBankId) {
			query = query.Where(tkexampaper.QuestionBankID(*req.QuestionBankId))
		}

		if !app2.IsNil(req.ExamPaperType) {
			query = query.Where(tkexampaper.ExamQuestionType(uint8(*req.ExamPaperType)))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithCourseExamPapers(func(query *ent.KcSmallCategoryExamPaperQuery) {
			query.SoftDelete().Where(kcsmallcategoryexampaper.Type(uint8(*req.Type))).
				Where(kcsmallcategoryexampaper.SmallCategoryID(*req.SmallCategoryId))
		}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete()
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SmallCourseExamPageDetail{}
			detail.ExamPaperId = v.ID
			detail.ExamPaperName = v.Name
			detail.ExamPaperType = int(v.ExamQuestionType)
			detail.QuestionBankId = v.QuestionBankID
			detail.QuestionBankName = ""
			detail.IsSelected = 0
			if !app2.IsNil(v.Edges.QuestionBank) {
				detail.QuestionBankName = v.Edges.QuestionBank.Name
			}
			if !app2.IsNil(v.Edges.CourseExamPapers) && len(v.Edges.CourseExamPapers) > 0 {
				detail.IsSelected = 1
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

//已选中 课时试卷、作业列表（分页）
func GetSelectedSmallCourseExamPageList(ctx *gin.Context) (interface{}, error) {
	var req request.SelectedSmallCourseExamPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SmallCourseExamPageListResponse{
		List: []response.SmallCourseExamPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).TkExamPaper.Query().SoftDelete().
			Where(tkexampaper.HasCourseExamPapersWith(
				kcsmallcategoryexampaper.SmallCategoryID(*req.SmallCategoryId),
				kcsmallcategoryexampaper.Type(uint8(*req.Type))),
			)

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete()
		}).ForPage(req.Page, req.PageSize).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SmallCourseExamPageDetail{}
			detail.ExamPaperId = v.ID
			detail.ExamPaperName = v.Name
			detail.ExamPaperType = int(v.ExamQuestionType)
			detail.QuestionBankId = v.QuestionBankID
			detail.QuestionBankName = ""
			detail.IsSelected = 1
			if !app2.IsNil(v.Edges.QuestionBank) {
				detail.QuestionBankName = v.Edges.QuestionBank.Name
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

//课时题目列表（分页）
func GetSmallCourseQuestionPageList(ctx *gin.Context) (interface{}, error) {
	var req request.SmallCourseQuestionPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SmallCourseQuestionPageListResponse{
		List: []response.SmallCourseQuestionPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).TkQuestion.Query().SoftDelete().
			Where(tkquestion.PidIsNil())

		if !app2.IsNil(req.QuestionBankId) {
			query = query.Where(tkquestion.QuestionBankID(*req.QuestionBankId))
		}

		if !app2.IsNil(req.QuestionType) {
			query = query.Where(tkquestion.Type(uint8(*req.QuestionType)))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithSmallCategoryQuestions(func(query *ent.KcSmallCategoryQuestionQuery) {
			query.SoftDelete().Where(kcsmallcategoryquestion.SmallCategoryID(*req.SmallCategoryId))
		}).WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete()
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SmallCourseQuestionPageDetail{}
			detail.QuestionId = v.ID
			detail.QuestionName = v.Name
			detail.QuestionType = int(v.Type)
			detail.QuestionBankId = v.QuestionBankID
			detail.QuestionBankName = ""
			detail.IsSelected = 0
			if !app2.IsNil(v.Edges.QuestionBank) {
				detail.QuestionBankName = v.Edges.QuestionBank.Name
			}
			if !app2.IsNil(v.Edges.SmallCategoryQuestions) && len(v.Edges.SmallCategoryQuestions) > 0 {
				detail.IsSelected = 1
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

//已选中 课时题目列表（分页）
func GetSelectedSmallCourseQuestionPageList(ctx *gin.Context) (interface{}, error) {
	var req request.SmallCourseQuestionPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SmallCourseQuestionPageListResponse{
		List: []response.SmallCourseQuestionPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).TkQuestion.Query().SoftDelete().
			Where(tkquestion.PidIsNil()).Where(tkquestion.HasSmallCategoryQuestionsWith(
			kcsmallcategoryquestion.SmallCategoryID(*req.SmallCategoryId)))

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithQuestionBank(func(query *ent.TkQuestionBankQuery) {
			query.SoftDelete()
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SmallCourseQuestionPageDetail{}
			detail.QuestionId = v.ID
			detail.QuestionName = v.Name
			detail.QuestionType = int(v.Type)
			detail.QuestionBankId = v.QuestionBankID
			detail.QuestionBankName = ""
			detail.IsSelected = 1
			if !app2.IsNil(v.Edges.QuestionBank) {
				detail.QuestionBankName = v.Edges.QuestionBank.Name
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

//已选中 课时资料列表（分页）
func GetSelectedSmallCourseAttachmentPageList(ctx *gin.Context) (interface{}, error) {
	var req request.SelectedSmallCourseAttachmentPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SmallCourseAttachmentPageListResponse{
		List: []response.SmallCourseAttachmentPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcSmallCategoryAttachment.Query().SoftDelete().
			Where(kcsmallcategoryattachment.SmallCategoryID(*req.SmallCategoryId))

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithAttachment().ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SmallCourseAttachmentPageDetail{}
			detail.AttachmentId = v.AttachmentID
			detail.AttachmentName = v.AttachmentName
			detail.AttachmentUrl = ""
			if !app2.IsNil(v.Edges.Attachment) {
				detail.AttachmentUrl = app2.GetOssHost() + v.Edges.Attachment.Filename
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

//获取回放列表
func GetLiveBackPageList(ctx *gin.Context) (interface{}, error) {
	var req request.CourseIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetLiveBackPageListResponse{
		List: ent.KcCourseSmallCategories{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourseSmallCategory.Query().SoftDelete().
			Where(kccoursesmallcategory.CourseID(*req.CourseId)).
			Where(kccoursesmallcategory.Type(3)).
			Where(kccoursesmallcategory.LiveSmallStatusIn(4, 5, 6, 7)).
			Select("id", "small_name", "updated_at", "live_small_status", "course_id")

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.List = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//回放重新录制
func UpdateLiveBack(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.UpdateLiveBack(ctx, *req.Id)
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

//点播替换回放
func ReplaceLiveBack(ctx *gin.Context) (interface{}, error) {
	var req request.ReplaceLiveBack
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.ReplaceLiveBack(ctx, *req.Id, *req.AttachmentId, req.VideoName)
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

//获取直播列表
func GetLivePageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetLiveSmallCatePageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetLiveSmallCatePageList{
		List: []response.LiveSmallCateDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourseSmallCategory.Query().SoftDelete().
			Where(kccoursesmallcategory.Type(3))

		if !app2.IsNil(req.Status) && len(req.Status) > 0 {
			query = query.Where(kccoursesmallcategory.LiveSmallStatusIn(req.Status...))
		}
		if !app2.IsNil(req.SmallName) {
			query = query.Where(kccoursesmallcategory.SmallNameContains(*req.SmallName))
		}
		if !app2.IsNil(req.LiveStartAt) && !app2.IsNil(req.LiveEndAt) {
			query = query.Where(kccoursesmallcategory.
				LiveSmallStartGT(*req.LiveStartAt), kccoursesmallcategory.LiveSmallStartLT(*req.LiveEndAt))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithCourse().ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.LiveSmallCateDetail{}
			detail.SmallName = v.SmallName
			detail.Id = v.ID
			detail.Status = int(v.LiveSmallStatus)
			detail.LiveSmallStart = v.LiveSmallStart
			detail.LiveSmallTime = v.LiveSmallTime
			detail.CourseId = v.CourseID
			detail.CourseName = ""
			if !app2.IsNil(v.Edges.Course) {
				detail.CourseName = v.Edges.Course.CourseName
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

//上下课回调处理
func ClassCallbackHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.BjyCallbackResponse{
			Code: 1,
		}
		var req request.ClassCallbackRequest
		err := ctx.Bind(&req)
		if err != nil {
			res.Msg = "参数错误"
			log.Error(res)
			log.Info(ctx.Request.MultipartForm)
			ctx.AbortWithStatusJSON(http.StatusOK, res)
			return
		}

		logRes, err := json.Marshal(req)
		if err != nil {
			log.Info(err.Error())
		} else {
			log.Info(string(logRes))
		}
		s := store.WithContext(ctx)
		smallCate, err := s.KcCourseSmallCategory.Query().
			Where(kccoursesmallcategory.LiveRoomID(req.RoomId)).
			First(ctx)
		if err != nil {
			res.Msg = "数据服务异常"
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusOK, res)
			return
		}
		status := 1
		switch req.Op {
		case "start":
			status = 2
			break
		case "end":
			status = 3
			break
		}

		_, err = s.KcCourseSmallCategory.
			UpdateOneID(smallCate.ID).
			SetLiveSmallStatus(int8(status)).
			Save(ctx)
		if err != nil {
			log.Error(err.Error())
		}
		res.Code = 0
		ctx.JSON(http.StatusOK, res)
	}
}

//点播，回放处理
func VideosCallbackHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.BjyCallbackResponse{
			Code: 1,
		}
		var req request.VideoCallbackRequest
		err := ctx.Bind(&req)
		if err != nil {
			res.Msg = "参数错误"
			log.Error(res)
			log.Error(1111, ctx.Request.Header)
			log.Error(2222, ctx.Request.PostForm)
			log.Info(ctx.Request.MultipartForm)
			log.Info(33333, err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		logRes, err := json.Marshal(req)
		if err != nil {
			log.Info(err.Error())
		} else {
			log.Info(string(logRes))
		}

		err = store.WithTx(ctx, func(ctx context.Context) error {
			c := app.Course{}
			if !app2.IsNil(req.RoomId) { //回放
				err := c.ProcessCallbackBack(ctx, req)
				if err != nil {
					return err
				}
			} else { //点播视频
				err := c.ProcessCallbackOrder(ctx, req)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			res.Msg = "数据服务异常"
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusOK, res)
			return
		}
		res.Code = 0
		ctx.JSON(http.StatusOK, res)
	}
}
