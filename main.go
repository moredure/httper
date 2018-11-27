package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(echo.New),
		fx.Provide(NewEnvironment),
		fx.Provide(NewRedisOptions),
		fx.Provide(redis.NewRing),
		fx.Provide(NewServer),
		fx.Invoke(Bootstrap),
	)
	app.Run()
}