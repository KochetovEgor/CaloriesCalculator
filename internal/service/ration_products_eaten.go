package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/utils"
	"CaloriesCalculator/internal/pkg/validate"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AddProductsToRation(ctx context.Context, user domain.User,
	date string, productsEaten []domain.ProductEaten) (domain.Ration, error) {
	logger := mylog.FromContext(ctx).With("user", user)
	ctx = mylog.NewContext(ctx, logger)

	if err := validate.ProductEatenSlice(productsEaten); err != nil {
		logger.Info(err.Error())
		return domain.Ration{}, err
	}

	products, err := s.productStorage.SelectByUser(ctx, user)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error selecting products", err)
		return domain.Ration{}, err
	}
	logger.Info("products selected")

	ration, productsEaten, err := utils.MakeRationFromProducts(
		products, productsEaten)
	if err != nil {
		logger.Info(err.Error())
		return domain.Ration{}, err
	}
	ration.Date = date
	logger = logger.With("ration", ration)
	ctx = mylog.NewContext(ctx, logger)

	id, err := s.rationStorage.AddRationToRation(ctx, user, ration)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error adding ration to ration", err)
		return domain.Ration{}, nil
	}
	logger.Info("ration updated")

	if err := s.rationStorage.AddProductsEaten(ctx, user, id, productsEaten); err != nil {
		err = convertErrAndLog(ctx, logger, "error adding products eaten", err)
		return domain.Ration{}, err
	}
	logger.Info("products eaten added")

	return ration, err
}
