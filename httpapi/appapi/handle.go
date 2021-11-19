package appapi

import (
	"context"
	"github.com/gin-gonic/gin"
	"gserver/httpapi/appapi/request"
	"gserver/httpapi/appapi/response"
	"gserver/internal/app"
	"gserver/internal/store"
	"gserver/pkg/log"
)

func Pong(ctx *gin.Context) (interface{}, error) {
	log.Info("on datbases", store.WithContext(ctx))
	var req request.Register
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	log.Info("pong logger!", "data", "123467")
	return "pong ok!", nil
}

//用户注册
func UserRegister(ctx *gin.Context) (interface{}, error) {
	var req request.Register
	var token string
	var userId int
	var res response.RegisterSuccess
	err := ctx.Bind(&req)
	if err != nil {
		return nil, err
	}
	uc := app.UserCenter{}
	err = store.WithTx(ctx, func(ctx context.Context) error {
		uid, err,user := uc.Create(ctx, req.Phone, req.Password)
		if err != nil {
			_ = user
			return err
		}
		token, err = uc.MakeToken(ctx, uid, struct {
			Ip       string
			City     string
			Province string
		}{})
		if err != nil {
			return err
		}
		log.Info("on user register:", uid)
		userId = uid
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.UserId = userId
	res.Token = token
	return res, nil
}
