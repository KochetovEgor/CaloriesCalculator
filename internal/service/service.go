// service is a package for core logic of the app.
package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
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
	logger := mylog.FromContext(ctx)

	hashPassword, _ := auth.HashPassword("123")
	if err := s.userStorage.Add(ctx, domain.User{Username: "Egor",
		HashPassword: hashPassword}); err != nil {
		convertErrAndLog(ctx, logger, "error adding user", err)
	}

	if err := s.userStorage.Delete(ctx, "egor"); err != nil {
		convertErrAndLog(ctx, logger, "error deleting user", err)
	}

	user, err := s.userStorage.Select(ctx, "Egor")
	if err != nil {
		convertErrAndLog(ctx, logger, "error selecting user", err)
	}

	fmt.Printf("%s\n", user)

	/*token, err := s.AuthUser(ctx, "egor", "123")
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println(token)*/

	/*logger := mylog.FromContext(ctx)

	errBase := errors.New("test error")
	fn := func(err error) error {
		return mylog.WrapError(err, slog.Bool("bool test", false), slog.Int("nested int test", 32))
	}
	err := fn(errBase)
	err = mylog.WrapError(err, slog.String("hello", "it's mi"), slog.Int("int test", 64))
	ctx = mylog.ErrToContext(ctx, err)
	logger.ErrorContext(ctx, err.Error())*/
}
