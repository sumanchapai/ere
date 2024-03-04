package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
)

var errNoWildCardSupportedOnCheckCommandDateFormat = errors.New("wildcard isn't supported in the date format of the check command")

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check your events",
	Long: `Check your events 

Check command checks for your events and reminders on a given date. You could pass the date
in any calendar system and it displays events on any calendar systems irrespective of the date

You could search for events by AD (Gregorian) calendar:
ere check 2022-10-29-AD 

OR by BS (Nepali) calendar: 
ere check 2080-7-12-BS.

Instead of typing the date manually, to see the events for the day after tomorrow, run:
ere check +2.

To see events for yesterday, run:
ere check -1.`,

	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			Check(getTodaysDateInAD())
		} else {
			// Note that we are ignoring the error returned as part of the
			// parseDateString because argument has been validated already
			// by another function
			date, _ := parseRelativeOrAbsoluteDate(args[0])
			Check(date)
		}
	},
	// Note that there cannot be any wildcard present in the date format of the
	// check command. This is because only then will we be able to show the events
	// from all calendars (both BS and AD for example.)
	// For example, if we allowed ere check *-7-12-BS, how do we convert that into
	// AD to be able to check whether any events from the AD calendar to display?
	// More something like the ls command seems appropriate for listing the events
	// as they are of different calendars where we could allow *-7-12.
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("expected at most one argument for route name")
		}
		if len(args) == 1 {
			// Validate date argument
			d, err := parseRelativeOrAbsoluteDate(args[0])
			if d.yearWildCard || d.monthWildCard || d.dayWildCard {
				return errNoWildCardSupportedOnCheckCommandDateFormat
			}
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

// Runs a check of the events from the vantage of the given date
func Check(d Date) {
	dates := datesForAllCalendar(d)
	// Todo get events from JSON file
	events := eventsFromEventsFile()
	todayEvents := make([]Event, 0)
	knockEvents := make(KnockEvents, 0)
	for _, date := range dates {
		matches := CheckEventsOnDate(date, events)
		todayEvents = append(todayEvents, matches.Today...)
		knockEvents = append(knockEvents, matches.Knock...)
	}
	fmt.Print(yellowAndUnderline("%v --- %v\n", d.toAD(), d.toBS()))
	fmt.Printf("\n")
	PrintEvents(todayEvents)
	fmt.Printf("\n")
	sort.Sort(knockEvents)
	PrintKnockEvents(knockEvents)
}

func PrintEvents(events []Event) {
	fmt.Print(green("Events:\n"))
	if len(events) == 0 {
		fmt.Println("no events")
	}
	for _, e := range events {
		fmt.Printf("%v\n", e.Title)
	}
}

func PrintKnockEvents(events []KnockEvent) {
	fmt.Print(yellow("Knocks:\n"))
	if len(events) == 0 {
		fmt.Println("nothing knocking")
		return
	}
	for _, e := range events {
		if e.ComingUpIn == 1 {
			fmt.Printf("in 1 day - %v\n", e.Title)
		} else {
			fmt.Printf("in %v days - %v\n", e.ComingUpIn, e.Title)
		}
	}
}

// Get the list of events from events file
// Creates the file it it doesn't exist and return empty list
func eventsFromEventsFile() []Event {
	configFolder := ereConfigFolder()
	eventsFile := filepath.Join(configFolder, ereEventsFileName)
	_, err := os.Stat(eventsFile)
	events := make([]Event, 0)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(eventsFile)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(eventsFile)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &events)
	if err != nil {
		log.Fatal(err)
	}
	return events
}
