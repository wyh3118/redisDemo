package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Hash(client *redis.Client) {
	ctx := context.TODO()

	// HSet
	result, err := client.HSet(ctx, "user:1", "id", 1, "name", "wyh").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// HGet
	res, err := client.HGet(ctx, "user:1", "name").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// HMGet
	resSlice, err := client.HMGet(ctx, "user:1", "id", "name").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resSlice)
	}

	// HGetAll
	m, err := client.HGetAll(ctx, "user:1").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}

	// HIncrBy
	i, err := client.HIncrBy(ctx, "user:1", "id", 2).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("after IncrBy: ", i)
	}

	// HKeys
	strings, err := client.HKeys(ctx, "user:1").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strings)
	}

	// HLen
	i2, err := client.HLen(ctx, "user:1").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("filed num: ", i2)
	}

	// HMSet
	user2 := map[string]any{
		"id":      2,
		"name":    "wyhh",
		"address": "HUE",
	}
	b, err := client.HMSet(ctx, "user:2", user2).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}

	// HDel
	i3, err := client.HDel(ctx, "user:1", "id", "name").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete filed num:", i3)
	}

	// HExists
	b2, err := client.HExists(ctx, "user:2", "phone").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("exists: ", b2)
	}
}
