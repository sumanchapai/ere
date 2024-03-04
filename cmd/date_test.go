package cmd

import "testing"

func TestDateAfterAddition(t *testing.T) {
	type TestCase struct {
		candidate string
		addDays   int
		expect    string
	}
	testCases := []TestCase{
		{candidate: "2080-1-12-BS", addDays: 5, expect: "2080-01-17-BS"},
		{candidate: "2080-10-29-BS", addDays: 5, expect: "2080-11-05-BS"},
		{candidate: "2024-2-28-AD", addDays: 3, expect: "2024-03-02-AD"},
	}
	for _, testCase := range testCases {
		candidateDate, err := parseAbsoluteDateString(testCase.candidate)
		if err != nil {
			panic(err)
		}
		gotDate := DateAfterAddition(candidateDate, testCase.addDays)
		if gotDate.String() != testCase.expect {
			t.Errorf("DateAfterAddition(%v, %v) got %v expected %v",
				testCase.candidate,
				testCase.addDays,
				gotDate.String(),
				testCase.expect)
		}
	}
}
