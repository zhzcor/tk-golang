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
	"gserver/internal/store/ent/activitytype"
	"gserver/internal/store/ent/hotsearch"
	"gserver/internal/store/ent/informationclassify"
	"gserver/internal/store/ent/messagetype"
	app2 "gserver/pkg/requestparam"
)

//添加（编辑）活动类型
func SetActivityType(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateActivityType(ctx, *req.Name, *req.Id, *req.Status)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddActivityType(ctx, *req.Name, *req.Status)
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

//删除活动类型
func DelActivityType(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelActivityType(ctx, req.Id)
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

//设置活动类型状态
func SetActivityTypeStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetActivityTypeStatus(ctx, *req.Id, *req.Status)
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

//活动类型查询(带分页)
func GetActivityTypeByPage(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.ActivityTypePageListResponse
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).ActivityType.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(activitytype.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(activitytype.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
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

//活动类型列表（无分页）
func GetActivityTypeAll(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var activityTypeList ent.ActivityTypes
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).ActivityType.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(activitytype.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(activitytype.Status(uint8(*req.Status)))
		}
		list, err := query.Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		activityTypeList = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return activityTypeList, nil
}

//添加（编辑）资讯分类
func SetInformationClassify(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateInformationClassify(ctx, *req.Name, *req.Id, *req.Status)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddInformationClassify(ctx, *req.Name, *req.Status)
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

//删除资讯分类
func DelInformationClassify(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelInformationClassify(ctx, req.Id)
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

//设置资讯分类状态
func SetInformationClassifyStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetInformationClassifyStatus(ctx, *req.Id, *req.Status)
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

//资讯分类查询(带分页)
func GetInformationClassifyByPage(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.InformationClassifyListResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).InformationClassify.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(informationclassify.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(informationclassify.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
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

//资讯分类列表（无分页）
func GetInformationClassifyAll(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	informationClassifies := ent.InformationClassifies{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).InformationClassify.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(informationclassify.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(informationclassify.Status(uint8(*req.Status)))
		}
		list, err := query.Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		informationClassifies = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return informationClassifies, nil
}

//添加（编辑）消息类型
func SetMsgType(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateMsgType(ctx, *req.Name, *req.Id, *req.Status)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddMsgType(ctx, *req.Name, *req.Status)
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

//删除消息类型
func DelMsgType(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelMsgType(ctx, req.Id)
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

//设置消息类型状态
func SetMsgTypeStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetMsgTypeStatus(ctx, *req.Id, *req.Status)
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

//消息类型查询(带分页)
func GetMsgTypeByPage(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.MsgTypeListResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).MessageType.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(messagetype.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(messagetype.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
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

//消息类型列表（无分页）
func GetMsgTypeAll(ctx *gin.Context) (interface{}, error) {
	var req request.ActivityTypePageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	msgTypes := ent.MessageTypes{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).MessageType.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(messagetype.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(messagetype.Status(uint8(*req.Status)))
		}
		list, err := query.Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		msgTypes = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return msgTypes, nil
}

//添加（编辑）热门搜索
func SetHotSearch(ctx *gin.Context) (interface{}, error) {
	var req request.SetHotSearchRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := c.UpdateHotSearch(ctx, *req.Name, *req.Id, *req.Status, *req.SearchCount)
			if err != nil {
				return err
			}
		} else {
			_, err := c.AddHotSearch(ctx, *req.Name, *req.Status, *req.SearchCount)
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

//删除热门搜索
func DelHotSearch(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.DelHotSearch(ctx, req.Id)
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

//设置热门搜索状态
func SetHotSearchStatus(ctx *gin.Context) (interface{}, error) {
	var req request.SetActivityTypeStatusRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	c := app.Community{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = c.SetHotSearchStatus(ctx, *req.Id, *req.Status)
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

//热门搜索查询(带分页)
func GetHotSearchByPage(ctx *gin.Context) (interface{}, error) {
	var req request.HotSearchPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.HotSearchListResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).HotSearch.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(hotsearch.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(hotsearch.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
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

//热门搜索列表（无分页）
func GetHotSearchAll(ctx *gin.Context) (interface{}, error) {
	var req request.HotSearchPageListRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	hotSearches := ent.HotSearches{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).HotSearch.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(hotsearch.Name(req.Name))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(hotsearch.Status(uint8(*req.Status)))
		}
		list, err := query.Order(ent.Desc("id")).All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		hotSearches = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return hotSearches, nil
}
