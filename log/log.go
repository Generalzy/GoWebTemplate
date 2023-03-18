package log

import (
	"github.com/Generalzy/GeneralSaaS/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

// InitLogger 初始化Logger
func InitLogger(c *conf.LogConf) (*zap.Logger, error) {
	writeSyncer := getLogWriter(c.Filename, c.MaxSize, c.MaxBackups, c.MaxAge)
	encoder := getEncoder()

	// 设置日志等级
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.Level(c.Level))

	lg := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return lg, nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   true,
	}

	// 终端和文件日志
	writer := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(writer)
}
