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
	result, err := s.Incr(ctx.Request().UserAgent()).Result()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, strconv.FormatInt(result, 10))
}

func NewServer(client *redis.Ring) Server {
	return &server{client}
}