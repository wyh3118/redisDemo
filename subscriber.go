package main

import (
	"redisDemo/Init"
	"redisDemo/demo"
)

func main() {
	redis := Init.Redis()
	demo.Subscriber(redis, "channel")
}
