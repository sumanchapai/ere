package cmd

import "time"

// Add the given number of days and return the Date
//
// Preconditions:
// date contains no wildcard in year, month, or day positions
func DateAfterAddition(date Date, noOfDays int) Date {
	// Precondition check
	if date.yearWildCard || date.monthWildCard || date.dayWildCard {
		panic("precondition violated in date addition")
	}
	calendar := date.calendar
	// Convert date to AD for easier date manipulation
	date = date.toAD()

	// Add the date
	t := time.Date(date.year, time.Month(date.month), date.day, 0, 0, 0, 0, time.Local)
	durationToAdd := time.Duration(noOfDays * int(time.Hour) * 24)
	t = t.Add(durationToAdd)
	date.year = t.Year()
	date.month = int(t.Month())
	date.day = t.Day()

	// Change the date to the original calendar system
	if calendar == BS {
		return date.toBS()
	} else {
		return date.toAD()
	}
}
