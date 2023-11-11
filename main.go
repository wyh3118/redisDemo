package main

import (
	"redisDemo/Init"
	"redisDemo/demo"
)

func main() {
	redis := Init.Redis()
	//demo.Func1(redis)
	//demo.String(redis)
	//demo.Hash(redis)
	//demo.List(redis)
	//demo.Set(redis)
	demo.SortedSet(redis)
}
