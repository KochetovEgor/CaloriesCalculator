package auth

import (
	"CaloriesCalculator/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func SetKey(key []byte) {
	secretKey = key
}

func GetKey() []byte {
	return secretKey
}

func CreateAccessToken(user domain.User) (string, error) {
	var claims = jwt.MapClaims{
		"iss":       domain.Issuer,
		"sub":       user.Username,
		"iat":       time.Now().Unix(),
		"user_name": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	return signedToken, err
}
