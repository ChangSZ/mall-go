package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/oms_portal_order"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 订单管理
func setOmsPortalOrderRouter(eng *gin.Engine) {
	handler := oms_portal_order.New()
	group := eng.Group("/order", middleware.CheckMemberToken())
	{
		group.POST("/generateConfirmOrder", handler.GenerateConfirmOrder) // 根据购物车信息生成确认单
		group.POST("/generateOrder", handler.GenerateOrder)               // 根据购物车信息生成订单
		group.POST("/paySuccess", handler.PaySuccess)                     // 用户支付成功的回调
		group.POST("/cancelTimeOutOrder", handler.CancelTimeOutOrder)     // 自动取消超时订单
		group.POST("/cancelOrder", handler.CancelOrder)                   // 取消单个超时订单
		group.GET("/list", handler.List)                                  // 按状态分页获取用户订单列表
		group.GET("/detail/:orderId", handler.Detail)                     // 根据ID获取订单详情
		group.POST("/cancelUserOrder", handler.CancelUserOrder)           // 用户取消订单
		group.POST("/confirmReceiveOrder", handler.ConfirmReceiveOrder)   // 用户确认收货
		group.POST("/deleteOrder", handler.DeleteOrder)                   // 用户删除订单
	}
}
