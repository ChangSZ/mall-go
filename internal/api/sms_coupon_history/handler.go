package sms_coupon_history

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_coupon_history"

	"github.com/gin-gonic/gin"
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
	smsCouponHistoryService sms_coupon_history.Service
}

func New() Handler {
	return &handler{
		smsCouponHistoryService: sms_coupon_history.New(),
	}
}

func (h *handler) i() {}
