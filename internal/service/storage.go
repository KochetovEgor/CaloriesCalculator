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
	Select(ctx context.Context, username string) (domain.User, error)
}

type ProductStorage interface {
	Close() error
	Init(ctx context.Context) error
	Add(ctx context.Context, user domain.User, product domain.Product) error
	Delete(ctx context.Context, user domain.User, productName string) error
	Update(ctx context.Context, user domain.User, product domain.Product) error
	SelectByUser(ctx context.Context, user domain.User) ([]domain.Product, error)
}
