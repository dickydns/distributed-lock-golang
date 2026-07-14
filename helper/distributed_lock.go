package helper

import (
	"errors"
	"fmt"
	"time"
)

type DistributedLock struct {
	redis *RedisHelper
}

func NewDistributedLock(redis *RedisHelper) *DistributedLock {
	return &DistributedLock{
		redis: redis,
	}
}

func (d *DistributedLock) Execute(
	key string,
	ttl time.Duration,
	callback func() (interface{}, error),
) (interface{}, error) {

	lockKey := fmt.Sprintf("lock:%s", key)

	locked, err := d.redis.SetNX(
		lockKey,
		"processing",
		ttl,
	)

	if err != nil {
		return nil, err
	}

	if !locked {
		return nil, errors.New("resource is locked")
	}

	fmt.Println("LOCK ACQUIRED")

	// defer func() {
	// 	d.redis.Delete(lockKey)
	// 	fmt.Println("LOCK RELEASED")
	// }()

	return callback()
}