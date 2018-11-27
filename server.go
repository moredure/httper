package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	Server interface {
		Index(echo.Context) error
	}
	server struct {
		*redis.Ring
	}
)

func (s *server) Index(ctx echo.Context) error {
	userAgent := ctx.Request().UserAgent()
	result, err := s.Incr(userAgent).Result()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, userAgent + ": " + strconv.FormatInt(result, 10))
}

func NewServer(client *redis.Ring) Server {
	return &server{client}
}