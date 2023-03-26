package migrate

import (
	"github.com/spf13/cobra"
	"go-es/database/migrations"
	"go-es/internal/pkg/migrate"
)

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run migrated",
	Run:   runUp,
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
