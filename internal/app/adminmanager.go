package app

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
	"tkserver/httpapi/admin/request"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/admin"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/kccourseteacher"
	"tkserver/internal/store/ent/permission"
	"tkserver/internal/store/ent/role"
	"tkserver/internal/store/ent/teacher"
	"tkserver/internal/store/ent/user"
	"tkserver/pkg/password"
	app "tkserver/pkg/requestparam"
)

const GlobalAdminId string = "GlobalAdminId"

type AdminManager struct {
}

func (a AdminManager) AddRole(ctx context.Context, roleName string, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Role.
		Query().
		Where(role.Name(roleName)).
		Where(role.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.RoleExistsError)
	}
	newRole, err := s.Role.Create().
		SetName(roleName).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newRole.ID, nil
}

//编辑角色
func (a AdminManager) UpdateRole(ctx context.Context, roleName string, id, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Role.
		Query().SoftDelete().
		Where(role.Name(roleName)).
		Where(role.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.Role.UpdateOneID(id).
		SetName(roleName).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除角色
func (a AdminManager) DelRole(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.RoleIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Role.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断角色是否存着
func (a AdminManager) RoleIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Role.
		Query().SoftDelete().
		Where(role.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "角色不存在",
		})
	}
	return nil
}

//为管理员添加角色
func (a AdminManager) AddAdminRole(ctx context.Context, adminId int, roleIds []int) error {
	s := store.WithContext(ctx)
	fined, err := s.Admin.
		Query().
		Where(admin.ID(adminId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.AdminNotExistsError)
	}
	err = s.Admin.UpdateOneID(adminId).ClearRoles().Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	_, err = s.Admin.UpdateOneID(adminId).AddRoleIDs(roleIds...).Save(ctx)

	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//为角色添加权限
func (a AdminManager) AddRolePermission(ctx context.Context, roleId int, permissionIds []string) error {
	s := store.WithContext(ctx)
	fined, err := s.Role.
		Query().
		Where(role.ID(roleId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.RoleNotExistsError)
	}
	err = s.Role.UpdateOneID(roleId).ClearRolePermission().Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	strPermissionIds, err := json.Marshal(permissionIds)
	if err != nil {
		return err
	}

	_, err = s.RolePermission.Create().
		SetRoleID(roleId).
		SetPermissionID(string(strPermissionIds)).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

func (a AdminManager) AddPermission(ctx context.Context, permissionName, desc string, pid uint32) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Permission.
		Query().
		Where(permission.Name(permissionName)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.PermissionExistsError)
	}
	newPermission, err := s.Permission.Create().
		SetName(permissionName).
		SetDesc(desc).
		SetPid(pid).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newPermission.ID, nil

}

// MakeToken 用户登录，直接用uid 就可以登录并生成token，
// 每次登录会增加一次日志记录
// 密码验证交由api层验证
func (a AdminManager) MakeAdminToken(ctx context.Context, uid int, media struct {
	Ip      string
	City    string
	Browser string
}) (string, error) {
	db := store.WithContext(ctx)
	ainfo, err := db.Admin.Query().Where(admin.ID(uid)).Only(ctx)
	if err != nil {
		return "", err
	}
	// 创建一个日志
	_, err = db.AdminLoginLog.Create().
		SetIP(media.Ip).
		SetCity(media.City).
		SetBrowser(media.Browser).
		SetAdmin(ainfo).
		Save(ctx)
	if err != nil {
		return "", errorno.NewStoreErr(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tk, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errorno.NewInternalErr(err)
	}
	return tk, nil
}

func (a AdminManager) CheckToken(tk string) (int, error) {
	t, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, errorno.NewErr(errorno.TokenError)
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return int(claims["uid"].(float64)), nil
	}
	return 0, errorno.NewErr(errorno.TokenError)
}

//从boss后台同步学员信息
func (a AdminManager) SyncUserFromBoss(ctx context.Context, req request.SyncUserFromBoss) error {
	s := store.WithContext(ctx)
	finded, err := s.User.Query().Where(user.BossUserID(req.BossUserId)).Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !finded {
		_, err := s.User.Create().
			SetUsername(req.UserName).
			SetBossUserID(req.BossUserId).
			SetCardType(uint8(req.CardType)).
			SetPhone(req.Phone).
			SetPassword(req.Phone).
			SetIDCard(req.IdCard).
			Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}

		return nil
	} else {
		_, err := s.User.Update().
			Where(user.BossUserID(req.BossUserId)).
			SetUsername(req.UserName).
			SetIDCard(req.IdCard).
			SetPhone(req.Phone).
			SetCardType(uint8(req.CardType)).
			Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}

		return nil
	}
}

//从boss后台同步管理员信息
func (a AdminManager) SyncAdminFromBoss(ctx context.Context, req request.SyncAdminFromBoss) error {
	s := store.WithContext(ctx)
	finded, err := s.Admin.Query().Where(admin.BossAdminID(*req.BossAdminId)).Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !finded {
		_, err := s.Admin.Create().
			SetRealName(*req.RealName).
			SetBossAdminID(*req.BossAdminId).
			SetEmail(*req.Email).
			SetPhone(*req.Phone).
			SetPassword(*req.Phone).
			SetIsActive(uint8(*req.IsActive)).
			SetThirdOpenid(*req.ThirdOpenId).
			SetPlatform(uint8(*req.Platform)).
			Save(ctx)
		if err != nil {
			return errorno.NewStoreErr(err)
		}
		return nil
	} else {
		switch req.SyncType {
		case 1:
			_, err := s.Admin.Update().
				Where(admin.BossAdminID(*req.BossAdminId)).
				SetIsActive(uint8(*req.IsActive)).
				Save(ctx)
			if err != nil {
				return errorno.NewStoreErr(err)
			}
			break
		case 2:
			_, err := s.Admin.Update().
				Where(admin.BossAdminID(*req.BossAdminId)).
				SetThirdOpenid(*req.ThirdOpenId).
				Save(ctx)
			if err != nil {
				return errorno.NewStoreErr(err)
			}
			break
		default:
			_, err := s.Admin.Update().
				Where(admin.BossAdminID(*req.BossAdminId)).
				SetRealName(*req.RealName).
				SetBossAdminID(*req.BossAdminId).
				SetEmail(*req.Email).
				SetPhone(*req.Phone).
				SetIsActive(uint8(*req.IsActive)).
				SetThirdOpenid(*req.ThirdOpenId).
				SetPlatform(uint8(*req.Platform)).
				Save(ctx)
			if err != nil {
				return errorno.NewStoreErr(err)
			}
			break
		}
		return nil
	}
}

//编辑管理员
func (a AdminManager) UpdateAdmin(ctx context.Context, req request.SetAdminRequest) (int, error) {
	s := store.WithContext(ctx)
	if errorno := a.AdminIdExist(ctx, *req.Id); errorno != nil {
		return 0, errorno
	}
	md := s.Admin.UpdateOneID(*req.Id).
		ClearRoles()
	if *req.RoleId > 0 {
		md = md.AddRoleIDs(*req.RoleId)
	}
	if !app.IsNil(req.Remark) {
		md = md.SetRemark(*req.Remark)
	}
	//if !app.IsNil(req.Password) {
	//	md = md.SetPassword(*req.Password)
	//}
	res, err := md.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//设置管理员状态
func (a AdminManager) SetAdminStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := a.AdminIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Admin.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置角色状态
func (a AdminManager) SetRoleStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := a.RoleIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Role.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断地区是否存着
func (a AdminManager) AdminIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Admin.
		Query().
		Where(admin.ID(id)).
		Where(admin.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "管理员不存在",
		})
	}
	return nil
}

//修改密码
func (a AdminManager) ResetPassword(ctx context.Context, req request.ResetPassword) error {
	s := store.WithContext(ctx)
	detail, err := s.Admin.Query().Where(admin.ID(*req.Id)).First(ctx)
	if err != nil {
		return err
	}

	if !password.Comp(*req.OldPassword+detail.Salt, detail.Password) {
		return errorno.NewErr(errorno.PasswordError)
	}
	//保存新密码
	_, err = s.Admin.UpdateOneID(*req.Id).SetPassword(*req.ConfirmPassword).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//上传头像
func (a AdminManager) UpdateAdminAvatar(ctx context.Context, req request.UpdateAdminAvatar) error {
	s := store.WithContext(ctx)
	if errorno := a.AttachmentIdExist(ctx, *req.AdminAvatarId); errorno != nil {
		return errorno
	}
	if errorno := a.AdminIdExist(ctx, *req.Id); errorno != nil {
		return errorno
	}
	_, err := s.Admin.UpdateOneID(*req.Id).SetAdminAvatarID(*req.AdminAvatarId).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断附件是否存着
func (a AdminManager) AttachmentIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Attachment.
		Query().SoftDelete().
		Where(attachment.ID(id)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "附件不存在",
		})
	}
	return nil
}

//添加老师
func (a AdminManager) AddTeacher(ctx context.Context, req request.SetTeacher) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Teacher.
		Query().
		Where(teacher.Name(*req.Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	creation := s.Teacher.Create().
		SetName(*req.Name).
		SetPhone(*req.Phone).
		SetDetail(req.Detail).
		SetSortOrder(req.SortOrder).
		SetEmail(req.Email).
		SetAvatarID(req.AvatarId).
		SetTeachingAge(uint8(req.TeachingAge)).
		SetSex(uint8(*req.Sex)).
		SetNickname(req.Nickname).
		SetSubTitle(req.SubTitle).
		SetStatus(uint8(req.Status))

	if !app.IsNil(req.MajorIds) && len(req.MajorIds) > 0 {
		creation = creation.AddMajorIDs(req.MajorIds...)
	}

	newTeacher, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	if !app.IsNil(req.CourseIds) && len(req.CourseIds) > 0 {
		courseBulk := make([]*ent.KcCourseTeacherCreate, len(req.CourseIds))
		for i, v := range req.CourseIds {
			courseBulk[i] = s.KcCourseTeacher.Create().
				SetTeacherID(newTeacher.ID).SetCourseID(v)
		}
		_, err := s.KcCourseTeacher.CreateBulk(courseBulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
	}

	if !app.IsNil(req.TeacherTags) && len(req.TeacherTags) > 0 {
		tagBulk := make([]*ent.TeacherTagCreate, len(req.TeacherTags))
		for i, v := range req.TeacherTags {
			tagBulk[i] = s.TeacherTag.Create().
				SetName(v).SetTeacherID(newTeacher.ID)
		}
		_, err := s.TeacherTag.CreateBulk(tagBulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
	}
	return newTeacher.ID, nil
}

//编辑老师
func (a AdminManager) UpdateTeacher(ctx context.Context, req request.SetTeacher) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Teacher.
		Query().
		Where(teacher.Name(*req.Name)).
		Where(teacher.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	creation := s.Teacher.UpdateOneID(*req.Id).
		ClearMajors().
		ClearTeacherCourses().
		ClearTeacherTags().
		SetName(*req.Name).
		SetPhone(*req.Phone).
		SetDetail(req.Detail).
		SetSortOrder(req.SortOrder).
		SetEmail(req.Email).
		SetAvatarID(req.AvatarId).
		SetTeachingAge(uint8(req.TeachingAge)).
		SetSex(uint8(*req.Sex)).
		SetNickname(req.Nickname).
		SetSubTitle(req.SubTitle).
		SetStatus(uint8(req.Status))

	if !app.IsNil(req.MajorIds) && len(req.MajorIds) > 0 {
		creation = creation.AddMajorIDs(req.MajorIds...)
	}

	newTeacher, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}

	if !app.IsNil(req.CourseIds) && len(req.CourseIds) > 0 {
		courseBulk := make([]*ent.KcCourseTeacherCreate, len(req.CourseIds))
		for i, v := range req.CourseIds {
			courseBulk[i] = s.KcCourseTeacher.Create().
				SetTeacherID(newTeacher.ID).SetCourseID(v)
		}
		_, err := s.KcCourseTeacher.CreateBulk(courseBulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
	}

	if !app.IsNil(req.TeacherTags) && len(req.TeacherTags) > 0 {
		tagBulk := make([]*ent.TeacherTagCreate, len(req.TeacherTags))
		for i, v := range req.TeacherTags {
			tagBulk[i] = s.TeacherTag.Create().
				SetName(v).SetTeacherID(newTeacher.ID)
		}
		_, err := s.TeacherTag.CreateBulk(tagBulk...).Save(ctx)
		if err != nil {
			return 0, err
		}
	}
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newTeacher.ID, nil
}

//删除老师
func (a AdminManager) DelTeacher(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := a.TeacherIdExist(ctx, id); errorno != nil {
		return errorno
	}
	fined, err := s.KcCourseTeacher.Query().
		Where(kccourseteacher.TeacherID(id)).
		Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	if fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "老师已管理课程，不可删除",
		})
	}
	_, err = s.Teacher.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置老师状态
func (a AdminManager) SetTeacherStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := a.TeacherIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Teacher.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断老师是否存着
func (a AdminManager) TeacherIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Teacher.
		Query().
		Where(teacher.ID(id)).
		Where(teacher.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "老师不存在",
		})
	}
	return nil
}

//后台添加用户
func (a AdminManager) AddUser(ctx context.Context, req request.SetUser) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.User.
		Query().SoftDelete().
		Where(user.Phone(*req.Phone)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	if app.IsNil(req.Password) || *req.Password != *req.ConfirmPassword {
		return 0, errorno.NewErr(errorno.PasswordError)
	}
	creation := s.User.Create().
		SetPhone(*req.Phone).
		SetStatus(uint8(*req.Status)).
		SetPassword(*req.Password).
		SetRegFrom(5)

	if !app.IsNil(req.Username) {
		creation = creation.SetUsername(*req.Username)
	}
	//if !app.IsNil(req.BossUserId) {
	//	creation = creation.SetBossUserID(*req.BossUserId)
	//}
	if !app.IsNil(req.IdCard) {
		creation = creation.SetIDCard(*req.IdCard)
	}
	if !app.IsNil(req.CardType) {
		creation = creation.SetCardType(uint8(*req.CardType))
	}
	if !app.IsNil(req.Sex) {
		creation = creation.SetSex(uint8(*req.Sex))
	}
	//if !app.IsNil(req.Birthday) {
	//	creation = creation.SetBirthday(*req.Birthday)
	//}
	//if req.CityId > 0 {
	//	creation = creation.SetFromCityID(req.CityId)
	//}
	//if req.CateId > 0 {
	//	creation = creation.SetFromItemCategoryID(req.CateId)
	//}
	if !app.IsNil(req.SignRemark) {
		creation = creation.SetSignRemark(*req.SignRemark)
	}
	if req.AvatarId > 0 {
		avatarInfo, err := s.Attachment.Query().Where(attachment.ID(req.AvatarId)).First(ctx)
		if err == nil {
			creation = creation.SetAvatar(avatarInfo.Filename)
		}
	}
	userDetail, err := creation.Save(ctx)

	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return userDetail.ID, nil
}

//编辑用户
func (a AdminManager) UpdateUser(ctx context.Context, req request.SetUser) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.User.
		Query().SoftDelete().
		Where(user.Phone(*req.Phone)).
		Where(user.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	creation := s.User.UpdateOneID(*req.Id).
		SetPhone(*req.Phone).
		SetStatus(uint8(*req.Status))

	if !app.IsNil(req.Username) {
		creation = creation.SetUsername(*req.Username)
	}
	if !app.IsNil(req.IdCard) {
		creation = creation.SetIDCard(*req.IdCard)
	}
	if !app.IsNil(req.CardType) {
		creation = creation.SetCardType(uint8(*req.CardType))
	}
	if !app.IsNil(req.Sex) {
		creation = creation.SetSex(uint8(*req.Sex))
	}
	//if !app.IsNil(req.Birthday) {
	//	creation = creation.SetBirthday(*req.Birthday)
	//}
	//if req.CityId > 0 {
	//	creation = creation.SetFromCityID(req.CityId)
	//}
	//if req.CateId > 0 {
	//	creation = creation.SetFromItemCategoryID(req.CateId)
	//}
	if !app.IsNil(req.SignRemark) {
		creation = creation.SetSignRemark(*req.SignRemark)
	}
	if req.AvatarId > 0 {
		avatarInfo, err := s.Attachment.Query().Where(attachment.ID(req.AvatarId)).First(ctx)
		if err == nil {
			creation = creation.SetAvatar(avatarInfo.Filename)
		}
	}
	userDetail, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return userDetail.ID, nil
}
