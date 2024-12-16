package cmd

import (
	"amis-base/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmdArgs struct {
	Port string
}

func init() {
	serverCmd.Flags().StringVarP(&serverCmdArgs.Port, "port", "p", "", "指定服务端口")

	registerCmd(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		if serverCmdArgs.Port != "" {
			viper.Set("app.port", serverCmdArgs.Port)
		}

		app.Start()
	},
}
