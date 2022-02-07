package eMysql

import "time"

type gormConf struct {
	LogName                 string
	LogSlowThreshold        time.Duration
	LogIgnoreRecordNotFound bool
}

func newGormConf(options ...GormConfigOption) gormConf {
	c := gormConf{
		LogName:                 "gorm",
		LogSlowThreshold:        100 * time.Millisecond,
		LogIgnoreRecordNotFound: false,
	}

	for _, option := range options {
		option(&c)
	}

	return c
}

type GormConfigOption func(conf *gormConf)

func WithLogName(name string) GormConfigOption {
	return func(conf *gormConf) {
		conf.LogName = name
	}
}

func WithLogSlowThreshold(duration time.Duration) GormConfigOption {
	return func(conf *gormConf) {
		conf.LogSlowThreshold = duration
	}
}

func WithLogIgnoreRecordNotFound(isIgnore bool) GormConfigOption {
	return func(conf *gormConf) {
		conf.LogIgnoreRecordNotFound = isIgnore
	}
}
