package demo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
// 区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们。
// 这样做的好处是节省了执行命令的网络往返时间（RTT）。
func Pipeline(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*500)
	defer cancel()

	pipe := client.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_count")
	pipe.Expire(ctx, "pipeline_count", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(incr.Val())

	//上面的代码相当于将以下两个命令一次发给 Redis Server 端执行，与不使用 Pipeline 相比能减少一次RTT。
	//INCR pipeline_counter
	//EXPIRE pipeline_counts 3600
}

// Pipelined 或者，你也可以使用Pipelined 方法，它会在函数退出时调用 Exec。
func Pipelined(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*500)
	defer cancel()

	var incr *redis.IntCmd
	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipeline_count")
		pipe.Expire(ctx, "pipeline_count", time.Hour)
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(incr.Val())
}

// TxPipeline Redis是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，
// 例如在它们之间交替执行。但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
// 在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。
func TxPipeline(client *redis.Client) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*500)
	defer cancel()

	pipe := client.TxPipeline()
	incr := pipe.Incr(ctx, "pipeline_count")
	pipe.Expire(ctx, "pipeline_count", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(incr.Val())
	//上面代码相当于在一个RTT下执行了下面的redis命令：
	//MULTI
	//INCR pipeline_counter
	//EXPIRE pipeline_counts 3600
	//EXEC
}
