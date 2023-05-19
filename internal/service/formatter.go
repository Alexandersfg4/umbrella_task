package service

import "fmt"

func MakeAmountOfDaysAnswer(daysLeft int) string {
	return fmt.Sprintf("%d days left until 1 Jan 2025", daysLeft)
}
