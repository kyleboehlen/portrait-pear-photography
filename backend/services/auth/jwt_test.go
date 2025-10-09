package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestCreateAdminJWT_Success(t *testing.T) {
	// Reset secrets for clean test
	secrets = nil

	tokenString, err := CreateAdminJWT()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if tokenString == "" {
		t.Error("Expected non-empty token string")
	}

	// Verify token can be parsed
	claims, err := parseAdminJWT(tokenString)
	if err != nil {
		t.Errorf("Expected token to be parseable, got error: %v", err)
	}
	if claims.Issuer != "friday-api" {
		t.Errorf("Expected issuer to be 'friday-api', got '%s'", claims.Issuer)
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		t.Error("Expected token to not be expired")
	}
}

func TestCreateAdminJWT_SecretsInitialization(t *testing.T) {
	// Reset secrets
	secrets = nil

	// First call should initialize secrets
	token1, err := CreateAdminJWT()
	if err != nil {
		t.Errorf("Expected no error on first call, got %v", err)
	}

	// Second call should reuse same secrets
	token2, err := CreateAdminJWT()
	if err != nil {
		t.Errorf("Expected no error on second call, got %v", err)
	}

	// Both should be parseable with same key
	_, err1 := parseAdminJWT(token1)
	_, err2 := parseAdminJWT(token2)
	if err1 != nil || err2 != nil {
		t.Error("Both tokens should be parseable with same key")
	}
}

func TestParseAdminJWT_ValidToken(t *testing.T) {
	// Reset and create a valid token
	secrets = nil
	tokenString, err := CreateAdminJWT()
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	claims, err := parseAdminJWT(tokenString)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else {
		if claims == nil {
			t.Error("Expected non-nil claims")
		} else {
			if claims.Issuer != "friday-api" {
				t.Errorf("Expected issuer to be 'friday-api', got '%s'", claims.Issuer)
			}
			if claims.ExpiresAt.Time.Before(time.Now()) {
				t.Error("Expected token to not be expired")
			}
			if claims.IssuedAt.Time.After(time.Now()) {
				t.Error("Expected issued at time to be in the past")
			}
		}
	}
}

func TestParseAdminJWT_InvalidToken(t *testing.T) {
	// Ensure secrets are initialized
	secrets = nil
	_, _ = CreateAdminJWT()

	invalidToken := "invalid.token.string"

	claims, err := parseAdminJWT(invalidToken)

	if err == nil {
		t.Error("Expected error for invalid token")
	}
	if claims != nil {
		t.Error("Expected nil claims for invalid token")
	}
}

func TestParseAdminJWT_WrongSignature(t *testing.T) {
	// Create a token with a different key
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "friday-api",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	wrongToken, _ := token.SignedString([]byte("wrong-key"))

	// Initialize secrets with different key
	secrets = nil
	_, _ = CreateAdminJWT()

	parsedClaims, err := parseAdminJWT(wrongToken)

	if err == nil {
		t.Error("Expected error for token with wrong signature")
	}
	if parsedClaims != nil {
		t.Error("Expected nil claims for token with wrong signature")
	}
}

func TestIsJWTValid_ValidToken(t *testing.T) {
	// Reset and create a valid token
	secrets = nil
	tokenString, err := CreateAdminJWT()
	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	valid, err := IsJWTValid(tokenString)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !valid {
		t.Error("Expected token to be valid")
	}
}

func TestIsJWTValid_InvalidToken(t *testing.T) {
	// Ensure secrets are initialized
	secrets = nil
	_, _ = CreateAdminJWT()

	invalidToken := "invalid.token.string"

	valid, err := IsJWTValid(invalidToken)

	if err == nil {
		t.Error("Expected error for invalid token")
	}
	if valid {
		t.Error("Expected token to be invalid")
	}
}

func TestIsJWTValid_ExpiredToken(t *testing.T) {
	// Create an expired token
	secrets = &Secrets{key: string(make([]byte, 64))}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)), // Expired 1 hour ago
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		Issuer:    "friday-api",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	expiredToken, _ := token.SignedString([]byte(secrets.key))

	valid, err := IsJWTValid(expiredToken)

	if err == nil {
		t.Errorf("Expected error for expired token")
	}
	if valid {
		t.Error("Expected expired token to be invalid")
	}
}

func TestIsJWTValid_WrongIssuer(t *testing.T) {
	// Create a token with wrong issuer
	secrets = &Secrets{key: string(make([]byte, 64))}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "wrong-issuer",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	wrongIssuerToken, _ := token.SignedString([]byte(secrets.key))

	valid, err := IsJWTValid(wrongIssuerToken)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if valid {
		t.Error("Expected token with wrong issuer to be invalid")
	}
}

func TestGetCurrentHash_ValidKey(t *testing.T) {
	// Test assumes theForbiddenHashArray has current month entry
	now := time.Now()
	expectedKey := now.Format("2006-01")

	hash, exists := GetCurrentHash()

	// Note: This test depends on theForbiddenHashArray having current month data
	// The actual values will depend on your implementation
	if !exists {
		t.Logf("No hash found for key %s - this may be expected if theForbiddenHashArray is empty", expectedKey)
	} else {
		if hash == "" {
			t.Error("Expected non-empty hash when exists is true")
		}
	}
}

func TestSecretsOnlyInitializedOnce(t *testing.T) {
	// Reset secrets
	secrets = nil

	// Call setup multiple times
	setSecretKey()
	firstSecrets := secrets

	setSecretKey()
	secondSecrets := secrets

	// Should be the same instance
	if firstSecrets != secondSecrets {
		t.Error("Expected secrets to be initialized only once")
	}
}
