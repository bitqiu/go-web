// Package cmd 命令行模式
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"go-web/pkg/helpers"
)

var rootCmd = &cobra.Command{
	Use:               "go-web",
	Short:             "go-web",
	SilenceUsage:      true,
	Long:              `go-web`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	//注册命令
	rootCmd.AddCommand(
		CmdServe,
	)
	// 配置默认运行 Web 服务
	registerDefaultCmd(rootCmd, CmdServe)

}

// registerGlobalFlags 注册全局选项（flag）
func registerGlobalFlags(rootCmd *cobra.Command) {
	// TODO 配置文件选项 --config /path/config.yaml or -c /path/config.yaml
}

// registerDefaultCmd 注册默认命令
func registerDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firtArg := helpers.FirstElement(os.Args[1:])

	if err == nil && cmd.Use == rootCmd.Use && firtArg != "-h" && firtArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

//Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
