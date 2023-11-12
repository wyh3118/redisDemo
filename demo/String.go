package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func String(client *redis.Client) {
	ctx := context.TODO()

	// set
	err := client.Set(ctx, "k1", "v2", 0).Err()
	if err == nil {
		fmt.Println("k1插入成功")
	}

	// get
	res, err := client.Get(ctx, "k1").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("get k1: " + res)
	}

	// GetAndSet
	res, err = client.GetSet(ctx, "k1", "v1-new").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("old value: " + res)
	}

	// Del
	num, err := client.Del(ctx, "k1").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除个数: " + strconv.Itoa(int(num)))
	}

	// MSet
	res, err = client.MSet(ctx, "k1", "v1", "k2", "v2", "k3", "v3").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// MGet
	resSlice, err := client.MGet(ctx, "k1", "k2", "k3").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resSlice)
	}

	// IncrBy
	err = client.Set(ctx, "number", 1, 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("set num ok")
	}

	num, err = client.IncrBy(ctx, "number", 2).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("after IncrBy: ", num)
	}

	// Expire
	err = client.Expire(ctx, "number", time.Second*20).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("expire number 20 second")
	}

	// Keys
	// 以下命令查找所有以k开头的key
	result, err := client.Keys(ctx, "k*").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	client.Del(ctx, "k1", "k2", "k3")
}
