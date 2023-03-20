package main

import (
	"fmt"
	"github.com/Generalzy/GeneralSaaS/conf"
	"github.com/Generalzy/GeneralSaaS/db"
	"github.com/Generalzy/GeneralSaaS/global"
	"github.com/Generalzy/GeneralSaaS/log"
	"github.com/Generalzy/GeneralSaaS/middleware"
	"github.com/Generalzy/GeneralSaaS/validators"
	"github.com/Generalzy/GeneralSaaS/web/urls"
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
		// 如果error不为nil则程序整体退出
		// InitMysql中泄露的协程也会被清除
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
	middleware.InitMiddleWare(engine)
	// urls
	urls.InitUrls(engine)
	// run
	runEnv := fmt.Sprintf("%s:%d", svcConfig.Host, svcConfig.Port)
	global.GlobalLogger.Info(runEnv + " 服务运行中....")
	err := engine.Run(runEnv)
	if err != nil {
		global.GlobalLogger.Error(err.Error())
		os.Exit(0)
	}
}

func main() {
	RunServer()
}
