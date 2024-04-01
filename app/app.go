package app

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/api"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/config"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/processors"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/repositories"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/infra/storage/postgres"
)

type App struct {
	s          scratch.Scratch
	storage    *postgres.DB
	storage2   *postgres.DB
	processors *processors.Processors
}

func New(cfg *config.Config, scratch scratch.Scratch) App {
	storage := scratch.Postgres
	repos := repositories.New()

	storage2, _ := postgres.NewDB(&cfg.PostgresLocal, scratch.Prometheus)

	return App{
		s:          scratch,
		storage:    storage,
		storage2:   storage2,
		processors: processors.New(storage, storage2, repos),
	}
}

func (a *App) InitAPI() {
	api.New(a.s.Server, a.processors)
}

func (a App) Run() error {
	a.InitAPI()

	return a.s.Run()
}
