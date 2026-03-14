package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AuthUser(ctx context.Context, username, password string) (string, error) {
	logger := mylog.FromContext(ctx)

	user, err := s.userStorage.Select(ctx, username)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error selecting user", err)
		return "", err
	}
	logger = logger.With("user", user)
	logger.Info("selected user")

	if err := auth.CheckPassword(user.HashPassword, password); err != nil {
		logger.Info(domain.ErrInvalidUserOrPassword.Error())
		return "", domain.ErrInvalidUserOrPassword
	}
	logger.Info("authorised")

	token, err := auth.CreateAccessToken(user)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error creating access token", err)
		return "", err
	}
	logger.Info("created access token")

	return token, nil
}
