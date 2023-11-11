package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func SortedSet(client *redis.Client) {
	ctx := context.TODO()

	// ZAdd
	result, err := client.ZAdd(ctx, "ZSet", redis.Z{
		Score:  5,
		Member: "wyh",
	}).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add: ", result)
	}

	//ZCard
	i, err := client.ZCard(ctx, "ZSet").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ZSet size: ", i)
	}

	// ZCount
	// 返回score>=1&&score<=5的元素的数量
	i2, err := client.ZCount(ctx, "ZSet", "1", "5").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1<=score<=5 count: ", i2)
	}

	// ZIncrBy
	// 添加某一元素的分数
	f, err := client.ZIncrBy(ctx, "ZSet", -3, "wyh").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("after score: ", f)
	}

	// ZRange
	// 返回索引范围内的元素，根据分数从小到大排序

	// ZRevRange
	// 返回索引范围内的元素，根据分数从大到小排序

	// ZRangeByScore
	// 根据分数段获得元素
	strings, err := client.ZRangeByScore(ctx, "ZSet", &redis.ZRangeBy{
		Min:    "1",
		Max:    "10",
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strings)
	}

	// ZRangeByScoreWithScore
	// 用法与ZRangeByScore相同，区别是会返回score

	// ZRem
	// 删除集合中的元素

	// ZRemRangeByRank
	// 根据索引删除元素
	// 0号位置的元素代表最低分，-1号位置的的元素代表最高分

	// ZRemRangeByScore
	// 根据score范围删除元素

	// ZScore
	// 根据元素值查找元素的score

	// ZRank
	// 根据元素值查找元素在集合中的排名
}
