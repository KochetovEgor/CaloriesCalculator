package storage

import "context"

type Storage interface {
	Close() error
	Init(context.Context) error
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
