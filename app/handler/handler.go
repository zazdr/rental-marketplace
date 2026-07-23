package handler

import (
	"app/handler/internal/dep"
	"app/handler/internal/handler/fail"
	"app/handler/internal/handler/user/signup"
	"app/handler/internal/util"
	"app/store"
	"app/ui"
)

type Handler struct {
	Fail fail.Fail
	User struct {
		Signup signup.Signup
	}
}

func New(store *store.Store, ui *ui.UI) *Handler {
	dep := dep.New(util.New(), store, ui)

	handler := &Handler{}

	handler.Fail = fail.New(dep)
	handler.User.Signup = signup.New(dep)

	return handler
}
