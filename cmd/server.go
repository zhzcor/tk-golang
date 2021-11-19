package cmd

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gserver/httpapi"
	_ "gserver/internal/app"
	"gserver/internal/config"
	"gserver/internal/store"
	"gserver/pkg/asynctask"
	"gserver/pkg/log"
	"os"
)

func ServerRun() *cobra.Command {
	var addr string
	var configFile string
	loggerOption := log.DefaultOptions()

	runServerCmd := &cobra.Command{
		Use: "server",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("on envs", os.Environ())

			// 加载配置
			err := config.LoadServerConfig(configFile)
			if err != nil {
				return err
			}
			ctx := context.Background()
			err = log.Configure(loggerOption)
			if err != nil {
				return err
			}
			// 初始化db
			ctx, err = store.NewContext(ctx, config.ServerConfig.Database.Driver, config.ServerConfig.Database.DSN)
			if err != nil {
				return err
			}

			//初始化异步队列
			err = asynctask.Background(ctx, config.ServerConfig.QueueRedisAddr)
			if err != nil {
				log.Fatal("async task server run error:", err)
			}
			defer asynctask.Stop()

			return httpapi.Run(ctx, addr)
		},
	}
	runServerCmd.Flags().StringVarP(&addr, "listen", "l", ":8080", "gpush server listen addr")
	runServerCmd.Flags().StringVarP(&configFile, "cfile", "f", "", "gpush load server config file ")
	loggerOption.AttachCobraFlags(runServerCmd)

	return runServerCmd
}
