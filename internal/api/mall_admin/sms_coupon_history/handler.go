package sms_coupon_history

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_admin/sms_coupon_history"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 根据优惠券id，使用状态，订单编号分页获取领取记录
	// @Tags SmsCouponHistoryController
	// @Router /couponHistory/list [get]
	List(*gin.Context)
}

type handler struct {
	service sms_coupon_history.Service
}

func New() Handler {
	return &handler{
		service: sms_coupon_history.New(),
	}
}

func (h *handler) i() {}
