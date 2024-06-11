package view

import (
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gotth-stack/internal/core"
	"github.com/ssss-tantalum/gotth-stack/internal/model"
	"github.com/ssss-tantalum/gotth-stack/internal/utils"

	pages "github.com/ssss-tantalum/gotth-stack/templates/pages/index"
)

type IndexHandler struct {
	app *core.App
}

func NewIndexHandler(app *core.App) IndexHandler {
	return IndexHandler{
		app: app,
	}
}

func (h *IndexHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	var todos []model.Todo
	h.app.DB().NewSelect().Model(&todos).Scan(ctx)

	return utils.Render(c, pages.Index(todos))
}
