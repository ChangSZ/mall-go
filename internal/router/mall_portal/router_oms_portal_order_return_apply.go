package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/oms_portal_order_return_apply"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 退货申请管理
func setOmsPortalOrderReturnApplyRouter(eng *gin.Engine) {
	handler := oms_portal_order_return_apply.New()
	group := eng.Group("/returnApply", middleware.CheckMemberToken())
	{
		group.POST("/create", handler.Create) // 申请退货
	}
}
