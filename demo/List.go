package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func List(client *redis.Client) {
	ctx := context.TODO()

	// LPush
	num, err := client.LPush(ctx, "list", 1, 2, 3, 4, 5).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("LPush number: ", num)
	}

	// RPop
	result, err := client.RPop(ctx, "list").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("RPop: ", result)
	}

	// LRange
	strings, err := client.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("list: ", strings)
	}

	// LRem
	i, err := client.LRem(ctx, "list", 1, 5).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	// 当count为正数时，代表从左边删，删除几个
	// 当count为负数时，代表从右边删，删除几个
	// 当count为0时，代表删除所有值为value的元素

	// LIndex
	// 返回列表位置在index的元素
	s, err := client.LIndex(ctx, "list", 0).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
