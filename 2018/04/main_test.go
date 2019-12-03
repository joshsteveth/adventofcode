package main

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	format := "01.02.2006 15-04"

	tt := map[string]struct {
		input          string
		options        []func(time.Time) time.Time
		expectedResult string
	}{
		"without-options": {"[1518-11-01 22:30]", nil, "11.01.1518 22-30"},
		"with-adjustTime": {"[1518-11-01 22:30]", []func(time.Time) time.Time{adjustTime}, "11.02.1518 00-00"},
	}

	for testname, tc := range tt {
		t.Run(testname, func(t *testing.T) {
			var a activity
			if err := a.parseTime(tc.input, tc.options...); err != nil {
				t.Errorf("failed to parse time: %v", err)
				return
			}

			result := a.timestamp.Format(format)
			if result != tc.expectedResult {
				t.Errorf("unexpected result. wanted: %s got %s", tc.expectedResult, result)
			}
		})
	}
}

func TestParseActivityType(t *testing.T) {
	tt := map[string]struct {
		str            string
		isError        bool
		expectedResult activity
	}{
		"guard-correct": {"[1518-11-01 00:00] Guard #10 begins shift", false,
			activity{typ: beginShift, guardID: 10}},
		"guard-false":    {"[1518-11-01 00:00] Guard #ab begins shift", true, activity{}},
		"falls":          {"[1518-11-01 00:05] falls asleep", false, activity{typ: fallsAsleep}},
		"wakes":          {"[1518-11-01 00:25] wakes up", false, activity{typ: wakesUp}},
		"false-category": {"[1518-11-01 00:05] foo bar", true, activity{}},
	}

	for testname, tc := range tt {
		t.Run(testname, func(t *testing.T) {
			var a activity
			err := a.parseActivityType(tc.str)
			if tc.isError {
				if err == nil {
					t.Errorf("expected an error, received nil error")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if a != tc.expectedResult {
				t.Errorf("unexpected result. wanted: %+v; got: %+v", tc.expectedResult, a)
			}

		})
	}
}

func TestCalculateActivities(t *testing.T) {
	str := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
	}

	var as []activity

	for _, s := range str {
		var a activity
		err := a.parseString(s, adjustTime)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		as = append(as, a)
	}

	guardID, sleepTime, listMinute, err := calculateActivities(as)
	if err != nil {
		t.Errorf("unexpected error while calculating activities: %v", err)
	}
	if guardID != 10 {
		t.Errorf("false guard ID. wanted 10, got %v", guardID)
	}
	if sleepTime != 45 {
		t.Errorf("false sleep time, wanted 45, got %v", sleepTime)
	}

	expectedListMinute := []int{}
	for i := 5; i < 25; i++ {
		expectedListMinute = append(expectedListMinute, i)
	}
	for i := 30; i < 55; i++ {
		expectedListMinute = append(expectedListMinute, i)
	}
	if len(listMinute) != len(expectedListMinute) {
		t.Errorf("false list minute, wanted %v, got %v", expectedListMinute, listMinute)
	}
}
