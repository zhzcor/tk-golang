package dingding

import (
	"fmt"
	"testing"
)

//func TestGetCodeUserInfo(t *testing.T) {
//
//	b := Dding{}
//	ss, err := b.GetUserInfoByCode("5532dec9ff3a30438bdd62370f786a39")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(ss)
//}

func TestGetCodeUserInfo1(t *testing.T) {

	b := Dding{}
	ss, err := b.LoginByQRcode("a3407faf328434e0bf87b7d6ccc79387")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ss)
}

func TestGetToken(t *testing.T) {

	ss, err := GetAccessToken()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ss)
}
