package utils

import (
	"context"
	"time"
)

const CacheTime = time.Minute * 1 // 20 minutes

func SetCache(key string, data string, ttl time.Duration) error {
	return Redis.Set(context.Background(), key, data, ttl).Err()
}

func GetCache(key string) (string, error) {
	return Redis.Get(context.Background(), key).Result()
}

func DeleteCache(key string) error {
	return Redis.Del(context.Background(), key).Err()
}

func ClearPostsCache() error {
	keys, err := Redis.Keys(Ctx, "post:*").Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		return Redis.Del(Ctx, keys...).Err()
	}
	return nil
}
