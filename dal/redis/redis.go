package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
)

func init() {
	if err := initClient(); err != nil {
		panic(err)
	}
	logrus.Info("redis 连接成功")
}

func initClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		logrus.Error("Redis 连接失败")
		return err
	}
	return nil
}
