package admin

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"tkserver/httpapi/admin/request"
	"tkserver/httpapi/admin/response"
	"tkserver/internal/app"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/advertise"
	"tkserver/internal/store/ent/message"
	"tkserver/internal/store/ent/shareposter"
	app2 "tkserver/pkg/requestparam"
)

//推广管理

//添加（编辑）广告
func SetAdvertise(ctx *gin.Context) (interface{}, error) {
	var req request.SetAdvertise
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateAdvertise(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddAdvertise(ctx, req)
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

//删除广告
func DelAdvertise(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelAdvertise(ctx, req.Id)
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

//设置广告状态
func SetAdvertiseStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetAdvertiseStatus(ctx, *req.Id, *req.Status)
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

//广告查询(带分页)
func GetAdvertiseByPage(ctx *gin.Context) (interface{}, error) {
	var req request.AdPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.AdPageListResponse{
		List: []response.AdPageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Advertise.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(advertise.NameContains(*req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(advertise.Status(uint8(*req.Status)))
		}
		if !app2.IsNil(req.Position) {
			query = query.Where(advertise.Position(uint8(*req.Position)))
		}

		count, err := query.Count(ctx)
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
			detail := response.AdPageDetail{}
			detail.ID = v.ID
			detail.Name = v.Name
			detail.CreatedAt = v.CreatedAt
			detail.Remark = v.Remark
			detail.SortOrder = v.SortOrder
			detail.Status = v.Status
			detail.Position = v.Position
			detail.AdURL = v.AdURL
			detail.ClickCount = v.ClickCount
			detail.AttachmentID = v.AttachmentID
			detail.StartAt = v.StartAt
			detail.EndAt = v.EndAt
			detail.AttachmentName = ""
			if !app2.IsNil(v.Edges.Attachment) {
				detail.AttachmentName = app2.GetOssHost() + v.Edges.Attachment.Filename
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

//添加（编辑）分享海报
func SetSharePoster(ctx *gin.Context) (interface{}, error) {
	var req request.SetSharePostRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateSharePoster(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddSharePoster(ctx, req)
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

//删除分享海报
func DelSharePoster(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelSharePoster(ctx, req.Id)
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

//设置分享海报状态
func SetSharePosterStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetSharePosterStatus(ctx, *req.Id, *req.Status)
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

//分享海报查询(带分页)
func GetSharePosterByPage(ctx *gin.Context) (interface{}, error) {
	var req request.SharePosterPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.SharePosterPageListResponse{
		List: []response.SharePosterDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).SharePoster.Query().SoftDelete()
		if !app2.IsNil(req.Status) {
			query = query.Where(shareposter.Status(uint8(*req.Status)))
		}
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.ForPage(req.Page, req.PageSize).
			Order(ent.Desc("id")).WithAttachment().All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.SharePosterDetail{}
			detail.ID = v.ID
			detail.Name = v.Name
			detail.SortOrder = v.SortOrder
			detail.Remark = v.Remark
			detail.URL = v.URL
			detail.CreatedAt = v.CreatedAt
			detail.Status = v.Status
			detail.SharePosterImgID = v.SharePosterImgID
			detail.ShareCount = v.ShareCount
			detail.ImgFileName = ""
			if !app2.IsNil(v.Edges.Attachment) {
				detail.ImgFileName = app2.GetOssHost() + v.Edges.Attachment.Filename
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

func SetMessage(ctx *gin.Context) (interface{}, error) {
	var req request.SetMessageRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	operatorAdminId := ctx.GetInt(app.GlobalAdminId)
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateMessage(ctx, req, operatorAdminId)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddMessage(ctx, req, operatorAdminId)
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

//删除消息
func DelMessage(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelMessage(ctx, req.Id)
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

//设置消息状态
func SetMessageStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.PopManage{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetMessageStatus(ctx, *req.Id, *req.Status)
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

//消息查询(带分页)
func GetMessageByPage(ctx *gin.Context) (interface{}, error) {
	var req request.MessagePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.MessagePageList{
		List: []response.MessageDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Message.Query().SoftDelete()
		if !app2.IsNil(req.Status) {
			query = query.Where(message.Status(uint8(*req.Status)))
		}
		if !app2.IsNil(req.MessageTypeId) {
			query = query.Where(message.MessageTypeID(*req.MessageTypeId))
		}
		if !app2.IsNil(req.Name) {
			query = query.Where(message.NameContains(*req.Name))
		}
		if !app2.IsNil(req.ClassId) {
			query = query.Where(message.ClassID(*req.ClassId))
		}
		if !app2.IsNil(req.CourseId) {
			query = query.Where(message.CourseID(*req.CourseId))
		}
		if app2.IsNil(req.ClassId) && app2.IsNil(req.CourseId) {
			query = query.Where(message.CourseIDIsNil()).
				Where(message.ClassIDIsNil())
		}
		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		list, err := query.WithAttachment(func(query *ent.AttachmentQuery) {
			query.Select("id", "filename")
		}).WithMessageType(func(query *ent.MessageTypeQuery) {
			query.Select("id", "name")
		}).WithAdmin(func(query *ent.AdminQuery) {
			query.Select("id", "real_name")
		}).ForPage(req.Page, req.PageSize).Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.MessageDetail{}
			detail.ID = v.ID
			detail.Name = v.Name
			detail.Status = v.Status
			detail.Detail = v.Detail
			detail.AutoPublishAt = v.AutoPublishAt
			detail.PublishType = v.PublishType
			detail.CreatedAt = v.CreatedAt
			detail.AttachmentID = v.AttachmentID
			detail.CreatedAdminID = v.CreatedAdminID
			detail.MessageTypeID = v.MessageTypeID
			detail.AttachmentName = ""
			detail.CreatedAdminName = ""
			detail.MessageTypeName = ""
			if !app2.IsNil(v.Edges.Attachment) {
				detail.AttachmentName = app2.GetOssHost() + v.Edges.Attachment.Filename
			}
			if !app2.IsNil(v.Edges.Admin) {
				detail.CreatedAdminName = v.Edges.Admin.RealName
			}
			if !app2.IsNil(v.Edges.MessageType) {
				detail.MessageTypeName = v.Edges.MessageType.Name
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
