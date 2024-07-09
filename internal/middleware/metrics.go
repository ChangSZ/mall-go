package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/metrics"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/pkg/env"
)

// Metrics 统计
func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		ts := time.Now()
		defer func() {
			var businessCode int
			if err := core.AbortError(ctx); err != nil {
				businessCode = err.BusinessCode()
			}
			path := ctx.Request.URL.Path
			if alias := core.Alias(ctx); alias != "" {
				path = alias
			}

			metrics.RecordHandler()(&proposal.MetricsMessage{
				ProjectName:  configs.ProjectName,
				Env:          env.Active().Value(),
				TraceID:      core.TraceID(ctx),
				HOST:         ctx.Request.Host,
				Path:         path,
				Method:       ctx.Request.Method,
				HTTPCode:     ctx.Writer.Status(),
				BusinessCode: businessCode,
				CostSeconds:  time.Since(ts).Seconds(),
				IsSuccess:    !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK),
			})
		}()

		ctx.Next()
	}
}
