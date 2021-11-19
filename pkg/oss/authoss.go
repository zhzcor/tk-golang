package oss

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gserver/internal/config"
	"io"
	"time"
)

//const (
//	ALIYUN_ACCESS_KEY_ID     string = "LTAI18KaoTkPneVv"
//	ALIYUN_ACCESS_KEY_SECRET string = "unyKNploEVz5mdjzqr9iCcVkMZEYz3"
//	ALIYUN_RAM_ROLE          string = "acs:ram::1627635423977150:role/aliyunosstokengeneratorrole"
//	ALIYUN_ROLE_SESSION_NAME string = "external-username"
//	ALIYUN_REGION_ID         string = "cn-hangzhou"
//	ALIYUN_OSS_HOST          string = "http://ssmd-devloop.oss-cn-shenzhen.aliyuncs.com"
//	//ALIYUN_OSS_CALLBACK      string = "http://pre-api.zhntu.com/api/v1/oss/callback/user_id/1"
//	ALIYUN_OSS_CALLBACK      string = "http://pre-study-api.zhntu.com/admin/other/oss_callback"
//	ALIYUN_OSS_ENDPOINT      string = "ssmd-devloop"
//	ALIYUN_OSS_ENDPOINT_HOST string = "oss-cn-shenzhen.aliyuncs.com"
//)

func GetRole() (*sts.Credentials, error) {
	aliyunConfig := config.ServerConfig.Aliyun
	client, err := sts.NewClientWithAccessKey(aliyunConfig.RegionId, aliyunConfig.AccessKeyId, aliyunConfig.AccessKeySecret)
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = aliyunConfig.RamRole
	request.RoleSessionName = aliyunConfig.RoleSessionName
	request.Policy = "{\n  \"Statement\": [\n    {\n      \"Action\": [\n        \"oss:*\"\n      ],\n      \"Effect\": \"Allow\",\n      \"Resource\": [\"acs:oss:*:*:*\"]\n    }\n  ],\n  \"Version\": \"1\"\n}"

	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}
	return &response.Credentials, nil
}

func OssCennect() (Bucket *oss.Bucket, err error) {
	aliyunConfig := config.ServerConfig.Aliyun
	client, err := oss.New(aliyunConfig.OssEndpointHost, aliyunConfig.AccessKeyId, aliyunConfig.AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	Bucket, err = client.Bucket(aliyunConfig.OssEndpoint)
	return

}

func LocalUrl(file io.Reader, folderName string) (url string, err error) {
	Bucket, err := OssCennect()
	if err != nil {
		return "", err
	}

	t := time.Now()
	folder := t.Format("2006-01-02")
	//fmt.Println(t.Year())
	//拼接文件名称
	fileName := fmt.Sprintf("/qs_%d", t.UnixNano())
	//fmt.Println("...s......................")
	//fmt.Println(fileName)
	str := folderName + "/" + folder + fileName + ".jpg"

	err = Bucket.PutObject(str, file)
	if err != nil {
		return "上传错误", err
	} else {
		url = fmt.Sprintf("%s%s%s", config.ServerConfig.Aliyun.OssHost, "/", str)
		return url, nil
	}
}

func Base64LoadUrl(base64String, folderName string) (string, error) {
	ddd, err := base64.StdEncoding.DecodeString(base64String) //成图片文件并把文件写入到buffer
	if err != nil {
		return "", err
	}
	bbb := bytes.NewBuffer(ddd) // 必须加一个buffer 不然没有read方法就会报错

	url, err := LocalUrl(bbb, folderName)

	return url, err
}
