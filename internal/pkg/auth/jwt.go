package auth

import (
	"CaloriesCalculator/internal/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Issuer for JWT claim "iss".
const Issuer = "calories_calculator"

// secretKey is a key for signing JWT.
var secretKey []byte

// SetKey sets secret key for signing JWT
func SetKey(key []byte) {
	secretKey = key
}

// GetKey returns secret key for signing JWT.
func GetKey() []byte {
	return secretKey
}

func keyFunc() jwt.Keyfunc {
	return func(_ *jwt.Token) (any, error) { return secretKey, nil }
}

// CreateAccessToken creates JWT with claims "iss", "sub", "iat", "exp", "user_name", "user_id".
func CreateAccessToken(user domain.User) (string, error) {
	var claims = jwt.MapClaims{
		"iss":       Issuer,
		"sub":       user.Username,
		"iat":       time.Now().Unix(),
		"exp":       (time.Now().Add(10 * time.Hour)).Unix(),
		"user_id":   user.Id,
		"user_name": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	return signedToken, err
}

// VerifyAccessToken returns non nil error if accessToken is not valid,
// else it returns claims of token.
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

	id, _ := claims["user_id"].(float64)
	user.Id = int(id)
	user.Username, _ = claims["user_name"].(string)
	return user, nil
}
