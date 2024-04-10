package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/router"
	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/logger"
	"github.com/ChangSZ/mall-go/pkg/shutdown"
	"github.com/ChangSZ/mall-go/pkg/timeutil"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.uber.org/zap"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/ChangSZ/mall-go/blob/master/LICENSE

// @securityDefinitions.apikey  LoginToken
// @in                          header
// @name                        token

// @BasePath /
func main() {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectCronLogFile),
	)

	if err != nil {
		panic(err)
	}

	// 初始化logger
	log.Init(configs.ProjectLogFile, configs.ProjectLogRotateMaxDays, configs.ProjectLogLevel)

	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)

	var opts = []http.ServerOption{ // 这里的ServerOption很多只适用于grpc protobuf
		http.Address(":8081"),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "UPDATE"}),
			handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding",
				"X-CSRF-Token", "Authorization", "X-Auth-Token", "X-Auth-UUID", "X-Auth-Openid",
				"referrer", "Authorization", "x-client-id", "x-client-version", "x-client-type"}),
			handlers.AllowCredentials(),
			handlers.ExposedHeaders([]string{"Content-Length"}),
		)),
	}

	// httpSrv := http.NewServer(opts...)
	// httpSrv.HandlePrefix("/", r)

	// app := kratos.New(kratos.Server(httpSrv))
	// fmt.Println("开始运行")
	// if err := app.Run(); err != nil {
	// 	panic(err)
	// }

	defer func() {
		_ = accessLogger.Sync()
		_ = cronLogger.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accessLogger, cronLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if mysql.DB() != nil {
				if err := mysql.DB().DbWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := mysql.DB().DbRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if redis.Cache() != nil {
				if err := redis.Cache().Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		func() {
			if s.CronServer != nil {
				s.CronServer.Stop()
			}
		},
	)
}
