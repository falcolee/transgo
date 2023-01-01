package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
	prefix string
}

func NewRedisStorage(client *redis.Client, prefix string) *RedisStorage {
	return &RedisStorage{
		client: client,
	}
}

func (i *RedisStorage) Store(id string, content []byte) error {
	return i.client.Set(context.Background(), i.prefix+id, string(content), redis.KeepTTL).Err()
}

func (i *RedisStorage) FetchOne(id string) ([]byte, error) {
	data := i.client.Get(context.Background(), i.prefix+id)
	if data.Err() != nil {
		return nil, data.Err()
	} else {
		return []byte(data.Val()), nil
	}
}

func (i *RedisStorage) KeyExists(id string) (bool, error) {
	result, err := i.client.Exists(context.Background(), i.prefix+id).Result()
	if result == 1 {
		return true, nil
	}
	return false, err
}

var _ Cache = (*RedisStorage)(nil)
