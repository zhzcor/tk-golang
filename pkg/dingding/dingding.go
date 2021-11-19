package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gserver/internal/cache"
	"gserver/internal/config"
	"gserver/internal/errorno"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Dding struct {
}

//const (
//	DingAccessKey string = "dingozow3ah7wrzmdfrd"
//	DingAppSecret string = "u5rEKiYWNCreLhyGNueELyW_IR-F25o-wym1ZfmkBXRAR46ogcjM3UJCe_WzYAiC"
//
//	DingCorSecret string = "GqiXdJ_2EVCE1Ia_NTnEQDR0dFiTnvE5qB5Hwu6lENTMwvdVdVlKLmOEiRuMUT5w"
//	DingCorAppId  string = "dingoazzymb712gjmdnqtc"
//)

const (
	Domain            string = "https://oapi.dingtalk.com"
	GetUserInfoByCode string = "/sns/getuserinfo_bycode" // 临时授权码获取授权用户的个人信息
	GetToken          string = "/gettoken"
	GetUserByUnionId  string = "/user/getUseridByUnionid"
)

func (d Dding) LoginByQRcode(code string) (userid string, err error) {
	var resp *http.Response
	//服务端通过临时授权码获取授权用户的个人信息
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)                // 毫秒时间戳
	signature := EncodeSHA256(timestamp, config.ServerConfig.Dingding.DingCorSecret) // 加密签名  加密算法见我另一个函数
	url2 := fmt.Sprintf(
		Domain+GetUserInfoByCode+"?accessKey=%s&timestamp=%s&signature=%s",
		config.ServerConfig.Dingding.DingCorAppId, timestamp, signature)
	p := struct {
		Tmp_auth_code string `json:"tmp_auth_code"`
	}{code} // post数据
	p1, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	p3 := strings.NewReader(string(p1)) //构建post数据
	//fmt.Println(string(p1))
	resp, err = http.Post(url2, "application/json;charset=UTF-8", p3)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i) ///返回的数据给i
	errcode := i["errcode"].(float64)
	if errcode != 0 {
		return "", errors.New(i["errmsg"].(string))
	}
	unionid := i["user_info"].(map[string]interface{})["unionid"].(string) // unionid 可以用来查询userinfo
	accessToken, err := GetAccessToken()                                   // 获取accessToken
	if err != nil {
		return "", errors.New(fmt.Sprintf("登录错误accesstoken获取失败: %s", err))
	}
	userid, err = GetUseridByUnionId(accessToken, unionid)
	if err != nil {
		return "", errors.New(fmt.Sprintf("登录错误userid获取失败: %s", err))
	}
	return userid, nil
}

func GetUseridByUnionId(accessToken, unionId string) (userid string, err error) {
	//根据unionId获取userid
	var resp *http.Response
	url := fmt.Sprintf(Domain+GetUserByUnionId+"?access_token=%s&unionid=%s",
		accessToken, unionId)
	resp, err = http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i)
	errcode := i["errcode"].(float64)
	if errcode != 0 {
		return "", errors.New(i["errmsg"].(string))
	}
	return i["userid"].(string), nil
}

func GetAccessToken() (string, error) {
	dingScanToken := cache.DingScanAccessToken
	returnToken := ""
	if token, ok := cache.MemoryCache.Get(dingScanToken); ok {
		switch token.(type) {
		case string:
			returnToken = token.(string)
			break
		default:
			return "", errorno.NewErr(errorno.Errno{
				Code:    20001,
				Message: "token缓存验证失败",
			})
		}

	} else {
		token, err := getAccessToken()
		if err != nil {
			return "", err
		}
		returnToken = token
		cache.MemoryCache.SetWithExpiration(dingScanToken, returnToken, 60*time.Minute)
	}

	return returnToken, nil
}

func getAccessToken() (string, error) {
	var resp *http.Response
	//获取access_token
	url := fmt.Sprintf(Domain+GetToken+"?appkey=%s&appsecret=%s",
		config.ServerConfig.Dingding.DingAccessKey, config.ServerConfig.Dingding.DingAppSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i)
	if i["errcode"].(float64) == 0 {
		return i["access_token"].(string), nil
	}
	return "", errors.New("accesstoken获取错误：" + i["errmsg"].(string))
}

func EncodeSHA256(message, secret string) string {
	// 钉钉签名算法实现
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sum := h.Sum(nil) // 二进制流
	message1 := base64.StdEncoding.EncodeToString(sum)
	uv := url.Values{}
	uv.Add("0", message1)
	message2 := uv.Encode()[2:]
	return message2

}
