package redis

import (
	"context"
	"os"

	"template/service/logger"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var RedisPool *redis.Client

func Init() {
	// Init Redis Config
	password := viper.GetString("Redis.Password")
	host := viper.GetString("Redis.Host")
	port := viper.GetString("Redis.Port")
	maxConnection := viper.GetInt("Redis.MaxConnection")

	// Init Redis Pool
	logger.Logger.Info("Redis pool is starting")
	RedisPool = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
		PoolSize: maxConnection,
	})

	// Test Redis Connection
	_, err := RedisPool.Ping(context.Background()).Result()
	if err != nil {
		logger.Logger.Errorf("Unable to connect to redis: %s\n", err)
		os.Exit(1)
	}
}

func ShutDown() {
	logger.Logger.Infof("Redis pool is shutting down")
	if RedisPool != nil {
		RedisPool.Close()
	}
}
