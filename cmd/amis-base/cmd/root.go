package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "amisBase",
	Short: "使用 GoFiber 和 amis 构建你的数据面板",
}

func registerCmd(command *cobra.Command) {
	rootCmd.AddCommand(command)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
