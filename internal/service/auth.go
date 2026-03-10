package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AuthUser(ctx context.Context, username, password string) (domain.User, error) {
	user, err := s.storage.SelectUser(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	if err := auth.CheckPassword(user.HashPassword, password); err != nil {
		return domain.User{}, domain.ErrInvalidUserOrPassword
	}

	logger := mylog.FromContext(ctx)
	logger.Info("succesfully authorised", "username", username)

	return user, nil
}
