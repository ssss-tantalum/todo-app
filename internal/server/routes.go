package server

import (
	"github.com/ssss-tantalum/gotth-stack/internal/core"
	"github.com/ssss-tantalum/gotth-stack/internal/handler/api"
	"github.com/ssss-tantalum/gotth-stack/internal/handler/view"
)

func initRoutes(app *core.App) {
	r := app.Router()

	// View routes
	indexHandler := view.NewIndexHandler(app)
	r.GET("/", indexHandler.Index)

	// API routes
	a := r.Group("/api")
	todoHandler := api.NewTodoHandler(app)
	a.GET("/todos", todoHandler.List)
	a.POST("/todos", todoHandler.Create)
	a.GET("/todos/:id", todoHandler.Edit)
	a.PUT("/todos/:id", todoHandler.Update)
	a.DELETE("/todos/:id", todoHandler.Delete)
}
