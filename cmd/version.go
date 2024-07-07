package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed .version
var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print ere version",
	Long:  "Print ere version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("%v\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
