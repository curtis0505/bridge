package logger

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	loggerconf "github.com/curtis0505/bridge/libs/logger/v2/conf"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/exp/slog"
	"reflect"
	"runtime"
	"strings"
)

type Config struct {
	UseTerminal        bool
	TerminalJSONOutput bool
	UseFile            bool
	VerbosityTerminal  int
	VerbosityFile      int
	FilePath           string

	LogStash LogStashConfig
}

type LogStashConfig struct {
	ConnectionType string
	Use            bool
	Verbosity      int
	Address        string
}

var (
	appNameKey = "appName"
	appName    string
)

var (
	errInvalidLogLength = fmt.Errorf("[elog] Invalid log format: ctx length must be even")
)

const PleaseCheckLog = "checkLog"

func getSlogLevelByLog15Level(level int) slog.Level {
	switch level {
	case 0, 1:
		//Crit, Error
		return slog.LevelError
	case 2:
		// Warn
		return slog.LevelWarn
	case 3:
		// Info
		return slog.LevelInfo
	case 4, 5:
		// Debug, Trace
		return slog.LevelDebug
	default:
		return slog.LevelDebug
	}
}

func InitLog(config Config) {
	//terminalLogLevel := getSlogLevelByLog15Level(config.VerbosityTerminal)
	//if config.TerminalJSONOutput == true {
	//	log.SetDefault(log.NewLogger(log.JSONHandler(os.Stdout)))
	//} else {
	//	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, terminalLogLevel, true)))
	//}

	logger.InitLog(loggerconf.Config{
		UseTerminal:        config.UseTerminal,
		TerminalJSONOutput: config.TerminalJSONOutput,
		//UseFile:            config.UseFile,
		VerbosityTerminal: config.VerbosityTerminal,
		//VerbosityFile:      config.VerbosityFile,
		//FilePath:           config.FilePath,
	})
	// TODO: support file & logstash log
	//	if config.LogStash.Use { // LogStash
	//		handler, err := log.NetHandler(config.LogStash.ConnectionType, config.LogStash.Address, log.JSONFormatEx(false, true))
	//		if err != nil {
	//			fmt.Println("LogStash Connect Failed", err)
	//		} else {
	//			h = append(h, log.LvlFilterHandler(log.Lvl(config.LogStash.Verbosity), handler))
	//		}
	//	}
	//	return h
	//}()
}

func SetAppName(name string) {
	appName = name
}

func checkValidCtxFormat(ctx ...interface{}) error {
	if len(ctx)%2 != 0 {
		if len(ctx) == 1 && reflect.TypeOf(ctx).Kind() == reflect.Slice {
		} else {
			return errInvalidLogLength
		}
	} else {
		for index, ctxVal := range ctx {
			// Validate key only
			if index%2 == 0 {
				switch ctxType := reflect.ValueOf(ctxVal); ctxType.Kind() {
				case reflect.String:
					if ctxType.String() == "" {
						return fmt.Errorf("[elog] Invalid log format: Key shoud not be empty")
					}
				default:
					return fmt.Errorf("[elog] Invalid log format: Key shoud be string type")
				}
			}
		}
	}
	return nil
}

func Trace(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			log.Trace(err.Error(), ctx...)
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	// log.Debug(msg, ctx...)
	logger.Debug(msg, logger.BuildLogInput().WithData(ctx...))
}

func Debug(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			log.Debug(err.Error(), ctx...)
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	// log.Debug(msg, ctx...)
	logger.Debug(msg, logger.BuildLogInput().WithData(ctx...))
}

func Info(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			// log.Info(err.Error(), ctx...)
			logger.Info(msg, logger.BuildLogInput().WithData(ctx...))
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	// log.Info(msg, ctx...)
	logger.Info(msg, logger.BuildLogInput().WithData(ctx...))
}

func Warn(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			// log.Warn(err.Error(), ctx...)
			logger.Warn(msg, logger.BuildLogInput().WithData(ctx...))
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	//log.Warn(msg, ctx...)
	logger.Warn(msg, logger.BuildLogInput().WithData(ctx...))
}

func Error(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			// log.Error(err.Error(), ctx...)
			logger.Error(msg, logger.BuildLogInput().WithData(ctx...))
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	// log.Error(msg, ctx...)
	logger.Error(msg, logger.BuildLogInput().WithData(ctx...))
}

func Crit(msg string, ctx ...interface{}) {
	if err := checkValidCtxFormat(ctx...); err != nil {
		switch err {
		case errInvalidLogLength:
			ctx = append([]interface{}{msg}, ctx...)
			// log.Crit(err.Error(), ctx...)
			logger.Crit(msg, logger.BuildLogInput().WithData(ctx...))
			return
		default:
			Error(err.Error(), PleaseCheckLog, msg)
			return
		}
	}
	if appName != "" {
		ctx = append([]interface{}{appNameKey, appName}, ctx...)
	}
	// log.Crit(msg, ctx...)
	logger.Crit(msg, logger.BuildLogInput().WithData(ctx...))
}

func TraceStack() (string, string) {
	pc := make([]uintptr, 15)
	i := 2
	functionNames := ""
	stacks := ""
	for i < len(pc) {
		n := runtime.Callers(i, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		functions := strings.Split(frame.Function, ".")
		functionName := functions[len(functions)-1]
		if functionName == "" || functionName == "call" || functionName == "Next" || functionName == "func1" {
			break
		}
		functionNames += fmt.Sprintf("%s\n ", functions[len(functions)-1])
		stacks += fmt.Sprintf("%s:%d\n ", frame.File, frame.Line)
		i += 1
	}

	return functionNames, stacks
}
