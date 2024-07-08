package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	tp := sdkTrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
}

// Tracing 链路追踪
func Tracing(name string) gin.HandlerFunc {
	return otelgin.Middleware(name)
}
