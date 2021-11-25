package app

import (
	"context"
	"net/url"
	"tkserver/httpapi/admin/request"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/appagreement"
	"tkserver/internal/store/ent/appversion"
	"tkserver/internal/store/ent/attachment"
	"tkserver/internal/store/ent/city"
	"tkserver/internal/store/ent/groupcard"
	"tkserver/internal/store/ent/itemcategory"
	"tkserver/internal/store/ent/level"
	"tkserver/internal/store/ent/major"
)

type BasicConfig struct {
}

//添加地区
func (b BasicConfig) AddCity(ctx context.Context, cityName string, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.City.
		Query().SoftDelete().
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
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//编辑地区
func (b BasicConfig) UpdateCity(ctx context.Context, cityName string, id, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.City.
		Query().SoftDelete().
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

//添加地区
func (b BasicConfig) AddLevel(ctx context.Context, levelName string, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Level.
		Query().SoftDelete().
		Where(level.Name(levelName)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newLevel, err := s.Level.Create().
		SetName(levelName).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newLevel.ID, nil
}

//编辑地区
func (b BasicConfig) UpdateLevel(ctx context.Context, levelName string, id, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Level.
		Query().SoftDelete().
		Where(level.Name(levelName)).
		Where(level.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newCity, err := s.Level.UpdateOneID(id).
		SetName(levelName).
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//删除地区
func (b BasicConfig) DelLevel(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.LevelIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Level.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断地区是否存着
func (b BasicConfig) LevelIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Level.
		Query().
		Where(level.ID(id)).
		Where(level.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "层次不存在",
		})
	}
	return nil
}

//添加项目
func (b BasicConfig) AddItemCategory(ctx context.Context, Name string, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ItemCategory.
		Query().SoftDelete().
		Where(itemcategory.Name(Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	addItem := s.ItemCategory.Create().
		SetName(Name).
		SetSortOrder(sortOrder)
	newItem, err := addItem.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//编辑项目
func (b BasicConfig) UpdateItemCategory(ctx context.Context, Name string, id, sortOrder int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ItemCategory.
		Query().SoftDelete().
		Where(itemcategory.Name(Name)).
		Where(itemcategory.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	updateItem := s.ItemCategory.UpdateOneID(id).
		SetName(Name).
		SetSortOrder(sortOrder)

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
func (b BasicConfig) AddMajor(ctx context.Context, Name string, sortOrder int) (int, error) {
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
		SetSortOrder(sortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newItem.ID, nil
}

//编辑专业
func (b BasicConfig) UpdateMajor(ctx context.Context, Name string, id, sortOrder int) (int, error) {
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

//添加地区
func (b BasicConfig) AddGroupCard(ctx context.Context, req request.SetGroupCard) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.GroupCard.
		Query().SoftDelete().
		Where(groupcard.Title(req.Title)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newGroupCard, err := s.GroupCard.Create().
		SetTitle(req.Title).
		SetAttachmentID(req.AttachmentId).
		SetSubTitle(req.SubTitle).
		SetSortOrder(req.SortOrder).
		SetStatus(uint8(req.Status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newGroupCard.ID, nil
}

//编辑群名片
func (b BasicConfig) UpdateGroupCard(ctx context.Context, req request.SetGroupCard) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.GroupCard.
		Query().SoftDelete().
		Where(groupcard.Title(req.Title)).
		Where(groupcard.IDNEQ(req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newGroupCard, err := s.GroupCard.UpdateOneID(req.Id).
		SetTitle(req.Title).
		SetAttachmentID(req.AttachmentId).
		SetSubTitle(req.SubTitle).
		SetSortOrder(req.SortOrder).
		SetStatus(uint8(req.Status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newGroupCard.ID, nil
}

//删除群名片
func (b BasicConfig) DelGroupCard(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := b.GroupCardIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.GroupCard.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置群名片状态
func (b BasicConfig) SetGroupCardStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := b.GroupCardIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.GroupCard.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断群名片是否存着
func (b BasicConfig) GroupCardIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.GroupCard.
		Query().
		Where(groupcard.ID(id)).
		Where(groupcard.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "群名片不存在",
		})
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
