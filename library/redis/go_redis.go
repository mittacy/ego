package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mittacy/ego/library/log"
	"sync"
	"time"
)

var ErrNoInit = errors.New("redis: please initialize with InitRedis() method")

var (
	isInit        bool                     // 是否初始化
	redisPool     map[string]*redis.Client // 单例池
	redisPoolLock sync.RWMutex             // 单例池锁
	redisLog      *log.Logger              // redis日志
)

func InitRedis(connectConf map[string]Conf) {
	// 初始化redis连接配置池
	InitRedisConf(connectConf)

	// 初始化单例池
	redisPool = make(map[string]*redis.Client, 0)
	redisPoolLock = sync.RWMutex{}

	// 初始化日志
	redisLog = log.New("redis")

	isInit = true
}

// GetRedis 获取redis连接
// @param name 配置名
// @param defaultDB 使用哪个数据库
// @return *redis.Client
func GetRedis(name string, defaultDB int) *redis.Client {
	if !isInit {
		panic(ErrNoInit)
	}

	cacheName := cachePoolName(name, defaultDB)

	redisPoolLock.RLock()
	if db, ok := redisPool[cacheName]; ok {
		redisPoolLock.RUnlock()
		return db
	}
	redisPoolLock.RUnlock()

	conf, isExist := connectConf[name]
	if !isExist {
		redisLog.Errorw("配置不存在", "name", name)
		return &redis.Client{}
	}

	db, err := NewRedisConnect(conf, defaultDB)
	if err != nil {
		redisLog.Errorw("连接数据库失败", "conf", conf, "err", err)
		return &redis.Client{}
	}

	redisPoolLock.Lock()
	redisPool[cacheName] = db
	redisPoolLock.Unlock()

	return db
}

// NewClient 获取新客户端
// @param conf 配置名
// @param db 使用哪个数据库
// @return *redis.Client
func NewRedisConnect(conf Conf, db int) (*redis.Client, error) {
	if !isInit {
		panic(ErrNoInit)
	}

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       db,
	}

	if conf.PoolSize > 0 { // 最大连接数
		options.PoolSize = conf.PoolSize
	}
	if conf.MinIdleConn > 0 { // 最小空闲连接数
		options.MinIdleConns = conf.MinIdleConn
	}
	if conf.IdleTimeout > 0 { // 空闲时间(秒)
		options.IdleTimeout = conf.IdleTimeout * time.Second
	}

	rdb := redis.NewClient(options)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return &redis.Client{}, err
	}

	return rdb, nil
}

// Redis 在结构体引入组合并赋值RedisConfName、RedisDB，即可通过Redis()获取redis连接
// Example
// type User struct {
// 	 Redis
// }
//
// var user = User{Redis{RedisConfName: "localhost", RedisDB:0}}
//
// func (u *User) GetUser(id int64) error {
// 	 u.GoRedis().Set(k, v)
// }
type GoRedis struct {
	RedisConfName string
	RedisDB       int
}

// Redis 获取redis连接
// @return *redis.Client
func (ctl *GoRedis) Redis() *redis.Client {
	return GetRedis(ctl.RedisConfName, ctl.RedisDB)
}

func cachePoolName(name string, db int) string {
	return fmt.Sprintf("%s:%d", name, db)
}
