package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AuthUser(ctx context.Context, username, password string) (string, error) {
	logger := mylog.FromContext(ctx)

	if err := validateUsernameAndPWD(username, password); err != nil {
		logger.Info(err.Error())
		return "", err
	}

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

func (s *Service) RegisterUser(ctx context.Context, username, password string) (domain.User, error) {
	logger := mylog.FromContext(ctx)

	if err := validateUsernameAndPWD(username, password); err != nil {
		logger.Info(err.Error())
		return domain.User{}, err
	}

	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error hashing password", err)
		return domain.User{}, err
	}

	user := domain.User{Username: username, HashPassword: hashPassword}
	logger = logger.With("user", user)

	if err := s.userStorage.Add(ctx, user); err != nil {
		err = convertErrAndLog(ctx, logger, "error adding user", err)
		return domain.User{}, err
	}
	logger.Info("user added")

	return user, nil
}
