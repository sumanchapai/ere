package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check",
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			Check(getTodaysDateInAD())
			Check(getTodaysDateInBS())
		} else {
			// Note that we are ignoring the error returned as part of the
			// parseDateString because argument has been validated already
			// by another function
			date, _ := parseDateString(args[0])
			Check(date)
		}
	},
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("expected at most one argument for route name")
		}
		if len(args) == 1 {
			// Validate date argument
			_, err := parseDateString(args[0])
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

// Runs a check of the events from the vantage of the given date
func Check(date Date) {
	// TODO
	fmt.Println("checking for", date)
}
