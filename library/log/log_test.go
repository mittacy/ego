package log

import (
	"context"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLog(t *testing.T) {
	bizLog := New("default")

	bizLog.Debug("this is SugarDebug")
	bizLog.Info("this is SugarInfo")
	bizLog.Warn("this is SugarWarn")
	bizLog.Error("this is SugarError")

	bizLog.Debugf("this is %s", "Debugf")
	bizLog.Infof("this is %s", "Infof")
	bizLog.Warnf("this is %s", "Warn")
	bizLog.Errorf("this is %s", "Errorf")

	bizLog.Debugw("this is Debugw", "k", "Debugw")
	bizLog.Infow("this is Infow", "k", "Infow")
	bizLog.Warnw("this is Warnw", "k", "Warnw")
	bizLog.Errorw("this is Errorw", "k", "Errorw")
}

func TestConf(t *testing.T) {
	field := zapcore.Field{
		Key:    "module_name",
		Type:   zapcore.StringType,
		String: "serverName",
	}
	bizLog := New("default",
		WithPath("./storage/logs"),
		WithTimeFormat("2006-01-02 15:04:05"),
		WithLevel(zapcore.InfoLevel),
		WithPreName("biz_"),
		WithEncoderJSON(false),
		WithFields(field))

	bizLog.Debug("this is SugarDebug")
	bizLog.Info("this is SugarInfo")
	bizLog.Warn("this is SugarWarn")
	bizLog.Error("this is SugarError")

	bizLog.Debugf("this is %s", "Debugf")
	bizLog.Infof("this is %s", "Infof")
	bizLog.Warnf("this is %s", "Warn")
	bizLog.Errorf("this is %s", "Errorf")

	bizLog.Debugw("this is Debugw", "k", "Debugw")
	bizLog.Infow("this is Infow", "k", "Infow")
	bizLog.Warnw("this is Warnw", "k", "Warnw")
	bizLog.Errorw("this is Errorw", "k", "Errorw")
}

func TestDefault(t *testing.T) {
	field := zapcore.Field{
		Key:    "module_name",
		Type:   zapcore.StringType,
		String: "serverName",
	}
	SetDefaultConf(WithPath("./storage/logs"),
		WithTimeFormat("2006-01-02 15:04:05"),
		WithLevel(zapcore.InfoLevel),
		WithPreName("global_"),
		WithEncoderJSON(false),
		WithFields(field))

	bizLog := New("default")

	bizLog.Debug("this is SugarDebug")
	bizLog.Info("this is SugarInfo")
	bizLog.Warn("this is SugarWarn")
	bizLog.Error("this is SugarError")

	bizLog.Debugf("this is %s", "Debugf")
	bizLog.Infof("this is %s", "Infof")
	bizLog.Warnf("this is %s", "Warn")
	bizLog.Errorf("this is %s", "Errorf")

	bizLog.Debugw("this is Debugw", "k", "Debugw")
	bizLog.Infow("this is Infow", "k", "Infow")
	bizLog.Warnw("this is Warnw", "k", "Warnw")
	bizLog.Errorw("this is Errorw", "k", "Errorw")
}

func TestWithRequestTrace(t *testing.T) {
	k := "request_id"
	c := context.WithValue(context.Background(), k, "r61f0ed0d70098_Zw8R1aoyl4tGeB4HMV")

	// 告知log如何从上下文获取请求id
	SetDefaultConf(WithRequestIdKey(k))

	l := New("trace")
	l.DebugWithTrace(c,"this is SugarDebug")
	l.InfoWithTrace(c,"this is SugarInfo")
	l.WarnWithTrace(c,"this is SugarWarn")
	l.ErrorWithTrace(c,"this is SugarError")

	l.DebugfWithTrace(c,"this is %s", "Debugf")
	l.InfofWithTrace(c,"this is %s", "Infof")
	l.WarnfWithTrace(c,"this is %s", "Warn")
	l.ErrorfWithTrace(c,"this is %s", "Errorf")

	l.DebugwWithTrace(c,"this is Debugw", "k", "Debugw")
	l.InfowWithTrace(c,"this is Infow", "k", "Infow")
	l.WarnwWithTrace(c,"this is Warnw", "k", "Warnw")
	l.ErrorwWithTrace(c,"this is Errorw", "k", "Errorw")
}
