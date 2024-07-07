package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/spf13/cobra"
)

var searchInArchived bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search an event by title (regex is valid)",
	Long: `Search an even by titlte (regex is allowed)

To search events with the word birthday in their title:
ere search "birthday"

To search events whose title begins with the world "deadline" and contains the word "urgent":
ere search "^deadline.*urgent"
`,
	Run: func(_ *cobra.Command, args []string) {
		searchRegex, err := regexp.Compile(fmt.Sprintf("(?i)%v", args[0]))
		if err != nil {
			fmt.Println("error in regular expression")
			log.Fatal(err)
		}
		var fileName string
		if searchInArchived {
			fileName = ereArchivedEventsFileName
		} else {
			fileName = ereActiveEventsFileName
		}
		allEvents := eventsFromEventsFile(fileName)

		filteredEvents := make([]Event, 0)
		for _, e := range allEvents {
			if searchRegex.Match([]byte(e.Title)) {
				filteredEvents = append(filteredEvents, e)
			}
		}
		lsTable(filteredEvents)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().BoolVar(&searchInArchived, "archive", false, "search through archived events")
}
