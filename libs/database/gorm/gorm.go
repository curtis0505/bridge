package gorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"os"
	"time"
)

func FormatDSN(host, user, password, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", user, password, host, database)
}

func NewGORMClient(host, user, password, database string) (*gorm.DB, error) {

	var _logger gormlogger.Interface
	switch os.Getenv("APPMODE") {
	case "live", "dq":
		_logger = &GormLogger{
			LogLevel:      gormlogger.Silent,
			SlowThreshold: time.Second,
		}
	default:
		_logger = &GormLogger{
			LogLevel:      gormlogger.Info,
			SlowThreshold: time.Second,
		}
	}

	return gorm.Open(mysql.Open(FormatDSN(host, user, password, database)), &gorm.Config{NowFunc: time.Now().UTC, Logger: _logger})
}

var (
	_ gormlogger.Interface = &GormLogger{}
)

type GormLogger struct {
	SlowThreshold time.Duration
	LogLevel      gormlogger.LogLevel
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.LogLevel = level
	return l
}

func (l GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	logger.Info(s, logger.BuildLogInput().WithData(i...))
}

func (l GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	logger.Warn(s, logger.BuildLogInput().WithData(i...))
}

func (l GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	logger.Error(s, logger.BuildLogInput().WithData(i...))
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Warn && errors.Is(err, gormlogger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logger.Warn("SQL", logger.BuildLogInput().WithError(err).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "sql", sql))
		} else {
			logger.Warn("SQL", logger.BuildLogInput().WithError(err).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql))
		}
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, rows := fc()
		if rows == -1 {
			logger.Error("SQL", logger.BuildLogInput().WithError(err).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "sql", sql))
		} else {
			logger.Error("SQL", logger.BuildLogInput().WithError(err).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql))
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		if rows == -1 {
			logger.Warn("SQL", logger.BuildLogInput().WithError(fmt.Errorf("slow query")).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "sql", sql))
		} else {
			logger.Warn("SQL", logger.BuildLogInput().WithError(fmt.Errorf("slow query")).WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql))
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			logger.Info("SQL", logger.BuildLogInput().WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "sql", sql))
		} else {
			logger.Info("SQL", logger.BuildLogInput().WithData("elapsed", float64(elapsed.Nanoseconds())/1e6, "rows", rows, "sql", sql))
		}
	}
}
