package admin

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/kcclass"
	"gserver/internal/store/ent/kcclassteacher"
	"gserver/internal/store/ent/kccourse"
	"gserver/internal/store/ent/kcuserclass"
	"gserver/internal/store/ent/user"
	app2 "gserver/pkg/requestparam"
)

//添加（编辑）班级
func SetKcClass(ctx *gin.Context) (interface{}, error) {
	var req request.SetClass
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	successRsp := response.IdSuccess{Id: 0}
	cl := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			id, err := cl.UpdateClass(ctx, req)
			if err != nil {
				return err
			}
			successRsp.Id = id
		} else {
			id, err := cl.AddClass(ctx, req)
			if err != nil {
				return err
			}
			successRsp.Id = id
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return successRsp, nil
}

//删除班级
func DelClass(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelClass(ctx, req.Id)
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

//设置班级状态
func SetClassStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetClassStatus(ctx, *req.Id, *req.Status)
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

//班级详情
func ClassDetailById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.ClassDetail{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		class, err := store.WithContext(ctx).KcClass.
			Query().SoftDelete().
			Where(kcclass.ID(*req.Id)).
			WithItem().
			WithCity().
			WithMajors().
			WithAttachment().
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		res.Id = class.ID
		res.ClassTitle = class.ClassTitle
		res.ClassDesc = class.ClassDesc
		res.ClassCode = class.ClassCode
		res.Price = class.Price
		res.IsDisplay = int(class.IsDisplay)
		res.ClassPeriodsType = int(class.ClassPeriodType)
		res.Status = int(class.Status)
		res.CateId = class.CateID
		res.StudentCount = class.StudentCount
		res.CourseCount = class.CourseCount
		res.Majors = ent.Majors{}
		res.CityName = ""
		res.CityId = class.CityID
		res.CateName = ""
		res.CateId = class.CateID
		res.ClassCoverImgId = class.ClassCoverImgID
		res.ClassCoverImgUrl = ""
		if !app2.IsNil(class.Edges.Majors) {
			res.Majors = class.Edges.Majors
		}
		if !app2.IsNil(class.Edges.City) {
			res.CityName = class.Edges.City.Name
		}
		if !app2.IsNil(class.Edges.Item) {
			res.CateName = class.Edges.Item.Name
		}
		if !app2.IsNil(class.Edges.Attachment) {
			res.ClassCoverImgUrl = app2.GetOssHost() + class.Edges.Attachment.Filename
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//班级列表（分页）
func ClassPageList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.ClassPageListResponse{
		List: []response.ClassPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcClass.Query().SoftDelete()
		if !app2.IsNil(req.ClassTitle) {
			query = query.Where(kcclass.ClassTitleContains(*req.ClassTitle))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(kcclass.Status(uint8(*req.Status)))
		}
		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		published, err := query.Clone().Where(kcclass.Status(1)).Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		unpublished, err := query.Clone().Where(kcclass.Status(2)).Count(ctx)
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
			}).ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.ClassPageDetail{}
			detail.ID = v.ID
			detail.ClassTitle = v.ClassTitle
			detail.Status = v.Status
			detail.ClassCode = v.ClassCode
			detail.Price = v.Price
			detail.CreatedAdminID = v.CreatedAdminID
			detail.CourseCount = v.CourseCount
			detail.StudentCount = v.StudentCount
			detail.CreatedAt = v.CreatedAt
			detail.CateID = v.CateID
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

//获取班级下课程列表（分页）
func GetClassCoursePageList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassCoursePageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.ClassCoursePageListResponse{
		List: []response.ClassCourseDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourse.Query().SoftDelete().
			Where(kccourse.HasClassesWith(kcclass.ID(*req.Id)))

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.WithTeacher(func(query *ent.TeacherQuery) {
				query.WithAttachment()
			})
		}).
			ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.ClassCourseDetail{}
			detail.CourseId = v.ID
			detail.CourseName = v.CourseName
			detail.Type = int(v.CourseType)
			detail.Price = v.CoursePrice
			detail.PeriodsType = 3
			detail.CourseCoverImgId = v.CourseCoverImgID
			detail.CourseCoverImgUrl = ""
			detail.TeacherId = 0
			detail.TeacherName = ""
			detail.TeacherAvatar = ""
			if !app2.IsNil(v.Edges.CourseTeachers) && len(v.Edges.CourseTeachers) > 0 {
				teacher := v.Edges.CourseTeachers[len(v.Edges.CourseTeachers)-1]
				detail.TeacherId = teacher.TeacherID
				if !app2.IsNil(teacher.Edges.Teacher) {
					detail.TeacherName = teacher.Edges.Teacher.Name
					if !app2.IsNil(teacher.Edges.Teacher.Edges.Attachment) {
						detail.TeacherAvatar = app2.GetOssHost() + teacher.Edges.Teacher.Edges.Attachment.Filename
					}
				}
			}
			if !app2.IsNil(v.Edges.Attachment) {
				detail.CourseCoverImgUrl = app2.GetOssHost() + v.Edges.Attachment.Filename
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

//获取班级下 已选中 课程列表（分页）
func GetSelectedClassCoursePageList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassCoursePageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.ClassCoursePageListResponse{
		List: []response.ClassCourseDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcCourse.Query().SoftDelete()
		if !app2.IsNil(req.CourseName) {
			query = query.Where(kccourse.CourseNameContains(*req.CourseName))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithClasses(func(query *ent.KcClassQuery) {
			query.SoftDelete().Where(kcclass.ID(*req.Id))
		}).WithCourseTeachers(func(query *ent.KcCourseTeacherQuery) {
			query.WithTeacher()
		}).
			ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.ClassCourseDetail{}
			detail.CourseId = v.ID
			detail.CourseName = v.CourseName
			detail.Type = int(v.CourseType)
			detail.Price = v.CoursePrice
			detail.PeriodsType = 3
			detail.CourseCoverImgId = v.CourseCoverImgID
			detail.CourseCoverImgUrl = ""
			detail.TeacherId = 0
			detail.TeacherName = ""
			detail.IsSelected = 0
			if !app2.IsNil(v.Edges.CourseTeachers) && len(v.Edges.CourseTeachers) > 0 {
				teacher := v.Edges.CourseTeachers[len(v.Edges.CourseTeachers)-1]
				detail.TeacherId = teacher.TeacherID
				if !app2.IsNil(teacher.Edges.Teacher) {
					detail.TeacherName = teacher.Edges.Teacher.Name

				}
			}
			if !app2.IsNil(v.Edges.Attachment) {
				detail.CourseCoverImgUrl = app2.GetOssHost() + v.Edges.Attachment.Filename
			}

			if !app2.IsNil(v.Edges.Classes) && len(v.Edges.Classes) > 0 {
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

//添加班级课程
func AddClassCourse(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassCourse
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddClassCourse(ctx, req)
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

//移除班级课程
func RemoveClassCourse(ctx *gin.Context) (interface{}, error) {
	var req request.RemoveClassCourse
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveClassCourse(ctx, req)
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

//添加班级老师
func AddClassTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddClassTeacher(ctx, req)
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

//删除班级老师
func RemoveClassTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveClassTeacher(ctx, req)
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

//获取班级老师列表（分页）
func GetClassTeacherPageList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.ClassTeacherPageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcClassTeacher.Query().
			Where(kcclassteacher.ClassID(*req.ClassId)).
			Where(kcclassteacher.TeacherIDNotNil())
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
		}).ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.ClassTeacherDetail{}
			detail.ID = v.ID
			detail.ClassID = v.ClassID
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

//获取班级已添加teacher id列表
func GetSelectedClassTeacherIdsList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetTeacherIdsInClass{
		TeacherIds: []int{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := s.KcClassTeacher.Query().
			Where(kcclassteacher.ClassID(*req.ClassId)).
			Where(kcclassteacher.TeacherIDNotNil()).
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

//设置班级老师
func SetClassTeacherStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetCourseTeacherStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetClassTeacherStatus(ctx, *req.Id, *req.ShowStatus)
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

//添加班级用户
func AddClassUser(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassUser
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddClassUser(ctx, req)
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

//删除班级用户
func RemoveClassUser(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassUser
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveClassUser(ctx, req)
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
func AddClassUserValidity(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassUserValidity
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err = c.SetClassUserValidity(ctx, req)
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

//获取班级用户列表（分页）
func GetClassUserPageList(ctx *gin.Context) (interface{}, error) {
	var req request.GetCourseUserPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.ClassUserPageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).KcUserClass.Query().
			Where(kcuserclass.ClassID(*req.Id)).
			Where(kcuserclass.UserIDNotNil())
		if !app2.IsNil(req.Username) {
			query = query.Where(kcuserclass.HasUserWith(
				user.UsernameContains(*req.Username)))
		}
		if !app2.IsNil(req.Phone) {
			query = query.Where(kcuserclass.HasUserWith(
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
			detail := response.ClassUserDetail{}
			detail.ID = v.ID
			detail.ClassID = v.ClassID
			detail.UserID = v.UserID
			detail.StudyRate = v.StudyRate
			detail.PeriodType = v.PeriodType
			detail.ClosingDate = v.ClosingDate
			detail.CreatedAt = v.CreatedAt
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

//获取班级已添加teacher id列表
func GetSelectedClassUserIdsList(ctx *gin.Context) (interface{}, error) {
	var req request.ClassIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetUserIdsInClass{
		UserIds: []int{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		err := s.KcUserClass.Query().
			Where(kcuserclass.ClassID(*req.ClassId)).
			Where(kcuserclass.UserIDNotNil()).
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

//添加班级班主任
func AddClassMasterTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetClassTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.AddClassMasterTeacher(ctx, req)
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

//删除班级班主任
func RemoveClassMaster(ctx *gin.Context) (interface{}, error) {
	var req request.ClassIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.KcClass{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.RemoveClassMasterTeacher(ctx, req)
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

//班级详情
func GetClassMasterTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.ClassIdRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetClassMasterTeacher{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		class, err := store.WithContext(ctx).KcClass.
			Query().SoftDelete().
			Where(kcclass.ID(*req.ClassId)).
			WithMasterTeachers(func(query *ent.TeacherQuery) {
				query.WithAttachment()
			}).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		res.ClassId = class.ID
		res.ClassHeadMasterName = ""
		res.ClassHeadMasterId = class.ClassHeadMasterID
		res.Avatar = ""

		if !app2.IsNil(class.Edges.MasterTeachers) {
			res.ClassHeadMasterName = class.Edges.MasterTeachers.Name
			if !app2.IsNil(class.Edges.MasterTeachers.Edges.Attachment) {
				res.Avatar = app2.GetOssHost() + class.Edges.MasterTeachers.Edges.Attachment.Filename
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
