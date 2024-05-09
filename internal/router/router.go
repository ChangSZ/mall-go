package router

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ChangSZ/mall-go/assets"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/middleware"
	"github.com/ChangSZ/mall-go/internal/repository/cron"
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

const _UI = `
███    ███  █████  ██      ██             ██████   ██████  
████  ████ ██   ██ ██      ██            ██       ██    ██ 
██ ████ ██ ███████ ██      ██      █████ ██   ███ ██    ██ 
██  ██  ██ ██   ██ ██      ██            ██    ██ ██    ██ 
██      ██ ██   ██ ███████ ███████        ██████   ██████ 
`

func RoutersInit(cronServer cron.Server) *gin.Engine {
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

	fmt.Println(color.Blue(_UI))

	eng.StaticFS("assets", http.FS(assets.Bootstrap))
	eng.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))

	// 设置 Render 路由
	setRenderRouter(eng)

	// 设置 API 路由
	setApiRouter(eng, cronServer)

	// 设置 Ums 路由
	setUmsAdminRouter(eng)
	setUmsMemberLevelRouter(eng)
	setUmsMenuRouter(eng)
	setUmsResourceCateRouter(eng)
	setUmsResourceRouter(eng)
	setUmsRoleRouter(eng)

	// 设置 Pms 路由
	setPmsBrandRouter(eng)
	setPmsProducteAttrCateRouter(eng)
	setPmsProducteAttrRouter(eng)
	setPmsProducteCateRouter(eng)
	setPmsProducteRouter(eng)
	setPmsSkuStockRouter(eng)

	// 设置 Cms 路由
	setCmsSubjectRouter(eng)
	setCmsPrefrenceAreaRouter(eng)

	// 设置 Sms 路由
	setSmsCouponRouter(eng)
	setSmsCouponHistoryRouter(eng)
	setSmsFlashPromotionRouter(eng)
	setSmsFlashPromotionProductRelationRouter(eng)
	setSmsFlashPromotionSessionRouter(eng)
	setSmsHomeAdvertiseRouter(eng)
	setSmsHomeBrandRouter(eng)
	setSmsHomeNewProductRouter(eng)
	setSmsHomeRecommendProductRouter(eng)
	setSmsHomeRecommendSubjectRouter(eng)

	// 设置 Oms 路由
	setOmsOrderRouter(eng)
	setOmsCompanyAddressRouter(eng)
	setOmsOrderReturnApplyRouter(eng)
	setOmsOrderReturnReasonRouter(eng)
	setOmsOrderSettingRouter(eng)

	// 设置 Socket 路由
	setSocketRouter(eng)

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
