package u

import (
	"app/handler/internal/dep"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func Render(c *echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response())
}

func IsHTMX(c *echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func RenderFull(
	d *dep.Dep,
	c *echo.Context,
	main templ.Component,
) error {
	if IsHTMX(c) {
		return Render(c, main)
	}
	return Render(c, d.UI.Layout(main))
}

func FailClientNew(code int) error {
	return &echo.HTTPError{
		Code:    code,
		Message: http.StatusText(code),
	}
}

func FailServerNew(message string, err error) error {
	if err != nil {
		message = message + ": " + err.Error()
	}
	return &echo.HTTPError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
