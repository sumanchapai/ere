/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive past events",
	Run: func(_ *cobra.Command, _ []string) {
		archivedEvents := eventsFromEventsFile(ereArchivedEventsFileName)
		activeEvents := eventsFromEventsFile(ereActiveEventsFileName)
		newActiveEvents := make([]Event, 0)
		archivedCount := 0
		for _, activeEvent := range activeEvents {
			// Check if the even is to be archived
			if canArchiveEvent(activeEvent) {
				archivedEvents = append(archivedEvents, activeEvent)
				archivedCount++
			} else {
				newActiveEvents = append(newActiveEvents, activeEvent)
			}
		}
		// Save the archived events file
		saveEvents(archivedEvents, ereArchivedEventsFileName)
		// Save the active events file
		saveEvents(newActiveEvents, ereActiveEventsFileName)

		switch archivedCount {
		case 0:
			fmt.Println("no events to archive")
		case 1:
			fmt.Println("archived 1 event")
		default:
			fmt.Printf("archived %v event\n", archivedCount)
		}
	},
}

// Reports whether the even can be archived
// Note that this might return some false positive in that it might
// return false for events that might be able to be archived
// More edge case test have to be written to fix that. The kind of things
// that aren't considered are that: some events targeted for end of the month
// of some month like 30, 31 might be archivable if the date on one month has
// passed and the next month isn't long enough.
func canArchiveEvent(event Event) bool {
	today := getTodaysDateInAD()
	eventDate, err := parseAbsoluteDateString(event.Date)
	if err != nil {
		panic(err)
	}
	if eventDate.calendar == BS {
		today = today.toBS()
	}
	// If year is a wildcard, the event will come up again in the future
	// thus cannot be archived
	if eventDate.yearWildCard {
		return false
	}
	if eventDate.year > today.year {
		return false
	}
	// If the year is < current year, then this event won't arrive in the future
	if eventDate.year < today.year {
		return true
	}
	// Assert: event year is same as current year
	// Check if the month is a wildcard
	if eventDate.monthWildCard {
		// If it's the final month, then we will have to check
		// further whether the event has already passed
		if today.month == 12 {
			// If the day is wildcard, it will match today/tomorrow thus cannot be archived
			if eventDate.dayWildCard {
				return false
			} else {
				// Otherwise archive only if it has already passed
				return eventDate.day < today.day
			}
		} else {
			// Otherwise, this date might come in the future months
			// Note that, here's we might return false positives, because
			// sometimes, the date might not come in the future months if the
			// future months aren't long enough. For example if someone is trying to
			// match around end of the month like: 30, 31, 32, etc.
			return false
		}
	}
	// Month has already passed, archive this event
	if eventDate.month < today.month {
		return true
	}
	// If month is yet to come, don't archive, although we might return some false positives
	if eventDate.month > today.month {
		return false
	}
	// Same month same year
	// If the day is a wildcard, don't archive because it'll match today
	if eventDate.dayWildCard {
		return false
	} else {
		// Otherwise archive only if the day has already passed
		return eventDate.day < today.day
	}
}

func init() {
	rootCmd.AddCommand(archiveCmd)
}
