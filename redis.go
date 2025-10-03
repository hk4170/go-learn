package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main(){
	redis := redis.NewClient(&redis.Options{
		Addr : "localhost:6379",
		Password: "",
		DB: 0,
	})
	if redis.Ping().Val() != "PONG" {
		log.Fatal("redis not run")
	}
	// 写入
	redis.Set("name","djdjd",0)
	// 读取
	println((redis.Get("name").String()))
	//删除
	redis.Del("name")
	redis.Append("comment","hello")
	redis.Set("server","192.168.1.1",0)
	redis.Set("proxy","192.168.1.2",0)
	redis.Set("port","8080",0)

	println(redis.Get("comment").Val())
	server := redis.Get("server").Val()
	proxy := redis.Get("proxy").Val()
	port := redis.Get("port").Val()
	println(server,proxy,port)
}