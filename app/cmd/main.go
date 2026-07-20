package main

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

func main() {
	config, err := config.New()
	exit(err)

	db, err := db.New(config)
	exit(err)
	querier := querier.New(db)
	store := store.New(config, querier)
	ui := ui.New()

	server := server.New(config)
	handler := handler.New(store, ui)
	router.New(ui, server, handler)

	exit(server.Start())
}
