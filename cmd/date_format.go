package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/opensource-nepal/go-nepali/dateConverter"
)

type Calendar string

const (
	AD            Calendar = "AD"
	BS            Calendar = "BS"
	dateSeparator          = "-"
)

var (
	errInvalidDate           = errors.New("invalid date string")
	errInvalidCalendarOnDate = errors.New("invalid calendar on date")
)

type Date struct {
	year  int
	month int
	day   int

	yearWildCard  bool
	monthWildCard bool
	dayWildCard   bool

	calendar Calendar
}

func (d Date) String() string {
	var yearString string
	var monthString string
	var dayString string
	if d.yearWildCard {
		yearString = "*"
	} else {
		yearString = fmt.Sprintf("%v", d.year)
	}
	if d.monthWildCard {
		monthString = "*"
	} else {
		monthString = fmt.Sprintf("%02d", d.month)
	}
	if d.dayWildCard {
		dayString = "*"
	} else {
		dayString = fmt.Sprintf("%02d", d.day)
	}
	return fmt.Sprintf("%v%v%v%v%v%v%v", yearString, dateSeparator, monthString, dateSeparator, dayString, dateSeparator, d.calendar)
}

func getTodaysDateInAD() Date {
	today := time.Now()
	var toReturn Date
	toReturn.year = today.Year()
	toReturn.month = int(today.Month())
	toReturn.day = today.Day()
	toReturn.calendar = AD
	return toReturn
}

func getTodaysDateInBS() Date {
	today := time.Now()
	var toReturn Date
	nepaliDate, err := dateConverter.EnglishToNepali(today.Year(), int(today.Month()), today.Day())
	if err != nil {
		log.Fatal(err)
	}
	toReturn.year = nepaliDate[0]
	toReturn.month = nepaliDate[1]
	toReturn.day = nepaliDate[2]
	toReturn.calendar = BS
	return toReturn
}

// Parses the date string into a date object
// error returned is non-nil if the date string is not parsable
func parseDateString(date string) (Date, error) {
	// Remove whitespaces
	date = strings.TrimSpace(date)
	var toReturn Date
	// Example dates:
	// 2020-10-29-AD : one off event
	// *-10-29-AD : every year, oct 29
	// *-1-10-BS : every year, baiskah 10
	// *-*-1-BS : the first of every nepali month
	parts := strings.Split(date, dateSeparator)
	if len(parts) != 4 {
		return toReturn, errInvalidDate
	}
	// Year
	if parts[0] == "*" {
		toReturn.yearWildCard = true
	} else {
		yearInt, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			return toReturn, err
		}
		if yearInt <= 0 {
			return toReturn, err
		}
		toReturn.year = int(yearInt)
	}
	// Month
	if parts[1] == "*" {
		toReturn.monthWildCard = true
		monthInt, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			return toReturn, err
		}
		if monthInt <= 0 || monthInt > 12 {
			return toReturn, err
		}
		toReturn.month = int(monthInt)
	}
	// Day
	if parts[2] == "*" {
		toReturn.dayWildCard = true
		dayInt, err := strconv.ParseInt(parts[2], 10, 32)
		if err != nil {
			return toReturn, err
		}
		if dayInt <= 0 || dayInt > 32 {
			return toReturn, err
		}
		toReturn.day = int(dayInt)
	}
	// Calendar
	if parts[3] == string(AD) {
		toReturn.calendar = AD
	} else if parts[3] == string(BS) {
		toReturn.calendar = BS
	} else {
		return toReturn, errInvalidCalendarOnDate
	}
	return toReturn, nil
}
