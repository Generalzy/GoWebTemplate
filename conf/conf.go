package conf

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	RedisSingle = iota + 1
	RedisSentinel
	RedisCluster
)

type ServiceConf struct {
	RedisConf  *RedisConf
	ServerConf *ServerConf
	MysqlConf  *MysqlConf
	LogConf    *LogConf
}

type ServerConf struct {
	Debug    bool
	Host     string
	Language string
	Port     int
}

type MysqlConf struct {
	Host     string
	Username string
	Password string
	DB       string
	Port     int
	Timeout  int
}

type RedisConf struct {
	Port       int
	Mod        int
	MasterName string
	Host       []string
}

type LogConf struct {
	Filename   string
	Level      uint8
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func InitConf() (*ServiceConf, error) {
	dir, _ := os.Getwd()
	viper.SetConfigFile(filepath.Join(dir, "service.toml"))

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		return nil, err
	}

	return &ServiceConf{
		RedisConf: &RedisConf{
			Port:       viper.GetInt("redis.port"),
			Host:       viper.GetStringSlice("redis.host"),
			Mod:        viper.GetInt("redis.model"),
			MasterName: viper.GetString("redis.mastername"),
		},
		ServerConf: &ServerConf{
			Debug:    viper.GetBool("server.debug"),
			Host:     viper.GetString("server.host"),
			Port:     viper.GetInt("server.port"),
			Language: viper.GetString("server.language"),
		},
		MysqlConf: &MysqlConf{
			Host:     viper.GetString("mysql.host"),
			Username: viper.GetString("mysql.username"),
			Password: viper.GetString("mysql.password"),
			DB:       viper.GetString("mysql.db"),
			Port:     viper.GetInt("mysql.port"),
			Timeout:  viper.GetInt("mysql.timeout"),
		},
		LogConf: &LogConf{
			Filename:   viper.GetString("log.filename"),
			Level:      uint8(viper.GetUint("log.level")),
			MaxSize:    viper.GetInt("log.maxsize"),
			MaxBackups: viper.GetInt("log.maxbackups"),
			MaxAge:     viper.GetInt("log.maxage"),
		},
	}, nil
}
