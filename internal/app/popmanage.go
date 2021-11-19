package app

import (
	"context"
	"gserver/httpapi/admin/request"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent/advertise"
	"gserver/internal/store/ent/message"
	"gserver/internal/store/ent/shareposter"
	app "gserver/pkg/requestparam"
)

type PopManage struct {
}

//添加广告
func (p PopManage) AddAdvertise(ctx context.Context, req request.SetAdvertise) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Advertise.
		Query().
		Where(advertise.Name(*req.Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newAd, err := s.Advertise.Create().
		SetName(*req.Name).
		SetStatus(uint8(*req.Status)).
		SetAttachmentID(*req.AttachmentId).
		SetStartAt(*req.StartAt).
		SetEndAt(*req.EndAt).
		SetAdURL(*req.AdUrl).
		SetRemark(req.Remark).
		SetPosition(uint8(*req.Position)).
		SetStatus(uint8(*req.Status)).
		SetSortOrder(req.SortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newAd.ID, nil
}

//编辑广告
func (p PopManage) UpdateAdvertise(ctx context.Context, req request.SetAdvertise) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.Advertise.
		Query().
		Where(advertise.Name(*req.Name)).
		Where(advertise.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	newAd, err := s.Advertise.UpdateOneID(*req.Id).
		SetName(*req.Name).
		SetStatus(uint8(*req.Status)).
		SetAttachmentID(*req.AttachmentId).
		SetStartAt(*req.StartAt).
		SetEndAt(*req.EndAt).
		SetAdURL(*req.AdUrl).
		SetRemark(req.Remark).
		SetPosition(uint8(*req.Position)).
		SetStatus(uint8(*req.Status)).
		SetSortOrder(req.SortOrder).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return newAd.ID, nil
}

//删除广告
func (p PopManage) DelAdvertise(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := p.AdvertiseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Advertise.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置广告状态
func (p PopManage) SetAdvertiseStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := p.AdvertiseIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Advertise.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断广告是否存着
func (p PopManage) AdvertiseIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Advertise.
		Query().
		Where(advertise.ID(id)).
		Where(advertise.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "广告id不存在",
		})
	}
	return nil
}

//添加活动海报
func (p PopManage) AddSharePoster(ctx context.Context, req request.SetSharePostRequest) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.SharePoster.
		Query().
		Where(shareposter.Name(*req.Name)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	remark := ""
	url := ""
	if !app.IsNil(req.Remark) {
		remark = *req.Remark
	}
	if !app.IsNil(req.Url) {
		url = *req.Url
	}
	result, err := s.SharePoster.Create().
		SetName(*req.Name).
		SetStatus(uint8(*req.Status)).
		SetAttachmentID(*req.AttachmentId).
		SetURL(url).
		SetSortOrder(*req.SortOrder).
		SetRemark(remark).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//编辑活动海报
func (p PopManage) UpdateSharePoster(ctx context.Context, req request.SetSharePostRequest) (int, error) {
	s := store.WithContext(ctx)
	fined, err := s.SharePoster.
		Query().
		Where(shareposter.Name(*req.Name)).
		Where(shareposter.IDNEQ(*req.Id)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	if fined {
		return 0, errorno.NewErr(errorno.DataExistsError)
	}
	remark := ""
	url := ""
	if !app.IsNil(req.Remark) {
		remark = *req.Remark
	}
	if !app.IsNil(req.Url) {
		url = *req.Url
	}
	result, err := s.SharePoster.UpdateOneID(*req.Id).
		SetName(*req.Name).
		SetStatus(uint8(*req.Status)).
		SetAttachmentID(*req.AttachmentId).
		SetURL(url).
		SetSortOrder(*req.SortOrder).
		SetRemark(remark).
		Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//删除活动海报
func (p PopManage) DelSharePoster(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := p.SharePosterIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.SharePoster.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置活动海报状态
func (p PopManage) SetSharePosterStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := p.SharePosterIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.SharePoster.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断活动海报是否存着
func (p PopManage) SharePosterIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Advertise.
		Query().
		Where(advertise.ID(id)).
		Where(advertise.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "广告id不存在",
		})
	}
	return nil
}

//添加消息
func (p PopManage) AddMessage(ctx context.Context, req request.SetMessageRequest, operatorAdminId int) (int, error) {
	s := store.WithContext(ctx)
	md := s.Message.Create().
		SetName(*req.Name).
		SetPublishType(uint8(*req.PublishType)).
		SetDetail(*req.Detail).
		SetMessageTypeID(*req.MessageTypeId)
	if !app.IsNil(req.CourseId) {
		md = md.SetCourseID(*req.CourseId)
	}
	if !app.IsNil(req.ClassId) {
		md = md.SetClassID(*req.ClassId)
	}
	if !app.IsNil(req.AttachmentId) {
		md = md.SetAttachmentID(*req.AttachmentId)
	}
	if !app.IsNil(req.AutoPublishAt) {
		md = md.SetAutoPublishAt(*req.AutoPublishAt)
	}
	result, err := md.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//编辑消息
func (p PopManage) UpdateMessage(ctx context.Context, req request.SetMessageRequest, operatorAdminId int) (int, error) {
	s := store.WithContext(ctx)
	if errorno := p.MessageIdExist(ctx, *req.Id); errorno != nil {
		return 0, errorno
	}
	md := s.Message.UpdateOneID(*req.Id).
		SetName(*req.Name).
		SetPublishType(uint8(*req.PublishType)).
		SetDetail(*req.Detail).
		SetMessageTypeID(*req.MessageTypeId)
	if !app.IsNil(req.CourseId) {
		md = md.SetCourseID(*req.CourseId)
	}
	if !app.IsNil(req.ClassId) {
		md = md.SetClassID(*req.ClassId)
	}
	if !app.IsNil(req.AttachmentId) {
		md = md.SetAttachmentID(*req.AttachmentId)
	}
	if !app.IsNil(req.AutoPublishAt) {
		md = md.SetAutoPublishAt(*req.AutoPublishAt)
	}
	result, err := md.Save(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err)
	}
	return result.ID, nil
}

//删除消息
func (p PopManage) DelMessage(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	if errorno := p.MessageIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Message.UpdateOneID(id).SoftDelete().Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//设置消息状态
func (p PopManage) SetMessageStatus(ctx context.Context, id, status int) error {
	s := store.WithContext(ctx)
	if errorno := p.MessageIdExist(ctx, id); errorno != nil {
		return errorno
	}
	_, err := s.Message.UpdateOneID(id).SetStatus(uint8(status)).Save(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	return nil
}

//id判断消息是否存着
func (p PopManage) MessageIdExist(ctx context.Context, id int) error {
	s := store.WithContext(ctx)
	fined, err := s.Message.
		Query().
		Where(message.ID(id)).
		Where(message.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return errorno.NewStoreErr(err)
	}
	if !fined {
		return errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "消息id不存在",
		})
	}
	return nil
}
