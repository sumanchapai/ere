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
	AD Calendar = "AD"
	BS Calendar = "BS"
	// dateSeparator          = "-"
)

var (
	errInvalidRelativeDate   = errors.New("invalid relative date string, has to be like: -23 or +34 or 5")
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
	return fmt.Sprintf("%v-%v-%v-%v", yearString, monthString, dayString, d.calendar)
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

// This function takes as input relative or absolute date,
// Relative date means date in relation to today, and it is written
// as +1 or just 1 to refer to tomorrow, -2 to refer to the day before
// yesterday. Absolute date is in the Date format which can be either
// in AD (Gregorian) or the BS (Nepali) calendar.
func parseRelativeOrAbsoluteDate(date string) (Date, error) {
	date = strings.TrimSpace(date)
	if date == "" {
		return *new(Date), errInvalidDate
	}
	// Check if date is like "55" or "2"
	i, err := strconv.ParseInt(date, 10, 64)
	if err == nil {
		thatDay := time.Now().Add(time.Duration(i * 24 * int64(time.Hour)))
		return Date{
			year:          thatDay.Year(),
			month:         int(thatDay.Month()),
			day:           thatDay.Day(),
			yearWildCard:  false,
			monthWildCard: false,
			dayWildCard:   false,
			calendar:      AD,
		}, nil
	}
	// Check if date is like "+1" or "-3"
	toMultiply := 1
	isRelativeDate := false
	// At this point, we can know that the date string is >= 1 in length, so we can
	// safely check the index 0
	switch date[0] {
	case '+':
		isRelativeDate = true
	case '-':
		isRelativeDate = true
		toMultiply = -1
	}
	if isRelativeDate {
		num, err := strconv.ParseInt(date[1:], 10, 64)
		if err != nil {
			return *new(Date), errInvalidRelativeDate
		}
		num = num * int64(toMultiply)
		thatDay := time.Now().Add(time.Duration(num * 24 * int64(time.Hour)))
		return Date{
			year:          thatDay.Year(),
			month:         int(thatDay.Month()),
			day:           thatDay.Day(),
			yearWildCard:  false,
			monthWildCard: false,
			dayWildCard:   false,
			calendar:      AD,
		}, nil
	}
	// At this point, date string is absolute, so we leave it upto the
	// parseDateString function to handle that
	return parseAbsoluteDateString(date)
}

// Parses the date string in format like YYYY-MM-DD-CALENDAR
// (for example: 2010-10-12-AD) into a date object
// error returned is non-nil if the date string is not parsable
func parseAbsoluteDateString(date string) (Date, error) {
	// Remove whitespaces
	date = strings.TrimSpace(date)
	var toReturn Date
	// Example dates:
	// 2020-10-29-AD : one off event
	// *-10-29-AD : every year, oct 29
	// *-1-10-BS : every year, baiskah 10
	// *-*-1-BS : the first of every nepali month
	parts := strings.Split(date, "-")
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
	} else {
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
	} else {
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

// Return date for all calendars from a given date of a certain calendar
func datesForAllCalendar(date Date) map[Calendar]Date {
	datesByCalendar := make(map[Calendar]Date)
	datesByCalendar[AD] = date.toAD()
	datesByCalendar[BS] = date.toBS()
	return datesByCalendar
}

// Returns the BS date for the given date
// Preconditions:
// 1. No wildcard can be present in year, month, or day
func (date Date) toBS() Date {
	if date.calendar == BS {
		return date
	}
	var bsDate Date
	bsDate.calendar = BS
	nepaliDate, err := dateConverter.EnglishToNepali(date.year, date.month, date.day)
	if err != nil {
		log.Fatal(err)
	}
	bsDate.year = nepaliDate[0]
	bsDate.month = nepaliDate[1]
	bsDate.day = nepaliDate[2]
	return bsDate
}

// Returns the AD date for the given date
// Preconditions:
// 1. No wildcard can be present in year, month, or day
func (date Date) toAD() Date {
	if date.calendar == AD {
		return date
	}
	var adDate Date
	adDate.calendar = AD
	englishDate, err := dateConverter.NepaliToEnglish(date.year, date.month, date.day)
	if err != nil {
		log.Fatal(err)
	}
	adDate.year = englishDate[0]
	adDate.month = englishDate[1]
	adDate.day = englishDate[2]
	return adDate
}
