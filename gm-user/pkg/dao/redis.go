package dao

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var Rdb *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	opt := redis.Options{
		Addr:     "192.168.254.128:16379",
		DB:       0,
		Username: "",
		Password: "milowang123",
		//PoolSize:        ,
		//MinIdleConns:    cfg.MaxIdle,
		//ConnMaxIdleTime: time.Second * time.Duration(cfg.MaxIdle),
		//ConnMaxLifetime: time.Second * time.Duration(cfg.MaxActive),
	}

	cli := redis.NewClient(&opt)

	Rdb = &RedisCache{
		rdb: cli,
	}
}

func (r *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	return r.rdb.Set(ctx, key, value, expire).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}
