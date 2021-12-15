package appapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tkserver/httpapi/appapi/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/app"
	"tkserver/internal/store"
	user2 "tkserver/internal/store/ent/user"
	app2 "tkserver/pkg/requestparam"

	"tkserver/internal/errorno"
	"tkserver/pkg/wechatapplet"
)

// 微信小程序登录
func AppletWeChatLogin(ctx *gin.Context) (interface{}, error) {
	var req request.WechatAppletLogin
	var res response.LoginSuccess
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := wechatapplet.WXLogin(req.Code)
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	dataInfo, err := wechatapplet.Decrypt(req.EncryptedData, wxLoginResp.SessionKey, req.Iv)

	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}
	fmt.Printf("%v ", dataInfo)

	s := store.WithContext(ctx)

	user, err := s.User.Query().Where(user2.Phone(dataInfo.PhoneNumber)).First(ctx)
	uc := app.UserCenter{}

	if errno := uc.CheckUserNotFound(err); errno != nil {
		//用户不存在创建角色
		_, err, info := uc.Create(ctx, dataInfo.PhoneNumber, dataInfo.PhoneNumber,4)

		if err != nil {
			return nil, err
		}

		user = info

	}
	var token string

	token, err = uc.MakeToken(ctx, user.ID, struct {
		Ip       string
		City     string
		Province string
	}{})

	res.UserId = user.ID
	res.Token = token
	res.Username = user.Username
	res.Nickname = user.Nickname
	res.UUID = user.UUID
	res.RegFrom = user.RegFrom
	res.Phone = user.Phone
	if user.Birthday != nil && !user.Birthday.IsZero() {
		res.Birthday = user.Birthday
		res.BirthdayStr = user.Birthday.Format("2006-01-02")
	}

	res.Sex = user.Sex
	res.CardType = user.CardType
	res.Status = user.Status
	res.FromItemCategoryID = user.FromItemCategoryID
	res.FromCityID = user.FromCityID
	res.Status = user.Status
	res.CreatedAt = user.CreatedAt
	if user.Avatar != "" {
		res.Avatar = app2.GetOssHost() + user.Avatar
	}
	res.SignRemark = user.SignRemark

	return res, nil

}
