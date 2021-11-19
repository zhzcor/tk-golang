package errorno

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

// 定义错误码的地方
var (
	storeError = Errno{
		Code:    500,
		Message: "数据服务异常，请稍后再试",
	}
	paramError = Errno{
		Code:    400,
		Message: "参数错误",
	}
	internalError = Errno{
		Code:    501,
		Message: "内部错误，请稍后再试",
	}

	TokenError = Errno{
		Code:    -1,
		Message: "token 错误或已过期，请重新登录",
	}

	UserNotfoundError = Errno{
		Code:    10001,
		Message: "用户不存在",
	}
	CreateUserExistsError = Errno{
		Code:    10003,
		Message: "创建用户失败，用户已存在",
	}
	LoginFailError = Errno{
		Code:    10004,
		Message: "账号或密码错误",
	}

	DataExistsError = Errno{
		Code:    10005,
		Message: "数据已存在",
	}

	LoginCodeFailError = Errno{
		Code:    10006,
		Message: "验证码错误或过期",
	}

	LoginPhoneFailError = Errno{
		Code:    10007,
		Message: "手机号错误",
	}

	DataNotExistsError = Errno{
		Code:    10008,
		Message: "数据不存在",
	}
	PasswordError = Errno{
		Code:    10009,
		Message: "密码错误",
	}
)

func (err Errno) Error() string {
	return err.Message
}

func (err Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, fmt.Sprintf("%+v", err.Errored))
}

func NewErr(errno Errno, errs ...error) *Err {
	newErr := Err{
		Code:    errno.Code,
		Message: errno.Message,
	}
	for _, e := range errs {
		newErr.Errored = append(newErr.Errored, errors.WithStack(e))
	}
	runtime.Callers(1, newErr.stack)
	return &newErr
}

func NewStoreErr(err error) *Err {
	errno := storeError
	newErr := Err{
		Code:    errno.Code,
		Message: errno.Message,
	}
	newErr.Errored = append(newErr.Errored, errors.WithStack(err))
	return &newErr
}

func NewParamErr(err error) *Err {
	errno := paramError
	newErr := Err{
		Code:    errno.Code,
		Message: errno.Message,
	}
	newErr.Errored = append(newErr.Errored, errors.WithStack(err))
	return &newErr
}

func NewInternalErr(err error) *Err {
	errno := internalError
	newErr := Err{
		Code:    errno.Code,
		Message: errno.Message,
	}
	newErr.Errored = append(newErr.Errored, errors.WithStack(err))
	return &newErr
}

func IsErr(err error) bool {
	return errors.Is(err, Err{})
}

func ErrAs(err error) *Err {
	var e *Err
	if errors.As(err, &e) {
		return e
	}
	return nil
}
