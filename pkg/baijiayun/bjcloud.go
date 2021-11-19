package baijiayun

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"tkserver/internal/config"
	"tkserver/internal/store"
	"tkserver/pkg/log"
)

//const (
//	PartnerId         int    = 49674498
//	PartnerKey        string = "WVGjvj1hd5jrRSD4oefxhdDU5pOALVeFAdmVb/ff/3jPWhI0FTyw9F9HWsKVTiPkDcSEwngBya6IyGeDA7YBbdBBYLVk"
//	Domain            string = "https://e49674498.at.baijiayun.com/"
//	NoHttpDomain      string = "e49674498"
//	VideoTranscodeUrl string = "http://pre-study-api.zhntu.com/admin/other/bjy_order_callback"
//	ClassCallbackUrl  string = "http://pre-study-api.zhntu.com/admin/other/bjy_class_callback"
//)

const (
	CreateLiveRoom          string = "/openapi/room/create"                           //创建直播教室
	UpdateLiveRoom          string = "/openapi/room/update"                           //更新直播教室
	DeleteLiveRoom          string = "/openapi/room/delete"                           //删除直播教室
	GetLiveRoomInfo         string = "/openapi/room/info"                             //获取直播教室信息
	GetOrderUploadUrl       string = "/openapi/video/getUploadUrl"                    //获取点播文件上传地址
	GetChangeStatus         string = "/openapi/video/getStatus"                       //获取点播视频转码状态
	GetLiveBackInfo         string = "/openapi/playback/getBasicInfo"                 //查询直播回放信息
	ReplaceLiveBack         string = "/openapi/playback/replacePlaybackByVideo"       //点播替换回放
	UpdateVideoInfo         string = "/openapi/video/update"                          //更新视频信息
	GetClassCallbackUrl     string = "/openapi/live_account/getClassCallbackUrl"      //查询直播上下课回调地址
	SetClassCallbackUrl     string = "/openapi/live_account/setClassCallbackUrl"      //设置直播上下课事件回调地址
	SetTranscodeCallbackUrl string = "/openapi/video_account/setTranscodeCallbackUrl" //设置转码回调地址（点播和回放）
	GetTranscodeCallbackUrl string = "/openapi/video_account/getTranscodeCallbackUrl" //查询转码回调地址（点播和回放）
	GetOnDemandTokenUrl     string = "/openapi/video/getPlayerToken"                  //获取播放器token（点播和回放）
	GetPlayBackUrl          string = "/openapi/playback/getPlayerToken"               //获取回放token（点播和回放）
	GetStudentCodeUrl       string = "/openapi/room/getcode"                          //生成用户参加码
	StartClass              string = "/openapi/live/startClass"                       //教室上课
	StopClass               string = "/openapi/live/stopClass"                        //教室下课
)

type CommonResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
	Ts   int                    `json:"ts"`
}
type BjyCloud struct {
	PartnerId int
}

type CreateLiveRoomResponse struct {
	RoomId      string `json:"room_id"`
	AdminCode   string `json:"admin_code"`
	TeacherCode string `json:"teacher_code"`
	StudentCode string `json:"student_code"`
}

type TokenList struct {
	Token   string `json:"token"`
	VideoId int    `json:"video_id"`
}

type ResStudentCode struct {
	StudentCode string `json:"student_code"`
}

//创建直播房间
func (b BjyCloud) CreateLiveRoom(title string, startAt, endAt int, classType, isMockLive, mockVideoId int) (CreateLiveRoomResponse, error) {
	params := map[string]interface{}{
		"title":            title,
		"start_time":       startAt,
		"end_time":         endAt,
		"type":             classType,
		"is_mock_live":     isMockLive,
		"mock_video_id":    mockVideoId,
		"mock_live_record": 1,
	}
	signedParams := b.getSignedParams(params)
	response := CreateLiveRoomResponse{}
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+CreateLiveRoom, signedParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

//更新直播房间
func (b BjyCloud) UpdateLiveRoom(title string, roomId, startAt, endAt, mockVideoId int) error {
	params := map[string]interface{}{
		"title":         title,
		"room_id":       roomId,
		"start_time":    startAt,
		"end_time":      endAt,
		"mock_video_id": mockVideoId,
	}
	signedParams := b.getSignedParams(params)

	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+UpdateLiveRoom, signedParams)
	if err != nil {
		return err
	}
	return nil
}

//删除直播房间
func (b BjyCloud) DeleteLiveRoom(roomId int) error {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	signedParams := b.getSignedParams(params)

	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+DeleteLiveRoom, signedParams)
	if err != nil {
		return err
	}
	return nil
}

type LiveRoomInfo struct {
	RoomId    int    `json:"room_id"`
	Title     string `json:"title"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Type      string `json:"type"`
}

//获取直播教室信息
func (b BjyCloud) GetLiveRoomInfo(roomId int) (LiveRoomInfo, error) {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	liveRoomInfo := LiveRoomInfo{}

	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetLiveRoomInfo, signedParams)
	if err != nil {
		return liveRoomInfo, err
	}
	err = json.Unmarshal(res, &liveRoomInfo)
	if err != nil {
		return liveRoomInfo, err
	}
	return liveRoomInfo, nil
}

type OrderUploadUrl struct {
	VideoId   int    `json:"video_id"`
	UploadUrl string `json:"upload_url"`
}

//获取点播视频上传地址
func (b BjyCloud) GetOrderUploadUrl(filename string, viewType int) (OrderUploadUrl, error) {

	params := map[string]interface{}{
		"file_name":       filename,
		"audio_with_view": viewType,
		"definition":      "16,1,2,4,8",
	}
	signedParams := b.getSignedParams(params)

	var orderUploadUrl OrderUploadUrl
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetOrderUploadUrl, signedParams)
	if err != nil {
		return OrderUploadUrl{}, err
	}

	err = json.Unmarshal(res, &orderUploadUrl)
	if err != nil {
		return OrderUploadUrl{}, err
	}
	return orderUploadUrl, nil
}

//点播视频/音频上传
func (b BjyCloud) OrderUpload(ctx context.Context, id int, videoUrl, uploadUrl, filename string) {
	b.UpdateUpdateVideoStatus(ctx, id, 2, "")

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	resp, err := http.Get(videoUrl)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error("order upload error : " + "resp status:" + fmt.Sprint(resp.StatusCode))
		b.UpdateUpdateVideoStatus(ctx, id, 4, "resp status:"+fmt.Sprint(resp.StatusCode))
		return
	}

	bin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}

	_, err = fw.Write(bin)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	w.Close()

	req, err := http.NewRequest("POST", uploadUrl, buf)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Error("order upload error : " + "resp status 1:" + fmt.Sprint(resp.StatusCode))
		b.UpdateUpdateVideoStatus(ctx, id, 4, "resp status 1:"+fmt.Sprint(resp.StatusCode))
		return
	}
	ff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}

	respData := CommonResponse{}
	err = json.Unmarshal(ff, &respData)
	if err != nil {
		log.Error("order upload error : " + err.Error())
		b.UpdateUpdateVideoStatus(ctx, id, 4, err.Error())
		return
	}
	if respData.Code != 1 {
		log.Error("order upload error : 点播文件上传失败")
		b.UpdateUpdateVideoStatus(ctx, id, 4, "点播文件上传失败")
		return
	}
	b.UpdateUpdateVideoStatus(ctx, id, 3, "")

}

//更新视频信息
func (b BjyCloud) UpdateVideoInfo(title string, videoId int) error {
	params := map[string]interface{}{
		"name":     title,
		"video_id": videoId,
	}
	signedParams := b.getSignedParams(params)

	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+UpdateVideoInfo, signedParams)
	if err != nil {
		return err
	}
	return nil
}

type VideoChangeCodeStatus struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

//获取视频转码状态
func (b BjyCloud) GetChangeCodeStatus(videoId int) (VideoChangeCodeStatus, error) {
	params := map[string]interface{}{
		"video_id": videoId,
	}
	signedParams := b.getSignedParams(params)

	var changeStatus VideoChangeCodeStatus
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetChangeStatus, signedParams)
	if err != nil {
		return VideoChangeCodeStatus{}, err
	}

	err = json.Unmarshal(res, &changeStatus)
	if err != nil {
		return VideoChangeCodeStatus{}, err
	}
	return changeStatus, nil
}

type LiveBackInfo struct {
	Status  int    `json:"status"`
	VideoId int    `json:"video_id"`
	Name    string `json:"name"`
}

//查询直播回放信息
func (b BjyCloud) GetLiveBackInfo(roomId int) (LiveBackInfo, error) {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetLiveBackInfo, signedParams)
	var liveBackInfo LiveBackInfo
	if err != nil {
		return LiveBackInfo{}, err
	}

	err = json.Unmarshal(res, &liveBackInfo)
	if err != nil {
		return LiveBackInfo{}, err
	}
	return liveBackInfo, nil
}

//点播替换回放
func (b BjyCloud) ReplaceLiveBack(roomId, videoId int) error {
	params := map[string]interface{}{
		"target_room_id": roomId,
		"video_id":       videoId,
	}
	signedParams := b.getSignedParams(params)
	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+ReplaceLiveBack, signedParams)
	if err != nil {
		return err
	}
	return nil
}

//设置视频转码回调地址
func (b BjyCloud) SetTranscodeCallbackUrl() error {
	params := map[string]interface{}{
		"url": config.ServerConfig.Baijiayun.VideoTranscodeUrl,
	}
	signedParams := b.getSignedParams(params)
	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+SetTranscodeCallbackUrl, signedParams)
	if err != nil {
		return err
	}
	return nil
}

//直播上课
func (b BjyCloud) SetStartClass(roomId int) error {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	signedParams := b.getSignedParams(params)
	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+StartClass, signedParams)
	if err != nil {
		return err
	}
	return nil
}

//直播下课
func (b BjyCloud) StopClass(roomId int) error {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	signedParams := b.getSignedParams(params)
	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+StopClass, signedParams)
	if err != nil {
		return err
	}
	return nil
}

type UrlInfo struct {
	Url string `json:"url"`
}

//获取视频转码回调地址
func (b BjyCloud) GetTranscodeCallbackUrl() (UrlInfo, error) {
	params := map[string]interface{}{}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetTranscodeCallbackUrl, signedParams)
	var urlInfo UrlInfo
	if err != nil {
		return UrlInfo{}, err
	}

	err = json.Unmarshal(res, &urlInfo)
	if err != nil {
		return UrlInfo{}, err
	}
	return urlInfo, nil
}

//设置上下课回调地址
func (b BjyCloud) SetClassCallbackUrl() error {
	params := map[string]interface{}{
		"url": config.ServerConfig.Baijiayun.ClassCallbackUrl,
	}
	signedParams := b.getSignedParams(params)
	_, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+SetClassCallbackUrl, signedParams)
	if err != nil {
		return err
	}
	return nil
}

//获取视频转码回调地址
func (b BjyCloud) GetClassCallbackUrl() (UrlInfo, error) {
	params := map[string]interface{}{}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetClassCallbackUrl, signedParams)
	var urlInfo UrlInfo
	if err != nil {
		return UrlInfo{}, err
	}

	err = json.Unmarshal(res, &urlInfo)
	if err != nil {
		return UrlInfo{}, err
	}
	return urlInfo, nil
}

//发送post请求
func (b BjyCloud) sendPost(bjyUrl string, params string) ([]byte, error) {
	req, _ := http.NewRequest("POST", bjyUrl, strings.NewReader(params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respData := CommonResponse{}
	err = json.Unmarshal(respByte, &respData)
	if err != nil {
		return nil, err
	}

	if respData.Code != 0 {
		return nil, errors.New(respData.Msg)
	}
	data, err := json.Marshal(respData.Data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//生成签名后参数
func (b BjyCloud) getSignedParams(params map[string]interface{}) string {
	var dataParams string
	var keys []string
	params["partner_id"] = config.ServerConfig.Baijiayun.PartnerId
	params["timestamp"] = int(time.Now().Unix())
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//拼接
	for _, k := range keys {
		pp := params[k]
		var ss string
		switch value := pp.(type) {
		case int:
			ss = strconv.Itoa(value)
			break
		case string:
			ss = value
			break
		}
		dataParams = dataParams + k + "=" + ss + "&"
	}
	md5DataParams := dataParams + "partner_key=" + config.ServerConfig.Baijiayun.PartnerKey

	//对字符串进行md5
	h := md5.New()
	h.Write([]byte(md5DataParams))
	md5Str := hex.EncodeToString(h.Sum(nil))

	return dataParams + "sign=" + md5Str
}

func (b BjyCloud) UpdateUpdateVideoStatus(ctx context.Context, id, status int, remark string) {
	_, err := store.WithContext(ctx).KcVideoUploadTask.
		UpdateOneID(id).
		SetStatus(status).
		SetRemark(remark).
		Save(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("update task err:%s", err.Error()))
	}
}

func (b BjyCloud) CreateUploadTask(ctx context.Context, videoId, attachmentId, courseId, taskType int, title string, name string) (int, error) {
	res, err := store.WithContext(ctx).KcVideoUploadTask.Create().
		SetStatus(1).
		SetVideoID(videoId).
		SetAttachmentID(attachmentId).
		SetCourseID(courseId).
		SetType(uint8(taskType)).
		SetTitle(title).
		SetVideoName(name).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return res.ID, err
}

//获取点播token
func (b BjyCloud) GetOnDemandToken(videoId int) (TokenList, error) {
	params := map[string]interface{}{
		"video_id": videoId,
	}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetOnDemandTokenUrl, signedParams)

	var token TokenList
	if err != nil {
		return TokenList{}, err
	}

	err = json.Unmarshal(res, &token)
	if err != nil {
		return TokenList{}, err
	}
	return token, nil
}

//获取回放token
func (b BjyCloud) GetPlayBackToken(roomId int) (TokenList, error) {
	params := map[string]interface{}{
		"room_id": roomId,
	}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetPlayBackUrl, signedParams)

	var token TokenList
	if err != nil {
		return TokenList{}, err
	}

	err = json.Unmarshal(res, &token)
	if err != nil {
		return TokenList{}, err
	}
	return token, nil
}

//获取回放token
func (b BjyCloud) GetStudentCode(roomId int, uid int, avatar string) (ResStudentCode, error) {
	if uid == 0 {
		uid = int(time.Now().Unix())
	}
	params := map[string]interface{}{
		"room_id":     roomId,
		"user_number": uid,
		"user_avatar": avatar,
	}
	signedParams := b.getSignedParams(params)
	res, err := b.sendPost(config.ServerConfig.Baijiayun.Domain+GetStudentCodeUrl, signedParams)

	var token ResStudentCode
	if err != nil {
		return ResStudentCode{}, err
	}

	err = json.Unmarshal(res, &token)
	if err != nil {
		return ResStudentCode{}, err
	}
	return token, nil
}
