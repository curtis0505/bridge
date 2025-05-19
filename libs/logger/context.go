package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"runtime"
	"strings"
)

func createLogWithGinContext(c *gin.Context, ctx ...interface{}) []interface{} {
	logs := make([]interface{}, 0)

	for _, log := range ctx {
		logs = append(logs, log)
	}

	userAgentStr := c.Request.Header.Get("userAgent")
	userAgent, _ := parseUserAgentHeader(userAgentStr)
	if userAgent == nil {
		userAgent = &UserAgent{}
	}
	requestId := userAgent.RequestId
	if requestId == "" {
		requestId = c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			// generate a new request id for internal call tracking
			requestId = uuid.NewV4().String()
		}
		// Header key is canonicalized when added to the header
		c.Request.Header.Add("X-Request-Id", requestId)
	}
	logs = append(logs, []interface{}{"requestId", requestId}...)
	appId := userAgent.AppId
	logs = append(logs, []interface{}{"appId", appId}...)
	reqMethod := c.Request.Method
	logs = append(logs, []interface{}{"method", reqMethod}...)
	reqPath := c.Request.URL.Path
	logs = append(logs, []interface{}{"path", reqPath}...)
	return logs
}

func InfoWithGinContext(c *gin.Context, msg string, ctx ...interface{}) {
	Info(msg, createLogWithGinContext(c, ctx...)...)
}

func DebugWithGinContext(c *gin.Context, msg string, ctx ...interface{}) {
	Debug(msg, createLogWithGinContext(c, ctx...)...)
}

func WarnWithGinContext(c *gin.Context, msg string, ctx ...interface{}) {
	Warn(msg, createLogWithGinContext(c, ctx...)...)
}

func ErrorWithGinContext(c *gin.Context, msg string, ctx ...interface{}) {
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

	// add function name
	ctx = append(ctx, "func")
	ctx = append(ctx, functionNames)

	// add stack
	ctx = append(ctx, "stack")
	ctx = append(ctx, stacks)

	Error(msg, createLogWithGinContext(c, ctx...)...)
}

func CritWithGinContext(c *gin.Context, msg string, ctx ...interface{}) {
	Crit(msg, createLogWithGinContext(c, ctx...)...)
}
