package cache

import (
	"github.com/falcolee/transgo/common"
	"github.com/go-redis/redis/v8"
)

var Storage Cache

func InitCache(options *common.TLOptions) {
	Storage = NewInMemoryStorage()
	if options.UseCache {
		switch options.CacheStorage {
		case "file":
			Storage = NewFileStorage(options.StorageDir)
			return
		case "redis":
			rdb := redis.NewClient(&redis.Options{
				Addr:     options.TLConfig.Cache.RedisAddr,
				Password: options.TLConfig.Cache.RedisPassword,
				DB:       0, // use default DB
			})
			Storage = NewRedisStorage(rdb, options.TLConfig.Cache.RedisPrefix)
			return
		}
	}
}
