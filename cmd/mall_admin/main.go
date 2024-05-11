package main

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/router/mall_admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/shutdown"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
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
	// 初始化logger
	log.Init(configs.ProjectLogFile, configs.ProjectLogRotateMaxDays, configs.ProjectLogLevel)

	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)

	var opts = []http.ServerOption{ // 这里的ServerOption很多只适用于grpc protobuf
		http.Address(configs.ProjectPort),
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

	// 初始化 DB
	mysql.Init()

	// 初始化 Cache
	redis.Init()

	// 初始化路由
	eng := mall_admin.RoutersInit()
	httpSrv := http.NewServer(opts...)
	httpSrv.HandlePrefix("/", eng)

	app := kratos.New(kratos.Server(httpSrv))
	go func() {
		log.Info("app Run...")
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			log.Info("app Stop...")
			if err := app.Stop(); err != nil {
				log.Error("app Stop err: ", err)
			}
		},

		// 关闭 db
		func() {
			if mysql.DB() != nil {
				if err := mysql.DB().DbWClose(); err != nil {
					log.Error("dbw close err: ", err)
				}

				if err := mysql.DB().DbRClose(); err != nil {
					log.Error("dbr close err: ", err)
				}
			}
		},

		// 关闭 cache
		func() {
			if redis.Cache() != nil {
				if err := redis.Cache().Close(); err != nil {
					log.Error("cache close err: ", err)
				}
			}
		},
	)
}
