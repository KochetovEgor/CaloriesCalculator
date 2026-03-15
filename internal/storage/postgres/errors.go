package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func isUniqueViolation(err error) bool {
	if err, ok := errors.AsType[*pgconn.PgError](err); ok {
		return err.Code == "23505"
	}
	return false
}

func isNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func isNoAffectedRows(ct pgconn.CommandTag) bool {
	return ct.RowsAffected() == 0
}
