package utils

import (
	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, component)
}
