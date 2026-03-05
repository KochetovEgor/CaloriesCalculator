package postgres

import (
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
)

const createTableLogin = `
CREATE TABLE IF NOT EXISTS login (
	username text PRIMARY KEY,
	hashed_password text NOT NULL
);
`

func (db *DB) Init(ctx context.Context) error {
	logger := mylog.FromContext(ctx)

	_, err := db.pool.Exec(ctx, createTableLogin)
	if err != nil {
		return fmt.Errorf("error creating table login: %w", err)
	}
	logger.Info("table login succesfully created")
	return nil
}
