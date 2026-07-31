package handler

import (
	"app/handler/internal/dep"
	"app/handler/internal/handler/fail"
	"app/handler/internal/handler/signup"
	"app/store"
	"app/ui"
)

type Handler struct {
	Fail   fail.Fail
	Signup signup.Signup
}

func New(store *store.Store, ui *ui.UI) *Handler {
	dep := dep.New(store, ui)

	handler := &Handler{}

	handler.Fail = fail.New(dep)
	handler.Signup = signup.New(dep)

	return handler
}
