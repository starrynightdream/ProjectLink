package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "1.0"

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Use:   "plink",
	Short: "项目介绍系统的后端",
	Long: `项目介绍系统 ProjectLink 的后端实现，
	内部包括一个服务器以及一些后端对数据库的操作命令`,
}

// Execute 初始化命令行工具
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
