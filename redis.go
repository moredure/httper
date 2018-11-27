package main

import "github.com/go-redis/redis"

func NewRedisOptions(env *Environment) (*redis.RingOptions, error) {
	redisUrl, err := redis.ParseURL(env.RedisURL)
	if err != nil {
		return nil, err
	}
	redisCloudUrl, err := redis.ParseURL(env.RedisCloudURL)
	if err != nil {
		return nil, err
	}
	return &redis.RingOptions{
		Addrs: map[string]string{
			"redis_cloud_url": redisCloudUrl.Addr,
			"redis_url": redisUrl.Addr,
		},
		DB: redisCloudUrl.DB,
		Password: redisCloudUrl.Password,
	}, nil
}