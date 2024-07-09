package router

import (
	"fmt"
	"time"

	"github.com/ChangSZ/golib/color"
	"github.com/ChangSZ/golib/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/middleware"
	"github.com/ChangSZ/mall-go/pkg/env"
)

func InitEngine(eng *gin.Engine, serverName, ui string) *gin.Engine {
	if env.Active().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 配置 CORS 中间件
	config := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding",
			"X-CSRF-Token", "Authorization", "X-Auth-Token", "X-Auth-UUID", "X-Auth-Openid",
			"referrer", "Authorization", "x-client-id", "x-client-version", "x-client-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	eng.Use(
		cors.New(config),
		middleware.Rate(),
		middleware.Metrics(),
		middleware.Tracing(serverName),
		middleware.AccessLog(log.GetLoggerWithTrace()),
		// middleware.AlertNotify(),
	)

	fmt.Println(color.Blue(ui))
	// 设置可信代理
	err := eng.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

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
