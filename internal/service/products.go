package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/service/validate"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AddProduct(ctx context.Context,
	user domain.User, product domain.Product) (domain.Product, error) {
	logger := mylog.FromContext(ctx).With("user", user, "product", product)
	ctx = mylog.NewContext(ctx, logger)

	if err := validate.User(user); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := validate.Product(product); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := s.productStorage.Add(ctx, user, product); err != nil {
		err = convertErrAndLog(ctx, logger, "error adding product", err)
		return domain.Product{}, err
	}
	logger.Info("product added")

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, user domain.User, productName string) error {
	logger := mylog.FromContext(ctx).With("user", user, "product", productName)
	ctx = mylog.NewContext(ctx, logger)

	if err := validate.User(user); err != nil {
		logger.Info(err.Error())
		return err
	}

	if err := s.productStorage.Delete(ctx, user, productName); err != nil {
		err = convertErrAndLog(ctx, logger, "error deleting product", err)
		return err
	}
	logger.Info("product deleted")

	return nil
}

func (s *Service) UpdateProduct(ctx context.Context,
	user domain.User, product domain.Product) (domain.Product, error) {
	logger := mylog.FromContext(ctx).With("user", user, "product", product)
	ctx = mylog.NewContext(ctx, logger)

	if err := validate.User(user); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := validate.Product(product); err != nil {
		logger.Info(err.Error())
		return domain.Product{}, err
	}

	if err := s.productStorage.Update(ctx, user, product); err != nil {
		err = convertErrAndLog(ctx, logger, "error updating product", err)
		return domain.Product{}, err
	}
	logger.Info("product updated")

	return product, nil
}

func (s *Service) SelectProductsByUser(ctx context.Context,
	user domain.User) ([]domain.Product, error) {
	logger := mylog.FromContext(ctx).With("user", user)
	ctx = mylog.NewContext(ctx, logger)

	if err := validate.User(user); err != nil {
		logger.Info(err.Error())
		return nil, err
	}

	products, err := s.productStorage.SelectByUser(ctx, user)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error selecting products", err)
		return nil, err
	}
	logger.Info("products selected")

	return products, nil
}
