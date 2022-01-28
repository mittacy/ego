package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/mittacy/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
)

var ErrNoInit = errors.New("gorm: please initialize with InitGorm() method")

var (
	isInit   bool                // 是否初始化
	gormPool map[string]*gorm.DB // 单例池
	gormLog  zapgorm2.Logger     // gorm慢日志
)

// InitGorm 函数功能说明
// @param connectConf
// Example:
// connectConf = map[string]Conf{
//		"localhost": {
//			Host:     viper.GetString("DB_CORE_RW_HOST"),
//			Port:     viper.GetInt("DB_CORE_RW_PORT"),
//			Database: viper.GetString("DB_CORE_RW_DATABASE"),
//			User:     viper.GetString("DB_CORE_RW_USERNAME"),
//			Password: viper.GetString("DB_CORE_RW_PASSWORD"),
//		},
//	}
func InitGorm(connectConf map[string]Conf) {
	// 初始化mysql连接配置
	InitMysqlConf(connectConf)

	// 初始化单例池、日志
	gormPool = make(map[string]*gorm.DB, 0)
	l := log.New("gorm")
	gormLog = zapgorm2.New(l.GetZap())

	if viper.GetDuration("GORM_SLOW_LOG_THRESHOLD") == 0 {
		gormLog.SlowThreshold = time.Millisecond * 100
	} else {
		gormLog.SlowThreshold = viper.GetDuration("GORM_SLOW_LOG_THRESHOLD") * time.Millisecond
	}
	gormLog.LogLevel = logger.Info
	gormLog.IgnoreRecordNotFoundError = true
	gormLog.SetAsDefault()

	isInit = true
}

// GetGorm 获取gorm连接
// @param name 配置名
// @return *gorm.DB
func GetGorm(name string) *gorm.DB {
	if db, ok := gormPool[name]; ok {
		return db
	}

	conf, isExist := connectConf[name]
	if !isExist {
		gormLog.Error(context.Background(), fmt.Sprintf("配置不存在, conf name: %s", name))
		return &gorm.DB{}
	}

	db, err := NewGormConnect(conf)
	if err != nil {
		gormLog.Error(context.Background(), fmt.Sprintf("连接数据库失败, conf: %+v", conf))
		return &gorm.DB{}
	}

	gormPool[name] = db

	return db
}

// NewGormConnect 获取新客户端
// @param conf 配置信息
// @return *gorm.DB gorm连接
// @return error
func NewGormConnect(conf Conf) (*gorm.DB, error) {
	if !isInit {
		panic(ErrNoInit)
	}

	dsn := fmt.Sprintf(dbDSNFormat, conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	if conf.Params != "" {
		dsn = fmt.Sprintf("%s&%s", dsn, conf.Params)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 是否禁用表名复数形式
		Logger:         gormLog,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Gorm 在结构体引入组合并赋值ConfName，即可通过DB()获取gorm连接
// Example
// type User struct {
// 	 Gorm
// }
//
// var user = User{Gorm{ConfName: "localhost"}}
//
// func (u *User) GetUser(id int64) error {
// 	 u.DB().Where("id = ?", id).First()
// }
type Gorm struct {
	ConfName string
}

// DB 获取DB连接
// @return *gorm.DB
func (ctl *Gorm) DB() *gorm.DB {
	return GetGorm(ctl.ConfName)
}
