package controllers

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

func (server *Server) InitializeRedis() *Server{
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis service " + err.Error())
	}
	server.c = c
	return server
}

//GetKey get key
func(server *Server) getKey(key string, src interface{}) error{

	val, err := server.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil{
		return err
	}
	return nil
}

func (server *Server) setKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil{
		return err
	}
	err = server.c.Set(key, cacheEntry, expiration).Err()
	if err != nil{
		return err
	}
	return nil
}
//https://golangbyexample.com/golang-redis-client-example/