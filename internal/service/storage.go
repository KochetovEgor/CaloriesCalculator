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

type RationStorage interface {
	Close() error
	Init(ctx context.Context) error
	AddNewRation(ctx context.Context, user domain.User, ration domain.Ration) (int, error)
	DeleteRation(ctx context.Context, user domain.User, date string) error
	UpdateRation(ctx context.Context, user domain.User, ration domain.Ration) (int, error)
	AddRationToRation(ctx context.Context, user domain.User, ration domain.Ration) (int, error)
	SelectRationsByUser(ctx context.Context, user domain.User) ([]domain.Ration, error)

	AddProductsEaten(ctx context.Context, user domain.User, rationId int, productsEaten []domain.ProductEaten) error
	DeleteProductsEatenByRation(ctx context.Context, rationId int) error
}
