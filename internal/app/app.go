package app

import (
	"umbrella_task/internal/handler"
	"umbrella_task/internal/middleware"

	"github.com/labstack/echo/v4"
)

type App struct {
	e *echo.Echo
}

func NewApp() *App {
	return &App{echo.New()}
}

func (app *App) Start() {
	app.e.Use(middleware.CheckUserRole)
	app.e.GET("/status", handler.GetDaysLeftHander)
	app.e.Logger.Fatal(app.e.Start(":1223"))
}
