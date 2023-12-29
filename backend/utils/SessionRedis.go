package utils

import (
	"context"
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
	cmd := RedisClient.Set(context.Background(), key, value, expiration)
	return cmd
}
