package db

import (
	"cloud/common/config"
	"cloud/common/logger"
	"github.com/go-redis/redis"
)

var client *redis.Client

func GetRedisClient() error {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     config.GetConf().Db.Addr,
			Password: config.GetConf().Db.Password,
			DB:       config.GetConf().Db.Database,
		})
	}

	_, err := client.Ping().Result()
	if err != nil {
		logger.Error("fail to connect redis, error:", err)
		return err
	}

	logger.Info("succeed connect to redis")
	return nil
}

func GetClient() *redis.Client {
	if client == nil {
		GetRedisClient()
	}
	return client
}

type RedisScanTriple struct {
	Cursor uint64
	Err    error
	Keys   []string
}

func (triple *RedisScanTriple) Reset() {
	triple.Cursor = 0
	triple.Keys = []string{}
	triple.Err = nil
}

func HGet(key, field string) (res string, err error) {
	return client.HGet(key, field).Result()
}
