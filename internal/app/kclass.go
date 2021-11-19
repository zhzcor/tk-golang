package app

import (
	"context"
	"tkserver/httpapi/admin/request"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent/kcclass"
	"tkserver/internal/store/ent/kcclassteacher"
	"tkserver/internal/store/ent/kcuserclass"
	app "tkserver/pkg/requestparam"
)

type KcClass struct {
}

//添加地区
func (b KcClass) AddClass(ctx context.Context, req request.SetClass) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcClass.
		Query().
		Where(kcclass.ClassTitle(*req.ClassTitle)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	creation := s.KcClass.Create().
		SetClassTitle(*req.ClassTitle)
	if !app.IsNil(req.CateId) && *req.CateId > 0 {
		creation = creation.SetCateID(*req.CateId)
	}
	if !app.IsNil(req.CityId) && *req.CityId > 0 {
		creation = creation.SetCityID(*req.CityId)
	}
	if !app.IsNil(req.MajorIds) {
		creation = creation.AddMajorIDs(req.MajorIds...)
	}
	if !app.IsNil(req.ClassCode) {
		creation = creation.SetClassCode(*req.ClassCode)
	}
	if !app.IsNil(req.ClassPeriodsType) {
		creation = creation.SetClassPeriodType(uint8(*req.ClassPeriodsType))
	}
	if !app.IsNil(req.ClassCoverImgId) {
		creation = creation.SetClassCoverImgID(*req.ClassCoverImgId)
	}
	if !app.IsNil(req.ClassDesc) {
		creation = creation.SetClassDesc(*req.ClassDesc)
	}
	if !app.IsNil(req.IsDisplay) {
		creation = creation.SetIsDisplay(uint8(*req.IsDisplay))
	}
	if !app.IsNil(req.Price) {
		creation = creation.SetPrice(*req.Price)
	}
	res, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//编辑地区
func (b KcClass) UpdateClass(ctx context.Context, req request.SetClass) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcClass.
		Query().
		Where(kcclass.ClassTitle(*req.ClassTitle)).
		Where(kcclass.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	creation := s.KcClass.UpdateOneID(*req.Id).
		ClearMajors().
		SetClassTitle(*req.ClassTitle)
	if !app.IsNil(req.CateId) && *req.CateId > 0 {
		creation = creation.SetCateID(*req.CateId)
	}
	if !app.IsNil(req.CityId) && *req.CityId > 0 {
		creation = creation.SetCityID(*req.CityId)
	}
	if !app.IsNil(req.MajorIds) {
		creation = creation.AddMajorIDs(req.MajorIds...)
	}
	if !app.IsNil(req.ClassCode) {
		creation = creation.SetClassCode(*req.ClassCode)
	}
	if !app.IsNil(req.ClassPeriodsType) {
		creation = creation.SetClassPeriodType(uint8(*req.ClassPeriodsType))
	}
	if !app.IsNil(req.ClassCoverImgId) {
		creation = creation.SetClassCoverImgID(*req.ClassCoverImgId)
	}
	if !app.IsNil(req.ClassDesc) {
		creation = creation.SetClassDesc(*req.ClassDesc)
	}
	if !app.IsNil(req.IsDisplay) {
		creation = creation.SetIsDisplay(uint8(*req.IsDisplay))
	}
	if !app.IsNil(req.Price) {
		creation = creation.SetPrice(*req.Price)
	}
	res, err := creation.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除地区
func (b KcClass) DelClass(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.KcClass.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置地区状态
func (b KcClass) SetClassStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.KcClass.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断班级是否存着
func (b KcClass) ClassIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcClass.
		Query().
		Where(kcclass.ID(id)).
		Where(kcclass.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "班级不存在",
		})
	}
	return nil
}

//添加班级课程
func (b KcClass) AddClassCourse(ctx context.Context, req request.SetClassCourse) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, *req.ClassId); errorno != nil {
		return errorno
	}

	if len(req.CourseIds) <= 0 {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "未选择课程",
		})
	}
	courseCount := len(req.CourseIds)

	_, err := s.KcClass.UpdateOneID(*req.ClassId).
		AddKcClassCourseIDs(req.CourseIds...).
		AddCourseCount(courseCount).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//移除班级课程
func (b KcClass) RemoveClassCourse(ctx context.Context, req request.RemoveClassCourse) error {
	s := store.WithContext(ctx)
	_, err := s.KcClass.UpdateOneID(*req.ClassId).
		RemoveKcClassCourseIDs(*req.CourseId).AddCourseCount(-1).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//添加班级老师
func (b KcClass) AddClassTeacher(ctx context.Context, req request.SetClassTeacher) error {
	s := store.WithContext(ctx)
	c := Course{}
	if errorno := b.ClassIdExist(ctx, *req.ClassId); errorno != nil {
		return errorno
	}
	if errorno := c.TeacherIdExist(ctx, *req.TeacherId); errorno != nil {
		return errorno
	}
	creation := s.KcClassTeacher.Create().
		SetClassID(*req.ClassId).SetTeacherID(*req.TeacherId)
	if !app.IsNil(req.ShowStatus) {
		creation.SetShowStatus(uint8(*req.ShowStatus))
	}
	_, err := creation.Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除班级老师
func (b KcClass) RemoveClassTeacher(ctx context.Context, req request.SetClassTeacher) error {
	s := store.WithContext(ctx)
	fined, err := s.KcClassTeacher.Query().
		Where(kcclassteacher.TeacherID(*req.TeacherId)).
		Where(kcclassteacher.ClassID(*req.ClassId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "班级老师不存在",
		})
	}
	_, err = s.KcClassTeacher.Delete().
		Where(kcclassteacher.ClassID(*req.ClassId)).
		Where(kcclassteacher.TeacherID(*req.TeacherId)).
		Exec(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置课程教师状态
func (b KcClass) SetClassTeacherStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	_, err := s.KcClassTeacher.UpdateOneID(id).SetShowStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//添加班级学员
func (b KcClass) AddClassUser(ctx context.Context, req request.SetClassUser) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, *req.ClassId); errorno != nil {
		return errorno
	}
	c := Course{}
	if errorno := c.UserIdExist(ctx, *req.UserId); errorno != nil {
		return errorno
	}
	if errorno := b.UserClassIdExist(ctx, *req.ClassId, *req.UserId); errorno != nil {
		return errorno
	}
	_, err := s.KcUserClass.Create().
		SetClassID(*req.ClassId).
		SetUserID(*req.UserId).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = s.KcClass.UpdateOneID(*req.ClassId).
		AddStudentCount(1).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除班级学员
func (b KcClass) RemoveClassUser(ctx context.Context, req request.SetClassUser) error {
	s := store.WithContext(ctx)
	fined, err := s.KcUserClass.Query().
		Where(kcuserclass.ClassID(*req.ClassId)).
		Where(kcuserclass.UserID(*req.UserId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "班级学员不存在",
		})
	}
	_, err = s.KcUserClass.Delete().
		Where(kcuserclass.ClassID(*req.ClassId)).
		Where(kcuserclass.UserID(*req.UserId)).
		Exec(ctx)
	if err != nil {
		return err
	}
	_, err = s.KcClass.UpdateOneID(*req.ClassId).
		AddStudentCount(-1).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置学员课程有效期
func (b KcClass) SetClassUserValidity(ctx context.Context, req request.SetClassUserValidity) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.KcUserClass.
		Query().
		Where(kcuserclass.UserID(*req.UserId)).
		Where(kcuserclass.ClassID(*req.ClassId)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if !fined {
		return 0, errorno.NewErr(errorno.DataNotExistsError)
	}
	_, err = s.KcUserClass.Update().
		Where(kcuserclass.UserID(*req.UserId)).
		Where(kcuserclass.ClassID(*req.ClassId)).
		SetPeriodType(2).
		SetClosingDate(*req.ClosingDate).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return 0, nil
}

//判断班级用户是否存着
func (b KcClass) UserClassIdExist(ctx context.Context, classId, userId int) error {
	s := store.WithContext(ctx)
	fined, err := s.KcUserClass.
		Query().
		Where(kcuserclass.ClassID(classId)).
		Where(kcuserclass.UserID(userId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "用户已存在",
		})
	}
	return nil
}

//添加班级班主任
func (b KcClass) AddClassMasterTeacher(ctx context.Context, req request.SetClassTeacher) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, *req.ClassId); errorno != nil {
		return errorno
	}

	fined, err := s.KcClassTeacher.
		Query().
		Where(kcclassteacher.TeacherID(*req.TeacherId)).
		Where(kcclassteacher.ClassID(*req.ClassId)).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.DataNotExistsError)
	}
	_, err = s.KcClass.UpdateOneID(*req.ClassId).
		SetClassHeadMasterID(*req.TeacherId).
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//删除班级班主任
func (b KcClass) RemoveClassMasterTeacher(ctx context.Context, req request.ClassIdRequest) error {
	s := store.WithContext(ctx)
	if errorno := b.ClassIdExist(ctx, *req.ClassId); errorno != nil {
		return errorno
	}
	_, err := s.KcClass.UpdateOneID(*req.ClassId).
		ClearClassHeadMasterID().
		Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}
