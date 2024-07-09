package oms_portal_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type cancelOrderRequest struct {
	OrderId int64 `form:"orderId" binding:"required"`
}

type cancelOrderResponse struct{}

// CancelOrder 取消单个超时订单
// @Summary 取消单个超时订单
// @Description 取消单个超时订单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=cancelOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/cancelOrder [post]
func (h *handler) CancelOrder(ctx *gin.Context) {
	req := new(cancelOrderRequest)
	_ = new(cancelOrderResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.SendDelayMessageCancelOrder(ctx, req.OrderId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, nil)
}
