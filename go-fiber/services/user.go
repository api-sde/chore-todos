package services

import (
	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
)

func GetUserById(userId string) (*models.User, error) {
	userJson, err := persistence.GetHashValue(store.Users, userId)

	if err != nil {
		return nil, err
	}

	user := new(models.User)
	models.ToModel(user, userJson)

	return user, nil
}

func GetUserByEmail(userEmail string) (*models.User, error) {
	userId, err := persistence.GetHashValue(store.Users, userEmail)
	userJson, err := persistence.GetHashValue(store.Users, userId)

	if err != nil {
		return nil, err
	}

	user := new(models.User)
	models.ToModel(user, userJson)

	return user, nil
}

func IsUserExisting(userEmail string) bool {
	return persistence.ExistInHash(store.Users, userEmail)
}

func IsUserExistingById(userId string) bool {
	return persistence.ExistInHash(store.Users, userId)
}
