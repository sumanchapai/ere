package cmd

import "testing"

func TestDateAfterSubtraction(t *testing.T) {
	type TestCase struct {
		candidate    string
		subtractDays int
		expect       string
	}
	testCases := []TestCase{
		{candidate: "2080-1-12-BS", subtractDays: 5, expect: "2080-01-07-BS"},
		{candidate: "2080-11-2-BS", subtractDays: 5, expect: "2080-10-26-BS"},
		{candidate: "2024-3-2-AD", subtractDays: 3, expect: "2024-02-28-AD"},
	}
	for _, testCase := range testCases {
		candidateDate, err := parseDateString(testCase.candidate)
		if err != nil {
			panic(err)
		}
		gotDate := DateAfterSubtraction(candidateDate, testCase.subtractDays)
		if gotDate.String() != testCase.expect {
			t.Errorf("DateAfterSubtraction(%v, %v) got %v expected %v",
				testCase.candidate,
				testCase.subtractDays,
				gotDate.String(),
				testCase.expect)
		}
	}
}
