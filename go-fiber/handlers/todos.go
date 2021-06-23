package handlers

import (
	"encoding/json"
	"time"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
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

func GetToDoById(ctx *fiber.Ctx) error {

	id := ctx.Params("todoId")

	toDoJson, err := persistence.GetHashValue(store.ToDos, id)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	model := new(models.ToDoItem)
	models.ToModel(model, toDoJson)

	return ctx.JSON(model)
}

func GetToDoByUser(ctx *fiber.Ctx) error {

	userId := ctx.Locals("LoggedUserId").(string)

	toDoIds, err := persistence.GetAllSet(store.ToDosByUserId + userId)
	allToDoMap, err := persistence.GetAllHash(store.ToDos)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	// Move into service
	var userToDosJson = make(map[string]string)

	for _, toDoId := range toDoIds {
		toDoItem := allToDoMap[toDoId]

		userToDosJson[toDoId] = toDoItem
	}

	userToDoList := models.ToCollectionModel(models.ToDoItem{}, userToDosJson)

	return ctx.JSON(userToDoList)
}

func CreateToDo(ctx *fiber.Ctx) error {

	userId := ctx.Locals("CurrentUserId").(string)

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
