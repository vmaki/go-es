package migrate

import "github.com/spf13/cobra"

var CmdMigrateRollback = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"}, // 设置别名 migrate down == migrate rollback
	Short:   "Reverse the up command",
	Run:     runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}
