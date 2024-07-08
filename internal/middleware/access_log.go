package middleware

import (
	"time"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

// withoutLogPaths 这些请求，默认不记录日志
var withoutLogPaths = map[string]bool{
	"/metrics":          true,
	"/assets/*filepath": true,

	"/debug/pprof/":             true,
	"/debug/pprof/cmdline":      true,
	"/debug/pprof/profile":      true,
	"/debug/pprof/symbol":       true,
	"/debug/pprof/trace":        true,
	"/debug/pprof/allocs":       true,
	"/debug/pprof/block":        true,
	"/debug/pprof/goroutine":    true,
	"/debug/pprof/heap":         true,
	"/debug/pprof/mutex":        true,
	"/debug/pprof/threadcreate": true,

	"/favicon.ico": true,

	"/system/health": true,
}

func AccessLog(logger log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		statusCode := ctx.Writer.Status()

		path := ctx.Request.URL.Path
		if _, ok := withoutLogPaths[path]; ok {
			if statusCode == 0 { // 非0时（报错）还是要打印下日志的
				return
			}
		}
		operation := path

		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			operation = operation + "?" + raw
		}

		ctx.Next()

		level := log.LevelInfo
		if statusCode >= 400 && statusCode <= 599 {
			level = log.LevelError
		}
		log.NewHelper(log.WithContext(ctx, logger)).Log(level,
			"ClientIP", ctx.ClientIP(),
			"Operation", operation,
			"Method", ctx.Request.Method,
			"StatusCode", statusCode,
			"ErrorMessage", ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
			"Latency", time.Since(startTime).Seconds(),
		)
	}
}
