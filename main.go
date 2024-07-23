package main

import (
	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/shutdown"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/cron"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/router"
	"github.com/ChangSZ/mall-go/pkg/browser"
	"github.com/ChangSZ/mall-go/pkg/file"
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

	var cronServer cron.Server
	var err error
	openBrowserUri := configs.ProjectDomain + configs.ProjectPort + "/render"
	_, ok := file.IsExists(configs.ProjectInstallMark)
	if !ok { // 未安装
		openBrowserUri += "/install"
	}

	// 初始化 DB
	mysql.Init()

	// 初始化 Cache
	redis.Init()

	// 初始化 CRON Server
	cronServer, err = cron.New()
	if err != nil {
		log.Fatal("new cron err: ", err)
	}
	cronServer.Start()

	// 初始化路由
	eng := router.RoutersInit(cronServer)
	log.Info("app Run...")
	if err := eng.Run(configs.ProjectPort); err != nil {
		panic(err)
	}

	// 自动打开浏览器
	if openBrowserUri != "" {
		_ = browser.Open(openBrowserUri)
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

		// 关闭 cron Server
		func() {
			if cronServer != nil {
				cronServer.Stop()
			}
		},
	)
}
