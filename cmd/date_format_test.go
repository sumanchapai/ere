package cmd

import "testing"

func TestParseDate(t *testing.T) {
	type TestCase struct {
		candidate string
		expect    Date
	}
	testCases := []TestCase{
		{candidate: "2020-*-*-BS", expect: Date{
			year:          2020,
			month:         0,
			day:           0,
			yearWildCard:  false,
			monthWildCard: true,
			dayWildCard:   true,
			calendar:      BS,
		}},
	}
	for _, testCase := range testCases {
		got, err := parseDateString(testCase.candidate)
		if err != nil {
			t.Errorf("got error %v when parsing date string parseDateString(%q)", err, testCase.candidate)
		} else {
			if got != testCase.expect {
				t.Errorf("parseDateString(%q) got %v expected %v", testCase.candidate, got, testCase.candidate)
			}
		}
	}
}
