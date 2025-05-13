package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 创建Redis连接池
	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // Redis服务器地址
		Password:     "",               // Redis密码，若没有则留空
		DB:           0,                // 选择要使用的数据库
		PoolSize:     100,              // 连接池的最大连接数
		MinIdleConns: 10,               // 最小空闲连接数
	})

	// 测试连接
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接Redis失败: %v", err)
	}
	fmt.Println("连接成功:", pong)

	// 设置键值对
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatalf("设置键值对失败: %v", err)
	}

	// 获取键值对
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatalf("获取键值对失败: %v", err)
	}
	fmt.Println("key:", val)

	// // 处理键不存在的情况
	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("键key2不存在")
	// } else if err != nil {
	// 	log.Fatalf("获取键值对失败: %v", err)
	// } else {
	// 	fmt.Println("key2:", val2)
	// }
}
