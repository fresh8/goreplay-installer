package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is terraplate's root command.
// Every other command attached to RootCmd is a child command to it.
var RootCmd = &cobra.Command{
	Use:   "goreplay-install",
	Short: "goreplay-install tool for Fresh8 Gaming.",
	Long:  `goreplay-install tool for Fresh8 Gaming.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionCmd.Run(cmd, args)
	},
}
