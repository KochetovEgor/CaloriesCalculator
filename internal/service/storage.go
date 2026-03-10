package service

import (
	"CaloriesCalculator/internal/domain"
	"context"
)

// Storage is an intarface for interaction with storage.
type Storage interface {
	Close() error
	Init(ctx context.Context) error
	AddUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, username string) error
	SelectUser(ctx context.Context, username string) (domain.User, error)
}
