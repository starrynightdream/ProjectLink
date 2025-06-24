package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/starrynightdream/ProjectLink/backend/conf"
	"github.com/starrynightdream/ProjectLink/backend/mock"
	"github.com/starrynightdream/ProjectLink/backend/router"
)

const (
	DefaultConfPath = "conf/env.yaml"
)

// StartConf 启动服务参数
type StartConf struct {
	ConfigPath string // 配置文件位置
	Mock       bool   // 是否启动 Mock 服务
}

var scof StartConf

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	Long: `Start Server with some flag:
	--conf, choose the config file
	--mock, is run the mock server`,
	Run: func(cmd *cobra.Command, args []string) {
		envConf, err := conf.DecodeEnvConfigFromFile(scof.ConfigPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var service router.ProjectsService
		if envConf.ServerConf.FourceMock || scof.Mock {
			fmt.Println("Mock")
			service = mock.NewMockService()
		} else {
			fmt.Println("Not such server")
			os.Exit(1)
		}

		r := router.SetupRouter(service)
		r.Run(fmt.Sprintf("%s:%d", envConf.ServerConf.Host, envConf.ServerConf.Port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().StringVar(&scof.ConfigPath, "conf", DefaultConfPath, "conf path")
	rootCmd.PersistentFlags().BoolVar(&scof.Mock, "mock", false, "is start mock server")
}
