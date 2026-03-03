package postgres

import (
	"CaloriesCalculator/internal/config"
	"CaloriesCalculator/internal/mylog"
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Storage) (*DB, error) {
	logger := mylog.FromContext(ctx)

	u, err := url.Parse(cfg.Url)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	q := u.Query()
	q.Set("pool_max_conns", strconv.Itoa(cfg.MaxConns))
	q.Set("pool_min_conns", strconv.Itoa(cfg.MinConns))
	u.RawQuery = q.Encode()

	pool, err := pgxpool.New(ctx, u.String())
	if err != nil {
		return nil, fmt.Errorf("error creating pool: %w", err)
	}
	logger.Info("storage pool created")

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging pool: %w", err)
	}
	logger.Info("storage pool succesfully pinged")

	return &DB{pool: pool}, nil
}

func (db *DB) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}
	return nil
}
