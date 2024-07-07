package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete event by given id",
	Run: func(_ *cobra.Command, args []string) {
		id := args[0]
		events := eventsFromEventsFile(ereActiveEventsFileName)
		newEvents := make([]Event, 0)
		found := false
		for _, event := range events {
			if event.Id == id {
				found = true
			} else {
				newEvents = append(newEvents, event)
			}
		}
		if !found {
			fmt.Println("no event found with id", id)
		} else {
			// Save new events
			saveEvents(newEvents, ereActiveEventsFileName)
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
