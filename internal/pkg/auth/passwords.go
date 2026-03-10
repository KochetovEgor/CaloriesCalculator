// auth is package for creating bcrypt hash of passwords and validating JWT
package auth

import (
	"CaloriesCalculator/internal/domain"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates bcrypt hash from password.
// If password is longer than 72 bytes it returns domain.ErrPasswordTooLong error.
func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		err = domain.ErrPaswordTooLong
	} else {
		err = fmt.Errorf("error generating bcrypt hash: %w", err)
	}
	return hash, err
}

// ChechPassword compares bcrypt hashPassword with given password.
func CheckPassword(hashPassword []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
}
