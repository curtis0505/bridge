package logger

import (
	"github.com/curtis0505/bridge/libs/logger/v2/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var logger *zap.Logger

var once = new(sync.Once)

func Trace(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Debug(msg, logs...)
}

func Debug(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Debug(msg, logs...)
}

func Info(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Info(msg, logs...)
}

func Warn(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Warn(msg, logs...)
}

func Error(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Error(msg, logs...)
}

func Crit(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}
	logger.Panic(msg, logs...)
}

func InitLog(config conf.Config) {
	once.Do(func() {
		terminalLogLevel := getLogLevel(config.VerbosityTerminal)
		var cfg zap.Config
		if config.TerminalJSONOutput == true {
			cfg = zap.NewProductionConfig()
			cfg.DisableCaller = true
			cfg.EncoderConfig.LevelKey = "lvl"
			cfg.EncoderConfig.TimeKey = "t"
			cfg.EncoderConfig.StacktraceKey = "at"
			cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		} else {
			cfg = zap.NewDevelopmentConfig()
			cfg.DisableCaller = true
			cfg.EncoderConfig.LevelKey = "lvl"
			cfg.EncoderConfig.TimeKey = "t"
			cfg.EncoderConfig.StacktraceKey = "at"
			cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		}
		//cfg.OutputPaths = []string{"/dev/stdout"}
		cfg.Level = zap.NewAtomicLevelAt(terminalLogLevel)
		var err error
		logger, err = cfg.Build(zap.AddStacktrace(zapcore.FatalLevel), zap.AddStacktrace(zapcore.DPanicLevel))
		if err != nil {
			panic(err)
		}
	})
}

func getLogLevel(level int) zapcore.Level {
	switch level {
	case 0, 1:
		//Crit, Error
		return zapcore.ErrorLevel
	case 2:
		// Warn
		return zapcore.WarnLevel
	case 3:
		// Info
		return zapcore.InfoLevel
	default:
		// Debug, Trace
		return zapcore.DebugLevel
	}
}
