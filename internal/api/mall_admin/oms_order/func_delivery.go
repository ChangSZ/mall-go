package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type deliveryRequest struct{}

type deliveryResponse struct {
	Count int64 `json:",inline"`
}

// Delivery 批量发货
// @Summary 批量发货
// @Description 批量发货
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deliveryRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/update/delivery [post]
func (h *handler) Delivery(ctx *gin.Context) {
	_ = new(deliveryRequest)
	req := make([]dto.OmsOrderDeliveryParam, 0)
	res := new(deliveryResponse)
	if err := ctx.ShouldBind(&req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Delivery(ctx, req)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
