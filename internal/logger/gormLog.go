package logger

import (
	"context"
	"errors"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type log struct {
	GormLoggerConfig
}

type GormLoggerConfig struct {
	SlowThreshold time.Duration
	TraceLog      bool
}

func NewGormLog(config GormLoggerConfig) gormLogger.Interface {
	return &log{
		GormLoggerConfig: config,
	}
}

// LogMode log mode
func (l *log) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info print info
func (l log) Info(ctx context.Context, msg string, data ...interface{}) {
	WithContext(ctx).Infof(msg, data)
}

// Warn print warn messages
func (l log) Warn(ctx context.Context, msg string, data ...interface{}) {
	WithContext(ctx).Warnf(msg, data)
}

// Error print error messages
func (l log) Error(ctx context.Context, msg string, data ...interface{}) {
	WithContext(ctx).Errorf(msg, data)
}

// Trace print sql message
func (l log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		WithContext(ctx).Errorf("gorm error:%v, sql=%v, affected_rows=%v duration=%v", err.Error(), sql, rows, float64(elapsed.Nanoseconds())/1e6)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		sql, rows := fc()
		WithContext(ctx).Warnf("gorm slowlog, sql=%v, affected_rows=%v duration=%v", sql, rows, float64(elapsed.Nanoseconds())/1e6)
	default:
		if l.TraceLog {
			sql, rows := fc()
			WithContext(ctx).Infof("gorm trace, sql=%v, affected_rows=%v duration=%v", sql, rows, float64(elapsed.Nanoseconds())/1e6)
		}
	}
}
