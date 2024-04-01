package scratch

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/config"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/infra/storage/postgres"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/server"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type Scratch struct {
	Cfg        *config.Config
	Server     *server.Server
	Postgres   *postgres.DB
	Prometheus *prometheus.Registry
	Log        *zap.Logger
}

func New(cfg *config.Config, version string) Scratch {
	var (
		prometheusRegistry *prometheus.Registry
		loggerInstance     *zap.Logger
	)
	srv := server.NewServer(cfg, version)
	prometheusRegistry = srv.Prometheus
	loggerInstance = srv.Log

	db, err := postgres.NewDB(&cfg.Postgres, prometheusRegistry)
	if err != nil {
		panic(err)
	}

	return Scratch{
		Cfg:        cfg,
		Server:     srv,
		Postgres:   db,
		Prometheus: prometheusRegistry,
		Log:        loggerInstance,
	}
}

func (s Scratch) Run() error {
	s.Server.Run()

	return nil
}
