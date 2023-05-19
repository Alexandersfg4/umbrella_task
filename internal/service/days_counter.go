package service

import "time"

func GetDaysLeftUntil2025() int {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	return GetDaysLeftUntilDate(date)
}

// Get days form today to date, but not including the end date.
func GetDaysLeftUntilDate(date time.Time) int {
	now := time.Now().UTC()
	nowWithoutHours := now.Truncate(time.Hour * 24)
	duration := date.Sub(nowWithoutHours)
	return int(duration.Hours() / 24)
}
