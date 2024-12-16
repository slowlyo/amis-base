package cmd

import (
	"amis-base/internal/pkg/db"
	"github.com/spf13/cobra"
)

func init() {
	registerCmd(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "运行数据库迁移",
	Run: func(cmd *cobra.Command, args []string) {
		db.Migration()
	},
}
