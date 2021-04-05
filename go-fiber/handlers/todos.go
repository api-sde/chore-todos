package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllToDos(ctx *fiber.Ctx) error {

	toDoMap, err := persistence.GetAllHash(store.ToDos)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	toDoList := models.ToCollectionModel(models.ToDoItem{}, toDoMap)

	return ctx.JSON(toDoList)
}

func GetToDos(ctx *fiber.Ctx) error {

	id := ctx.Params("todoId")

	toDoJson, err := persistence.GetHashValue(store.ToDos, id)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	model := new(models.ToDoItem)
	models.ToModel(model, toDoJson)

	return ctx.JSON(model)
}

func CreateToDo(ctx *fiber.Ctx) error {

	/* Move this to auth handler, pass data to ctx.Locals(...) ? */
	authBearer := ctx.Get(fiber.HeaderAuthorization)
	jwtToken := strings.Fields(authBearer)[1]
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userid"].(string)
	// --

	newToDo := new(models.ToDoItem)
	if err := ctx.BodyParser(newToDo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "Couldn't parse to do", "error": err})
	}

	newToDo.ItemId = uuid.New().String()
	newToDo.CreatedBy = userId
	newToDo.OwnedBy = userId
	newToDo.CreationTime, newToDo.LastUpdateTime = time.Now(), time.Now()

	newToDoJson, err := json.Marshal(newToDo)

	if err != nil {
		return ctx.SendStatus(500)
	}

	persistence.InsertInHash(store.ToDos, newToDo.ItemId, newToDoJson)
	persistence.InsertInSet(store.ToDosByUserId+userId, newToDo.ItemId)

	return ctx.Status(fiber.StatusCreated).JSON(newToDo)
}
