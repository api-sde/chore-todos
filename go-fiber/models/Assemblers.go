package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
)

func TypeName(model interface{}) string {
	return reflect.TypeOf(model).Name()
}

func ToModel(model interface{}, jsonValue string) interface{} {
	json.Unmarshal([]byte(jsonValue), model)

	return model
}

func ToCollectionModel(modelTarget interface{}, jsonByUUID map[string]string) interface{} {
	var collectionResult []interface{}
	
	for key, json := range jsonByUUID {

		_, err := uuid.Parse(key)

		if err != nil {
			continue
		}

		var modelType = GetNewModelType(modelTarget)
		ToModel(modelType, json)
		collectionResult = append(collectionResult, modelType)
	}

	return collectionResult
}

func GetNewModelType(modelType interface{}) interface{} {

	var newInstance interface{}

	switch TypeName(modelType) {

	case TypeName(User{}):
		newInstance = new(User)

	case TypeName(ToDoItem{}):
		newInstance = new(ToDoItem)

	default:
		newInstance = nil
	}

	return newInstance
}
