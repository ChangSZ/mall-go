package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/oms_order"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 订单管理
func setOmsOrderRouter(eng *gin.Engine) {
	handler := oms_order.New()
	group := eng.Group("/order", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.GET("/list", handler.List)                               // 查询订单
		group.POST("/update/delivery", handler.Delivery)               // 批量发货
		group.POST("/update/close", handler.Close)                     // 批量关闭订单
		group.POST("/delete", handler.Delete)                          // 批量删除订单
		group.GET("/:id", handler.GetItem)                             // 获取订单详情：订单信息、商品信息、操作记录
		group.POST("/update/receiverInfo", handler.UpdateReceiverInfo) // 修改收货人信息
		group.POST("/update/moneyInfo", handler.UpdateMoneyInfo)       // 修改订单费用信息
		group.POST("/update/note", handler.UpdateNote)                 // 备注订单
	}
}
