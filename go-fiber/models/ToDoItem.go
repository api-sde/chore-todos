package models

type ToDoItem struct {
	Title       string
	Description string
	IsCompleted bool

	OwnedBy string
}
