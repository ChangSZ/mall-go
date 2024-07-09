package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/alipay"
)

// 支付宝支付
func setAlipayRouter(eng *gin.Engine) {
	handler := alipay.New()
	group := eng.Group("/alipay")
	{
		group.GET("/pay", handler.Pay)        // 支付宝电脑网站支付
		group.GET("/webPay", handler.WebPay)  // 支付宝手机网站支付
		group.POST("/notify", handler.Notify) // 支付宝异步回调
		group.GET("/query", handler.Query)    // 支付宝统一收单线下交易查询
	}
}
