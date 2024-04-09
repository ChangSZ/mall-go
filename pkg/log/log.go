package log

import (
	"context"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

var myLogger log.Logger

const (
	// LevelDebug is logger debug level.
	LevelDebug log.Level = iota - 1
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level
	LevelFatal
)

// GetLogger returns global logger appliance as logger in current process.
func GetLoggerWithTrace() log.Logger {
	return log.With(log.GetLogger(), "trace.id", tracing.TraceID())
}

// Log Print log by level and keyvals.
func Log(level log.Level, keyvals ...interface{}) {
	log.Log(level, keyvals...)
}

func WithTrace(ctx context.Context) *log.Helper {
	if c, ok := ctx.(*gin.Context); ok {
		return log.NewHelper(log.WithContext(c.Request.Context(), log.With(log.GetLogger(), "trace.id", tracing.TraceID())))
	}
	return log.NewHelper(log.WithContext(ctx, log.With(log.GetLogger(), "trace.id", tracing.TraceID())))
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	log.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	log.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func Debugw(keyvals ...interface{}) {
	log.Debugw(keyvals...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	log.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	log.Infof(format, a...)
}

// Infow logs a message at info level.
func Infow(keyvals ...interface{}) {
	log.Infow(keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	log.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	log.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(keyvals ...interface{}) {
	log.Warnw(keyvals...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	log.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	log.Errorf(format, a...)
}

// Errorw logs a message at error level.
func Errorw(keyvals ...interface{}) {
	log.Errorw(keyvals...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	log.Fatal(a...)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	log.Fatalf(format, a...)
}

// Fatalw logs a message at fatal level.
func Fatalw(keyvals ...interface{}) {
	log.Fatalw(keyvals...)
}

func Init(filePath string, maxDays int) {
	writer, err := RotateDailyLog(filePath, maxDays)
	if err != nil {
		panic("创建日志文件失败")
	}
	multiWriter := io.MultiWriter(os.Stdout, writer)
	myLogger = log.With(log.NewStdLogger(multiWriter),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(5),
	)
	log.SetLogger(myLogger)
}
