package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/onepaas/onepaas/pkg/viper"
	"github.com/rs/zerolog/log"
)

var ctx = context.Background()
var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
        Password: viper.GetString("redis.password"),
        DB:       viper.GetInt("redis.db"),
    })

    _, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Panic().Err(err).Msgf("OnePaaS can't ping the Redis service (%s:%d)", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	}
}

func GetRedis() *redis.Client {
	return redisClient
}
