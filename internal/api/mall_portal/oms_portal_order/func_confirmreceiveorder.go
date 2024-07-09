package oms_portal_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type confirmReceiveOrderRequest struct {
	OrderId int64 `form:"orderId" binding:"required"`
}

type confirmReceiveOrderResponse struct{}

// ConfirmReceiveOrder 用户确认收货
// @Summary 用户确认收货
// @Description 用户确认收货
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body confirmReceiveOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=confirmReceiveOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/confirmReceiveOrder [post]
func (h *handler) ConfirmReceiveOrder(ctx *gin.Context) {
	req := new(confirmReceiveOrderRequest)
	_ = new(confirmReceiveOrderResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.ConfirmReceiveOrder(ctx, req.OrderId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, nil)
}
