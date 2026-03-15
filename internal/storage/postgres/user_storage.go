package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserStorage implements servise.UserStorage interface in postgres.
type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{pool: pool}
}

// Close closes connection pool.
func (s *UserStorage) Close() error {
	if s.pool != nil {
		s.pool.Close()
	}
	return nil
}

const tableUsersName = "users"

const createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username text NOT NULL,
	hashed_password text NOT NULL,
	CONSTRAINT unique_username UNIQUE (username)
);
`

// Init initialises table for working with users.
func (s *UserStorage) Init(ctx context.Context) error {
	attrs := []any{
		"table", tableUsersName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	_, err := s.pool.Exec(ctx, createTableUsers)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error creating table: %w", err)
	}

	logger.Debug("table created")
	return nil
}

const addUserToUsers = `
INSERT INTO users (username, hashed_password)
VALUES ($1, $2);
`

// Add adds user to storage.
func (s *UserStorage) Add(ctx context.Context, user domain.User) error {
	attrs := []any{
		"table", tableUsersName,
		"user", user,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	_, err := s.pool.Exec(ctx, addUserToUsers, user.Username, user.HashPassword)
	if err != nil {
		if isUniqueViolation(err) {
			return domain.ErrUserAlreadyExists
		}
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error adding user to table: %w", err)
	}

	logger.Debug("user added")
	return nil
}

const deleteUserFromUsers = `
DELETE FROM users
WHERE username = $1;
`

// Delete deletes user from storage.
func (s *UserStorage) Delete(ctx context.Context, user domain.User) error {
	attrs := []any{
		"table", tableUsersName,
		"username", user.Username,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	_, err := s.pool.Exec(ctx, deleteUserFromUsers, user.Username)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error deleting user from table: %w", err)
	}
	logger.Debug("user deleted")
	return nil
}

const selectUserFromUsers = `
SELECT
	username, hashed_password
FROM users
WHERE username = $1;
`

// Select selects user from storage and returns it.
func (s *UserStorage) Select(ctx context.Context, username string) (domain.User, error) {
	attrs := []any{
		"table", tableUsersName,
		"username", username,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	var user domain.User
	err := s.pool.QueryRow(ctx, selectUserFromUsers, username).Scan(
		&user.Username, &user.HashPassword)
	if err != nil {
		if isNoRows(err) {
			return user, domain.ErrInvalidUserOrPassword
		}
		err = mylog.WrapError(err, attrs...)
		return user, fmt.Errorf("error selecting user from table: %w", err)
	}
	logger.Debug("selected user")
	return user, nil
}
