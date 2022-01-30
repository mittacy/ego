package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

type Logger struct {
	l    *zap.Logger
	name string
}

func (l *Logger) GetZap() *zap.Logger {
	return l.l
}

func (l *Logger) GetName() string {
	return l.name
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var logPool = map[string]*Logger{}

// New 创建新日志文件句柄，使用默认配置
// @param name 日志名
// @param options 日志配置，将覆盖默认配置
// @return *Logger
func New(name string, options ...ConfigOption) *Logger {
	name = strings.TrimSpace(name)
	if name == "" {
		log.Println("create: the input log name cannot empty")
		return nil
	}

	// 检查共用同名日志句柄
	if l, ok := logPool[name]; ok && l != nil {
		return l
	}

	// 创建新日志
	c := getDefaultConf()
	for _, option := range options {
		option(&c)
	}
	c.Name = name

	l := newWithConf(c)
	logPool[name] = l
	return l
}

func newWithConf(conf Conf) *Logger {
	// 打开日志文件
	writer, err := os.OpenFile(conf.LogPath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:     "log_at",
		LevelKey:    "level",
		NameKey:     "logger",
		MessageKey:  "context",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(conf.TimeFormat))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 是否为JSON格式日志
	var encoder zapcore.Encoder
	if conf.IsJSONEncode {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var syncer zapcore.WriteSyncer
	// 是否输出控制台
	if conf.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer))
	} else {
		syncer = zapcore.AddSync(writer)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		conf.LowLevel,
	)

	// 全局字段添加到每个日志中
	l := &Logger{
		l:    zap.New(core, zap.Fields(conf.Fields...)),
		name: conf.Name,
	}

	return l
}
