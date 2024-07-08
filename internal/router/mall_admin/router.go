package mall_admin

import (
	_ "github.com/ChangSZ/mall-go/docs"
	"github.com/ChangSZ/mall-go/internal/router"

	"github.com/gin-gonic/gin"
)

const _UI = `
███    ███  █████  ██      ██             █████  ██████  ███    ███ ██ ███    ██ 
████  ████ ██   ██ ██      ██            ██   ██ ██   ██ ████  ████ ██ ████   ██ 
██ ████ ██ ███████ ██      ██      █████ ███████ ██   ██ ██ ████ ██ ██ ██ ██  ██ 
██  ██  ██ ██   ██ ██      ██            ██   ██ ██   ██ ██  ██  ██ ██ ██  ██ ██ 
██      ██ ██   ██ ███████ ███████       ██   ██ ██████  ██      ██ ██ ██   ████ 
`

func RoutersInit() *gin.Engine {
	eng := gin.Default()
	router.InitEngine(eng, _UI)

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

	// 设置 Minio 路由
	setMinioRouter(eng)

	return eng
}
