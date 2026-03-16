package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AddProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	logger := mylog.FromContext(ctx)

	if err := validateProduct(product); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := s.productStorage.Add(ctx, product); err != nil {
		err = convertErrAndLog(ctx, logger, "error adding product", err)
		return domain.Product{}, err
	}
	logger = logger.With("product", product)
	logger.Info("product added")

	return product, nil
}
