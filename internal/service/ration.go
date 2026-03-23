package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/utils"
	"CaloriesCalculator/internal/pkg/validate"
	"CaloriesCalculator/pkg/mylog"
	"context"
)

func (s *Service) AddRation(ctx context.Context, user domain.User,
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

	id, err := s.rationStorage.AddNewRation(ctx, user, ration)
	if err != nil {
		err = convertErrAndLog(ctx, logger, "error adding ration", err)
		return domain.Ration{}, err
	}
	logger.Info("ration added")

	if err := s.rationStorage.AddProductsEaten(ctx, user, id, productsEaten); err != nil {
		err = convertErrAndLog(ctx, logger, "error adding products eaten", err)
		return domain.Ration{}, err
	}
	logger.Info("products eaten added")

	return ration, nil
}

func (s *Service) DeleteRation(ctx context.Context, user domain.User, date string) error {
	logger := mylog.FromContext(ctx).With("user", user, "ration", date)
	ctx = mylog.NewContext(ctx, logger)

	if err := s.rationStorage.DeleteRation(ctx, user, date); err != nil {
		err = convertErrAndLog(ctx, logger, "error deleting ration", err)
		return err
	}
	logger.Info("ration deleted")

	return nil
}
