package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *App) initRouter() {
	app.router = echo.New()

	// Serve static files
	app.router.Static("/static", "static")

	// Middlewares
	app.router.Use(middleware.Logger())
}
