package helper

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisHelper struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisHelper() *RedisHelper {
	client := redis.NewClient(&redis.Options{
		Addr:fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:0,
	})
	ctx := context.Background()

	return &RedisHelper{
		Client: client,
		Ctx:    ctx,
	}
}

//set redis data
func (r *RedisHelper) Set(key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, ttl).Err()
}

//get redis data
func (r *RedisHelper) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

//delete redis data
func (r *RedisHelper) Delete(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

//check connection
func (r *RedisHelper) Ping() error {
	_, err := r.Client.Ping(r.Ctx).Result()
	return err
}

//Set nx
func (r *RedisHelper) SetNX(
	key string,
	value interface{},
	ttl time.Duration,
) (bool, error) {
	return r.Client.SetNX(
		r.Ctx,
		key,
		value,
		ttl,
	).Result()
}
