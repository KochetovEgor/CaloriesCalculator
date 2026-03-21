package validate

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/myerrors"
)

func UsernameAndPassword(username string, password string) error {
	var errs []error
	if err := Username(username); err != nil {
		errs = append(errs, err)
	}
	if err := Password(password); err != nil {
		errs = append(errs, err)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}

func User(user domain.User) error {
	return Username(user.Username)
}

func Product(product domain.Product) error {
	var errs []error
	if err := Weight(product.BaseWeight); err != nil {
		errs = append(errs, err)
	}
	if err := Portion(product.BasePortion); err != nil {
		errs = append(errs, err)
	}
	if err := Calories(product.Calories); err != nil {
		errs = append(errs, err)
	}
	if err := Fats(product.Fats); err != nil {
		errs = append(errs, err)
	}
	if err := Proteins(product.Proteins); err != nil {
		errs = append(errs, err)
	}
	if err := Carbohydrates(product.Carbohydrates); err != nil {
		errs = append(errs, err)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}
