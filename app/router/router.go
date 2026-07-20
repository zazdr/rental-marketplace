package router

import (
	"app/handler"
	"app/server"
	"app/shared"
	"app/ui"

	"github.com/labstack/echo/v5"
)

func New(
	ui *ui.UI,
	server *server.Server,
	handler *handler.Handler,
) {
	server.Echo.StaticFS(
		shared.RouterStatic,
		echo.MustSubFS(ui.Static.File, ui.Static.Dir),
	)

	server.Echo.GET(shared.RouterUserSignup, handler.User.Signup.Get)
}
