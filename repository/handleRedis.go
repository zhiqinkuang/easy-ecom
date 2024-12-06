package repository

import (
	"github.com/redis/go-redis/v9"
	"time"
)

// Set 设置 Redis 键值对
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(rctx, key, value, expiration).Err()
}

// Get 获取 Redis 键值
func Get(key string) (string, error) {
	return rdb.Get(rctx, key).Result()
}

// ZAdd 添加元素到 ZSET
func ZAdd(key string, score float64, member interface{}) error {
	return rdb.ZAdd(rctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// ZRange 获取 ZSET 范围内的元素
func ZRange(key string, start, stop int64) ([]redis.Z, error) {
	return rdb.ZRangeWithScores(rctx, key, start, stop).Result()
}

// Delete 删除 Redis 键
func Delete(key string) error {
	return rdb.Del(rctx, key).Err()
}

// SetExpire 设置 Redis 键的过期时间
func SetExpire(key string, expiration time.Duration) error {
	return rdb.Expire(rctx, key, expiration).Err()
}
