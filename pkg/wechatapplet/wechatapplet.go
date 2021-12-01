package wechatapplet

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	WxAppletAppId    = "wx326001aaef2af162"
/*	WxAppletAppId    = "wx33f640141e02040e"
*/	WxWxAppletSecret = "db07f9355ef03f416a0dd48be2dbbdd0"
)

var (
	ErrAppIDNotMatch       = errors.New("app id not match")
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)


type WXUserDataCrypt struct {
	appID, sessionKey string
}

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//手机号
type WxUserInfo struct {
	PhoneNumber    string `json:"phone_number"`
	PurePhoneNumber   string `json:"pure_phone_number"`
	CountryCode  string `json:"country_code`
	/*Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`*/
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

func WXLogin(code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	url = fmt.Sprintf(url, WxAppletAppId, WxWxAppletSecret, code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}


// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

func Decrypt(encryptedData, sessionKey,iv string) (*WxUserInfo, error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	var userInfo WxUserInfo
	err = json.Unmarshal(cipherText, &userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo.Watermark.AppID != WxAppletAppId {
		return nil, ErrAppIDNotMatch
	}
	return &userInfo, nil
}

/*func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 校验微信返回的用户数据
func ValidateUserInfo(rawData, sessionKey, signature string) bool {
	signature2 := GetSha1(rawData + sessionKey)
	return signature == signature2
}

// SHA-1 加密
func GetSha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	res := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return res
}

func DecryptWXOpenData(sessionKey string, encryptData string, iv string) (map[string]interface{}, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	fmt.Println(string(dataBytes))
	m := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	temp := m["watermark"].(map[string]interface{})
	appid := temp["appid"].(string)
	if appid != WxAppletAppId {
		return nil, fmt.Errorf("invalid appid, get !%s!", appid)
	}
	if err != nil {
		return nil, err
	}
	return m, nil

}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}*/
