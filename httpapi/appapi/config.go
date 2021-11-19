package appapi

import (
	"context"
	"tkserver/internal/store/ent/appversion"

	"github.com/gin-gonic/gin"
	"tkserver/httpapi/appapi/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/store"
	"tkserver/internal/store/ent"
	"tkserver/internal/store/ent/appagreement"
)

//app协议
func GetAppAgreement(ctx *gin.Context) (interface{}, error) {
	var req request.GetAgreement
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}

	res := response.GetAgreementResponse{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		agreement, err := store.WithContext(ctx).AppAgreement.Query().
			Where(appagreement.Type(uint8(req.Type))).First(ctx)
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

//获取版本
func GetAppVersion(ctx *gin.Context) (interface{}, error) {
	var req request.PhoneType
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	s := store.WithContext(ctx)
	data := s.AppVersion.Query().SoftDelete().Where(appversion.PhoneType(uint8(req.Type))).Order(ent.Desc("id")).FirstX(ctx)

	if data != nil {
		var res response.AppVersion
		res.PhoneType = int(data.PhoneType)
		res.Url = data.URL
		res.Name = data.Name
		res.IsForceUpdate = int(data.IsForceUpdate)
		res.Sn = data.Sn

		return data, nil
	}

	return nil, nil
}
