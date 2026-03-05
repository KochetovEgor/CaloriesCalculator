package storage

import (
	"CaloriesCalculator/internal/domain"
	"context"
)

type Storage interface {
	Close() error
	Init(context.Context) error
	AddUser(context.Context, domain.User) error
	DeleteUser(context.Context, string) error
	SelectUser(context.Context, string) (domain.User, error)
}

var defaultStorage Storage

func Default() Storage {
	return defaultStorage
}

func SetDefault(storage Storage) {
	defaultStorage = storage
}

func Close() error {
	return defaultStorage.Close()
}

func Init(ctx context.Context) error {
	return defaultStorage.Init(ctx)
}

func AddUser(ctx context.Context, user domain.User) error {
	return defaultStorage.AddUser(ctx, user)
}

func DeleteUser(ctx context.Context, username string) error {
	return defaultStorage.DeleteUser(ctx, username)
}

func SelectUser(ctx context.Context, username string) (domain.User, error) {
	return defaultStorage.SelectUser(ctx, username)
}
