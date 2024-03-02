package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	Title string
	Knock string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an event",
	Run: func(_ *cobra.Command, args []string) {
		title, err := parseTitle(Title)
		if err != nil {
			log.Fatal(err)
		}
		knocks, err := parseKnock(Knock)
		if err != nil {
			log.Fatal(err)
		}
		date, err := parseDateString(args[0])
		if err != nil {
			log.Fatal(err)
		}
		runAddCommand(date, title, knocks)
	},
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requres date argument")
		}
		_, err := parseDateString(args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Title, "title", "t", "", "title of the event")
	addCmd.Flags().StringVarP(&Knock, "knock", "k", "", "knock")
	addCmd.MarkFlagRequired("title")
}

// Returns the title after parsing, if the parsing fails, the error returned is non-nil
func parseTitle(s string) (string, error) {
	return s, nil
}

// Parse the knock flag and return the list of integers
// as the knock value. If the parsing fails, the error returned is non-nil
func parseKnock(knock string) ([]int, error) {
	knock = strings.TrimSpace(knock)
	parts := strings.Split(knock, ",")
	knocks := make([]int, 0)
	for _, part := range parts {
		val, err := strconv.ParseInt(part, 10, 32)
		if err != nil {
			return knocks, err
		}
		knocks = append(knocks, int(val))
	}
	return knocks, nil
}

// Add the event
// Prints the error and exists if encountered any error
func runAddCommand(date Date, title string, knock []int) {
	var event Event
	event.Id = fmt.Sprint(time.Now().UnixNano())
	event.Date = date.String()
	event.Title = title
	event.Knock = knock
	// Get events list, and append to the events list
	events := eventsFromEventsFile()
	events = append(events, event)
	// Save the events
	configFolder := ereConfigFolder()
	eventsFile := filepath.Join(configFolder, ereEventsFileName)
	// Save the events
	bytes, err := json.MarshalIndent(events, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(eventsFile, bytes, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
