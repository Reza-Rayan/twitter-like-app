package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Redis = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
