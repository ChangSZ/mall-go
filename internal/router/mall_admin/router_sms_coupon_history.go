package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_coupon_history"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 优惠券领取记录管理
func setSmsCouponHistoryRouter(eng *gin.Engine) {
	handler := sms_coupon_history.New()
	group := eng.Group("/couponHistory", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.GET("/list", handler.List) // 根据优惠券id，使用状态，订单编号分页获取领取记录
	}
}
