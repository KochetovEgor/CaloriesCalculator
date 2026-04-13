package validate

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/myerrors"
	"fmt"
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
	var errs []error
	if err := Username(user.Username); err != nil {
		errs = append(errs, err)
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}

func Product(product domain.Product) error {
	var errs []error
	if err := BaseWeight(product.BaseWeight); err != nil {
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

func ProductEatenSlice(productsEaten []domain.ProductEaten) error {
	var errs []error
	for _, p := range productsEaten {
		if err := Weight(p.Weight); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
		if err := Portion(p.Portion); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
		if err := Calories(p.Calories); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
		if err := Fats(p.Fats); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
		if err := Proteins(p.Proteins); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
		if err := Carbohydrates(p.Carbohydrates); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", p.Name, err))
		}
	}

	if errs != nil {
		return myerrors.Join(errs...)
	}
	return nil
}
