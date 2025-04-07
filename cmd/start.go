package cmd

import (
	"compro/internal/app"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run: func(cmd *cobra.Command, args []string) {
		// call function route API
		app.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(StartCmd)
}
