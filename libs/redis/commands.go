package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var ErrNoEntry = errors.New("redis: entry does not exit")

const (
	Timeout = 3 // Seconds
)

func (r *Redis) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	str, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return ErrNoEntry
	}
	if err2 := json.Unmarshal(([]byte)(str), value); err2 != nil {
		return err2
	}
	return nil
}

func (r *Redis) Set(key string, value interface{}, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	byteData, _ := json.Marshal(value)
	err := r.Client.Set(ctx, key, byteData, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Delete(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	if err := r.Client.Del(ctx, keys...).Err(); err != nil {
		return err
	}
	return nil

}
