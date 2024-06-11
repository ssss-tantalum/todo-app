package utils

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, content string) error {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	c.Response().Write([]byte(content))

	return htmx.NewResponse().Write(c.Response().Writer)
}
