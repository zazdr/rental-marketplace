package store

import (
	"app/config"
	"app/sqlc/querier"
	"app/store/internal/dep"
)

type Store struct{}

func New(config *config.Config, querier *querier.Queries) *Store {
	_ = dep.New(config, querier)

	store := &Store{}

	return store
}
