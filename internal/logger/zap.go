package logger

import (
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/petermattis/goid"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"mall/internal/core"
	"strconv"
	"time"
)

var logging *zap.SugaredLogger

var prefix int64 = 10000000000000000

func WithGoID() *zap.SugaredLogger {
	gid := goid.Get()
	return logging.With("goid", strconv.FormatInt(prefix+gid, 10))
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return logging
	}
	duration := (time.Now().UnixNano() - cast.ToInt64(ctx.Value("startTime"))) / int64(time.Millisecond)
	return logging.With("duration", duration).With("traceId", ctx.Value("traceId"))
}

func InitLogger() error {
	logWriter, err := GetWriter(core.GlobalConfig.Logger.LogFile)
	if err != nil {
		fmt.Println(fmt.Sprintf("get logger writer error:%v", err))
		return err
	}
	c := zapcore.NewCore(GetEncoder(), zapcore.AddSync(logWriter), LogLevel(core.GlobalConfig.Logger.LogLevel))
	log := zap.New(c, zap.AddCaller())
	logging = log.Sugar()

	return nil

}

func GetWriter(filename string) (logf io.Writer, err error) {
	logf, err = rotatelogs.New(filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	return
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "content",
		LevelKey:       "level",
		TimeKey:        "ts",
		CallerKey:      "file",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
	})
}

func LogLevel(level string) (LogLevel zapcore.Level) {
	switch level {
	case "debug", "Debug":
		LogLevel = zapcore.DebugLevel
	case "info", "Info":
		LogLevel = zapcore.InfoLevel
	case "warn", "Warn":
		LogLevel = zapcore.WarnLevel
	case "error", "Error":
		LogLevel = zapcore.ErrorLevel
	}
	return LogLevel
}
