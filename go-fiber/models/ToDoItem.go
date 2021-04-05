package models

import "time"

type ToDoItem struct {
	ItemId      string
	Title       string
	Description string
	IsCompleted bool

	CreatedBy   string
	OwnedBy     string
	GroupIds    []int

	CreationTime time.Time
	LastUpdateTime time.Time
}
