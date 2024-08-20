/*
Copyright © 2024 Suman Chapai <sumanchapai@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// laCmd represents the la command
var laCmd = &cobra.Command{
	Use:   "la",
	Short: "look ahead <n> days",
	Run: func(cmd *cobra.Command, args []string) {
		numberOfDaysToLookAhead, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("number of days to look ahead must be a positive integer")
			log.Fatal(err)
		}
		if numberOfDaysToLookAhead < 0 {
			fmt.Println("number of days to look ahead must be a positive integer")
			log.Fatal(err)
		}

		knockEvents := make(KnockEvents, 0)
		for i := 0; i < numberOfDaysToLookAhead; i++ {
			dateInAD := getDateRelativeToTodayInAD(24 * time.Hour * time.Duration(i))
			// Dates for all calendar
			dates := datesForAllCalendar(dateInAD)
			events := eventsFromEventsFile(ereActiveEventsFileName)
			for _, date := range dates {
				matches := CheckEventsOnDate(date, events)

				// Each event is a knock when using the lookahead command
				// Note here that we're only looking for events that are coming
				// up within the specified number of days not their knocks
				for _, event := range matches.Today {
					knockEvents = append(knockEvents, KnockEvent{event, i})
				}

			}
		}
		sort.Sort(knockEvents)
		PrintLookAheadEvents(knockEvents)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(laCmd)
}
