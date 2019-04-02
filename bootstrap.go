package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"go.uber.org/multierr"
)

func Bootstrap(lc fx.Lifecycle, e *echo.Echo, server Server, env *Environment, ring *redis.Ring) {
	e.GET("/", server.Index)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := e.Start(":" + env.Port); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return multierr.Combine(e.Shutdown(ctx), ring.Close())
		},
	})
}