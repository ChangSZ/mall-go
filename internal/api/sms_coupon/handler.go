package sms_coupon

import (
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加优惠券
	// @Tags SmsCouponController
	// @Router /coupon [post]
	Create(*gin.Context)

	// Update 修改优惠券
	// @Tags SmsCouponController
	// @Router /coupon/update/{id} [post]
	Update(*gin.Context)

	// Delete 删除优惠券
	// @Tags SmsCouponController
	// @Router /coupon/delete/{id} [post]
	Delete(*gin.Context)

	// List 根据优惠券名称和类型分页获取优惠券列表
	// @Tags SmsCouponController
	// @Router /coupon/list [get]
	List(*gin.Context)

	// GetItem 获取单个优惠券的详细信息
	// @Tags SmsCouponController
	// @Router /coupon/{id} [get]
	GetItem(*gin.Context)
}

type handler struct {
	// smsCouponService sms_coupon.Service
}

func New() Handler {
	return &handler{
		// smsCouponService: sms_coupon.New(),
	}
}

func (h *handler) i() {}
