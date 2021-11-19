package appapi

import (
	"github.com/gin-gonic/gin"
	"gserver/httpapi/appapi/request"
	"gserver/httpapi/appapi/response"
	"gserver/internal/app"
	"gserver/internal/errorno"
	"gserver/internal/store"
	user2 "gserver/internal/store/ent/user"
	"gserver/pkg/password"
)

func UserLoginL(ctx *gin.Context) (interface{}, error) {
	s := store.WithContext(ctx)
	var req request.Register
	var token string
	var userId int
	var res response.LoginSuccess
	uc := app.UserCenter{}
	err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	user, err := s.User.Query().Where(user2.Username(req.Phone)).First(ctx)
	if errno := uc.CheckUserNotFound(err); errno != nil {
		return nil, errno
	}
	userId = user.ID
	token, err = uc.MakeToken(ctx, userId, struct {
		Ip       string
		City     string
		Province string
	}{})
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	if !password.Comp(req.Password+user.Salt, user.Password) {
		return nil, errorno.NewErr(errorno.LoginFailError)
	}
	res.UserId = userId
	res.Token = token

	return res, nil
}
