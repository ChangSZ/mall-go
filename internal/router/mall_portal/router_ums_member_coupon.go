package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/ums_member_coupon"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 会员优惠券管理
func setUmsMemberCouponRouter(eng *gin.Engine) {
	handler := ums_member_coupon.New()
	group := eng.Group("/member/coupon", middleware.CheckMemberToken())
	{
		group.POST("/add/:couponId", handler.Add)                     // 领取指定优惠券
		group.GET("/listHistory", handler.ListHistory)                // 获取会员优惠券历史列表
		group.GET("/list", handler.List)                              // 获取会员优惠券列表
		group.GET("/list/cart/:type", handler.ListCart)               // 获取登录会员购物车的相关优惠券
		group.GET("/listByProduct/:productId", handler.ListByProduct) // 获取当前商品相关优惠券
	}
}
