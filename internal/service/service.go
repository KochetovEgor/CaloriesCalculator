// service is a package for core logic of the app.
package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"context"
	"fmt"
	"log/slog"
)

// Service contains storages.
type Service struct {
	userStorage UserStorage
}

func New(userStorage UserStorage) *Service {
	return &Service{userStorage: userStorage}
}

// Close closes all storages.
func (s *Service) Close() error {
	if err := s.userStorage.Close(); err != nil {
		return err
	}
	return nil
}

// Init initializes all storages.
func (s *Service) Init(ctx context.Context) error {
	if err := s.userStorage.Init(ctx); err != nil {
		return fmt.Errorf("error initializing service: %w", err)
	}
	return nil
}

func (s *Service) Test(ctx context.Context) {
	hashPassword, _ := auth.HashPassword("123")
	if err := s.userStorage.Add(ctx, domain.User{Username: "Egor",
		HashPassword: hashPassword}); err != nil {
		slog.Error(err.Error())
	}

	if err := s.userStorage.Delete(ctx, "egor"); err != nil {
		slog.Error(err.Error())
	}

	user, err := s.userStorage.Select(ctx, "Egor")
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Printf("%s\n", user)

	token, err := s.AuthUser(ctx, "egor", "123")
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(token)
}
