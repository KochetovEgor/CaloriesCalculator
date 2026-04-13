package validate

import (
	"CaloriesCalculator/internal/domain"
	"slices"
	"testing"
)

func extractWrappedErrors(err error) []error {
	if wrappedErr, ok := err.(interface{ Unwrap() []error }); ok && wrappedErr != nil {
		return wrappedErr.Unwrap()
	}
	return nil
}

func TestUsernameAndPassword(t *testing.T) {
	tests := []struct {
		username string
		password string
		want     []error
	}{
		{
			username: "Raiden Shogun",
			password: "qwerty",
			want:     []error{},
		},
		{
			username: "Чжун Ли",
			password: "Моракс",
			want:     nil,
		},
		{
			username: "Ra",
			password: "12345",
			want:     []error{domain.ErrUsernameTooShort},
		},
		{
			username: "Mavuika",
			password: "Ho",
			want:     []error{domain.ErrPasswordTooShort},
		},
		{
			username: "Furina",
			password: "Тестовый длинный текст ффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффффф",
			want:     []error{domain.ErrPaswordTooLong},
		},
		{
			username: "Fu",
			password: "12",
			want:     []error{domain.ErrUsernameTooShort, domain.ErrPasswordTooShort},
		},
	}

	for _, test := range tests {
		gotErr := UsernameAndPassword(test.username, test.password)
		got := extractWrappedErrors(gotErr)

		if len(got) != len(test.want) {
			t.Errorf("\nUsernameAndPassword(%v, %v) = %v\nWant: %v",
				test.username, test.password, got, test.want)
			return
		}

		for _, err := range test.want {
			if !slices.Contains(got, err) {
				t.Errorf("\nUsernameAndPassword(%v, %v) = %v\nWant: %v",
					test.username, test.password, got, test.want)
				return
			}
		}
	}
}

func TestUser(t *testing.T) {
	tests := []struct {
		input domain.User
		want  []error
	}{
		{
			input: domain.User{Id: 2, Username: "Raiden", HashPassword: nil},
			want:  []error{},
		},
		{
			input: domain.User{Username: "Testing"},
			want:  nil,
		},
		{
			input: domain.User{Id: 1, Username: "Ra"},
			want:  []error{domain.ErrUsernameTooShort},
		},
	}

	for _, test := range tests {
		gotErr := User(test.input)
		got := extractWrappedErrors(gotErr)

		if len(got) != len(test.want) {
			t.Errorf("\nUser(%v) = %v\nWant: %v",
				test.input, got, test.want)
			return
		}

		for _, err := range test.want {
			if !slices.Contains(got, err) {
				t.Errorf("\nUser(%v) = %v\nWant: %v",
					test.input, got, test.want)
				return
			}
		}
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		input domain.Product
		want  []error
	}{
		{
			input: domain.Product{Name: "Burger", BaseWeight: 100, BasePortion: 250.4,
				Calories: 350, Fats: 10.5, Proteins: 26, Carbohydrates: 30},
			want: nil,
		},
		{
			input: domain.Product{BaseWeight: 0},
			want:  []error{domain.ErrBaseWeightMustBeGreaterThanZero},
		},
		{
			input: domain.Product{Name: "", BaseWeight: -5, BasePortion: -1,
				Calories: -10, Fats: -5, Proteins: -0.5, Carbohydrates: -100.6},
			want: []error{domain.ErrBaseWeightMustBeGreaterThanZero, domain.ErrPortionMustBePositive,
				domain.ErrCaloriesMustBePostitve, domain.ErrFatsMustBePositive,
				domain.ErrProteinsMustBePositive, domain.ErrCarbohydratesMustBePositive},
		},
	}

	for _, test := range tests {
		gotErr := Product(test.input)
		got := extractWrappedErrors(gotErr)

		if len(got) != len(test.want) {
			t.Errorf("\nProduct(%v) = %v\nWant: %v",
				test.input, got, test.want)
			return
		}

		for _, err := range test.want {
			if !slices.Contains(got, err) {
				t.Errorf("\nProduct(%v) = %v\nWant: %v",
					test.input, got, test.want)
				return
			}
		}
	}
}
