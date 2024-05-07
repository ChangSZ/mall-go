package router

import (
	"github.com/ChangSZ/mall-go/internal/api/sms_coupon_history"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 优惠券领取记录管理
func setSmsCouponHistoryRouter(eng *gin.Engine) {
	handler := sms_coupon_history.New()
	group := eng.Group("/couponHistory", middleware.CheckToken())
	{
		group.GET("/list", handler.List) // 根据优惠券id，使用状态，订单编号分页获取领取记录
	}
}
