package eMongo

import (
	"github.com/mittacy/ego/library/log"
)

const (
	UriFormat    = "mongodb://%s:%d"       // 无加密
	UriPSWFormat = "mongodb://%s:%s@%s:%d" // 加密
)

type Conf struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

var (
	connectConf map[string]Conf // 连接配置池
	l           *log.Logger
)

// InitConnectConf 初始化连接配置
// Example:
// c = map[string]Conf{
//		"localhost": {
//			Host:     viper.GetString("MONGO_RW_HOST"),
//			Port:     viper.GetInt("MONGO_RW_PORT"),
//			Database: viper.GetString("MONGO_RW_DATABASE"),
//			User:     viper.GetString("MONGO_RW_USERNAME"),
//			Password: viper.GetString("MONGO_RW_PASSWORD"),
//		},
//	}
// @param c
func InitConnectConf(c map[string]Conf) {
	connectConf = c
}
