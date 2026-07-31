package ui

import (
	"app/dto"
	"app/ui/internal/layout"
	"app/ui/internal/page/fail"
	"app/ui/internal/page/signup"
	"embed"

	"github.com/a-h/templ"
)

//go:embed static
var static embed.FS

type UI struct {
	Static struct {
		File embed.FS
		Dir  string
	}
	Layout func(main templ.Component) templ.Component
	Fail   func(code int, message string) templ.Component
	Signup func(form *dto.FormSignup) templ.Component
}

func New() *UI {
	ui := &UI{}

	ui.Static.File = static
	ui.Static.Dir = "static"

	ui.Layout = layout.Layout

	ui.Fail = fail.Fail
	ui.Signup = signup.Signup

	return ui
}
