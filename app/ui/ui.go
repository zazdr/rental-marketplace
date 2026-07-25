package ui

import (
	"app/dto"
	"app/ui/internal/templ/page/fail"
	"app/ui/internal/templ/page/user/signup"
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
	Fail func(state *dto.State, code int, message string) templ.Component
	User struct {
		Signup func(state *dto.State, form *dto.FormUserSignup) templ.Component
	}
}

func New() *UI {
	ui := &UI{}

	ui.Static.File = static
	ui.Static.Dir = "static"

	ui.Fail = fail.Fail

	ui.User.Signup = signup.Signup

	return ui
}
