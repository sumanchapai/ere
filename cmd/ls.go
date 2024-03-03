package cmd

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all events",
	Run: func(_ *cobra.Command, _ []string) {
		allEvents := eventsFromEventsFile()
		lsTable(allEvents)
	},
}

// Colors
var (
	headerFmt          = color.New(color.FgGreen, color.Underline).SprintfFunc()
	yellow             = color.New(color.FgYellow).SprintfFunc()
	yellowAndUnderline = color.New(color.FgYellow, color.Underline).SprintfFunc()
	green              = color.New(color.FgGreen).SprintfFunc()
)

func lsTable(events []Event) {
	tbl := table.New("ID", "Title", "Date")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(yellow)
	for _, e := range events {
		tbl.AddRow(e.Id, e.Title, e.Date)
	}
	tbl.Print()
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
