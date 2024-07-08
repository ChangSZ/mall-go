package router

import (
	"html/template"
	"net/http"

	"github.com/ChangSZ/mall-go/assets"
	_ "github.com/ChangSZ/mall-go/docs"
	"github.com/ChangSZ/mall-go/internal/repository/cron"

	"github.com/gin-gonic/gin"
)

const _UI = `
███    ███  █████  ██      ██             ██████   ██████  
████  ████ ██   ██ ██      ██            ██       ██    ██ 
██ ████ ██ ███████ ██      ██      █████ ██   ███ ██    ██ 
██  ██  ██ ██   ██ ██      ██            ██    ██ ██    ██ 
██      ██ ██   ██ ███████ ███████        ██████   ██████ 
`

func RoutersInit(cronServer cron.Server) *gin.Engine {
	eng := gin.Default()
	eng.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.Templates, "templates/**/*")))
	eng.StaticFS("assets", http.FS(assets.Bootstrap))

	InitEngine(eng, _UI)

	// 设置 Render 路由
	setRenderRouter(eng)

	// 设置 API 路由
	setApiRouter(eng, cronServer)

	// 设置 Socket 路由
	setSocketRouter(eng)
	return eng
}
