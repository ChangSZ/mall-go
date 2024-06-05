package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type cancelUserOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

type cancelUserOrderResponse struct{}

// CancelUserOrder 用户取消订单
// @Summary 用户取消订单
// @Description 用户取消订单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelUserOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=cancelUserOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/cancelUserOrder [post]
func (h *handler) CancelUserOrder(ctx *gin.Context) {
	req := new(cancelUserOrderRequest)
	_ = new(cancelUserOrderResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.CancelOrder(ctx, req.OrderId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, nil)
}
