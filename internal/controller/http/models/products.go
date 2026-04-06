package models

import "CaloriesCalculator/internal/domain"

type Product struct {
	Name          string  `json:"name"`
	BaseWeight    float64 `json:"base_weight"`
	BasePortion   float64 `json:"base_portion"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func ProductToModel(product domain.Product) Product {
	return Product(product)
}

type ProductEaten struct {
	Name          string  `json:"name"`
	Weight        float64 `json:"weight"`
	Portion       float64 `json:"portion"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func ProductEatenToDomain(product ProductEaten) domain.ProductEaten {
	return domain.ProductEaten(product)
}
