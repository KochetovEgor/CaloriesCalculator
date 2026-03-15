package service

import "CaloriesCalculator/internal/domain"

func validatePassword(password string) error {
	if len(password) > 72 {
		return domain.ErrPaswordTooLong
	}
	return nil
}

func validateUsername(username string) error {
	if len([]rune(username)) < 3 {
		return domain.ErrUsernameTooShort
	}
	return nil
}

func validateUsernameAndPWD(username string, password string) error {
	if len([]rune(username)) < 3 {
		return domain.ErrUsernameTooShort
	}
	if len(password) > 72 {
		return domain.ErrPaswordTooLong
	}
	return nil
}
