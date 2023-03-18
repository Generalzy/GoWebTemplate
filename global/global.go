package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GlobalLogger      *zap.Logger
	GlobalMysqlClient *gorm.DB
	GlobalRedisClient *redis.Client
)
