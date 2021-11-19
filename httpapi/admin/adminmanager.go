package admin

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/config"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/admin"
	"gserver/internal/store/ent/role"
	"gserver/internal/store/ent/rolepermission"
	"gserver/internal/store/ent/teacher"
	"gserver/internal/store/ent/user"
	"gserver/pkg/dingding"
	"gserver/pkg/password"
	app2 "gserver/pkg/requestparam"
	"io/ioutil"
	"net/http"
)

//管理员登录
func LoginAdmin(ctx *gin.Context) (interface{}, error) {
	s := store.WithContext(ctx)
	var req request.LoginAdmin
	var token string
	var res response.AdminLoginSuccess
	am := app.AdminManager{}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	adminDetail, err := s.Admin.Query().
		Where(admin.Phone(req.Phone)).
		Where(admin.IsActive(1)).
		WithRoles(func(query *ent.RoleQuery) {
			query.SoftDelete()
		}).WithAdminAttachments().First(ctx)
	if errno := app.CheckDataNotFound(err); errno != nil {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "账号不存在",
		})
	}

	token, err = am.MakeAdminToken(ctx, adminDetail.ID, struct {
		Ip      string
		City    string
		Browser string
	}{Ip: app2.RemoteIp(ctx.Request), City: req.CityName, Browser: ctx.Request.UserAgent()})
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	if !password.Comp(req.Password+adminDetail.Salt, adminDetail.Password) {
		return nil, errorno.NewErr(errorno.LoginFailError)
	}
	res.AdminId = adminDetail.ID
	res.Token = token
	res.Email = adminDetail.Email
	res.Phone = adminDetail.Phone
	res.RealName = adminDetail.RealName
	res.RoleId = 0
	res.AdminAvatarId = adminDetail.AdminAvatarID
	res.AvatarUrl = ""
	res.PermissionIds = []string{}
	if !app2.IsNil(adminDetail.Edges.Roles) && len(adminDetail.Edges.Roles) > 0 {
		roleInfo := adminDetail.Edges.Roles[len(adminDetail.Edges.Roles)-1]
		res.RoleId = roleInfo.ID
		permission, err := s.RolePermission.Query().
			Where(rolepermission.RoleID(roleInfo.ID)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		if ent.IsNotFound(err) {
			res.PermissionIds = []string{}
		} else {
			err := json.Unmarshal([]byte(permission.PermissionID), &res.PermissionIds)
			if err != nil {
				return nil, err
			}
		}
	}
	if !app2.IsNil(adminDetail.Edges.AdminAttachments) {
		res.AvatarUrl = app2.GetOssHost() + adminDetail.Edges.AdminAttachments.Filename
	}
	return res, nil
}

//从boss后台同步学员信息
func SyncUserFromBoss(ctx *gin.Context) (interface{}, error) {
	var req request.SyncUserFromBoss
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = am.SyncUserFromBoss(ctx, req)
		return err
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//从boss后台同步管理员信息
func SyncAdminFromBoss(ctx *gin.Context) (interface{}, error) {
	var req request.SyncAdminFromBoss
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = am.SyncAdminFromBoss(ctx, req)
		return err
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//添加角色
func AddRole(ctx *gin.Context) (interface{}, error) {
	var req request.AddRole
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) {
			_, err := am.UpdateRole(ctx, req.RoleName, *req.Id, req.Status)
			if err != nil {
				return err
			}
		} else {
			_, err := am.AddRole(ctx, req.RoleName, req.Status)
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

//管理员添加角色
func AddAdminRole(ctx *gin.Context) (interface{}, error) {
	var req request.AddAdminRole
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := am.AddAdminRole(ctx, *req.AdminId, req.RoleIds)
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

//添加权限
func AddPermission(ctx *gin.Context) (interface{}, error) {
	var req request.AddPermission
	var id int
	var res response.IdSuccess
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		roleId, err := am.AddPermission(ctx, *req.PermissionName, *req.Desc, uint32(*req.Pid))
		if err != nil {
			return err
		}
		id = roleId
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.Id = id
	return res, nil
}

//角色添加权限
func AddRolePermission(ctx *gin.Context) (interface{}, error) {
	var req request.AddRolePermission
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err := am.AddRolePermission(ctx, *req.RoleId, req.PermissionIds)
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

//删除角色
func DelRole(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.DelRole(ctx, req.Id)
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

//根据条件获取学员姓名手机号列表
func GetSimpleUserInfo(ctx *gin.Context) (interface{}, error) {
	var req request.GetUserInfo
	res := response.SimpleUserInfoByPage{
		List: []response.SimpleUserInfoResponse{},
	}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.User.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Phone) {
			query = query.Where(user.PhoneContains(req.Phone))
		}
		if !app2.StringIsEmpty(req.Username) {
			query = query.Where(user.UsernameContains(req.Username))
		}
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return err
		}
		res.Page.Total = count

		resp, err := query.WithCate(func(query *ent.ItemCategoryQuery) {
			query.WithParent()
		}).WithCity().ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range resp {
			resDetail := response.SimpleUserInfoResponse{}
			resDetail.Id = v.ID
			resDetail.Username = v.Username
			resDetail.Phone = v.Phone
			resDetail.AvatarUrl = app2.GetOssHost() + v.Avatar
			resDetail.IdCard = v.IDCard
			resDetail.CardType = int(v.CardType)
			resDetail.CreatedAt = *v.CreatedAt
			resDetail.Sex = int(v.Sex)
			if !app2.IsNil(v.Birthday) {
				resDetail.Birthday = *v.Birthday
			}
			resDetail.CityId = v.FromCityID
			resDetail.CateId = v.FromItemCategoryID
			resDetail.ParentCateId = 0
			resDetail.CityName = ""
			resDetail.CateName = ""
			resDetail.ParentCateName = ""
			if !app2.IsNil(v.Edges.City) {
				resDetail.CityName = v.Edges.City.Name
			}
			if !app2.IsNil(v.Edges.Cate) {
				resDetail.CateName = v.Edges.Cate.Name
				if !app2.IsNil(v.Edges.Cate.Edges.Parent) {
					resDetail.ParentCateName = v.Edges.Cate.Edges.Parent.Name
					resDetail.ParentCateId = v.Edges.Cate.Edges.Parent.ID
				}
			}
			res.List = append(res.List, resDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//根据条件获取老师姓名手机号列表
func GetSimpleTeacherInfo(ctx *gin.Context) (interface{}, error) {
	var req request.GetTeacherInfo
	res := response.SimpleTeacherInfoByPage{
		List: []response.SimpleTeacherInfoResponse{},
	}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)

		query := s.Teacher.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Phone) {
			query = query.Where(teacher.PhoneContains(req.Phone))
		}
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(teacher.NameContains(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(teacher.Status(uint8(*req.Status)))
		}

		count, err := query.Clone().Count(ctx)
		if err != nil {
			return err
		}
		res.Page.Total = count

		resp, err := query.WithAttachment().
			WithMajors().
			WithTeacherCourses(func(query *ent.KcCourseTeacherQuery) {
				query.WithCourse()
			}).ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range resp {
			resDetail := response.SimpleTeacherInfoResponse{}
			resDetail.Id = v.ID
			resDetail.Name = v.Name
			resDetail.Status = int(v.Status)
			resDetail.Phone = v.Phone
			resDetail.Sex = int(v.Sex)
			resDetail.CreatedAt = v.CreatedAt
			resDetail.Majors = ent.Majors{}
			resDetail.Courses = []ent.KcCourse{}
			resDetail.AvatarUrl = ""
			if !app2.IsNil(v.Edges.Attachment) {
				resDetail.AvatarUrl = app2.GetOssHost() + v.Edges.Attachment.Filename
			}

			if !app2.IsNil(v.Edges.Majors) && len(v.Edges.Majors) > 0 {
				resDetail.Majors = v.Edges.Majors
			}

			if !app2.IsNil(v.Edges.TeacherCourses) {
				if !app2.IsNil(v.Edges.TeacherCourses) {
					for _, c := range v.Edges.TeacherCourses {
						resDetail.Courses = append(resDetail.Courses, *c.Edges.Course)
					}

				}
			}

			res.List = append(res.List, resDetail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//管理员列表（分页）
func GetAdminListByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetAdminPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetAdminPageListResponse{
		List: []response.AdminDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Admin.Query().SoftDelete()
		if !app2.IsNil(req.Status) {
			query = query.Where(admin.Status(uint8(*req.Status)))
		}
		if !app2.IsNil(req.Phone) {
			query = query.Where(admin.PhoneContains(*req.Phone))
		}
		if !app2.IsNil(req.RealName) {
			query = query.Where(admin.RealNameContains(*req.RealName))
		}
		if !app2.IsNil(req.RoleId) {
			query = query.Where(admin.HasRolesWith(role.ID(*req.RoleId)))
		}

		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		query = query.WithRoles().ForPage(req.Page, req.PageSize).Order(ent.Desc("id"))

		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			detail := response.AdminDetail{}
			detail.ID = v.ID
			detail.RealName = v.RealName
			detail.Phone = v.Phone
			detail.Email = v.Email
			detail.Status = v.Status
			detail.Platform = v.Platform
			detail.CreatedAt = v.CreatedAt
			detail.Remark = v.Remark
			detail.Password = "******"
			detail.RoleId = 0
			detail.RoleName = ""
			if !app2.IsNil(v.Edges.Roles) && len(v.Edges.Roles) > 0 {
				role := v.Edges.Roles[len(v.Edges.Roles)-1]
				detail.RoleId = role.ID
				detail.RoleName = role.Name
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

//（编辑）管理员
func SetAdmin(ctx *gin.Context) (interface{}, error) {
	var req request.SetAdminRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err := a.UpdateAdmin(ctx, req)
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

//设置管理员状态
func SetAdminStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetAdminStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.SetAdminStatus(ctx, *req.Id, *req.Status)
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

//设置角色状态
func SetRoleStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetAdminStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.SetRoleStatus(ctx, *req.Id, *req.Status)
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

//获取角色列表（无分页）
func GetRoleList(ctx *gin.Context) (interface{}, error) {
	res := []response.RoleDetail{}
	result, err := store.WithContext(ctx).Role.
		Query().SoftDelete().
		Select("id", "name", "status").Order(ent.Desc("id")).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		detail := response.RoleDetail{}
		detail.Id = v.ID
		detail.Name = v.Name
		detail.Status = int(v.Status)
		res = append(res, detail)
	}
	return res, nil
}

//角色列表（分页）
func GetRoleListByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetRoleList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.RolePageList{
		List: []response.RoleDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Role.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(role.NameContains(*req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(role.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		query = query.ForPage(req.Page, req.PageSize).Order(ent.Desc("id"))

		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		for _, v := range list {
			detail := response.RoleDetail{}
			detail.Id = v.ID
			detail.Name = v.Name
			detail.Status = int(v.Status)
			res.List = append(res.List, detail)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//角色列表（分页）
func GetRolePermission(ctx *gin.Context) (interface{}, error) {
	var req request.RoleId
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.RolePermissionsResponse{
		PermissionIds: []string{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query, err := store.WithContext(ctx).RolePermission.Query().
			Where(rolepermission.RoleID(*req.RoleId)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				fmt.Println("xxxx")
				return nil
			}
			return err
		}
		fmt.Println(query)

		err = json.Unmarshal([]byte(query.PermissionID), &res.PermissionIds)
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

//修改密码
func ResetPassword(ctx *gin.Context) (interface{}, error) {
	var req request.ResetPassword
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.ResetPassword(ctx, req)
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

//上传头像
func UpdateAdminAvatar(ctx *gin.Context) (interface{}, error) {
	var req request.UpdateAdminAvatar
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.UpdateAdminAvatar(ctx, req)
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

//添加（编辑）老师
func SetTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.SetTeacher
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := a.UpdateTeacher(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := a.AddTeacher(ctx, req)
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

//删除老师
func DelTeacher(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.DelTeacher(ctx, req.Id)
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

//设置老师状态
func SetTeacherStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	a := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = a.SetTeacherStatus(ctx, req.Id, req.Status)
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

//id获取老师信息
func GetTeacherDetailById(ctx *gin.Context) (interface{}, error) {
	var req request.IdPtrOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	resDetail := response.TeacherDetailResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		teacherInfo, err := store.WithContext(ctx).Teacher.
			Query().SoftDelete().
			Where(teacher.ID(*req.Id)).
			WithAttachment().
			WithMajors().
			WithTeacherTags().
			WithTeacherCourses(
				func(query *ent.KcCourseTeacherQuery) {
					query.WithCourse()
				}).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		if err != nil {
			return err
		}

		resDetail.ID = teacherInfo.ID
		resDetail.Name = teacherInfo.Name
		resDetail.Status = teacherInfo.Status
		resDetail.Phone = teacherInfo.Phone
		resDetail.Sex = teacherInfo.Sex
		resDetail.CreatedAt = teacherInfo.CreatedAt
		resDetail.Detail = teacherInfo.Detail
		resDetail.SortOrder = teacherInfo.SortOrder
		resDetail.Nickname = teacherInfo.Nickname
		resDetail.Email = teacherInfo.Email
		resDetail.AvatarID = teacherInfo.AvatarID
		resDetail.TeachingAge = teacherInfo.TeachingAge
		resDetail.SubTitle = teacherInfo.SubTitle
		resDetail.Majors = ent.Majors{}
		resDetail.Courses = []ent.KcCourse{}
		resDetail.TeacherTags = ent.TeacherTags{}
		resDetail.AvatarUrl = ""
		if !app2.IsNil(teacherInfo.Edges.Attachment) {
			resDetail.AvatarUrl = app2.GetOssHost() + teacherInfo.Edges.Attachment.Filename
		}

		if !app2.IsNil(teacherInfo.Edges.Majors) && len(teacherInfo.Edges.Majors) > 0 {
			resDetail.Majors = teacherInfo.Edges.Majors
		}
		if !app2.IsNil(teacherInfo.Edges.TeacherTags) && len(teacherInfo.Edges.TeacherTags) > 0 {
			resDetail.TeacherTags = teacherInfo.Edges.TeacherTags
		}

		if !app2.IsNil(teacherInfo.Edges.TeacherCourses) {
			if !app2.IsNil(teacherInfo.Edges.TeacherCourses) {
				for _, c := range teacherInfo.Edges.TeacherCourses {
					resDetail.Courses = append(resDetail.Courses, *c.Edges.Course)
				}

			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resDetail, nil
}

//扫码登录
func ThirdLoginAdmin(ctx *gin.Context) (interface{}, error) {
	s := store.WithContext(ctx)
	var req request.ScanLogin
	var token string
	var res response.AdminLoginSuccess
	am := app.AdminManager{}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	ding := dingding.Dding{}
	ThirdId, err := ding.LoginByQRcode(*req.Code)
	if err != nil {
		return nil, err
	}

	adminDetail, err := s.Admin.Query().
		Where(admin.ThirdOpenid(ThirdId)).
		Where(admin.IsActive(1)).
		WithRoles(func(query *ent.RoleQuery) {
			query.SoftDelete()
		}).WithAdminAttachments().First(ctx)
	if errno := app.CheckDataNotFound(err); errno != nil {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "账号不存在",
		})
	}
	token, err = am.MakeAdminToken(ctx, adminDetail.ID, struct {
		Ip      string
		City    string
		Browser string
	}{Ip: app2.RemoteIp(ctx.Request), City: "", Browser: ctx.Request.UserAgent()})
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	res.AdminId = adminDetail.ID
	res.Token = token
	res.Email = adminDetail.Email
	res.Phone = adminDetail.Phone
	res.RealName = adminDetail.RealName
	res.RoleId = 0
	res.AdminAvatarId = adminDetail.AdminAvatarID
	res.AvatarUrl = ""
	res.PermissionIds = []string{}
	if !app2.IsNil(adminDetail.Edges.Roles) && len(adminDetail.Edges.Roles) > 0 {
		roleInfo := adminDetail.Edges.Roles[len(adminDetail.Edges.Roles)-1]
		res.RoleId = roleInfo.ID
		permission, err := s.RolePermission.Query().
			Where(rolepermission.RoleID(roleInfo.ID)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		if ent.IsNotFound(err) {
			res.PermissionIds = []string{}
		} else {
			err := json.Unmarshal([]byte(permission.PermissionID), &res.PermissionIds)
			if err != nil {
				return nil, err
			}
		}
	}
	if !app2.IsNil(adminDetail.Edges.AdminAttachments) {
		res.AvatarUrl = app2.GetOssHost() + adminDetail.Edges.AdminAttachments.Filename
	}
	return res, nil
}

//id获取工作台学员信息
func GetBossUserByPhone(ctx *gin.Context) (interface{}, error) {
	var req request.PhoneRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	resDetail := response.UserInfoDetail{}

	reqDetail := make(map[string]string)
	reqDetail["phone"] = *req.Phone
	bytesData, err := json.Marshal(reqDetail)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	url := config.ServerConfig.BossHost + "/api/v3/user_sync/by_phone"
	requestInfo, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	requestInfo.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(requestInfo)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tempMap map[string]interface{}
	err = json.Unmarshal(respBytes, &tempMap)
	if err != nil {
		return nil, err
	}

	if tempMap["res_info"] != "ok" {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "数据不存在",
		})
	}

	userDetailData, err := json.Marshal(tempMap["data"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userDetailData, &resDetail)
	if err != nil {
		return nil, err
	}
	return resDetail, nil
}

//添加（编辑）用户
func SetUser(ctx *gin.Context) (interface{}, error) {
	var req request.SetUser
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	am := app.AdminManager{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) {
			_, err := am.UpdateUser(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := am.AddUser(ctx, req)
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
