package app

import (
	"context"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	"tkserver/internal/store/ent/activitytype"
	"tkserver/internal/store/ent/hotsearch"
	"tkserver/internal/store/ent/informationclassify"
	"tkserver/internal/store/ent/messagetype"
)

type Community struct {
}

//添加活动类型
func (c Community) AddActivityType(ctx context.Context, name string, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ActivityType.
		Query().
		Where(activitytype.Name(name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newCity, err := s.ActivityType.Create().
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//编辑活动类型
func (c Community) UpdateActivityType(ctx context.Context, name string, id, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.ActivityType.
		Query().
		Where(activitytype.Name(name)).
		Where(activitytype.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newCity, err := s.ActivityType.UpdateOneID(id).
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newCity.ID, nil
}

//删除活动类型
func (c Community) DelActivityType(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.ActivityTypeIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.ActivityType.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置活动类型状态
func (c Community) SetActivityTypeStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := c.ActivityTypeIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.ActivityType.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断活动类型是否存着
func (c Community) ActivityTypeIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.ActivityType.
		Query().
		Where(activitytype.ID(id)).
		Where(activitytype.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "活动类型id不存在",
		})
	}
	return nil
}

//添加咨询分类
func (c Community) AddInformationClassify(ctx context.Context, name string, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.InformationClassify.
		Query().
		Where(informationclassify.Name(name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.InformationClassify.Create().
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//编辑咨询分类
func (c Community) UpdateInformationClassify(ctx context.Context, name string, id, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.InformationClassify.
		Query().
		Where(informationclassify.Name(name)).
		Where(informationclassify.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.InformationClassify.UpdateOneID(id).
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除咨询分类
func (c Community) DelInformationClassify(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.InformationClassifyIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.InformationClassify.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置咨询分类状态
func (c Community) SetInformationClassifyStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := c.InformationClassifyIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.InformationClassify.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断咨询分类是否存着
func (c Community) InformationClassifyIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.InformationClassify.
		Query().
		Where(informationclassify.ID(id)).
		Where(informationclassify.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "资讯分类id不存在",
		})
	}
	return nil
}

//添加消息类型
func (c Community) AddMsgType(ctx context.Context, name string, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.MessageType.
		Query().
		Where(messagetype.Name(name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.MessageType.Create().
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//编辑消息类型
func (c Community) UpdateMsgType(ctx context.Context, name string, id, status int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.MessageType.
		Query().
		Where(messagetype.Name(name)).
		Where(messagetype.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.MessageType.UpdateOneID(id).
		SetName(name).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除消息类型
func (c Community) DelMsgType(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.MsgTypeIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.MessageType.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置消息类型状态
func (c Community) SetMsgTypeStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := c.MsgTypeIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.MessageType.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断消息类型是否存着
func (c Community) MsgTypeIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.MessageType.
		Query().
		Where(messagetype.ID(id)).
		Where(messagetype.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "消息类型id不存在",
		})
	}
	return nil
}

//添加热门搜索
func (c Community) AddHotSearch(ctx context.Context, name string, status, searchCount int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.HotSearch.
		Query().
		Where(hotsearch.Name(name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.HotSearch.Create().
		SetName(name).
		SetSearchCount(searchCount).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//编辑热门搜索
func (c Community) UpdateHotSearch(ctx context.Context, name string, id, status, searchCount int) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.HotSearch.
		Query().
		Where(hotsearch.Name(name)).
		Where(hotsearch.IDNEQ(id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	res, err := s.HotSearch.UpdateOneID(id).
		SetName(name).
		SetSearchCount(searchCount).
		SetStatus(uint8(status)).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return res.ID, nil
}

//删除热门搜索
func (c Community) DelHotSearch(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := c.HotSearchIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.HotSearch.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置热门搜索状态
func (c Community) SetHotSearchStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := c.HotSearchIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.HotSearch.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断热门搜索是否存着
func (c Community) HotSearchIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.HotSearch.
		Query().
		Where(hotsearch.ID(id)).
		Where(hotsearch.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "热门搜索id不存在",
		})
	}
	return nil
}
