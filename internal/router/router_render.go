package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/render/admin"
	"github.com/ChangSZ/mall-go/internal/render/authorized"
	"github.com/ChangSZ/mall-go/internal/render/config"
	"github.com/ChangSZ/mall-go/internal/render/cron"
	"github.com/ChangSZ/mall-go/internal/render/dashboard"
	generator_handler "github.com/ChangSZ/mall-go/internal/render/generator"
	"github.com/ChangSZ/mall-go/internal/render/index"
	"github.com/ChangSZ/mall-go/internal/render/install"
	"github.com/ChangSZ/mall-go/internal/render/tool"
	"github.com/ChangSZ/mall-go/internal/render/upgrade"
)

func setRenderRouter(eng *gin.Engine) {

	renderInstall := install.New()
	renderIndex := index.New()
	renderDashboard := dashboard.New()
	renderGenerator := generator_handler.New()
	renderConfig := config.New()
	renderAuthorized := authorized.New()
	renderTool := tool.New()
	renderAdmin := admin.New()
	renderUpgrade := upgrade.New()
	renderCron := cron.New()

	eng.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/render")
	})
	// 无需 RBAC 权限验证
	render := eng.Group("/render")
	{
		// 首页
		render.GET("", renderIndex.Index)

		// 仪表盘
		render.GET("/dashboard", renderDashboard.View)

		// 安装
		render.GET("/install", renderInstall.View)
		render.POST("/install/execute", renderInstall.Execute)

		// 管理员
		render.GET("/login", renderAdmin.Login)
		render.GET("/admin/modify_password", renderAdmin.ModifyPassword)
		render.GET("/admin/modify_info", renderAdmin.ModifyInfo)

		// 配置信息
		render.GET("/config/email", renderConfig.Email)
		render.GET("/config/code", renderConfig.Code)

		// 代码生成器
		render.GET("/generator/gorm", renderGenerator.GormView)
		render.POST("/generator/gorm/execute", renderGenerator.GormExecute)

		render.GET("/generator/handler", renderGenerator.HandlerView)
		render.POST("/generator/handler/execute", renderGenerator.HandlerExecute)

		// 调用方
		render.GET("/authorized/list", renderAuthorized.List)
		render.GET("/authorized/add", renderAuthorized.Add)
		render.GET("/authorized/api/:id", renderAuthorized.Api)
		render.GET("/authorized/demo", renderAuthorized.Demo)

		// 管理员
		render.GET("/admin/list", renderAdmin.List)
		render.GET("/admin/add", renderAdmin.Add)
		render.GET("/admin/menu", renderAdmin.Menu)
		render.GET("/admin/menu_action/:id", renderAdmin.MenuAction)
		render.GET("/admin/action/:id", renderAdmin.AdminMenu)

		// 升级
		render.GET("/upgrade", renderUpgrade.UpgradeView)
		render.POST("/upgrade/execute", renderUpgrade.UpgradeExecute)

		// 工具箱
		render.GET("/tool/hashids", renderTool.HashIds)
		render.GET("/tool/logs", renderTool.Log)
		render.GET("/tool/cache", renderTool.Cache)
		render.GET("/tool/data", renderTool.Data)
		render.GET("/tool/websocket", renderTool.Websocket)

		// 后台任务
		render.GET("/cron/list", renderCron.List)
		render.GET("/cron/add", renderCron.Add)
		render.GET("/cron/edit/:id", renderCron.Edit)
	}
}
