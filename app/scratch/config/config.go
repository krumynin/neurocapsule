package config

import (
	"time"
)

type Config struct {
	LogIndex      string `envconfig:"LOG_INDEX"`
	Postgres      DB     `envconfig:"POSTGRES"`
	PostgresLocal DB     `envconfig:"POSTGRES_LOCAL"`
}

type DB struct {
	User            string        `envconfig:"USER" required:"true"`
	Password        string        `envconfig:"PASSWORD" required:"true"`
	Host            string        `envconfig:"HOST" required:"true"`
	Database        string        `envconfig:"DATABASE" required:"true"`
	MaxIdleConnTime time.Duration `envconfig:"MAX_IDLE_CONN_TIME" default:"5m"`
	MaxConns        int           `envconfig:"MAX_CONNS" default:"20"`
	ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"10m"`
}
