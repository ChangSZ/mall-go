package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type cancelTimeOutOrderRequest struct{}

type cancelTimeOutOrderResponse struct{}

// CancelTimeOutOrder 自动取消超时订单
// @Summary 自动取消超时订单
// @Description 自动取消超时订单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelTimeOutOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=cancelTimeOutOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/cancelTimeOutOrder [post]
func (h *handler) CancelTimeOutOrder(ctx *gin.Context) {
	_ = new(cancelTimeOutOrderRequest)
	_ = new(cancelTimeOutOrderResponse)

	_, err := h.service.CancelTimeOutOrder(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, nil)
}
