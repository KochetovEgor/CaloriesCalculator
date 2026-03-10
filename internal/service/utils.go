package service

import (
	"CaloriesCalculator/internal/domain"
	"log/slog"
)

func logErr(logger *slog.Logger, msg string, err error) {
	if err, ok := domain.ExtractErr(err); ok {
		logger.Info(err.Error())
		return
	}
	logger.Error(msg, "error", err.Error())
}
