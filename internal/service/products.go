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

func (s *Service) DeleteProduct(ctx context.Context, username, productName string) error {
	logger := mylog.FromContext(ctx)

	if err := validateUsername(username); err != nil {
		logger.Info(err.Error())
		return err
	}

	if err := s.productStorage.Delete(ctx, username, productName); err != nil {
		err = convertErrAndLog(ctx, logger, "error deleting product", err)
		return err
	}
	logger.Info("product deleted")

	return nil
}

func (s *Service) UpdateProduct(ctx context.Context,
	product domain.Product) (domain.Product, error) {
	logger := mylog.FromContext(ctx)

	if err := validateProduct(product); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := s.productStorage.Update(ctx, product); err != nil {
		err = convertErrAndLog(ctx, logger, "error updating product", err)
		return domain.Product{}, err
	}
	logger = logger.With("product", product)
	logger.Info("product updated")

	return product, nil
}

func (s *Service) SelectProductsByUser(ctx context.Context,
	username string) ([]domain.Product, error) {
	logger := mylog.FromContext(ctx)

	if err := validateUsername(username); err != nil {
		logger.Info(err.Error())
		return nil, err
	}

	products, err := s.productStorage.SelectByUser(ctx, username)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error selecting products", err)
		return nil, err
	}
	logger.Info("products selected")

	return products, nil
}
