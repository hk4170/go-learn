package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main(){
	redis := redis.NewClient(&redis.Options{
		Addr : "localhost:6379",
		Password: "",
		DB: 0,
	})
	if redis.Ping().Val() != "PONG" {
		fmt.Printf("error: %v\n")
	}
	// 写入
	redis.Set("name","gavin",0)
	// 读取
	println((redis.Get("name").String()))
	//删除
	redis.Del("name")
	redis.Append("comment","hello")
	println(redis.Get("comment").String())
}