package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Set(client *redis.Client) {
	ctx := context.TODO()

	// SAdd
	result, err := client.SAdd(ctx, "set", 1, 2, 3, 4).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// SCard
	// 统计set中的元素个数
	i, err := client.SCard(ctx, "set").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("set count: ", i)
	}

	// SIsMember
	// 判断是否存在

	// SMember
	// 获取集合中所有元素

	// SRem
}
