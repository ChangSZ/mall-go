package core

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"

	"github.com/ChangSZ/mall-go/internal/proposal"
)

const (
	_Alias           = "_alias_"
	_SessionUserInfo = "_session_user_info"
	_UmsUserInfo     = "_ums_user_info_"
)

func SessionUserInfo(ctx context.Context) proposal.SessionUserInfo {
	val, ok := ctx.Value(_SessionUserInfo).(proposal.SessionUserInfo)
	if !ok {
		return proposal.SessionUserInfo{}
	}
	return val
}

func SetSessionUserInfo(ctx *gin.Context, info proposal.SessionUserInfo) {
	ctx.Set(_SessionUserInfo, info)
}

func GetUmsUserInfo(ctx context.Context) proposal.UmsUserInfo {
	val, ok := ctx.Value(_UmsUserInfo).(proposal.UmsUserInfo)
	if !ok {
		return proposal.UmsUserInfo{}
	}
	return val
}

func SetUmsUserInfo(ctx *gin.Context, info proposal.UmsUserInfo) {
	ctx.Set(_UmsUserInfo, info)
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
