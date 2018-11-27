package main

import "github.com/caarlos0/env"

type Environment struct {
	RedisURL string `env:"REDIS_URL"`
	RedisCloudURL string `env:"REDISCLOUD_URL"`
	Port string `env:"PORT"`
}

func NewEnvironment() (*Environment, error) {
	cfg := new(Environment)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}