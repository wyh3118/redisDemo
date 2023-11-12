package Init

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.138.73.97:6379", // 数据库地址
		Username: "default",
		Password: "redis654321",
		DB:       0,
		PoolSize: 20,
	})

	_, err := rdb.Ping(context.TODO()).Result()
	if err == nil {
		fmt.Println("redis connect!")
		return rdb
	} else {
		log.Fatalln(err)
		return nil
	}
}
