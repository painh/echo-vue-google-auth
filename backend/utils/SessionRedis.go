package utils

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"painh.com/echo-vue-google-auth/config"
	"strconv"
	"time"
)

var RedisClient *redis.Client
var ttl int

func init() {
	cfg := config.Config
	redisConfig := cfg["SessionRedis"].(map[string]interface{})
	portInt := redisConfig["port"].(int)
	port := strconv.Itoa(portInt)
	ttl = redisConfig["ttl"].(int)

	// set default ttl

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig["host"].(string) + ":" + port,
		Password: redisConfig["password"].(string),
		DB:       int(redisConfig["db"].(int)),
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

func SessionRedisSet(key string, value interface{}) *redis.StatusCmd {
	expiration := time.Duration(ttl) * time.Second
	// value 를 json string으로 변환
	valueJson, _ := json.Marshal(value)

	cmd := RedisClient.Set(context.Background(), key, valueJson, expiration)
	return cmd
}

func SessionRedisGet(key string) *redis.StringCmd {
	cmd := RedisClient.Get(context.Background(), key)
	return cmd
}
