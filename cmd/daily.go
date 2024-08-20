package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Daily updates",
	Long:  "Print updates for today if it hasn't already been run",
	Run: func(_ *cobra.Command, _ []string) {
		if !HasRunTodaysUpdates() {
			// Run the check command with today's date
			Check(getDateRelativeToTodayInAD(0))
			// Then mark today's updates as read
			MarkTodaysUpdatesAsRead()
		}
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}

// Return the filename where the list of checked dates are kept
// Creates the file if it doesn't exist
func checkedDatesFilesName() string {
	configFolder := ereConfigFolder()
	checkedFile := filepath.Join(configFolder, ereCheckedFileName)
	_, err := os.Stat(checkedFile)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(checkedFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	return checkedFile
}

func getTodaysDateForCheckedFile() string {
	return time.Now().Format("2006-01-02")
}

// Function that checks whether today's update checking
// has been completed.
func HasRunTodaysUpdates() bool {
	todayDate := getTodaysDateForCheckedFile()
	checkedDatesFile := checkedDatesFilesName()
	// Read the last line in the file and see if the
	lastLine := getLastLineWithSeek(checkedDatesFile)
	lastLine = strings.TrimSpace(lastLine)
	return lastLine == todayDate
}

func MarkTodaysUpdatesAsRead() {
	// Do nothing it alreay marked today's update as read
	if HasRunTodaysUpdates() {
		return
	}
	todayDate := getTodaysDateForCheckedFile()
	checkedDatesFile := checkedDatesFilesName()
	f, err := os.OpenFile(checkedDatesFile, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// _, err = f.WriteString(fmt.Sprintf("%v\n", todayDate))
	_, err = fmt.Fprintf(f, "%v\n", todayDate)
	if err != nil {
		panic(err)
	}
}
