package appapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"tkserver/httpapi/appapi/request"
	/*"tkserver/httpapi/appapi/response"
	"tkserver/internal/app"*/
	"tkserver/internal/errorno"
	/*	"tkserver/internal/store"
	 */"tkserver/pkg/wechatapplet"
)

// 微信小程序登录
func AppletWeChatLogin(ctx *gin.Context) (interface{}, error) {
	var req request.WechatAppletLogin
	/*	var res response.LoginSuccess
	 */err := ctx.Bind(&req)
	if err != nil {
		return nil, errorno.NewParamErr(err)
	}
	// 根据code获取 openID 和 session_key
	/*	wxLoginResp,err := wechatapplet.WXLogin(req.Code)
		if err != nil {
			return nil, errorno.NewInternalErr(err)
		}*/

	/*	dataInfo,err := wechatapplet.DecryptWXOpenData(wxLoginResp.SessionKey,req.EncryptedData,req.Iv)
	 */
	dataInfo, err := wechatapplet.Decrypt(req.EncryptedData, "SE/BLocg+sMlvcKmxm8vQA==", req.Iv)
	fmt.Printf("%v ", dataInfo)
	if err != nil {
		return nil, errorno.NewInternalErr(err)
	}

	return nil, nil
	/*	var token string

		s := store.WithContext(ctx)
		uc := app.UserCenter{}*/

	/*	// 保存登录态
		session := sessions.Default(c)
		session.Set("openid", wxLoginResp.OpenId)
		session.Set("sessionKey", wxLoginResp.SessionKey )

		// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
		mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
		// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
		// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
		c.String(200, mySession)*/

}

// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
