package demo

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Func1(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接获取命令结果
	res, err := client.Get(ctx, "k1").Result()
	fmt.Println(res, err)

	// 先获取命令对象
	cmd := client.Get(ctx, "k1")
	fmt.Println(cmd.Val())
	fmt.Println(cmd.Err())

	// 直接获得命令的值
	res2 := client.Get(ctx, "k2").Val()
	if res2 == "" {
		fmt.Println("空")
	} else {
		fmt.Println(res2)
	}
}

func Func2(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接运行redis任意命令
	fmt.Println(client.Do(ctx, "keys", "*").Val())

	client.Set(ctx, "k1", "v1", time.Minute)
}

func Func3(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 用redis.nil来表示key不存在的错误
	res, err := client.Get(ctx, "backup1").Result()
	if err != nil {
		// key值不存在
		if errors.Is(err, redis.Nil) {
			fmt.Println("key值不存在")
		} else {
			// 其它错误
			fmt.Println("其它错误")
		}
	} else {
		fmt.Println(res)
	}
}
