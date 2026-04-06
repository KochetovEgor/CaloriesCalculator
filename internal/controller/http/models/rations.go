package models

import "CaloriesCalculator/internal/domain"

type Ration struct {
	Date          string  `json:"date"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func RationToModel(ration domain.Ration) Ration {
	return Ration(ration)
}

type RationWithProducts struct {
	Date     string         `json:"date"`
	Products []ProductEaten `json:"products"`
}
