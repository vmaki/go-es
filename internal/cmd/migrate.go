package cmd

import (
	"github.com/spf13/cobra"
	"go-es/internal/cmd/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

func init() {
	CmdMigrate.AddCommand(
		migrate.CmdMigrateUp,       // 执行迁移，生成表
		migrate.CmdMigrateRollback, // 回退上一次操作
		migrate.CmdMigrateRefresh,  // reset 后重新 migrate up
		migrate.CmdMigrateReset,    // 回溯所有迁移

		migrate.CmdMigrateFresh, // 重置数据后重新 migrate up
	)
}
