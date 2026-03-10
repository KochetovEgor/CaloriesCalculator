package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const tableLoginName = "login"

const addUserToLogin = `
INSERT INTO login (username, hashed_password)
VALUES ($1, $2);
`

func (db *DB) AddUser(ctx context.Context, user domain.User) error {
	_, err := db.pool.Exec(ctx, addUserToLogin, user.Username, user.HashPassword)
	if err != nil {
		return fmt.Errorf("error adding user to table login: %w", mapPgError(err))
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("user added",
		"table", tableLoginName, "username", user.Username)
	return nil
}

const deleteUserFromLogin = `
DELETE FROM login
WHERE username = $1;
`

func (db *DB) DeleteUser(ctx context.Context, username string) error {
	_, err := db.pool.Exec(ctx, deleteUserFromLogin, username)
	if err != nil {
		return fmt.Errorf("error deleting user %s from table login: %w", username, err)
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("user deleted",
		"table", tableLoginName, "username", username)
	return nil
}

const selectUserFromLogin = `
SELECT 
	username, hashed_password
FROM login
WHERE username = $1;
`

func (db *DB) SelectUser(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := db.pool.QueryRow(ctx, selectUserFromLogin, username).Scan(
		&user.Username, &user.HashPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, domain.ErrInvalidUserOrPassword
		}
		return user, fmt.Errorf("error selecting user %s from table login: %w",
			username, err)
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("selected user",
		"table", tableLoginName, "username", username)
	return user, nil
}
