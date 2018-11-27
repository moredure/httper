package main

import "github.com/go-redis/redis"

func NewRedisOptions(env *Environment) (*redis.Options, error) {
	return redis.ParseURL(env.RedisURL)
}