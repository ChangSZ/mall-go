package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/sms_coupon"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 优惠券管理
func setSmsCouponRouter(eng *gin.Engine) {
	handler := sms_coupon.New()
	group := eng.Group("/coupon", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("", handler.Create)            // 添加优惠券
		group.POST("/update/:id", handler.Update) // 修改优惠券
		group.POST("/delete/:id", handler.Delete) // 修改优惠券
		group.GET("/list", handler.List)          // 根据优惠券名称和类型分页获取优惠券列表
		group.GET("/:id", handler.GetItem)        // 获取单个优惠券的详细信息
	}
}
