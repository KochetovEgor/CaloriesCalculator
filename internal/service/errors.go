package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"log/slog"
)

func convertErrAndLog(ctx context.Context, logger *slog.Logger,
	logMsg string, err error) error {
	if err, ok := domain.ExtractErr(err); ok {
		logger.Info(err.Error())
		return err
	}

	logger.ErrorContext(mylog.ErrToContext(ctx, err), logMsg, "error", err.Error())
	return domain.ErrInternal
}
