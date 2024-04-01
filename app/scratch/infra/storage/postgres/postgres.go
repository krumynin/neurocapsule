package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	lpgx "stash.lamoda.ru/gotools/pgx"

	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/config"
)

type DB struct {
	Conn  lpgx.PoolMeasured
	Query *squirrel.StatementBuilderType
}

func NewDB(conf *config.DB, registerer prometheus.Registerer) (*DB, error) {
	dsn := formDbURI(conf)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig error: %w", err)
	}

	// using binary parameters query mode for compatibility with PgBouncer
	// https://blog.bullgare.com/2019/06/pgbouncer-and-prepared-statements/
	cfg.ConnConfig.PreferSimpleProtocol = true

	cfg.MaxConns = int32(conf.MaxConns)
	cfg.MaxConnLifetime = conf.ConnMaxLifetime
	cfg.MaxConnIdleTime = conf.MaxIdleConnTime

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("ConnectConfig error: %w", err)
	}

	conn := lpgx.MustPoolWithMetrics(pool, registerer, prometheus.Labels{"db_name": conf.Database}, conf.Database)

	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{Conn: conn, Query: &sq}, nil
}

func formDbURI(conf *config.DB) string {
	pass := url.QueryEscape(conf.Password)
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable&connect_timeout=10",
		conf.User, pass, conf.Host, conf.Database,
	)
}
