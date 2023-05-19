package handler

import (
	"net/http"
	"umbrella_task/internal/service"

	"github.com/labstack/echo/v4"
)

func GetDaysLeftHander(c echo.Context) error {
	daysLeft := service.GetDaysLeftUntil2025()
	answer := service.MakeAmountOfDaysAnswer(daysLeft)
	return c.String(http.StatusOK, answer)
}
