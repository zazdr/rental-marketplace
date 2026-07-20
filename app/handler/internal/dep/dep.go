package dep

import (
	"app/handler/internal/util"
	"app/store"
	"app/ui"
)

type Dep struct {
	Util  *util.Util
	Store *store.Store
	UI    *ui.UI
}

func New(util *util.Util, store *store.Store, ui *ui.UI) *Dep {
	return &Dep{
		Util:  util,
		Store: store,
		UI:    ui,
	}
}
