package appapi

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	"time"
	"tkserver/httpapi/appapi/request"
	"tkserver/httpapi/appapi/response"
	"tkserver/internal/app"
	"tkserver/internal/app/apiservice"
	"tkserver/internal/app/sms"
	"tkserver/internal/config"
	"tkserver/internal/errorno"
	"tkserver/internal/store"
	user2 "tkserver/internal/store/ent/user"
	"tkserver/pkg/oss"
	"tkserver/pkg/password"
	app2 "tkserver/pkg/requestparam"

	"tkserver/internal/cache"
	"tkserver/pkg/log"
)

const (
	regFromSynchronization = 4 //系统同步过来的用户
)

//用户登陆
func UserPwdLogin(ctx *gin.Context) (interface{}, error) {
	var req request.Register
	var token string
	var res response.LoginSuccess
	s := store.WithContext(ctx)
	uc := app.UserCenter{}
	err := ctx.Bind(&req)

	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	user, err := s.User.Query().Where(user2.Phone(req.Phone)).First(ctx)
	if errno := uc.CheckUserNotFound(err); errno != nil {
		return nil, errorno.NewErr(errorno.UserNotfoundError)
	} else {
		if !password.Comp(req.Password+user.Salt, user.Password) {
			return nil, errorno.NewErr(errorno.LoginFailError)
		}
	}

	token, err = uc.MakeToken(ctx, user.ID, struct {
		Ip       string
		City     string
		Province string
	}{})
	if err != nil {
		return nil, err
	}
	res.UserId = user.ID
	res.Token = token
	res.Username = user.Username
	res.Nickname = user.Nickname
	res.UUID = user.UUID
	res.Phone = user.Phone
	if user.Birthday != nil {
		res.Birthday = user.Birthday
		res.BirthdayStr = user.Birthday.Format("2006-01-02")
	}

	res.Sex = user.Sex
	res.CardType = user.CardType
	res.Status = user.Status
	res.CreatedAt = user.CreatedAt
	res.FromItemCategoryID = user.FromItemCategoryID
	res.FromCityID = user.FromCityID

	if user.Avatar != "" {
		res.Avatar = app2.GetOssHost() + user.Avatar
	}
	res.SignRemark = user.SignRemark
	return res, nil
}

//pc修改密码
func UserUpdatePwd(ctx *gin.Context) (interface{}, error) {
	var req request.OldPwd
	err := ctx.Bind(&req)
	uid, _ := ctx.Get("uid")
	uc := app.UserCenter{}

	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	s := store.WithContext(ctx)

	user, err := s.User.Query().Where(user2.ID(uid.(int))).First(ctx)

	if errno := uc.CheckUserNotFound(err); errno != nil {
		return nil, errorno.NewErr(errorno.UserNotfoundError)
	} else {

		if !password.Comp(req.Password+user.Salt, user.Password) {
			return nil, errorno.NewErr(errorno.LoginFailError)
		}
		s.User.Update().Where(user2.ID(uid.(int))).SetPassword(req.NewPassword).SaveX(ctx)
	}

	return nil, nil
}

//验证验证码
func UserCheckCode(ctx *gin.Context) (interface{}, error) {
	var req request.CheckCode
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	service := sms.NewSmsService()
	if ok := service.CheckCode(req.SmsCode, req.Phone); ok != true {
		return nil, errorno.NewErr(errorno.LoginCodeFailError)
	}

	return nil, nil
}

//用户登陆
func UserSmsLogin(ctx *gin.Context) (interface{}, error) {
	var req request.SmsLogin
	var token string
	var res response.LoginSuccess
	s := store.WithContext(ctx)
	uc := app.UserCenter{}
    var userAgent = ctx.Request.Header.Get("User-Agent")
/*	boss := app.BossRequest{}
*/	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}

	c := app.Common{}

	reform := c.DoClientAgent(userAgent)
	RegFrom := 0
	if reform == "iPhone" {
		RegFrom = 1
	}else if reform == "Android"{
		RegFrom = 2
	}

	fmt.Printf("%v ", reform)

	if ok := sms.NewSmsService().CheckCode(req.SmsCode, req.Phone); ok != true {
		return nil, errorno.NewErr(errorno.LoginCodeFailError)
	}

	user, err := s.User.Query().Where(user2.Phone(req.Phone)).First(ctx)

	isNewUser := 0
	if errno := uc.CheckUserNotFound(err); errno != nil {
		//用户不存在创建角色
		_, err, info := uc.Create(ctx, req.Phone, req.Phone,RegFrom)

		if err != nil {
			return nil, err
		}
		log.Info("on user register:", info.ID)

		user = info
		//访问boss是否有用户
		/*bossUser, err := boss.BossUserByPhone(req.Phone, config.ServerConfig.BossHost)
		if err != nil {
			log.Info("request boss by phone:", err)
		} else {
			//绑定id并且跟新信息
			if bossUser != nil && bossUser.BossUserId > 0 {
				user = s.User.UpdateOneID(user.ID).SetBossUserID(bossUser.BossUserId).SetIDCard(bossUser.IdCard).
					SetUsername(bossUser.Username).SetCardType(uint8(bossUser.CardType)).SetPhone(bossUser.Phone).SaveX(ctx)
			}
		}*/
		isNewUser = 1
	}

	token, err = uc.MakeToken(ctx, user.ID, struct {
		Ip       string
		City     string
		Province string
	}{})
	if err != nil {
		return nil, err
	}
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

	//是否发送新人短信
	if isNewUser == 1 {
		smsInfo := sms.NewSmsService()
		var param = map[string]interface{}{
			"phone": user.Phone,
		}
		_, err := smsInfo.SendSmsCode(&smsInfo.Client, user.Phone, param, 1)
		if err != nil {
			log.Info("新人注册欢迎短信:", err)
		}
	}

	return res, nil
}

//发送短信登陆
func SendSms(ctx *gin.Context) (interface{}, error) {
	p := request.PhoneSms{}
	s := response.SmsPhone{}
	if err := ctx.Bind(&p); err != nil {
		return nil, errorno.NewParamErr(err)
	}
	smsInfo := sms.NewSmsService()
	if err := smsInfo.IsMobile(p.Phone); err != true {
		return nil, errorno.NewErr(errorno.LoginPhoneFailError)
	}
	code := smsInfo.RandCode()
	smsKey := cache.SmsCacheKey + p.Phone
	cache.MemoryCache.Set(smsKey, code)
	var param = map[string]interface{}{
		"code": code,
	}
	isSend, err := smsInfo.SendSmsCode(&smsInfo.Client, p.Phone, param, 0)

	if !isSend {
		return nil, err
	}
	s.Phone = p.Phone
	s.SmsCode = code
	return s, nil
}

//设置密码
func SetPwd(ctx *gin.Context) (interface{}, error) {
	p := request.SetPwd{}
	if err := ctx.Bind(&p); err != nil {
		return nil, errorno.NewParamErr(err)
	}
	smsInfo := sms.NewSmsService()
	if err := smsInfo.IsMobile(p.Phone); err != true {
		return nil, errorno.NewErr(errorno.LoginPhoneFailError)
	}
	fmt.Printf("%v ", p)

	if ok := sms.NewSmsService().CheckCode(p.SmsCode, p.Phone); ok != true {
		return nil, errorno.NewErr(errorno.LoginCodeFailError)
	}
	s := store.WithContext(ctx)
	_, err := s.User.Update().Where(user2.Phone(p.Phone)).SetPassword(p.Password).Save(ctx)
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return nil, nil
}

//个人资料项目地区筛选
func ScreenItems(ctx *gin.Context) (interface{}, error) {

	cityList, err := apiservice.CityService().GetCityList(ctx)
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}
	//项目
	list := apiservice.NewItemCate().GetItemCateList(ctx)

	uid, _ := ctx.Get("uid")

	user, err := store.WithContext(ctx).User.Get(ctx, uid.(int))

	cityName := ""

	itemName := ""

	if err != nil {
		return nil, nil
	}

	for _, v := range cityList {
		if user.FromCityID == v.Id {
			cityName = v.Name
		}
	}

	for _, v := range list {
		if user.FromItemCategoryID == v.Id {
			if v.Parent.Name != "" {
				itemName = v.Parent.Name + "/" + v.Name
			} else {
				itemName = v.Name
			}
		}
	}

	var data = map[string]interface{}{
		"user_status":    user.RegFrom,
		"user_city_name": cityName,
		"user_item_name": itemName,
		"cityList":       cityList,
		"itemList":       list,
	}

	return data, nil
}

//修改个人资料
func UpdateUserInfo(ctx *gin.Context) (interface{}, error) {
	req := request.UpdateUserInfo{}
	uid, ok := ctx.Get("uid")
	id, _ := uid.(int)
	if ok == false {
		return nil, errorno.NewInternalErr(errorno.TokenError)
	}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	req.BirthdayTime, _ = time.ParseInLocation("2006-01-02", req.Birthday, time.Local)

	info, _ := app.UserCenter{}.UpdateUser(ctx, id, &req)
	if info.Avatar != "" {
		info.Avatar = app2.GetOssHost() + info.Avatar
	}

	return info, nil
}

func UpdatePhone(ctx *gin.Context) (interface{}, error) {
	req := request.UpdatePhone{}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	uid, ok := ctx.Get("uid")
	if ok == false {
		return nil, errorno.NewParamErr(errorno.TokenError)
	}
	//短信验证码
	if ok := sms.NewSmsService().CheckCode(req.SmsCode, req.Phone); ok != true {
		return nil, errorno.NewErr(errorno.LoginCodeFailError)
	}
	s := store.WithContext(ctx)

	user, err := s.User.Get(ctx, uid.(int))

	if err != nil {
		return nil, errorno.TokenError
	}

	q := s.User.UpdateOneID(uid.(int)).
		SetPhone(req.NewPhone)
	if user.Username == user.Phone {
		q.SetUsername(req.NewPhone)
	}
	/*	q.SetPassword(req.Password)
	 */
	_, err = q.Save(ctx)
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return nil, nil
}

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
