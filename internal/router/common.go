package router

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/middleware"
	"github.com/ChangSZ/mall-go/pkg/color"
	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitEngine(ui string) *gin.Engine {
	if env.Active().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	eng := gin.New()
	eng.Use(
		middleware.Rate(),
		middleware.Metrics(),
		// middleware.AlertNotify(),
		kgin.Middlewares(tracing.Server(), middleware.Logging(log.GetLoggerWithTrace()), middleware.AddTraceCtx),
	)

	fmt.Println(color.Blue(ui))

	system := eng.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx *gin.Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Active().Value(),
				Host:        ctx.Request.Host,
				Status:      "ok",
			}
			api.ResponseOK(ctx, resp)
		})
	}

	var enablePProf = true
	if enablePProf {
		if !env.Active().IsPro() {
			pprof.Register(eng) // register pprof to gin
		}
	}

	if !env.Active().IsPro() {
		eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // register swagger
	}

	var enablePrometheus = true
	if enablePrometheus {
		eng.GET("/metrics", gin.WrapH(promhttp.Handler())) // register prometheus
	}
	return eng
}
