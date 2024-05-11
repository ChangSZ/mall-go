package core

import (
	stdctx "context"
	"net/http"
	"strings"

	"github.com/ChangSZ/mall-go/internal/proposal"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

const (
	_Alias           = "_alias_"
	_SessionUserInfo = "_session_user_info"
	_AbortErrorName  = "_abort_error_"
	_UmsUserInfo     = "_ums_user_info_"
)

func SessionUserInfo(ctx stdctx.Context) proposal.SessionUserInfo {
	val, ok := ctx.Value(_SessionUserInfo).(proposal.SessionUserInfo)
	if !ok {
		return proposal.SessionUserInfo{}
	}
	return val
}

func SetSessionUserInfo(ctx *gin.Context, info proposal.SessionUserInfo) {
	ctx.Set(_SessionUserInfo, info)
}

func GetUmsUserInfo(ctx stdctx.Context) proposal.UmsUserInfo {
	val, ok := ctx.Value(_UmsUserInfo).(proposal.UmsUserInfo)
	if !ok {
		return proposal.UmsUserInfo{}
	}
	return val
}

func SetUmsUserInfo(ctx *gin.Context, info proposal.UmsUserInfo) {
	ctx.Set(_UmsUserInfo, info)
}

func AbortWithError(ctx *gin.Context, err BusinessError) {
	if err != nil {
		httpCode := err.HTTPCode()
		if httpCode == 0 {
			httpCode = http.StatusInternalServerError
		}

		ctx.AbortWithStatus(httpCode)
		ctx.Set(_AbortErrorName, err)
	}
}

func AbortError(ctx *gin.Context) BusinessError {
	err, ok := ctx.Get(_AbortErrorName)
	if !ok {
		return nil
	}
	return err.(BusinessError)
}

func Alias(ctx *gin.Context) string {
	path, ok := ctx.Get(_Alias)
	if !ok {
		return ""
	}
	return path.(string)
}

func SetAlias(ctx *gin.Context, path string) {
	if path = strings.TrimSpace(path); path != "" {
		ctx.Set(_Alias, path)
	}
}

func TraceID(ctx *gin.Context) string {
	var traceId string
	if span := trace.SpanContextFromContext(ctx.Request.Context()); span.HasTraceID() {
		traceId = span.TraceID().String()
	}
	return traceId
}
