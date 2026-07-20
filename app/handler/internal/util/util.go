package util

import (
	"app/dto"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

type Util struct{}

func New() *Util {
	return &Util{}
}

func (u *Util) Render(c *echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response())
}

func (u *Util) State(c *echo.Context) *dto.State {
	return &dto.State{
		Htmx: c.Request().Header.Get("HX-Request") == "true",
	}
}
