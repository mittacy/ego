package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

type Conf struct {
	Name         string
	Path         string          // 日志路径，默认为 .
	LowLevel     zapcore.Level   // 最低打印级别，默认为 debug
	LogInConsole bool            // 是否打印到控制台，默认为 false
	Fields       []zapcore.Field // 该日志都加的字段，默认为空
	IsJSONEncode bool            // 是否为json格式日志，默认为 true
	TimeFormat   string          // 时间格式，默认为 2006-01-02T15:04:05Z07:00
	PreName      string          // 日志前缀，默认为 biz_
}

func (c *Conf) LogPath() string {
	return fmt.Sprintf("%s/%s%s.log", c.Path, c.PreName, c.Name)
}

type ConfigOption func(conf *Conf)

// WithPath 设置日志路径, 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param path 路径
func WithPath(path string) ConfigOption {
	return func(conf *Conf) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0700); err != nil {
				log.Printf("log: create %s directory err: %s\n", path, err)
			}
		}

		conf.Path = path
	}
}

// WithLevel 设置服务记录的最低日志级别
// 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param l 日志级别(-1:debug、0:info、1:warn、2:error)
func WithLevel(l zapcore.Level) ConfigOption {
	return func(conf *Conf) {
		conf.LowLevel = l
	}
}

// WithLogInConsole 是否输出到控制台
// 修改后，新建的日志将会是新配置，已经建立的日志配置不变
// @param isLogInConsole
func WithLogInConsole(isLogInConsole bool) ConfigOption {
	return func(conf *Conf) {
		conf.LogInConsole = isLogInConsole
	}
}

// WithFields 添加全局日志的新字段, 新建的日志将会是新配置，已经建立的日志配置不变
// @param field 日志字段
func WithFields(fields ...zapcore.Field) ConfigOption {
	return func(conf *Conf) {
		conf.Fields = append(conf.Fields, fields...)
	}
}

// WithEncoderJSON 是否设置为json格式日志
// @param isJSONEncoder
// @return ConfigOption
func WithEncoderJSON(isJSONEncoder bool) ConfigOption {
	return func(conf *Conf) {
		conf.IsJSONEncode = isJSONEncoder
	}
}

// WithTimeFormat 设置时间格式
// @param format 时间格式
// @return ConfigOption
func WithTimeFormat(format string) ConfigOption {
	return func(conf *Conf) {
		conf.TimeFormat = format
	}
}

// WithPreName 设置日志前缀
// @param pre 前缀
// @return ConfigOption
func WithPreName(pre string) ConfigOption {
	return func(conf *Conf) {
		conf.PreName = pre
	}
}

// WithRequestIdKey 设置请求上下文请求id键名，记录日志时，将从上下文中取请求id并记录
// @param k 键名
// @return ConfigOption
func WithRequestIdKey(k string) ConfigOption {
	requestIdKey = k
	return func(conf *Conf) {
		return
	}
}

// defaultConf 全局配置，新的无配置日志默认使用该配置
var defaultConf = Conf{
	Path:         ".",
	LowLevel:     zapcore.DebugLevel,
	LogInConsole: false,
	Fields:       make([]zapcore.Field, 0),
	IsJSONEncode: true,
	TimeFormat:   time.RFC3339,
	PreName:      "biz_",
}

func getDefaultConf() Conf {
	fields := make([]zapcore.Field, len(defaultConf.Fields))
	for i, v := range defaultConf.Fields {
		fields[i].Key = v.Key
		fields[i].Type = v.Type
		fields[i].Integer = v.Integer
		fields[i].String = v.String
		fields[i].Interface = v.Interface
	}

	return Conf{
		Path:         defaultConf.Path,
		LowLevel:     defaultConf.LowLevel,
		LogInConsole: defaultConf.LogInConsole,
		Fields:       fields,
		IsJSONEncode: defaultConf.IsJSONEncode,
		TimeFormat:   defaultConf.TimeFormat,
		PreName:      defaultConf.PreName,
	}
}

// SetDefaultConf 设置默认日志配置
// @param options
func SetDefaultConf(options ...ConfigOption) {
	for _, option := range options {
		option(&defaultConf)
	}
}
