package cmd

import (
	"blobdev.com/pandaroll/internal/logger"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Not implemented yet :(")
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
