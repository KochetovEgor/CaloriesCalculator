package auth

import (
	"CaloriesCalculator/internal/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const Issuer = "calories_calculator"

var secretKey []byte

func SetKey(key []byte) {
	secretKey = key
}

func GetKey() []byte {
	return secretKey
}

func keyFunc() jwt.Keyfunc {
	return func(_ *jwt.Token) (any, error) { return secretKey, nil }
}

func CreateAccessToken(user domain.User) (string, error) {
	var claims = jwt.MapClaims{
		"iss":       Issuer,
		"sub":       user.Username,
		"iat":       time.Now().Unix(),
		"exp":       (time.Now().Add(time.Minute * 10)).Unix(),
		"user_name": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	return signedToken, err
}

func VerifyAccessToken(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken,
		keyFunc(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(Issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, fmt.Errorf("parse token failed: %w", err)
	}

	if !token.Valid {
		return nil, domain.ErrInvalidAccessToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.ErrInvalidAccessToken
	}

	return claims, nil
}

func GetUserFromToken(accessToken string) (domain.User, error) {
	var user domain.User
	claims, err := VerifyAccessToken(accessToken)
	if err != nil {
		return user, err
	}

	user.Username = claims["user_name"].(string)
	return user, nil
}
