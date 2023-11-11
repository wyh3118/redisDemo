package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Subscriber(client *redis.Client, channel string) {
	subscriber := client.Subscribe(context.TODO(), channel)
	fmt.Println("Subscriber running!")
	for result := range subscriber.Channel() {
		fmt.Println("from: ", result.Channel)
		fmt.Println("message: ", result.Payload)
	}
}

func PSubscriber(client *redis.Client, channel string) {
	subscriber := client.PSubscribe(context.TODO(), fmt.Sprintf("%s-*", channel))
	fmt.Println("PSubscriber running!")
	for result := range subscriber.Channel() {
		fmt.Println("from: ", result.Channel)
		fmt.Println("message: ", result.Payload)
	}
}

func Publisher(client *redis.Client, channel string) {
	numSub, _ := client.PubSubNumSub(context.TODO(), channel).Result()
	fmt.Printf("订阅数量: %d\n", numSub[channel])
	for i := 0; i < 10; i++ {
		result, err := client.Publish(context.TODO(), channel, fmt.Sprintf("message-%d", i)).Result()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

func Publisher2(client *redis.Client, channel string) {
	for i := 0; i < 10; i++ {
		result, err := client.Publish(context.TODO(), fmt.Sprintf("%s-%d", channel, i), fmt.Sprintf("message-%d", i)).Result()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
