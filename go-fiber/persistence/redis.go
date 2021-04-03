package persistence

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
)

func GetHashValue(hashKey string, fieldKey string) (string, error) {

	value, err := Redis.HGet(context.Background(), hashKey, fieldKey).Result()

	if err == redis.Nil || value == "" {
		return "", errors.New("No value found for the key")
	}

	return value, nil
}

func ToModel(model interface{}, jsonValue string) {
	json.Unmarshal([]byte(jsonValue), model)
}
