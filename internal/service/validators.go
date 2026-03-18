package service

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/myerrors"
)

func validateUsername(username string) error {
	if len([]rune(username)) < 3 {
		return domain.ErrUsernameTooShort
	}
	return nil
}

func validatePassword(password string) error {
	if len([]rune(password)) < 3 {
		return domain.ErrPasswordTooShort
	}
	if len(password) > 72 {
		return domain.ErrPaswordTooLong
	}
	return nil
}

func validateUsernameAndPWD(username string, password string) error {
	var errs []error
	if err := validateUsername(username); err != nil {
		errs = append(errs, err)
	}
	if err := validatePassword(password); err != nil {
		errs = append(errs, err)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}

func validateProduct(product domain.Product) error {
	var errs []error
	if err := validateUsername(product.Username); err != nil {
		errs = append(errs, err)
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
