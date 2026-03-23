// service is a package for core logic of the app.
package service

import (
	"context"
	"fmt"
)

// Service contains storages.
type Service struct {
	userStorage    UserStorage
	productStorage ProductStorage
	rationStorage  RationStorage
}

func New(userStorage UserStorage, productStorage ProductStorage,
	rationStorage RationStorage) *Service {
	return &Service{userStorage: userStorage,
		productStorage: productStorage,
		rationStorage:  rationStorage}
}

// Close closes all storages.
func (s *Service) Close() error {
	if err := s.userStorage.Close(); err != nil {
		return err
	}
	if err := s.productStorage.Close(); err != nil {
		return err
	}
	if err := s.rationStorage.Close(); err != nil {
		return err
	}
	return nil
}

// Init initializes all storages.
func (s *Service) Init(ctx context.Context) error {
	if err := s.userStorage.Init(ctx); err != nil {
		return fmt.Errorf("error initializing service: %w", err)
	}
	if err := s.productStorage.Init(ctx); err != nil {
		return fmt.Errorf("error initializing service: %w", err)
	}
	if err := s.rationStorage.Init(ctx); err != nil {
		return fmt.Errorf("error initializing service: %w", err)
	}
	return nil
}
