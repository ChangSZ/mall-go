package router

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/middleware"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/cron"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/router/interceptor"
	"github.com/ChangSZ/mall-go/pkg/file"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
)

type resource struct {
	mux          core.Mux
	logger       *zap.Logger
	interceptors interceptor.Interceptor
	cronServer   cron.Server
}

type Server struct {
	Mux        core.Mux
	CronServer cron.Server
}

// func NewHTTPServer(logger *zap.Logger, cronLogger *zap.Logger) (*Server, error) {
// 	if logger == nil {
// 		return nil, errors.New("logger required")
// 	}

// 	r := new(resource)
// 	r.logger = logger

// 	openBrowserUri := configs.ProjectDomain + configs.ProjectPort

// 	_, ok := file.IsExists(configs.ProjectInstallMark)
// 	if !ok { // 未安装
// 		openBrowserUri += "/install"
// 		// _ = browser.Open(openBrowserUri)
// 	} else { // 已安装
// 		// 初始化 DB
// 		mysql.Init()

// 		// 初始化 Cache
// 		redis.Init()

// 		// 初始化 CRON Server
// 		cronServer, err := cron.New(cronLogger)
// 		if err != nil {
// 			logger.Fatal("new cron err", zap.Error(err))
// 		}
// 		cronServer.Start()
// 		r.cronServer = cronServer
// 	}

// 	mux, err := core.New(logger,
// 		core.WithEnableOpenBrowser(openBrowserUri),
// 		core.WithEnableCors(),
// 		core.WithEnableRate(),
// 		core.WithAlertNotify(alert.NotifyHandler(logger)),
// 		core.WithRecordMetrics(metrics.RecordHandler(logger)),
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	r.mux = mux
// 	r.interceptors = interceptor.New(logger)

// 	// 设置 Render 路由
// 	setRenderRouter(r)

// 	// 设置 API 路由
// 	setApiRouter(r)

// 	// 设置 Ums 路由
// 	setUmsAdminRouter(r)
// 	setUmsMemberLevelRouter(r)
// 	setUmsMenuRouter(r)
// 	setUmsResourceCateRouter(r)
// 	setUmsResourceRouter(r)
// 	setUmsRoleRouter(r)

// 	// 设置 Socket 路由
// 	setSocketRouter(r)

// 	s := new(Server)
// 	s.Mux = mux
// 	s.CronServer = r.cronServer

// 	return s, nil
// }

func RoutersInit() *gin.Engine {
	eng := gin.New()
	eng.Use(kgin.Middlewares(recovery.Recovery(), tracing.Server(), logging.Server(log.GetLoggerWithTrace()), middleware.AddTraceCtx))

	openBrowserUri := configs.ProjectDomain + configs.ProjectPort

	_, ok := file.IsExists(configs.ProjectInstallMark)
	if !ok { // 未安装
		openBrowserUri += "/install"
	} else { // 已安装
		// 初始化 DB
		mysql.Init()

		// 初始化 Cache
		redis.Init()
	}

	// 设置 Render 路由
	setRenderRouter(eng)

	// 设置 API 路由
	setApiRouter(eng)

	// 设置 Ums 路由
	setUmsAdminRouter(eng)
	setUmsMemberLevelRouter(eng)
	setUmsMenuRouter(eng)
	setUmsResourceCateRouter(eng)
	setUmsResourceRouter(eng)
	setUmsRoleRouter(eng)

	// 设置 Socket 路由
	setSocketRouter(eng)

	s := new(Server)
	s.Mux = mux
	s.CronServer = eng.cronServer

	return s, nil
}
