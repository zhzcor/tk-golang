package app

import (
	"context"
	"gserver/httpapi/admin/request"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/appagreement"
	"gserver/internal/store/ent/appversion"
	"gserver/internal/store/ent/attachment"
	"gserver/internal/store/ent/city"
	"gserver/internal/store/ent/itemcategory"
	"gserver/internal/store/ent/major"
	"net/url"
)

type BasicConfig struct {
}

//添加地区
func (b BasicConfig) AddCity(ctx context.Context, cityName, code, desc string, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.City.
		Query().
		Where(city.Name(cityName)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newCity, err := s.City.Create().
		SetName(cityName).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//编辑地区
func (b BasicConfig) UpdateCity(ctx context.Context, cityName, code, desc string, id, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.City.
		Query().
		Where(city.Name(cityName)).
		Where(city.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newCity, err := s.City.UpdateOneID(id).
		SetName(cityName).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//删除地区
func (b BasicConfig) DelCity(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.CityIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.City.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置地区状态
func (b BasicConfig) SetCityStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := b.CityIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.City.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断地区是否存着
func (b BasicConfig) CityIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.City.
		Query().
		Where(city.ID(id)).
		Where(city.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "地区不存在",
		})
	}
	return nil
}

//添加项目
func (b BasicConfig) AddItemCategory(ctx context.Context, Name, code, desc string, sortOrder, pid int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ItemCategory.
		Query().
		Where(itemcategory.Name(Name)).
		Where(itemcategory.Pid(pid)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	addItem := s.ItemCategory.Create().
		SetName(Name).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder)
	if pid != 0 {
		addItem = addItem.SetPid(pid)
	}
	newItem, err := addItem.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//编辑项目
func (b BasicConfig) UpdateItemCategory(ctx context.Context, Name, code, desc string, id, sortOrder, pid int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ItemCategory.
		Query().
		Where(itemcategory.Name(Name)).
		Where(itemcategory.IDNEQ(id)).
		Where(itemcategory.Pid(pid)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	updateItem := s.ItemCategory.UpdateOneID(id).
		SetName(Name).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder)
	if pid != 0 {
		updateItem = updateItem.SetPid(pid)
	}
	newItem, err := updateItem.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//删除项目
func (b BasicConfig) DelItemCategory(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.ItemIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.ItemCategory.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置项目状态
func (b BasicConfig) SetItemStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := b.ItemIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.ItemCategory.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断项目是否存在
func (b BasicConfig) ItemIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.ItemCategory.
		Query().
		Where(itemcategory.ID(id)).
		Where(itemcategory.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "项目不存在",
		})
	}
	return nil
}

//添加专业
func (b BasicConfig) AddMajor(ctx context.Context, Name, code, desc string, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Major.
		Query().
		Where(major.Name(Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newItem, err := s.Major.Create().
		SetName(Name).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//编辑专业
func (b BasicConfig) UpdateMajor(ctx context.Context, Name, code, desc string, id, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Major.
		Query().
		Where(major.Name(Name)).
		Where(major.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newItem, err := s.Major.UpdateOneID(id).
		SetName(Name).
		SetCode(code).
		SetDesc(desc).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//删除专业
func (b BasicConfig) DelMajor(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.MajorIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Major.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置专业状态
func (b BasicConfig) SetMajorStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := b.MajorIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Major.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断专业是否存在
func (b BasicConfig) MajorIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Major.
		Query().
		Where(major.ID(id)).
		Where(major.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "专业不存在",
		})
	}
	return nil
}

//查询数据不存在
func CheckDataNotFound(err error) error {
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return errorno.NewErr(errorno.DataNotExistsError)
	}
	return nil
}

//添加协议
func (b BasicConfig) SetAgreement(ctx context.Context, req request.SetAgreement) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.AppAgreement.
		Query().
		Where(appagreement.Type(uint8(*req.Type))).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	name := ""
	switch *req.Type {
	case 1:
		name = "服务协议"
		break
	case 2:
		name = "隐私政策"
		break
	case 3:
		name = "ios充值协议"
		break
	case 4:
		name = "账户注销协议"
		break
	case 5:
		name = "App温馨提示"
		break
	case 6:
		name = "优惠卷使用规则"
		break
	case 7:
		name = "关于我们"
		break
	case 8:
		name = "加入我们"
		break
	case 9:
		name = "版权声明"
		break
	case 10:
		name = "联系我们"
		break
	case 11:
		name = "常见问题"
		break

	}
	if !fined {
		_, err := s.AppAgreement.Create().
			SetName(name).
			SetType(uint8(*req.Type)).
			SetDetail(*req.Detail).
			Save(ctx)
		if err != nil {
			return 0, err
		} else {
			return 0, nil
		}
	}
	_, err = s.AppAgreement.Update().
		Where(appagreement.Type(uint8(*req.Type))).
		SetDetail(*req.Detail).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return 0, nil
}

//添加app版本
func (b BasicConfig) AddAppVersion(ctx context.Context, req request.SetAppVersion) (int, error) {
	s := store.WithContext(ctx)
	var fined, err = s.AppVersion.
		Query().
		Where(appversion.Name(*req.Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	result, err := s.AppVersion.Create().
		SetName(*req.Name).
		SetSn(*req.Sn).
		SetURL(*req.Url).
		SetRemark(*req.Remark).
		SetIsForceUpdate(uint8(*req.IsForceUpdate)).
		SetPhoneType(uint8(*req.PhoneType)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//编辑app版本
func (b BasicConfig) UpdateAppVersion(ctx context.Context, req request.SetAppVersion) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.AppVersion.
		Query().
		Where(appversion.Name(*req.Name)).
		Where(appversion.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	result, err := s.AppVersion.UpdateOneID(*req.Id).
		SetName(*req.Name).
		SetSn(*req.Sn).
		SetURL(*req.Url).
		SetRemark(*req.Remark).
		SetIsForceUpdate(uint8(*req.IsForceUpdate)).
		SetPhoneType(uint8(*req.PhoneType)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//删除app版本
func (b BasicConfig) DelAppVersion(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.AppVersionIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.AppVersion.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断app版本是否存在
func (b BasicConfig) AppVersionIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.AppVersion.
		Query().
		Where(appversion.ID(id)).
		Where(appversion.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "app版本不存在",
		})
	}
	return nil
}

//oss回调处理
func (b BasicConfig) ProcessOssCallback(ctx context.Context, req request.OssCallBack) (interface{}, error) {
	s := store.WithContext(ctx)
	filename := url.QueryEscape(req.Filename)
	fined, err := s.Attachment.
		Query().SoftDelete().
		Where(attachment.Filename(filename)).
		Exist(ctx)
	if err != nil {
		return nil, errorno.NewStoreErr(err)
	}
	if fined {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "文件已存在",
		})
	}

	creation := s.Attachment.Create().
		SetFilename(filename)
	if req.AdminId > 0 {
		creation = creation.SetAdminID(req.AdminId)
	}
	if req.UserId > 0 {
		creation = creation.SetUserID(req.UserId)
	}
	if req.Size > 0 {
		creation = creation.SetFileSize(uint32(req.Size))
	}
	if req.MimeType != "" {
		creation = creation.SetMimeType(req.MimeType)
	}
	attachment, err := creation.Save(ctx)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}
