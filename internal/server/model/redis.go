package model

import (
	"github.com/redis/go-redis/v9"
)

func InitRedisPool(addr string, password string, db int, poolSize int, minIdleConns int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
	})
	return rdb
}
