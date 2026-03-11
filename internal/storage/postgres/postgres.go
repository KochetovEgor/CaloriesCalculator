package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/config"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// mapPgError maps typical postgres error into domain.errors, but doesn't change other errors.
func mapPgError(err error) error {
	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		switch pgErr.Code {
		case "23505":
			return domain.ErrUserAlreadyExists
		}
	}
	return err
}

// NewPool creates new postgres connection pool with given config cfg.
func NewPool(ctx context.Context, cfg config.Storage) (*pgxpool.Pool, error) {
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

	return pool, nil
}
