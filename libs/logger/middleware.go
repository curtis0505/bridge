package logger

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type UserAgent struct {
	AppName        string
	OsVersion      string
	DeviceModel    string
	Manufacturer   string
	DeviceLanguage string
	AppId          string
	RequestId      string
}

func parseUserAgentHeader(userAgentStr string) (*UserAgent, error) {
	// 앱이름|OS버전|디바이스모델|제조사|디바이스언어|(고정)app ID|(갱신)request id
	// i.e) neopin|14.8|iPhone X|apple|ko-KR|1066567D-4FAE-48AC-94B8-9D642AE36A33|1482E5B0-1B5B-4F5A-8F5C-8B9F5F5F5F5F
	slices := strings.Split(userAgentStr, "|")
	if len(slices) != 7 {
		return nil, errors.New("invalid user agent format")
	}
	userAgent := &UserAgent{
		AppName:        slices[0],
		OsVersion:      slices[1],
		DeviceModel:    slices[2],
		Manufacturer:   slices[3],
		DeviceLanguage: slices[4],
		AppId:          slices[5],
		RequestId:      slices[6],
	}
	return userAgent, nil
}

func (u *UserAgent) ToCtx(params *gin.LogFormatterParams) []interface{} {
	logs := make([]interface{}, 0)
	logs = append(logs, []interface{}{"appName", u.AppName}...)
	logs = append(logs, []interface{}{"osVersion", u.OsVersion}...)
	logs = append(logs, []interface{}{"deviceModel", u.DeviceModel}...)
	logs = append(logs, []interface{}{"manufacturer", u.Manufacturer}...)
	logs = append(logs, []interface{}{"deviceLanguage", u.DeviceLanguage}...)
	logs = append(logs, []interface{}{"appId", u.AppId}...)
	logs = append(logs, []interface{}{"requestId", u.RequestId}...)
	logs = append(logs, []interface{}{"statusCode", params.StatusCode}...)
	logs = append(logs, []interface{}{"latency", params.Latency.String()}...)
	logs = append(logs, []interface{}{"clientIp", params.ClientIP}...)
	logs = append(logs, []interface{}{"method", params.Method}...)
	logs = append(logs, []interface{}{"path", params.Path}...)
	return logs
}

func GinElogMiddleWare() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		userAgentStr := params.Request.Header.Get("userAgent")
		userAgent, _ := parseUserAgentHeader(userAgentStr)
		if userAgent == nil {
			userAgent = &UserAgent{}
		}
		if userAgent.RequestId == "" {
			requestId := params.Request.Header.Get("X-Request-Id")
			if requestId == "" {
				requestId = uuid.NewV4().String()
			}
			userAgent.RequestId = requestId
		}

		// skip health check log
		if strings.Contains(params.Path, "health") {
			return ""
		}

		// Access log formatting
		if params.StatusCode >= 200 && params.StatusCode < 300 {
			Info("Gin Request", userAgent.ToCtx(&params)...)
		} else if params.StatusCode >= 300 && params.StatusCode < 400 {
			Info("Gin Request", userAgent.ToCtx(&params)...)
		} else if params.StatusCode >= 400 && params.StatusCode < 500 {
			Warn("Gin Request", userAgent.ToCtx(&params)...)
		} else if params.StatusCode >= 500 && params.StatusCode < 600 {
			Error("Gin Request", userAgent.ToCtx(&params)...)
		} else {
			Info("Gin Request", userAgent.ToCtx(&params)...)
		}
		return ""
	})
}
