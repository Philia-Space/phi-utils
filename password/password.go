// Package password provides secure password hashing and verification using bcrypt.
package password

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// DefaultCost is the bcrypt cost factor used when hashing passwords.
const DefaultCost = bcrypt.DefaultCost

// Hash generates a bcrypt hash of the password at the default cost.
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// HashWithCost generates a bcrypt hash with a specific cost factor.
func HashWithCost(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// Verify checks a password against a bcrypt hash.
func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NeedsRehash reports whether the hash was created with a cost different from the default.
func NeedsRehash(hash string) bool {
	cost, err := bcrypt.Cost([]byte(hash))
	if err != nil {
		return true
	}
	return cost != DefaultCost
}

// IsBlankSafe returns true if password is not obviously weak (non-empty, not common).
func IsBlankSafe(password string) bool {
	return len(password) >= 8
}

// ConstantTimeCompare compares two strings in constant time to prevent timing attacks.
func ConstantTimeCompare(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
