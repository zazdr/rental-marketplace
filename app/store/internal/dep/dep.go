package dep

import (
	"app/config"
	"app/sqlc/querier"
)

type Dep struct {
	Config  *config.Config
	Querier *querier.Queries
}

func New(config *config.Config, querier *querier.Queries) *Dep {
	return &Dep{
		Config:  config,
		Querier: querier,
	}
}
