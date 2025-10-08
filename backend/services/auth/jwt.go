package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

type Secrets struct {
	key string
}

var (
	secrets *Secrets
	once    sync.Once
)

func setup() {
	once.Do(func() {
		// Generate a 512-bit (64-byte) random key
		secrets = &Secrets{
			key: string(make([]byte, 64)),
		}
	})
}

func CreateAdminJWT() (string, error) {
	if secrets == nil {
		setup()
	}

	// Only claims we care about are timestamps/issuer. We're only utilizing AuthN, not AuthZ. AuthZ is implied.
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "friday-api",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secrets.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseAdminJWT(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrets.key), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func IsJWTValid(tokenString string) (bool, error) {
	claims, err := parseAdminJWT(tokenString)
	if err != nil {
		return false, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) || claims.Issuer != "friday-api" {
		return false, nil
	}

	return true, nil
}
