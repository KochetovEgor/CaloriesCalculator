package validate

import "CaloriesCalculator/internal/domain"

func Username(username string) error {
	if len([]rune(username)) < 3 {
		return domain.ErrUsernameTooShort
	}
	return nil
}

func Password(password string) error {
	if len([]rune(password)) < 3 {
		return domain.ErrPasswordTooShort
	}
	if len(password) > 72 {
		return domain.ErrPaswordTooLong
	}
	return nil
}

func BaseWeight(baseWeight float64) error {
	if baseWeight <= 0 {
		return domain.ErrBaseWeightMustBeGreaterThanZero
	}
	return nil
}

func Weight(weight float64) error {
	if weight < 0 {
		return domain.ErrWeightMustBePositive
	}
	return nil
}

func Portion(portion float64) error {
	if portion < 0 {
		return domain.ErrPortionMustBePositive
	}
	return nil
}

func Calories(calories float64) error {
	if calories < 0 {
		return domain.ErrCaloriesMustBePostitve
	}
	return nil
}

func Fats(fats float64) error {
	if fats < 0 {
		return domain.ErrFatsMustBePositive
	}
	return nil
}

func Proteins(proteins float64) error {
	if proteins < 0 {
		return domain.ErrProteinsMustBePositive
	}
	return nil
}

func Carbohydrates(carbohydrates float64) error {
	if carbohydrates < 0 {
		return domain.ErrCarbohydratesMustBePositive
	}
	return nil
}
