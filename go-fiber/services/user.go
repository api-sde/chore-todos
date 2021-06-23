package services

import (
	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
)

func GetUserById(userId string) *models.User {
	userJson, _ := persistence.GetHashValue(store.Users, userId)
	// todo log error

	user := new(models.User)
	models.ToModel(user, userJson)

	return user
}
