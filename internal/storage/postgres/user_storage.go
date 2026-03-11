package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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
	username text PRIMARY KEY,
	hashed_password text NOT NULL
);
`

// Init initialises table for working with users.
func (s *UserStorage) Init(ctx context.Context) error {
	logger := mylog.FromContext(ctx)

	_, err := s.pool.Exec(ctx, createTableUsers)
	if err != nil {
		return fmt.Errorf("error creating table users: %w", err)
	}
	logger.Info("table created", "table", tableUsersName)
	return nil
}

const addUserToUsers = `
INSERT INTO users (username, hashed_password)
VALUES ($1, $2);
`

// Add adds user to storage.
func (s *UserStorage) Add(ctx context.Context, user domain.User) error {
	_, err := s.pool.Exec(ctx, addUserToUsers, user.Username, user.HashPassword)
	if err != nil {
		return fmt.Errorf("error adding user to table users: %w", mapPgError(err))
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("user added",
		"table", tableUsersName, "username", user.Username)
	return nil
}

const deleteUserFromUsers = `
DELETE FROM users
WHERE username = $1;
`

// Delete deletes user from storage.
func (s *UserStorage) Delete(ctx context.Context, username string) error {
	_, err := s.pool.Exec(ctx, deleteUserFromUsers, username)
	if err != nil {
		return fmt.Errorf("error deleting user %s from table users: %w", username, err)
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("user deleted",
		"table", tableUsersName, "username", username)
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
	var user domain.User
	err := s.pool.QueryRow(ctx, selectUserFromUsers, username).Scan(
		&user.Username, &user.HashPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, domain.ErrInvalidUserOrPassword
		}
		return user, fmt.Errorf("error selecting user %s from table users: %w",
			username, err)
	}
	logger := mylog.FromContext(ctx)
	logger.Debug("selected user",
		"table", tableUsersName, "username", username)
	return user, nil
}
