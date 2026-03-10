// service is a package for core logic of the app.
package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"context"
	"fmt"
	"log/slog"
)

// Service contains storage of type Storage.
type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{storage: storage}
}

// Close closes Service.storage.
func (s *Service) Close() error {
	return s.storage.Close()
}

// Init initializes storage.
func (s *Service) Init(ctx context.Context) error {
	if err := s.storage.Init(ctx); err != nil {
		return fmt.Errorf("error initializing service: %w", err)
	}
	return nil
}

func (s *Service) Test(ctx context.Context) {
	hashPassword, _ := auth.HashPassword("123")
	if err := s.storage.AddUser(ctx, domain.User{Username: "Egor",
		HashPassword: hashPassword}); err != nil {
		slog.Error(err.Error())
	}

	if err := s.storage.DeleteUser(ctx, "egor"); err != nil {
		slog.Error(err.Error())
	}

	user, err := s.storage.SelectUser(ctx, "Egor")
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Printf("%s\n", user)

	user, err = s.AuthUser(ctx, "egor", "123")
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Printf("%s\n", user)

	token, err := auth.CreateAccessToken(user)
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(token)
}
