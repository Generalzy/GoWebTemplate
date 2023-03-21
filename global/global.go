package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GlobalLogger        *zap.Logger
	GlobalMysqlClient   *gorm.DB
	GlobalRedisClient   *redis.Client
	GlobalTranslator    ut.Translator
	GlobalMongoDBClient *mongo.Client
)
