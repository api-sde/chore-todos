package services

import "golang.org/x/crypto/bcrypt"

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
