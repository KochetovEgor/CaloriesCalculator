package postgres

import (
	"CaloriesCalculator/internal/pkg/config"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"))
}

// NewPool creates new postgres connection pool with given config cfg.
func NewPool(ctx context.Context, cfg config.Storage) (*pgxpool.Pool, error) {
	logger := mylog.FromContext(ctx)

	u, err := url.Parse(createURL())
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

	return pool, nil
}
