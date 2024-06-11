package api

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gotth-stack/internal/core"
	"github.com/ssss-tantalum/gotth-stack/internal/model"
	"github.com/ssss-tantalum/gotth-stack/internal/utils"
	"github.com/ssss-tantalum/gotth-stack/templates/components"
)

type TodoHandler struct {
	app *core.App
}

func NewTodoHandler(app *core.App) TodoHandler {
	return TodoHandler{
		app: app,
	}
}

func (h TodoHandler) List(c echo.Context) error {
	ctx := c.Request().Context()

	var todos []model.Todo
	h.app.DB().NewSelect().Model(&todos).Scan(ctx)

	return utils.Render(c, components.Todos(todos))
}

func (h TodoHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	var todo model.Todo
	h.app.DB().NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx)

	return utils.Render(c, components.Todo(todo))
}

func (h TodoHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	content := c.FormValue("content")
	todo := model.Todo{
		ID:      uuid.NewString(),
		Content: content,
	}

	_, err := h.app.DB().NewInsert().Model(&todo).Exec(ctx)
	if err != nil {
		return err
	}

	return h.List(c)
}

func (h TodoHandler) Edit(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	var todo model.Todo
	h.app.DB().NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx)

	return utils.Render(c, components.EditTodo(todo))
}

func (h TodoHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	content := c.FormValue("content")

	todo := model.Todo{
		ID:      id,
		Content: content,
	}
	h.app.DB().NewUpdate().Model(&todo).WherePK().Exec(ctx)

	return h.List(c)
}

func (h TodoHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	var todo model.Todo
	_, err := h.app.DB().NewDelete().Model(&todo).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return h.List(c)
}
