package db

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/conf"
	"github.com/go-redis/redis/v8"
)

func InitRedis(c *conf.RedisConf) *redis.Client {
	switch c.Mod {
	case conf.RedisSingle:
		return redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
			DB:   0,
		})
	case conf.RedisSentinel:
		return redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    c.MasterName,
			SentinelAddrs: c.Host,
		})
	}
	return nil
}
