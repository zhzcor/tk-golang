package baijiayun

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestGetSign(t *testing.T) {
	params := map[string]interface{}{
		"name":             "第一章",
		"timestamp":        strconv.Itoa(int(time.Now().Unix())),
		"question_bank_id": strconv.Itoa(1),
	}
	b := BjyCloud{}
	ss := b.getSignedParams(params)
	fmt.Println(ss)
}

func TestCreateRoom(t *testing.T) {
	title := "for testing"
	m, _ := time.ParseDuration("1h")
	e, _ := time.ParseDuration("3600s")
	startAt := int(time.Now().Add(m).Unix())
	endAt := int(time.Now().Add(e).Unix())
	b := BjyCloud{}
	resp, err := b.CreateLiveRoom(title, startAt, endAt, 2, 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestBjyCloud_DeleteLiveRoom(t *testing.T) {
	roomId := 21060580061210
	b := BjyCloud{}
	err := b.DeleteLiveRoom(roomId)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")

}

func TestBjyCloud_GetLiveRoomInfo(t *testing.T) {
	roomId := 21070248266200
	b := BjyCloud{}
	resp, err := b.GetLiveRoomInfo(roomId)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.RoomId, resp.StartTime, resp.EndTime, resp.Type, resp.Title)

}

func TestBjyCloud_GetOrderUploadUrl(t *testing.T) {
	filename := "testingFile.avi"
	b := BjyCloud{}
	resp, err := b.GetOrderUploadUrl(filename, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.UploadUrl, resp.VideoId)
}

//func TestBjyCloud_OrderUpload(t *testing.T) {
//	filename := "http://ssmd-release.oss-cn-shenzhen.aliyuncs.com/ipad_media/2019516/5841557994900374@2x.mp4"
//	b := BjyCloud{}
//	b.OrderUpload(filename,1,
//		"http://latest-upload-video.baijiayun.com/upload?fid=151056010&ts=1623032428143&token=066693aac8292ae1014d74be6f2b73ff&st=jd&tbIdx=1",
//		"ffffhhh")
//
//	fmt.Println("ok")
//}

func TestBjyCloud_ChangeStatus(t *testing.T) {
	videoId := 151056010
	b := BjyCloud{}
	resp, err := b.GetChangeCodeStatus(videoId)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Info, resp.Status)
}

func TestGetTranscodeCallbackUrl(t *testing.T) {
	b := BjyCloud{}
	resp, err := b.GetTranscodeCallbackUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestStartClass(t *testing.T) {
	b := BjyCloud{}
	roomId := 21062836005090
	err := b.SetStartClass(roomId)
	if err != nil {
		panic(err)
	}
	fmt.Println("0kkk")
}

func TestStopClass(t *testing.T) {
	b := BjyCloud{}
	roomId := 21062836005090
	err := b.StopClass(roomId)
	if err != nil {
		panic(err)
	}
	fmt.Println("0kkk1111")
}
