package cmd

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"os"
	"tkserver/internal/config"
	"tkserver/internal/store"
	"tkserver/pkg/log"
)

func MgRun() *cobra.Command {
	var configFile string
	loggerOption := log.DefaultOptions()

	runServerCmd := &cobra.Command{
		Use: "migrate",
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
			err = store.Migration(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return runServerCmd
}
