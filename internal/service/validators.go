package service

import "CaloriesCalculator/internal/domain"

func validatePassword(password string) error {
	if len(password) > 72 {
		return domain.ErrPaswordTooLong
	}
	return nil
}
