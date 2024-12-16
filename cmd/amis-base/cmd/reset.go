package cmd

import (
	"amis-base/internal/pkg/db"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	registerCmd(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "重置数据",
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range args {
			switch item {
			case "pages":
				db.ResetPages()

				fmt.Println("pages 重置成功")
				break
			}
		}
	},
}
