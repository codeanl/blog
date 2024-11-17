package cache

import (
	"time"

	"github.com/go-redis/redis"
)

func Set(rdb *redis.Client, key string, value interface{}, expire time.Duration) error {
	return rdb.Set(key, value, expire).Err()
}

func Get(rdb *redis.Client, key string) string {
	return rdb.Get(key).Val()
}
