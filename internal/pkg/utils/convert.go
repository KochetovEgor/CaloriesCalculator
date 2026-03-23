package utils

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/myerrors"
	"fmt"
)

func MakeRationFromProducts(products []domain.Product,
	productsEaten []domain.ProductEaten) (domain.Ration, []domain.ProductEaten, error) {

	productMap := make(map[string]domain.Product, len(products))
	for _, p := range products {
		productMap[p.Name] = p
	}

	var errs []error
	var ration domain.Ration
	for i, pe := range productsEaten {
		p, ok := productMap[pe.Name]
		if !ok {
			errs = append(errs, fmt.Errorf("%s: %w", pe.Name, domain.ErrProductNotExists))
			continue
		}

		k := (pe.Weight + pe.Portion*p.BasePortion) / p.BaseWeight
		pe.Calories = p.Calories * k
		pe.Fats = p.Fats * k
		pe.Proteins = p.Proteins * k
		pe.Carbohydrates = p.Carbohydrates * k
		productsEaten[i] = pe

		ration.Calories += pe.Calories
		ration.Fats += pe.Fats
		ration.Proteins += pe.Proteins
		ration.Carbohydrates += pe.Carbohydrates
	}

	if errs != nil {
		return domain.Ration{}, nil, myerrors.Join(errs...)
	}
	return ration, productsEaten, nil
}
