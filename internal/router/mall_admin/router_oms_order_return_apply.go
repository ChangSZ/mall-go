package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/oms_order_return_apply"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 订单退货申请管理
func setOmsOrderReturnApplyRouter(eng *gin.Engine) {
	handler := oms_order_return_apply.New()
	group := eng.Group("/returnApply", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.GET("/list", handler.List)                       // 分页查询退货申请
		group.POST("/delete", handler.Delete)                  // 批量删除退货申请
		group.GET("/:id", handler.GetItem)                     // 获取退货申请详情
		group.POST("/update/status/:id", handler.UpdateStatus) // 修改退货申请状态
	}
}
