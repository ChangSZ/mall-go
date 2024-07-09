package ums_member_coupon

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_portal/ums_member_coupon"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Add 领取指定优惠券
	// @Tags UmsMemberCouponController
	// @Router /member/coupon/add/{couponId} [post]
	Add(*gin.Context)

	// ListHistory 获取会员优惠券历史列表
	// @Tags UmsMemberCouponController
	// @Router /member/coupon/listHistory [get]
	ListHistory(*gin.Context)

	// List 获取会员优惠券列表
	// @Tags UmsMemberCouponController
	// @Router /member/coupon/list [get]
	List(*gin.Context)

	// ListCart 获取登录会员购物车的相关优惠券
	// @Tags UmsMemberCouponController
	// @Router /member/coupon/list/cart/{type} [get]
	ListCart(*gin.Context)

	// ListByProduct 获取当前商品相关优惠券
	// @Tags UmsMemberCouponController
	// @Router /member/coupon/listByProduct/{productId} [get]
	ListByProduct(*gin.Context)
}

type handler struct {
	service ums_member_coupon.Service
}

func New() Handler {
	return &handler{
		service: ums_member_coupon.New(),
	}
}

func (h *handler) i() {}
