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

func lsTable(events []Event) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Title", "Date")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, e := range events {
		tbl.AddRow(e.Id, e.Title, e.Date)
	}
	tbl.Print()
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
