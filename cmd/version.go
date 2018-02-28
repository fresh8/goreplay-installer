package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.0.1"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of goreplay-installer",
	Long:  `All software has versions. This is goreplay-installer's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("goreplay-installer %s", version))
	},
}
