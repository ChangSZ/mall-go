package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/oms_order_setting"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 订单设置管理
func setOmsOrderSettingRouter(eng *gin.Engine) {
	handler := oms_order_setting.New()
	group := eng.Group("/orderSetting", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.GET("/:id", handler.GetItem)        // 获取指定订单设置
		group.POST("/update/:id", handler.Update) // 修改指定订单设置
	}
}
