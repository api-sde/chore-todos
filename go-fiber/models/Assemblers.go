package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

func ToModel(model interface{}, jsonValue string) {
	json.Unmarshal([]byte(jsonValue), model)
}

func ToCollectionModel(model interface{}, jsonByUUID map[string]string) interface{} {
	var collectionResult []interface{}
	
	for key, json := range jsonByUUID {

		_, err := uuid.Parse(key)

		if err != nil {
			continue
		}

		var modelType = GetModelType(model)
		ToModel(modelType, json)
		collectionResult = append(collectionResult, modelType)
	}

	return collectionResult
}

func GetModelType(model interface{}) interface{} {

	var modelType interface{}

	switch model.(type) {
	case User:
		modelType = new(User)
	case ToDoItem:
		modelType = new(ToDoItem)
	default:
		modelType = nil
	}

	return modelType

}
