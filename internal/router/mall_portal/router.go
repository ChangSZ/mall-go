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
	eng := gin.Default()
	router.InitEngine(eng, _UI)

	setUmsMemberRouter(eng)
	setUmsMemberCouponRouter(eng)
	setUmsMemberReceiveAddressRouter(eng)

	setPmsPortalBrandRouter(eng)
	setPmsPortalProductRouter(eng)

	setOmsCartItemBrandRouter(eng)
	setOmsPortalOrderReturnApplyRouter(eng)
	setOmsPortalOrderRouter(eng)

	setHomeRouter(eng)

	setAlipayRouter(eng)

	setMemberAttentionRouter(eng)
	setMemberCollectionRouter(eng)
	setMemberReadHistoryRouter(eng)

	return eng
}
