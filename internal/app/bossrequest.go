package app

import (
	"bytes"
	"encoding/json"
	"gserver/internal/errorno"
	"io/ioutil"
	"net/http"
	"time"
)

type BossRequest struct {
}

type UserInfoDetail struct {
	BossUserId int    `json:"boss_user_id"`
	Phone      string `json:"phone"`
	IdCard     string `json:"id_card"`
	CardType   int    `json:"card_type"`
	Username   string `json:"username"`
}

func (b *BossRequest) BossUserByPhone(p string,bossHost string) (*UserInfoDetail,error) {
	resDetail := UserInfoDetail{}

	reqDetail := make(map[string]string)
	reqDetail["phone"] = p
	bytesData, err := json.Marshal(reqDetail)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	url := bossHost + "/api/v3/user_sync/by_phone"
	requestInfo, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	requestInfo.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: time.Second *3,
	}
	resp, err := client.Do(requestInfo)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tempMap map[string]interface{}
	err = json.Unmarshal(respBytes, &tempMap)
	if err != nil {
		return nil, err
	}

	if tempMap["res_info"] != "ok" {
		return nil, errorno.NewErr(errorno.Errno{
			Code:    20001,
			Message: "数据不存在",
		})
	}
	userDetailData, err := json.Marshal(tempMap["data"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userDetailData, &resDetail)
	if err != nil {
		return nil, err
	}
	return &resDetail, nil
}
