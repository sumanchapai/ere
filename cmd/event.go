package cmd

import "fmt"

type Event struct {
	Id    string // Usually, time.UnixNano is used for Id
	Date  string
	Title string
	Knock []int
}

type KnockEvent struct {
	Event
	ComingUpIn int
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Id: %v, Date: %v, Title: %v, Knock: %v}\n", e.Id, e.Date, e.Title, e.Knock)
}

// A collection of matching events
type MatchingEvents struct {
	Today []Event
	Knock []KnockEvent
}

// Returns matching events for the given date
// Note that it only matches events on the same calendar
// This means that if you would like to get matching events from
// all calendars, you'll have to run this dates with different calendars
func CheckEventsOnDate(date Date, events []Event) MatchingEvents {
	var matches MatchingEvents
	matches.Today = make([]Event, 0)
	matches.Knock = make([]KnockEvent, 0)
	for _, event := range events {
		eventDate, err := parseDateString(event.Date)
		if err != nil {
			panic(err)
		}
		if eventDate.calendar != date.calendar {
			continue
		}
		// Check if the event date matches today
		if DateMatches(eventDate, date) {
			matches.Today = append(matches.Today, event)
			continue
		}
		// Check if the event is to be knocked
	innner:
		for _, knock := range event.Knock {
			addedDate := DateAfterAddition(date, knock)
			if DateMatches(eventDate, addedDate) {
				knockEvent := KnockEvent{Event: event, ComingUpIn: knock}
				matches.Knock = append(matches.Knock, knockEvent)
				break innner
			}
		}
	}
	return matches
}

// Preconditions:
// reference date has no wildcard in year, month or day position
func DateMatches(candidate Date, reference Date) bool {
	if reference.yearWildCard || reference.monthWildCard || reference.dayWildCard {
		panic("precondition violated date matches")
	}

	if candidate.yearWildCard || candidate.year == reference.year {
		if candidate.monthWildCard || candidate.month == reference.month {
			if candidate.dayWildCard || candidate.day == reference.day {
				return true
			}
		}
	}
	return false
}
