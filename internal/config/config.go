package config

import (
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"gserver/pkg/log"
	"gserver/pkg/version"
)

// config like bash> export gserver_ServerPort=8080;gserver_Database_Driver=sqlite3;gserver_Database_DSN=file:ent?mode=memory&cache=shared&_fk=1
// logger env bash> export gserver_LoggerLevel=DEBUG
// AllowOrigins like bash> export gserver_AllowOrigins=[http://127.0.0.1:8000,http://localhost:8000]
// ApiUrl required like bash> export gserver_ApiUrl=http://127.0.0.1:8000
type serverConfig struct {
	Pprof    bool `json:"pprof" `
	Database struct {
		Driver string `json:"driver" required:"true"`
		DSN    string `json:"dsn" required:"true"`
	}
	LoggerLevel    string
	StaticDir      string `json:"static_dir" required:"true"`       // 静态文件存放地址
	QueueRedisAddr string `json:"queue_redis_addr" required:"true"` // 队列的redis地址
	BossHost       string `json:"boss_host" required:"true"`        // 工作台接口地址
	Aliyun         struct {
		AccessKeyId     string `json:"access_key_id" required:"true"`
		AccessKeySecret string `json:"access_key_secret" required:"true"`
		RamRole         string `json:"ram_role" required:"true"`
		RoleSessionName string `json:"role_session_name" required:"true"`
		RegionId        string `json:"region_id" required:"true"`
		OssHost         string `json:"oss_host" required:"true"`
		OssCallback     string `json:"oss_callback" required:"true"`
		OssEndpoint     string `json:"oss_endpoint" required:"true"`
		OssEndpointHost string `json:"oss_endpoint_host" required:"true"`
	}
	Dingding struct {
		DingAccessKey string `json:"ding_access_key" required:"true"`
		DingAppSecret string `json:"ding_app_secret" required:"true"`
		DingCorSecret string `json:"ding_cor_secret" required:"true"`
		DingCorAppId  string `json:"ding_cor_app_id" required:"true"`
	}
	Baijiayun struct {
		PartnerId         string `json:"partner_id" required:"true"`
		PartnerKey        string `json:"partner_key" required:"true"`
		Domain            string `json:"domain" required:"true"`
		NoHttpDomain      string `json:"no_http_domain" required:"true"`
		VideoTranscodeUrl string `json:"video_transcode_url" required:"true"`
		ClassCallbackUrl  string `json:"class_callback_url" required:"true"`
	}
}

var ServerConfig serverConfig

func LoadServerConfig(file ...string) error {
	if err := configor.New(&configor.Config{
		ENVPrefix: version.AppName,
		Verbose:   true,
	}).Load(&ServerConfig, file...); err != nil {
		return err
	} else {
		log.Info("config loaded", zap.Any("ServerConfig", ServerConfig))
	}
	return nil
}
