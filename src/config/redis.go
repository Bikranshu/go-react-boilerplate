package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func RedisClient() *redis.Client {
	var redisPrefix = "redis"
	var redisHost = viper.GetString(redisPrefix + ".host")
	var redisPort = viper.GetString(redisPrefix + ".port")
	var redisPassword = viper.GetString(redisPrefix + ".password")
	var redisDB = viper.GetInt(redisPrefix + ".db")

	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})
}
