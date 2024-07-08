package oms_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type updateMoneyInfoRequest struct {
	dto.OmsMoneyInfoParam `json:",inline"`
}

type updateMoneyInfoResponse struct {
	Count int64 `json:",inline"`
}

// UpdateMoneyInfo 修改订单费用信息
// @Summary 修改订单费用信息
// @Description 修改订单费用信息
// @Tags OmsOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateMoneyInfoRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /order/update/moneyInfo [post]
func (h *handler) UpdateMoneyInfo(ctx *gin.Context) {
	req := new(updateMoneyInfoRequest)
	res := new(updateMoneyInfoResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.UpdateMoneyInfo(ctx, req.OmsMoneyInfoParam)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
