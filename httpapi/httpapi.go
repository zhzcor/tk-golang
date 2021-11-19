package httpapi

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"gserver/internal/store"
	"gserver/pkg/log"
	"time"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Run(ctx context.Context, addr string) error {
	srv := gin.New()
	zhTrans := zh.New()
	uni = ut.New(zhTrans, zhTrans)
	trans, _ = uni.GetTranslator("zh")
	fmt.Println(trans)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)

	}

	srv.Use(func(c *gin.Context) {
		// 注入框架依赖
		// 注入数据库等依赖
		// 生成一个新的ctx
		c.Set(store.CtxDatabaseKey, store.WithContext(ctx))
	})

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	srv.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		log.Infof("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
		return ""
	}))
	srv.Use(gin.Recovery())
	// 默认的CORS
	config := cors.DefaultConfig()
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	config.AddAllowHeaders("Authorization", "X-Requested-With", "Accept")
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	srv.Use(cors.New(config))
	// 载入路由以及中间件
	Use(srv)
	UseStatic(srv)
	return srv.Run(addr)
}
