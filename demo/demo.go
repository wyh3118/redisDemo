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

func ScanKeyDemo1(client *redis.Client) {
	ctx := context.TODO()
	var cursor uint64
	for {
		var keys []string
		var err error
		// 将redis中所有以prefix:为前缀的key都扫描出来
		keys, cursor, err = client.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

// ScanKeysDemo2 针对这种需要遍历大量key的场景，go-redis中提供了一个简化方法——Iterator，其使用示例如下。
func ScanKeysDemo2(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 按前缀扫描key
	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	if iter.Err() != nil {
		panic(iter.Err())
	}
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	//此外，对于 Redis 中的 set、hash、zset 数据类型，go-redis 也支持类似的遍历方法。

	//iter := client.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
	//iter := client.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
	//iter := client.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()
}

// DelKeysByMatch 删除匹配的字段
func DelKeysByMatch(client *redis.Client, match string) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*500)
	defer cancel()

	iterator := client.Scan(ctx, 0, match, 0).Iterator()
	for iterator.Next(ctx) {
		result, err := client.Del(ctx, iterator.Val()).Result()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("delete: ", result)
		}
	}

	if iterator.Err() != nil {
		fmt.Println(iterator.Err())
	}
}
