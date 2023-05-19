package service

import (
	"testing"
	"time"
)

func TestGetDaysLeftUntilDate(t *testing.T) {
	cases := []struct {
		Duration     time.Duration
		ExpectedDays int
	}{
		{Duration: time.Duration(time.Hour * 24), ExpectedDays: 1},
		{Duration: time.Duration(time.Hour * 1), ExpectedDays: 0},
		{Duration: time.Duration(time.Hour * 14592), ExpectedDays: 608},
	}

	for _, test := range cases {
		today := time.Now()
		nextDay := today.Add(test.Duration)

		gotDays := GetDaysLeftUntilDate(nextDay)

		if gotDays != test.ExpectedDays {
			t.Errorf("Got days: %d and expected: %d days are not matched", gotDays, test.ExpectedDays)
		}

	}
}
