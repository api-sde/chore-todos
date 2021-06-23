package services

import (
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ValidatePassword(password string, currentHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currentHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func InvalidateToken(validationTrace string) bool {
	success := persistence.InsertInSet(store.InvalidatedTokens, validationTrace)
	persistence.SetExpiration(validationTrace, time.Duration(time.Duration.Hours(72)))
	return success
}

func VerifyTokenValidity(validationTrace string) bool {
	isPresent, err := persistence.GetAllSet(store.InvalidatedTokens)

	if err != nil {
		return true
	}

	// If the token trace is here, user has log out
	for _, invalidated := range isPresent {
		if (invalidated == validationTrace) {
			return false
		}
	}

	return true
}