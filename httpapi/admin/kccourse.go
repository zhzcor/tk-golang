package admin

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"tkserver/httpapi/admin/request"
	"tkserver/httpapi/admin/response"
	"tkserver/internal/app"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/kccourse"
	"tkserver/internal/store/ent/kccoursesmallcategory"
	"tkserver/internal/store/ent/kccourseteacher"
	"tkserver/internal/store/ent/kcusercourse"
	"tkserver/internal/store/ent/kcvideouploadtask"
	"tkserver/internal/store/ent/tkquestionbank"
	"tkserver/internal/store/ent/user"
	"tkserver/pkg/baijiayun"
	app2 "tkserver/pkg/requestparam"
)

//添加（编辑）课程
func SetCourse(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourse
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	cs := app.Course{}
	successRsp := response.IdSuccess{Id: 0}
	operatorAdminId := ctx.GetInt(app.GlobalAdminId)
	cm := app.Common{}
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			id, err := cs.UpdateCourse(ctx, req, operatorAdminId)
			if err != nil {
				return err
			}
			successRsp.Id = id
			remark = fmt.Sprintf("%s:%s", "编辑课程", *req.CourseName)
		} else {
			id, err := cs.AddCourse(ctx, req, operatorAdminId)
			if err != nil {
				return err
			}
			successRsp.Id = id
			remark = fmt.Sprintf("%s:%s", "添加课程", *req.CourseName)
		}
		return nil
	})
	if err != nil {
		return successRsp, err
	}
	_ = cm.SetOperationLog(ctx, operatorAdminId, remark)
	return successRsp, nil
}

//删除课程
func DelCourse(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelCourse(ctx, req.Id)
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

//设置课程状态
func SetCourseStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetCourseStatus(ctx, *req.Id, *req.PushStatus)
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

//课程列表（分页）
func CoursePageList(ctx *gin.Context) (interface{}, error) {
	var req request.CoursePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.CoursePageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourse.Query().SoftDelete()
		if !app2.IsNil(req.CourseName) {
			query = query.Where(kccourse.CourseName(*req.CourseName))
		}
		if !app2.IsNil(req.PushStatus) {
			query = query.Where(kccourse.PushStatus(uint8(*req.PushStatus)))
		}
		if !app2.IsNil(req.CourseType) && len(req.CourseType) > 0 {
			query = query.Where(kccourse.CourseTypeIn(req.CourseType...))
		}
		if !app2.IsNil(req.CateId) {
			query = query.Where(kccourse.CateID(*req.CateId))
		}
		if !app2.IsNil(req.CreatedAdminName) {
			query = query.Where(kccourse.HasAdminWith(admin.RealName(*req.CreatedAdminName)))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		published, err := query.Clone().Where(kccourse.PushStatus(1)).Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		unpublished, err := query.Clone().Where(kccourse.PushStatus(2)).Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Statistic.Total = count
		res.Statistic.Published = published
		res.Statistic.Unpublished = unpublished
		res.Statistic.Closed = count - published - unpublished
		res.Page.Total = count
		list, err := query.WithAdmin().
			WithItem(func(query *ent.ItemCategoryQuery) {
				query.WithParent()
			}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.CoursePageDetail{}
			detail.ID = v.ID
			detail.CourseName = v.CourseName
			detail.CourseType = v.CourseType
			detail.CoursePrice = v.CoursePrice
			detail.CreatedAdminID = v.CreatedAdminID
			detail.PushStatus = v.PushStatus
			detail.PeopleNum = v.PeopleNum
			detail.CourseDesc = v.CourseDesc
			detail.CreatedAt = v.CreatedAt
			detail.ItemId = v.CateID
			detail.CityID = v.CityID
			detail.CreatedAdminName = ""
			detail.ItemCategoryName = ""
			if !app2.IsNil(v.Edges.Admin) {
				detail.CreatedAdminName = v.Edges.Admin.RealName
			}
			if !app2.IsNil(v.Edges.Item) {
				detail.ItemCategoryName = v.Edges.Item.Name
				if !app2.IsNil(v.Edges.Item.Edges.Parent) {
					detail.ItemCategoryName = v.Edges.Item.Edges.Parent.Name + "-" + v.Edges.Item.Name
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

//课程id name列表
func CourseSimpleList(ctx *gin.Context) (interface{}, error) {
	res := []response.CourseSimple{}
	err := store.WithTx(ctx, func(ctx context.Context) error {
		list, err := store.WithContext(ctx).KcCourse.
			Query().SoftDelete().
			Select("id", "course_name").
			Order(ent.Desc("id")).
			All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.CourseSimple{}
			detail.Id = v.ID
			detail.CourseName = v.CourseName
			res = append(res, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//添加课程老师
func AddCourseTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddCourseTeacher(ctx, req)
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

//设置课程状态
func SetCourseTeacherStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseTeacherStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetCourseTeacherStatus(ctx, *req.Id, *req.ShowStatus)
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

//删除课程老师
func RemoveCourseTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveCourseTeacher(ctx, req)
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

//获取课程老师列表（分页）
func GetCourseTeacherPageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetCourseTeacherPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.CourseTeacherPageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourseTeacher.Query().
			Where(kccourseteacher.CourseID(*req.Id)).
			Where(kccourseteacher.TeacherIDNotNil())
		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithTeacher(func(query *ent.TeacherQuery) {
			query.WithAttachment()
		}).
			ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.CourseTeacherDetail{}
			detail.ID = v.ID
			detail.CourseID = v.CourseID
			detail.TeacherID = v.TeacherID
			detail.SortOrder = v.SortOrder
			detail.ShowStatus = v.ShowStatus
			detail.TeacherName = ""
			detail.Avatar = ""
			if !app2.IsNil(v.Edges.Teacher) {
				detail.TeacherName = v.Edges.Teacher.Name
				if !app2.IsNil(v.Edges.Teacher.Edges.Attachment) {
					detail.Avatar = app2.GetOssHost() + v.Edges.Teacher.Edges.Attachment.Filename
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

//添加课程用户
func AddCourseUser(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseUser
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddCourseUser(ctx, req)
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

//删除课程用户
func RemoveCourseUser(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseUser
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveCourseUser(ctx, req)
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

//修改学员有效期
func AddCourseUserValidity(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseUserValidity
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetCourseUserValidity(ctx, req)
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

//获取课程用户列表（分页）
func GetCourseUserPageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetCourseUserPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.CourseUserPageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcUserCourse.Query().
			Where(kcusercourse.CourseID(*req.Id)).
			Where(kcusercourse.UserIDNotNil())
		if !app2.IsNil(req.Username) {
			query = query.Where(kcusercourse.HasUserWith(
				user.UsernameContains(*req.Username)))
		}
		if !app2.IsNil(req.Phone) {
			query = query.Where(kcusercourse.HasUserWith(
				user.PhoneContains(*req.Phone)))
		}
		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithUser(func(query *ent.UserQuery) {
			query.SoftDelete().Select("id", "username", "phone", "avatar")
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.CourseUserDetail{}
			detail.ID = v.ID
			detail.CourseID = v.CourseID
			detail.UserID = v.UserID
			detail.StudyRate = v.StudyRate
			detail.CreatedAt = v.CreatedAt
			detail.PeriodType = v.PeriodType
			detail.ClosingDate = v.ClosingDate
			detail.Username = ""
			detail.Phone = ""
			detail.Avatar = ""
			if !app2.IsNil(v.Edges.User) {
				detail.Username = v.Edges.User.Username
				detail.Phone = v.Edges.User.Phone
				detail.Avatar = app2.GetOssHost() + v.Edges.User.Avatar
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

//获取课程用户列表（分页）
func GetUserDetailById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.UserDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		user, err := store.WithContext(ctx).User.Query().
			Where(user.ID(*req.Id)).
			SoftDelete().WithCity().
			WithCate(func(query *ent.ItemCategoryQuery) {
				query.WithParent()
			}).
			First(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.ID = user.ID
		res.Phone = user.Phone
		res.Username = user.Username
		res.Nickname = user.Nickname
		res.Birthday = user.Birthday
		res.Email = user.Email
		res.CardType = user.CardType
		res.IDCard = user.IDCard
		res.Sex = user.Sex
		res.SignRemark = user.SignRemark
		res.CityName = ""
		res.ItemCategoryName = ""
		res.ParentItemCategoryName = ""
		if !app2.IsNil(user.Edges.City) {
			res.CityName = user.Edges.City.Name
		}
		if !app2.IsNil(user.Edges.Cate) {
			res.ItemCategoryName = user.Edges.Cate.Name
			if !app2.IsNil(user.Edges.Cate.Edges.Parent) {
				res.ParentItemCategoryName = user.Edges.Cate.Edges.Parent.Name
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//设置用户课程有效期
func SetCourseUserValidity(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseUserValidity
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetCourseUserValidity(ctx, req)
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

//添加课程用户
func AddCourseQuestionBank(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseQuestionBank
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddCourseQuestionBank(ctx, req)
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

//删除课程用户
func RemoveCourseQuestionBank(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseQuestionBank
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveCourseQuestionBank(ctx, req)
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

//获取课程题库详情
func GetCourseQuestionBankDetail(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseQuestionBank
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.CourseQuestionBank{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		course, err := store.WithContext(ctx).KcCourse.
			Query().SoftDelete().
			Where(kccourse.ID(*req.CourseId)).
			Select("id", "question_bank_id").
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		res.CourseId = course.ID
		res.QuestionBankId = 0
		res.QuestionBankName = ""
		res.ItemCategoryName = ""
		res.ParentItemCategoryName = ""
		if !app2.IsNil(course.QuestionBankID) {
			res.QuestionBankId = course.QuestionBankID
			bank, err := store.WithContext(ctx).TkQuestionBank.Query().
				Where(tkquestionbank.ID(course.QuestionBankID)).
				WithItemCategory(func(query *ent.ItemCategoryQuery) {
					query.WithParent()
				}).First(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					return nil
				}
				return err
			}
			res.QuestionBankName = bank.Name
			if !app2.IsNil(bank.Edges.ItemCategory) {
				res.ItemCategoryName = bank.Edges.ItemCategory.Name
				if !app2.IsNil(bank.Edges.ItemCategory.Edges.Parent) {
					res.ParentItemCategoryName = bank.Edges.ItemCategory.Edges.Parent.Name
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

//添加(编辑)课程 章
func SetKcCourseChapter(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseChapter
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateKcCourseChapter(ctx, *req.Id, *req.Title, *req.CourseId)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddKcCourseChapter(ctx, *req.Title, *req.CourseId)
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
func DelKcCourseChapter(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelKcCourseChapter(ctx, req.Id)
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

//添加(编辑)课程 节
func SetKcCourseSection(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseSection
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateKcCourseSection(ctx, *req.Id, *req.Title, *req.ChapterId)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddKcCourseSection(ctx, *req.Title, *req.ChapterId)
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
func DelKcCourseSection(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelKcCourseSection(ctx, req.Id)
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

//获取课程已添加teacher id列表
func GetSelectedCourserTeacherIdsList(ctx *gin.Context) (interface{}, error) {
	var req request.CourseIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetTeacherIdsInCourse{
		TeacherIds: []int{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := s.KcCourseTeacher.Query().
			Where(kccourseteacher.CourseID(*req.CourseId)).
			Where(kccourseteacher.TeacherIDNotNil()).
			Select("teacher_id").Scan(ctx, &res.TeacherIds)

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

//获取课程已添加user id列表
func GetSelectedCourserUserIdsList(ctx *gin.Context) (interface{}, error) {
	var req request.CourseIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetUserIdsInCourse{
		UserIds: []int{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := s.KcUserCourse.Query().
			Where(kcusercourse.CourseID(*req.CourseId)).
			Where(kcusercourse.UserIDNotNil()).
			Select("user_id").Scan(ctx, &res.UserIds)

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

//添加伪直播
func SetFalseVideo(ctx *gin.Context) (interface{}, error) {
	var req request.SetFalseVideoRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			err := c.UpdateFalseVideo(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddFalseVideo(ctx, req)
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

//删除伪直播
func DelFalseVideo(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Course{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelFalseVideo(ctx, req.Id)
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

//获取课程用户列表（分页）
func GetFalseVideoPageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetFalseVideoPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetFalseVideoPageListResponse{
		List: ent.KcVideoUploadTasks{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcVideoUploadTask.Query().SoftDelete().
			Where(kcvideouploadtask.CourseID(*req.CourseId)).
			Where(kcvideouploadtask.Type(2))
		if !app2.IsNil(req.Title) {
			query = query.Where(kcvideouploadtask.TitleContains(*req.Title))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(kcvideouploadtask.Status(*req.Status))
		}
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

//设置百家云回调地址
func SetBjyCallbackUrl(ctx *gin.Context) (interface{}, error) {
	err := store.WithTx(ctx, func(ctx context.Context) error {
		bjy := baijiayun.BjyCloud{}
		//设置视频转码回调地址
		err := bjy.SetTranscodeCallbackUrl()
		if err != nil {
			return err
		}
		//设置上下课回调地址
		err = bjy.SetClassCallbackUrl()
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

//获取课时详情
func GetSmallCateById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.CourseSmallCategoryById{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		smallCate, err := s.KcCourseSmallCategory.Query().
			Where(kccoursesmallcategory.ID(*req.Id)).
			SoftDelete().WithCsAttachment().
			First(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Id = smallCate.ID
		res.CourseId = smallCate.CourseID
		res.SectionId = smallCate.SectionID
		res.ChapterId = smallCate.ChapterID
		res.Type = int(smallCate.Type)
		res.SmallName = smallCate.SmallName
		res.LiveSmallTime = smallCate.LiveSmallTime
		res.LiveSmallStart = smallCate.LiveSmallStart
		res.LiveSmallRemark = smallCate.LiveSmallRemark
		res.LiveSmallRemark = smallCate.LiveSmallRemark
		res.FinishType = int(smallCate.FinishType)
		res.ViewingTime = smallCate.ViewingTime
		res.FalseVideoId = smallCate.FalseVideoID
		res.OrderVideoAttachId = smallCate.OrderVideoAttachID
		res.CoursewareAttachId = smallCate.CoursewareAttachID
		res.FalseVideoName = ""
		res.OrderVideoName = ""
		res.CoursewareAttachUrl = ""
		res.CoursewareName = smallCate.CoursewareName

		if !app2.IsNil(smallCate.Edges.CsAttachment) {
			res.CoursewareAttachUrl = app2.GetOssHost() + smallCate.Edges.CsAttachment.Filename
		}

		FalseTask, err := s.KcVideoUploadTask.Query().SoftDelete().
			Where(kcvideouploadtask.CourseID(smallCate.CourseID)).
			Where(kcvideouploadtask.Type(2)).
			Where(kcvideouploadtask.VideoID(smallCate.FalseVideoID)).
			Order(ent.Desc("id")).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		if !ent.IsNotFound(err) {
			res.FalseVideoName = FalseTask.VideoName
		}

		OrderTask, err := s.KcVideoUploadTask.Query().SoftDelete().
			Where(kcvideouploadtask.CourseID(smallCate.CourseID)).
			Where(kcvideouploadtask.Type(1)).
			Where(kcvideouploadtask.VideoID(smallCate.OrderVideoID)).
			Order(ent.Desc("id")).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		if !ent.IsNotFound(err) {
			res.OrderVideoName = OrderTask.VideoName
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
