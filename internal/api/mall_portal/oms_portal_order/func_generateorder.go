package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type generateOrderRequest struct {
	dto.OrderParam `json:",inline"`
}

type generateOrderResponse struct {
	dto.Order `json:",inline"`
}

// GenerateOrder 根据购物车信息生成订单
// @Summary 根据购物车信息生成订单
// @Description 根据购物车信息生成订单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body generateOrderRequest true "请求信息"
// @Success 200 {object} code.Success{data=generateOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/generateOrder [post]
func (h *handler) GenerateOrder(ctx *gin.Context) {
	req := new(generateOrderRequest)
	res := new(generateOrderResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.GenerateOrder(ctx, req.OrderParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Order = *data
	api.Success(ctx, res)
}
