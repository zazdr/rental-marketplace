package dep

import (
	"app/store"
	"app/ui"
)

type Dep struct {
	Store *store.Store
	UI    *ui.UI
}

func New(store *store.Store, ui *ui.UI) *Dep {
	return &Dep{
		Store: store,
		UI:    ui,
	}
}
