package main

import "github.com/go-redis/redis"

func NewRedisOptions(env *Environment) (*redis.RingOptions, error) {
	redisToGoUrl, err := redis.ParseURL(env.RedisToGoURL)
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
			"redis_to_go_url": redisToGoUrl.Addr,
		},
		Password: redisCloudUrl.Password,
	}, nil
}