package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetHashValue(hashKey string, fieldKey string) (string, error) {

	value, err := Redis.HGet(context.Background(), hashKey, fieldKey).Result()

	if err == redis.Nil || value == "" {
		return "", errors.New("No value found for the key")
	}

	return value, nil
}

func GetAllHash(hashKey string) (map[string]string, error) {

	allHashMap, err := Redis.HGetAll(context.Background(), hashKey).Result()

	if err != nil {
		return nil, errors.New("No value found for the key")
	}

	return allHashMap, nil
}

func ExistInHash(hashKey string, fieldKey string) bool {

	exist, err := Redis.HExists(context.Background(), hashKey, fieldKey).Result()

	if err != nil {
		panic(errors.New("Failure while trying to verify " + fieldKey + "existence into " + hashKey))
	}

	return exist
}

func InsertInHash(hashKey string, values ...interface{}) bool {

	_, err := Redis.HSet(context.Background(), hashKey, values).Result()

	if err != nil {
		//fmt.Log (errors.New("Failure while trying to insert in hash " + hashKey))

		return false
	}

	// TODO: chan -> send to pgsql service

	return true
}

func InsertInSet(setKey string, setValues ...interface{}) bool {

	_, err := Redis.SAdd(context.Background(), setKey, setValues).Result()

	if err != nil {
		//log
		return false
	}

	return true
}

func SetExpiration(setKey string, expiration time.Duration) {
	_, err := Redis.Expire(context.Background(), setKey, expiration).Result()

	if err != nil {
		return
	}
}

func GetAllSet(setKey string) ([]string, error) {

	setValues, err := Redis.SMembers(context.Background(), setKey).Result()

	if err != nil {
		return nil, errors.New("No value found for the key")
	}

	return setValues, nil
}
/// ToDos: HMGET / HSET / HGET / SADD / SMEMBERS
