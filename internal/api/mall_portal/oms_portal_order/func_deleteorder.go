package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type deleteOrderRequest struct {
	OrderId int64 `json:"orderId"`
}

type deleteOrderResponse struct{}

// DeleteOrder 用户删除订单
// @Summary 用户删除订单
// @Description 用户删除订单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/deleteOrder [post]
func (h *handler) DeleteOrder(ctx *gin.Context) {
	req := new(deleteOrderRequest)
	_ = new(deleteOrderResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.DeleteOrder(ctx, req.OrderId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, nil)
}
