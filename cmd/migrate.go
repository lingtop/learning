package cmd

import (
	"template/service/database/postgresql/migration"
	"template/service/logger"

	"github.com/spf13/cobra"
)

var forceMigrate bool = false

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Onboard Database",
	Run: func(cmd *cobra.Command, args []string) {
		// Init Logger
		logger.InitLogger()

		initTimezone()

		migration.Migrate(false, -1, forceMigrate)

		logger.SyncLogger()
	},
}

func init() {
	rootCmd.AddCommand(MigrateCmd)
	MigrateCmd.PersistentFlags().BoolVar(&forceMigrate, "force", false, "force migrate (default is false)")
}
