package service

import "testing"

func TestMakeAmountOfDaysAnswer(t *testing.T) {
	daysLeft := 30
	got := MakeAmountOfDaysAnswer(daysLeft)
	expect := "30 days left until 1 Jan 2025"

	if got != expect {
		t.Errorf("Incorrect answer\ngot: %s\nexpect: %s", got, expect)
	}
}
