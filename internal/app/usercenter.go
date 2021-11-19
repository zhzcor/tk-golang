package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"gserver/httpapi/appapi/request"
	"gserver/internal/errorno"
	"gserver/internal/store/ent/user"
	"gserver/pkg/password"
	"time"

	"gserver/internal/store"
	"gserver/internal/store/ent"
)

var jwtKey = []byte("jwtkey1234")

type UserCenter struct {
}

func (c UserCenter) Create(ctx context.Context, userName string, pwd string) (int, error, *ent.User) {
	s := store.WithContext(ctx)
	fined, err := store.WithContext(ctx).User.
		Query().
		Where(user.Phone(userName)).
		Exist(ctx)
	if err != nil {
		return 0, errorno.NewStoreErr(err), nil
	}
	if !fined {
		var newUser *ent.User
		var err error
		//err = validation.ValidateStruct(&info, validation.Field(&info.Email, validation.Required, is.Email))
		//if err != nil {
		//	return "", errorno.NewParamErr(err)
		//}
		newUser, err = s.User.Create().
			SetUsername(userName).
			SetPassword(pwd).
			SetPhone(userName).
			/*			SetEmail(info.Email).
						SetNickname(info.NickName).*/
			Save(ctx)
		if err != nil {
			return 0, errorno.NewStoreErr(err), nil
		}

		return newUser.ID, nil, newUser
	} else {
		return 0, errorno.NewErr(errorno.CreateUserExistsError), nil
	}
}

// MakeToken 用户登录，直接用uid 就可以登录并生成token，
// 每次登录会增加一次日志记录
// 密码验证交由api层验证
func (u UserCenter) MakeToken(ctx context.Context, uid int, media struct {
	Ip       string
	City     string
	Province string
}) (string, error) {
	db := store.WithContext(ctx)
	uinfo, err := db.User.Query().Where(user.ID(uid)).Only(ctx)
	if errno := u.CheckUserNotFound(err); errno != nil {
		return "", errno
	}
	// 创建一个日志
	_, err = db.UserLoginLog.Create().
		SetIP(media.Ip).
		SetCity(media.City).
		SetProvince(media.Province).
		SetUser(uinfo).
		Save(ctx)
	if err != nil {
		return "", errorno.NewStoreErr(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tk, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errorno.NewInternalErr(err)
	}
	return tk, nil
}

//登录验证
func (u UserCenter) CheckPassword(ctx context.Context, userId, pwd string) error {
	uinfo, err := store.WithContext(ctx).User.
		Query().
		Where(user.UUID(userId)).
		First(ctx)
	if errno := u.CheckUserNotFound(err); errno != nil {
		return errno
	}
	if password.Comp(pwd+uinfo.Salt, uinfo.Password) {
		return nil
	}
	return errorno.NewErr(errorno.LoginFailError)
}

func (u UserCenter) CheckUserNotFound(err error) error {
	if err != nil && !ent.IsNotFound(err) {
		return errorno.NewStoreErr(err)
	}
	if ent.IsNotFound(err) {
		return errorno.NewErr(errorno.UserNotfoundError)
	}
	return nil
}

func (c UserCenter) CheckToken(tk string) (int, error) {
	t, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, errorno.NewErr(errorno.TokenError)
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		str, _ := claims["uid"].(float64)
		s := int(str)
		return s, nil
	}
	return 0, errorno.NewErr(errorno.TokenError)
}

//修改用户资料
func (c UserCenter) UpdateUser(ctx context.Context, id int, data *request.UpdateUserInfo) (*ent.User,error) {
	query := store.WithContext(ctx).User.
		UpdateOneID(id)
	if data.Birthday != "" {
		query.SetBirthday(data.BirthdayTime)
	}
	if data.FromCityID != 0 {
		query.SetFromCityID(data.FromCityID)
	}
	if data.FromItemCategoryID != 0 {
		query.SetFromItemCategoryID(data.FromItemCategoryID)
	}
	if data.Username != "" {
		query.SetUsername(data.Username)
	}
	if data.Avatar != "" {
		query.SetAvatar(data.Avatar)
	}
	if data.SignRemark != "" {
		query.SetSignRemark(data.SignRemark)
	}
	if data.Sex != 0 {
		query.SetSex(uint8(data.Sex))
	}

	info := query.SaveX(ctx)

	return info,nil

}
