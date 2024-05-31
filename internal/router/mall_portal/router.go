package mall_portal

import (
	_ "github.com/ChangSZ/mall-go/docs"
	"github.com/ChangSZ/mall-go/internal/router"

	"github.com/gin-gonic/gin"
)

const _UI = `
███    ███  █████  ██      ██              ██████   ██████  ██████  ████████  █████  ██      
████  ████ ██   ██ ██      ██              ██   ██ ██    ██ ██   ██    ██    ██   ██ ██      
██ ████ ██ ███████ ██      ██              ██████  ██    ██ ██████     ██    ███████ ██      
██  ██  ██ ██   ██ ██      ██              ██      ██    ██ ██   ██    ██    ██   ██ ██      
██      ██ ██   ██ ███████ ███████ ███████ ██       ██████  ██   ██    ██    ██   ██ ███████ 
`

func RoutersInit() *gin.Engine {
	eng := router.InitEngine(_UI)

	setUmsMemberRouter(eng)
	setHomeRouter(eng)

	return eng
}
