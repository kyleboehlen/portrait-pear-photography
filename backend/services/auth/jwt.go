package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type Secrets struct {
	key string
}

var secrets *Secrets

func setSecretKey() {
	if secrets == nil {
		secrets = &Secrets{
			key: string(make([]byte, 64)),
		}
	}
}

func CreateAdminJWT() (string, error) {
	setSecretKey()

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

func GetCurrentHash() (string, bool) {
	now := time.Now()
	key := now.Format("2006-01")

	var test, exists = os.LookupEnv("TEST")
	if exists && test == "true" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"+key), 14)
		return string(hash), true
	}

	hash, exists := theForbiddenHashArray[key]
	return hash, exists
}

func ResetSecrets() {
	secrets = nil
}
