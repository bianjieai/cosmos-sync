package cache

import (
	"context"
	"github.com/bianjieai/cosmos-sync/config"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func Init(conf *config.Config) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addrs,
		Username: conf.Redis.User,
		Password: conf.Redis.Password,
		DB:       conf.Redis.Db,
	})

	res := rdb.Ping(context.Background())
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func Close() {
	_ = rdb.Close()
}
