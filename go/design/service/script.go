package service

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisConfig struct {
	Host     string
	Port     int
	Password string
}

func RedisClient() *redis.Client {
	config := redisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	options := redis.Options{
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           config.Password,
		DB:                 0,
		MaxRetries:         1,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	// 新建一个client
	client := redis.NewClient(&options)
	return client
}

func LuaScript() {
	ctx := context.Background()
	rdb := RedisClient()
	script := "return redis.call('GET', KEYS[1])"
	// aaa已经存储在redis中
	res, err := rdb.Eval(ctx, script, []string{"aaa"}).Result()
	if err != nil {
		fmt.Println("Err :", err)
	}

	fmt.Println("Result :", res)
}
