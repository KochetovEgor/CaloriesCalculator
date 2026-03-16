package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/myerrors"
)

func validateUsernameAndPWD(username string, password string) error {
	var errs []error
	if len([]rune(username)) < 3 {
		errs = append(errs, domain.ErrUsernameTooShort)
	}
	if len(password) > 72 {
		errs = append(errs, domain.ErrPaswordTooLong)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}

func validateProduct(product domain.Product) error {
	var errs []error
	if len([]rune(product.Username)) < 3 {
		errs = append(errs, domain.ErrUsernameTooShort)
	}
	if product.BaseWeight < 0 {
		errs = append(errs, domain.ErrBaseWeightMustBePositive)
	}
	if product.BasePortion < 0 {
		errs = append(errs, domain.ErrBasePortionMustBePositive)
	}
	if product.Fats < 0 {
		errs = append(errs, domain.ErrFatsMustBePositive)
	}
	if product.Proteins < 0 {
		errs = append(errs, domain.ErrProteinsMustBePositive)
	}
	if product.Carbohydrates < 0 {
		errs = append(errs, domain.ErrCarbohydratesMustBePositive)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}
