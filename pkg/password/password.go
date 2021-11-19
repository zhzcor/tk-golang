package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Gen(originPwd string) string {
	buf := []byte(originPwd)
	hashpwd, _ := bcrypt.GenerateFromPassword(buf, bcrypt.DefaultCost)
	return string(hashpwd)
}

func Comp(pwd, hashedPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)); err != nil {
		return false
	} else {
		return true
	}
}
