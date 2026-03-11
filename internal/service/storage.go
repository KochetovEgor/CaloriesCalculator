package service

import (
	"CaloriesCalculator/internal/domain"
	"context"
)

// UserStorage is an interface for interaction with user storage.
type UserStorage interface {
	Close() error
	Init(ctx context.Context) error
	Add(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, username string) error
	Select(ctx context.Context, username string) (domain.User, error)
}
