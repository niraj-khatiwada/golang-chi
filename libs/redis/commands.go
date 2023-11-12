package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrNoEntry = errors.New("redis: entry does not exit")

const (
	Timeout = 3 // Seconds
)

func (r *Redis) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	keysWithNamespace := r.generateKeys(key)
	str, err := r.Client.Get(ctx, keysWithNamespace[0]).Result()
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
	keysWithNamespace := r.generateKeys(key)
	err := r.Client.Set(ctx, keysWithNamespace[0], byteData, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Delete(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	keysWithNamespace := r.generateKeys(keys...)
	if err := r.Client.Del(ctx, keysWithNamespace...).Err(); err != nil {
		return err
	}
	return nil

}

func (r *Redis) generateKeys(keys ...string) []string {
	var keysWithNamespace []string
	for _, key := range keys {
		keysWithNamespace = append(keysWithNamespace, fmt.Sprintf("%s:%s", r.GlobalNamespace, key))
	}
	fmt.Println(keysWithNamespace)
	return keysWithNamespace
}
