package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/oms_order_return_reason"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 退货原因管理
func setOmsOrderReturnReasonRouter(eng *gin.Engine) {
	handler := oms_order_return_reason.New()
	group := eng.Group("/returnReason", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("/create", handler.Create)              // 添加退货原因
		group.POST("/update/:id", handler.Update)          // 修改退货原因
		group.GET("/list", handler.List)                   // 分页查询退货原因
		group.POST("/delete", handler.Delete)              // 批量删除退货原因
		group.GET("/:id", handler.GetItem)                 // 获取单个退货原因详情信息
		group.POST("/update/status", handler.UpdateStatus) // 修改退货原因启用状态
	}
}
