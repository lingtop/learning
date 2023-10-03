package cmd

import (
	"template/service/logger"

	"github.com/spf13/cobra"
)

var OnboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Start Onboard Server",
	Run: func(cmd *cobra.Command, args []string) {

		// Init Logger
		logger.InitLogger()

		// Init Listen OS Signal
		initListenOsSignal()

		// Init Database
		initPostgresql()

		// Init Redis
		initRedis()

		// Init Interface
		initListenInterface()

		logger.Logger.Info("Onboard service is running")

		wg.Wait()

		// Flush Log
		logger.SyncLogger()
	},
}

func init() {
	rootCmd.AddCommand(OnboardCmd)
}
