/*
Copyright © 2022 Pereg Hergoualc'h <pereg.hergoualch@revolve.team>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use: "deactivate",

	Short: "Restart the scaling down of the runners",

	Long: "This CLI is used to restart the scaling down of the GitHub Actions runners",

	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
	deactivateCmd.Flags().StringP("accounts", "a", "", "accounts numbers separated by a comma")
}