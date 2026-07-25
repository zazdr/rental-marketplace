package bootstrap

import (
	"app/config"
	"app/db"
	"app/handler"
	"app/router"
	"app/server"
	"app/sqlc/querier"
	"app/store"
	"app/ui"
)

func New() (*server.Server, error) {
	config, err := config.New()
	if err != nil {
		return nil, err
	}

	db, err := db.New(config)
	if err != nil {
		return nil, err
	}
	querier := querier.New(db)
	store := store.New(config, querier)
	ui := ui.New()

	server := server.New(config)
	handler := handler.New(store, ui)
	router.New(ui, server, handler)

	return server, nil
}
