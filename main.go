package main

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/GeneralSaaS/urls"
	"github.com/Generalzy/GeneralSaaS/conf"
	"github.com/Generalzy/GeneralSaaS/db"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/Generalzy/GeneralSaaS/log"
	"github.com/Generalzy/GeneralSaaS/middleware"
	"github.com/Generalzy/GeneralSaaS/validators"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	engine    *gin.Engine
	svcConfig *conf.ServerConf
)

func init() {
	config, err := conf.InitConf()
	if err != nil {
		panic(err)
	}

	global.GlobalLogger, err = log.InitLogger(config.LogConf)
	if err != nil {
		global.GlobalLogger.Error(err.Error())
		os.Exit(1)
	} else {
		global.GlobalLogger.Info("初始化日志")
	}

	global.GlobalMysqlClient, err = db.InitMysql(config.MysqlConf)
	if err != nil {
		global.GlobalLogger.Error(err.Error())
		os.Exit(1)
	} else {
		global.GlobalLogger.Info("初始化mysql")
	}

	global.GlobalRedisClient = db.InitRedis(config.RedisConf)
	global.GlobalLogger.Info("初始化redis")

	svcConfig = config.ServerConf

	// validators
	global.GlobalTranslator, err = validators.InitTrans(svcConfig.Language)
	if err != nil {
		global.GlobalLogger.Error(err.Error())
		os.Exit(1)
	} else {
		global.GlobalLogger.Info("初始化翻译器")
	}
}

func RunServer() {
	switch svcConfig.Debug {
	case true:
		gin.SetMode(gin.DebugMode)
	case false:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	engine = gin.New()
	// middleware
	engine.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// urls
	urls.InitUrls(engine)

	err := engine.Run(fmt.Sprintf("%s:%d", svcConfig.Host, svcConfig.Port))
	if err != nil {
		global.GlobalLogger.Error(err.Error())
		os.Exit(0)
	}
}

func main() {
	RunServer()
}
