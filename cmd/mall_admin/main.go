package main

import (
	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/shutdown"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/router/mall_admin"
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
	logCfg := log.Config{
		FilePath: configs.ProjectLogFile,
		MaxDays:  configs.ProjectLogRotateMaxDays,
		LogLevel: configs.ProjectLogLevel,
	}
	log.Init(logCfg)

	// 初始化 DB
	mysql.Init()

	// 初始化 Cache
	redis.Init()

	// 初始化路由
	eng := mall_admin.RoutersInit()
	log.Info("app Run...")
	if err := eng.Run(configs.MallAdminPort); err != nil {
		panic(err)
	}

	// 优雅关闭
	shutdown.NewHook().Close(
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
