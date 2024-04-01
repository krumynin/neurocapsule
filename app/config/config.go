package config

import (
	"github.com/kelseyhightower/envconfig"
	scratchcfg "stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/config"
)

const PREFIX = "NEUROCAPSULE"

type Config struct {
	*scratchcfg.Config
}

func NewFromEnv() *Config {
	c := Config{}
	envconfig.MustProcess(PREFIX, &c)

	return &c
}
