package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {}

func (r Redis) Connect() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	return client
}

func(s Redis) Set(){
	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s Redis) Get(){
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
}
//https://golangbyexample.com/golang-redis-client-example/