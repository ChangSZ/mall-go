package oms_portal_order

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type detailRequest struct {
	OrderId int64 `uri:"orderId"`
}

type detailResponse struct {
	dto.OrderDetail `json:",inline"`
}

// Detail 根据ID获取订单详情
// @Summary 根据ID获取订单详情
// @Description 根据ID获取订单详情
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} code.Success{data=detailResponse}
// @Failure 400 {object} code.Failure
// @Router /order/detail/{orderId} [get]
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.Detail(ctx, req.OrderId)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.OrderDetail = *data
	api.Success(ctx, res)
}
