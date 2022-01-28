package redis

import "time"

type Conf struct {
	Host        string
	Password    string
	Port        int
	PoolSize    int
	MinIdleConn int
	IdleTimeout time.Duration
}

var connectConf map[string]Conf // 连接配置池

// InitMysqlConf 函数功能说明
// Example:
// c = map[string]Conf{
//		"localhost": {
//			Host:        viper.GetString("REDIS_LOCALHOST_RW_HOST"),
//			Password:    viper.GetString("REDIS_LOCALHOST_RW_PASSWORD"),
//			Port:        viper.GetInt("REDIS_LOCALHOST_RW_PORT"),
//			PoolSize:    viper.GetInt("REDIS_LOCALHOST_POOL_SIZE"),
//			MinIdleConn: viper.GetInt("REDIS_LOCALHOST_MIN_IDLE_CONN"),
//			IdleTimeout: viper.GetDuration("REDIS_LOCALHOST_IDLE_TIMEOUT"),
//		},
//	}
// @param c
func InitRedisConf(c map[string]Conf) {
	connectConf = c
}
