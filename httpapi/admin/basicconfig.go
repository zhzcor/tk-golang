package admin

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/admin/request"
	"gserver/httpapi/admin/response"
	"gserver/internal/app"
	"gserver/internal/cache"
	"gserver/internal/config"
	"gserver/internal/errorno"
	"gserver/internal/store"
	"gserver/internal/store/ent"
	"gserver/internal/store/ent/admin"
	"gserver/internal/store/ent/adminoperationlog"
	"gserver/internal/store/ent/appagreement"
	"gserver/internal/store/ent/appversion"
	"gserver/internal/store/ent/city"
	"gserver/internal/store/ent/itemcategory"
	"gserver/internal/store/ent/kccourse"
	"gserver/internal/store/ent/kccoursesmallcategory"
	"gserver/internal/store/ent/major"
	"gserver/internal/store/ent/tkquestion"
	"gserver/internal/store/ent/user"
	"gserver/internal/store/ent/userloginlog"
	"gserver/pkg/log"
	"gserver/pkg/oss"
	app2 "gserver/pkg/requestparam"
	"time"
)

//添加（编辑）地区
func SetCity(ctx *gin.Context) (interface{}, error) {
	var req request.SetBasicTag
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	cm := app.Common{}
	adminId := ctx.GetInt(app.GlobalAdminId)
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if req.Id > 0 { //编辑
			_, err := bc.UpdateCity(ctx, req.Name, req.Code, req.Desc, req.Id, req.SortOrder)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑地区", req.Name)
		} else {
			_, err := bc.AddCity(ctx, req.Name, req.Code, req.Desc, req.SortOrder)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加地区", req.Name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除地区
func DelCity(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.DelCity(ctx, req.Id)
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

//设置地区状态
func SetCityStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.SetCityStatus(ctx, req.Id, req.Status)
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

//地区查询(带分页)
func GetCityByPage(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.CityListSuccess
	err = store.WithTx(ctx, func(ctx context.Context) error {
		//offset, pageSize := helper.GetPageOffset(req.Page, req.PageSize)
		cityQuery := store.WithContext(ctx).City.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			cityQuery = cityQuery.Where(city.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			cityQuery = cityQuery.Where(city.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			cityQuery = cityQuery.Where(city.Status(uint8(*req.Status)))
		}

		count, err := cityQuery.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count

		cityQuery = cityQuery.ForPage(req.Page, req.PageSize).Order(ent.Desc("id"))
		list, err := cityQuery.All(ctx)
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

//地区查询（无分页）
func GetCityAll(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var cityList ent.Cities
	err = store.WithTx(ctx, func(ctx context.Context) error {
		cityQuery := store.WithContext(ctx).City.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			cityQuery = cityQuery.Where(city.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			cityQuery = cityQuery.Where(city.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			cityQuery = cityQuery.Where(city.Status(uint8(*req.Status)))
		}
		list, err := cityQuery.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		cityList = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cityList, nil
}

//添加（编辑）项目
func SetItemCategory(ctx *gin.Context) (interface{}, error) {
	var req request.SetItemCategory
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	cm := app.Common{}
	adminId := ctx.GetInt(app.GlobalAdminId)
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if req.Pid > 0 {
			finded, err := store.WithContext(ctx).ItemCategory.Query().SoftDelete().
				Where(itemcategory.ID(req.Pid)).Exist(ctx)
			if err != nil {
				return err
			}
			if !finded {
				return errorno.NewErr(errorno.PidNotfoundError)
			}
		}
		if req.Id > 0 { //编辑
			_, err := bc.UpdateItemCategory(ctx, req.Name, req.Code, req.Desc, req.Id, req.SortOrder, req.Pid)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑项目", req.Name)
		} else {
			_, err := bc.AddItemCategory(ctx, req.Name, req.Code, req.Desc, req.SortOrder, req.Pid)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加项目", req.Name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除项目
func DelItemCategory(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.DelItemCategory(ctx, req.Id)
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

//设置项目状态
func SetItemStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.SetItemStatus(ctx, req.Id, req.Status)
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

//项目查询(带分页)
func GetItemCategoryByPage(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.ItemCategoryListSuccess
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).ItemCategory.Query().SoftDelete().
			Where(itemcategory.PidIsNil())
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(itemcategory.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			query = query.Where(itemcategory.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(itemcategory.Status(uint8(*req.Status)))
		}

		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		query = query.ForPage(req.Page, req.PageSize).WithChildren().Order(ent.Desc("id"))

		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.ItemDetail{}
			detail.ID = v.ID
			detail.Name = v.Name
			detail.SortOrder = v.SortOrder
			detail.CreatedAt = v.CreatedAt
			detail.Status = v.Status
			detail.Pid = v.Pid
			detail.Code = v.Code
			detail.Children = ent.ItemCategories{}
			if !app2.IsNil(v.Edges.Children) && len(v.Edges.Children) > 0 {
				detail.Children = v.Edges.Children
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

//地区查询（无分页）
func GetItemCategoryAll(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var itemCategoryList ent.ItemCategories
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).ItemCategory.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(itemcategory.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			query = query.Where(itemcategory.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(itemcategory.Status(uint8(*req.Status)))
		}
		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		itemCategoryList = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return itemCategoryList, nil
}

//添加（编辑）专业
func SetMajor(ctx *gin.Context) (interface{}, error) {
	var req request.SetBasicTag
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	cm := app.Common{}
	adminId := ctx.GetInt(app.GlobalAdminId)
	remark := ""
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if req.Id > 0 { //编辑
			_, err := bc.UpdateMajor(ctx, req.Name, req.Code, req.Desc, req.Id, req.SortOrder)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "编辑专业", req.Name)
		} else {
			_, err := bc.AddMajor(ctx, req.Name, req.Code, req.Desc, req.SortOrder)
			if err != nil {
				return err
			}
			remark = fmt.Sprintf("%s:%s", "添加专业", req.Name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	_ = cm.SetOperationLog(ctx, adminId, remark)
	return nil, nil
}

//删除专业
func DelMajor(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.DelMajor(ctx, req.Id)
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

//设置专业状态
func SetMajorStatus(ctx *gin.Context) (interface{}, error) {
	var req request.BasicConfigStatus
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.SetMajorStatus(ctx, req.Id, req.Status)
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

//专业查询(带分页)
func GetMajorByPage(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var res response.MajorListSuccess
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Major.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(major.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			query = query.Where(major.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(major.Status(uint8(*req.Status)))
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
		res.List = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//专业查询（无分页）
func GetMajorAll(ctx *gin.Context) (interface{}, error) {
	var req request.BasicTagPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	var majorList ent.Majors
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).Major.Query().SoftDelete()
		if !app2.StringIsEmpty(req.Name) {
			query = query.Where(major.Name(req.Name))
		}
		if !app2.StringIsEmpty(req.Code) {
			query = query.Where(major.Code(req.Code))
		}
		if !app2.IsNil(req.Status) {
			query = query.Where(major.Status(uint8(*req.Status)))
		}
		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		majorList = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return majorList, nil
}

//阿里云oss
func AuthOss(ctx *gin.Context) (interface{}, error) {
	var req request.AuthOssRequest
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	authOssResponse := response.AuthOssResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		dir := ""
		switch *req.Action {
		case "tk_question":
			dir = "tk_question"
			break
		case "avatar":
			dir = "avatar"
			break
		case "kc_course":
			dir = "kc_course"
			break
		case "questions_and_answers":
			dir = "questions_and_answers"
			break
		case "comment":
			dir = "comment"
			break
		case "pop_img":
			dir = "pop_img"
			break
		case "common":
			dir = "common"
			break
		default:
			return errorno.NewErr(errorno.Errno{
				Code:    20001,
				Message: "上传路径不存在",
			})
		}

		policyToken, err := oss.GetPolicyToken(dir)
		if err != nil {
			return err
		}
		authOssResponse.Accessid = policyToken.AccessKeyId
		authOssResponse.Host = policyToken.Host
		authOssResponse.Callback = policyToken.Callback
		authOssResponse.Policy = policyToken.Policy
		authOssResponse.Signature = policyToken.Signature
		authOssResponse.Expire = policyToken.Expire
		authOssResponse.CallbackInfo.CallbackBody = policyToken.CallbackParam.CallbackBody
		authOssResponse.CallbackInfo.CallbackBodyType = policyToken.CallbackParam.CallbackBodyType
		authOssResponse.CallbackInfo.CallbackUrl = policyToken.CallbackParam.CallbackUrl
		authOssResponse.Key = dir
		authOssResponse.Endpoint = config.ServerConfig.Aliyun.OssEndpoint
		authOssResponse.EndpointHost = config.ServerConfig.Aliyun.OssEndpointHost

		cacheOssCredentials := cache.OssCredentialsKey
		if credentials, ok := cache.MemoryCache.Get(cacheOssCredentials); ok {
			if stringCredentials, ok := credentials.(*sts.Credentials); ok {
				authOssResponse.Token.AccessKeyId = stringCredentials.AccessKeyId
				authOssResponse.Token.AccessKeySecret = stringCredentials.AccessKeySecret
				authOssResponse.Token.Expiration = stringCredentials.Expiration
				authOssResponse.Token.SecurityToken = stringCredentials.SecurityToken
			} else {
				return errorno.NewErr(errorno.Errno{
					Code:    20001,
					Message: "oss验证失败",
				})
			}

		} else {
			credentials, err := oss.GetRole()
			if err != nil {
				return err
			}
			authOssResponse.Token.AccessKeyId = credentials.AccessKeyId
			authOssResponse.Token.AccessKeySecret = credentials.AccessKeySecret
			authOssResponse.Token.Expiration = credentials.Expiration
			authOssResponse.Token.SecurityToken = credentials.SecurityToken
			cache.MemoryCache.SetWithExpiration(cacheOssCredentials, credentials, 60*time.Minute)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return authOssResponse, nil
}

//oss处理
func OssCallbackHandler(ctx *gin.Context) (interface{}, error) {
	var req request.OssCallBack
	err := ctx.Bind(&req)
	if err != nil {
		log.Info(ctx.Request.MultipartForm)
		//oss.ResponseFailed(ctx.Writer)
		return nil, err
	}

	logRes, err := json.Marshal(req)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	} else {
		log.Info(string(logRes))
	}
	if req.Filename == "" {
		log.Info("参数错误")
		//oss.ResponseFailed(ctx.Writer)
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "参数错误",
		})
	}
	var attachmentFile interface{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		c := app.BasicConfig{}
		attachment, err := c.ProcessOssCallback(ctx, req)
		if err != nil {
			return err
		}
		attachmentFile = attachment
		return nil
	})
	if err != nil {
		log.Error(err.Error())
		//oss.ResponseFailed(ctx.Writer)
		return nil, err
	}
	//oss.ResponseSuccess(ctx.Writer)
	return attachmentFile, nil
}

//添加app协议
func SetAppAgreement(ctx *gin.Context) (interface{}, error) {
	var req request.SetAgreement
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		_, err := bc.SetAgreement(ctx, req)
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

func GetAppAgreement(ctx *gin.Context) (interface{}, error) {
	var req request.GetAgreement
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetAgreementResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		agreement, err := store.WithContext(ctx).AppAgreement.Query().
			Where(appagreement.Type(uint8(*req.Type))).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		res.Type = int(agreement.Type)
		res.Detail = agreement.Detail
		return nil
	})
	return res, nil
}

//添加（编辑）app版本
func SetAppVersion(ctx *gin.Context) (interface{}, error) {
	var req request.SetAppVersion
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		if !app2.IsNil(req.Id) { //编辑
			_, err := bc.UpdateAppVersion(ctx, req)
			if err != nil {
				return err
			}
		} else {
			_, err := bc.AddAppVersion(ctx, req)
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

//删除app版本
func DelAppVersion(ctx *gin.Context) (interface{}, error) {
	var req request.IdOnly
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	bc := app.BasicConfig{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		err = bc.DelAppVersion(ctx, req.Id)
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

//app版本列表（分页）
func GetAppVersionByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetAppVersionPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.AppVersionPageListSuccess{
		List: ent.AppVersions{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).AppVersion.Query().SoftDelete()
		if !app2.IsNil(req.Name) {
			query = query.Where(appversion.NameContains(*req.Name))
		}
		if !app2.IsNil(req.Sn) {
			query = query.Where(appversion.SnContains(*req.Sn))
		}
		if !app2.IsNil(req.PhoneType) {
			query = query.Where(appversion.PhoneType(uint8(*req.PhoneType)))
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
		res.List = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//后台登录日志
func GetAdminLoginLogByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetLoginLogPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.GetLoginLogPageListResponse{
		List: ent.AdminLoginLogs{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).AdminLoginLog.
			Query().SoftDelete()

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
		res.List = list
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//后台操作日志
func GetAdminOperationLogByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetOperateLogPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.OperationLogListResponse{
		List: []response.OperationLogDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).AdminOperationLog.
			Query().SoftDelete()

		if !app2.IsNil(req.AdminName) {
			query = query.Where(adminoperationlog.HasAdminWith(
				admin.RealNameContains(*req.AdminName)))
		}

		if !app2.IsNil(req.StartAt) && !app2.IsNil(req.EndAt) {
			query = query.Where(adminoperationlog.CreatedAtLT(*req.EndAt)).
				Where(adminoperationlog.CreatedAtGT(*req.StartAt))
		}

		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		query = query.WithAdmin().ForPage(req.Page, req.PageSize).Order(ent.Desc("id"))

		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.OperationLogDetail{}
			detail.Id = v.ID
			detail.Ip = v.IP
			detail.Record = v.Record
			detail.AdminId = v.AdminID
			detail.CreatedAt = v.CreatedAt
			detail.AdminName = ""
			if !app2.IsNil(v.Edges.Admin) {
				detail.AdminName = v.Edges.Admin.RealName
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

//首页统计
func GetIndexStatistic(ctx *gin.Context) (interface{}, error) {
	res := []response.IndexStatistic{}
	indexStatistic := response.IndexStatistic{}
	err := store.WithTx(ctx, func(ctx context.Context) error {
		s := store.WithContext(ctx)
		todayTime, err := time.Parse("2006-01-02", time.Now().
			Format("2006-01-02"))
		if err != nil {
			return err
		}

		dd, err := time.ParseDuration("24h")
		if err != nil {
			return err
		}
		tomorrowTime := todayTime.Add(dd)

		LoginCount, err := s.UserLoginLog.Query().
			Where(userloginlog.CreatedAtGT(todayTime)).Count(ctx)
		if err != nil {
			return err
		}
		indexStatistic.Label = "今日登录用户"
		indexStatistic.Value = LoginCount
		res = append(res, indexStatistic)

		regCount, err := s.User.Query().
			Where(user.CreatedAtGT(todayTime)).Count(ctx)
		if err != nil {
			return err
		}
		indexStatistic.Label = "今日注册用户"
		indexStatistic.Value = regCount
		res = append(res, indexStatistic)

		LiveCount, err := s.KcCourseSmallCategory.Query().
			Where(kccoursesmallcategory.LiveSmallStartGT(todayTime),
				kccoursesmallcategory.LiveSmallStartLT(tomorrowTime)).
			Count(ctx)
		if err != nil {
			return err
		}
		indexStatistic.Label = "今日直播课"
		indexStatistic.Value = LiveCount
		res = append(res, indexStatistic)

		CourseCount, err := s.KcCourse.Query().
			Where(kccourse.CreatedAtGT(todayTime)).Count(ctx)
		if err != nil {
			return err
		}
		indexStatistic.Label = "今日新增课程"
		indexStatistic.Value = CourseCount
		res = append(res, indexStatistic)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}

		QuestionCount, err := s.TkQuestion.Query().
			Where(tkquestion.CreatedAtGT(todayTime)).
			Where(tkquestion.PidIsNil()).
			Count(ctx)
		if err != nil {
			return err
		}
		indexStatistic.Label = "今日新增题目"
		indexStatistic.Value = QuestionCount
		res = append(res, indexStatistic)

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//app登录日志
func GetLoginUserByPage(ctx *gin.Context) (interface{}, error) {
	var req request.GetLoginUserPageList
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	res := response.LoginUserListResponse{
		List: []response.LoginUserDetail{},
	}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		query := store.WithContext(ctx).UserLoginLog.
			Query().SoftDelete()

		if !app2.IsNil(req.Username) {
			query = query.Where(userloginlog.HasUserWith(
				user.UsernameContains(*req.Username)))
		}

		if !app2.IsNil(req.StartAt) && !app2.IsNil(req.EndAt) {
			query = query.Where(userloginlog.CreatedAtLT(*req.EndAt)).
				Where(userloginlog.CreatedAtGT(*req.StartAt))
		}

		count, err := query.Count(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		res.Page.Total = count
		query = query.WithUser().ForPage(req.Page, req.PageSize).Order(ent.Desc("id"))

		list, err := query.All(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		for _, v := range list {
			detail := response.LoginUserDetail{}
			detail.ID = v.ID
			detail.Platform = v.Platform
			detail.IP = v.IP
			detail.CreatedAt = v.CreatedAt
			detail.Version = v.Version
			detail.Cid = v.Cid
			detail.Device = v.Device
			detail.LatestLoginAt = v.LatestLoginAt
			detail.Username = ""
			detail.UserId = 0
			detail.LatestLoginAt = v.LatestLoginAt
			if !app2.IsNil(v.Edges.User) {
				detail.Username = v.Edges.User.Username
				detail.UserId = v.Edges.User.ID
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
